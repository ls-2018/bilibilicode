package lancerlogstream

import (
	"net"
	"time"
	"sync"
	"errors"
	"math/rand"
	"expvar"

	"go-common/library/log"
	"go-common/app/service/ops/log-agent/pkg/bufio"
	xtime "go-common/library/time"
)

var (
	ErrAddrListNil = errors.New("addrList can't be nil")
	ErrPoolSize    = errors.New("Pool size should be no greater then length of addr list")
)

type LancerBufConn struct {
	conn    net.Conn
	wr      *bufio.Writer
	enabled bool
	ctime   time.Time
}

type connPool struct {
	c                    *ConnPoolConfig
	invalidUpstreams     map[string]interface{}
	invalidUpstreamsLock sync.RWMutex
	validBufConnChan     chan *LancerBufConn
	invalidBufConnChan   chan *LancerBufConn
	connCounter          map[string]int
	connCounterLock      sync.RWMutex
	newConnLock          sync.Mutex
}

type ConnPoolConfig struct {
	Name        string         `tome:"name"`
	AddrList    []string       `tome:"addrList"`
	DialTimeout xtime.Duration `tome:"dialTimeout"`
	IdleTimeout xtime.Duration `tome:"idleTimeout"`
	BufSize     int            `tome:"bufSize"`
	PoolSize    int            `tome:"poolSize"`
}

func (c *ConnPoolConfig) ConfigValidate() (error) {
	if c == nil {
		return errors.New("Config of pool is nil")
	}

	if len(c.AddrList) == 0 {
		return errors.New("pool addr list can't be empty")
	}
	if time.Duration(c.DialTimeout) == 0 {
		c.DialTimeout = xtime.Duration(time.Second * 5)
	}

	if time.Duration(c.IdleTimeout) == 0 {
		c.IdleTimeout = xtime.Duration(time.Minute * 15)
	}

	if c.BufSize == 0 {
		c.BufSize = 1024 * 1024 * 2 // 2M by default
	}

	if c.PoolSize == 0 {
		c.PoolSize = len(c.AddrList)
	}
	return nil
}

// newConn make a connection to lancer
func (cp *connPool) newConn() (conn net.Conn, err error) {
	cp.newConnLock.Lock()
	defer cp.newConnLock.Unlock()
	for {
		if addr, err := cp.randomOneUpstream(); err == nil {
			if conn, err := net.DialTimeout("tcp", addr, time.Duration(cp.c.DialTimeout)); err == nil && conn != nil {
				log.Info("connect to %s success", addr)
				cp.connCounterAdd(addr)
				return conn, nil
			} else {
				cp.markUpstreamInvalid(addr)
				continue
			}
		} else {
			return nil, err
		}
	}
}

// newBufConn ????????????buf??????, buf??????????????????conn(????????????????????????)
func (cp *connPool) newBufConn() (bufConn *LancerBufConn, err error) {
	bufConn = new(LancerBufConn)
	bufConn.wr = bufio.NewWriterSize(nil, cp.c.BufSize)
	if err := cp.setConn(bufConn); err == nil {
		bufConn.enabled = true
	} else {
		bufConn.enabled = false
	}

	return bufConn, nil
}

// flushBufConn ??????flush buffer
func (cp *connPool) flushBufConn() {
	for {
		bufConn, _ := cp.getBufConn()
		bufConn.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
		if err := bufConn.wr.Flush(); err != nil {
			log.Error("Error when flush to %s: %s", bufConn.conn.RemoteAddr().String(), err)
			bufConn.enabled = false
		}
		cp.putBufConn(bufConn)
		time.Sleep(time.Second * 5)
	}
}

// initConnPool ?????????conn pool??????
func initConnPool(c *ConnPoolConfig) (cp *connPool, err error) {
	if err = c.ConfigValidate(); err != nil {
		return nil, err
	}
	if len(c.AddrList) == 0 {
		return nil, ErrAddrListNil
	}

	if c.PoolSize > len(c.AddrList) {
		return nil, ErrPoolSize
	}

	rand.Seed(time.Now().Unix())
	cp = new(connPool)
	cp.c = c
	cp.validBufConnChan = make(chan *LancerBufConn, cp.c.PoolSize)
	cp.invalidBufConnChan = make(chan *LancerBufConn, cp.c.PoolSize)
	cp.invalidUpstreams = make(map[string]interface{})
	cp.connCounter = make(map[string]int)
	cp.initPool()
	go cp.maintainUpstream()
	go cp.flushBufConn()
	go cp.maintainBufConnPool()
	expvar.Publish("conn_pool"+cp.c.Name, expvar.Func(cp.connPoolStatus))
	return cp, nil
}

// connableUpstreams ???????????????????????????upstream??????
func (cp *connPool) connableUpstreams() ([]string) {
	list := make([]string, 0)
	cp.invalidUpstreamsLock.RLock()
	defer cp.invalidUpstreamsLock.RUnlock()
	for _, addr := range cp.c.AddrList {
		if _, ok := cp.invalidUpstreams[addr]; !ok {
			if count, ok := cp.connCounter[addr]; ok && count == 0 {
				list = append(list, addr)
			}
		}
	}
	return list
}

// write write []byte to BufConn
func (bc *LancerBufConn) write(p []byte) (int, error) {
	bc.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	return bc.wr.Write(p)
}

// randomOneUpstream ???????????????????????????????????????upstream
func (cp *connPool) randomOneUpstream() (s string, err error) {
	list := cp.connableUpstreams()
	if len(list) == 0 {
		err = errors.New("No valid upstreams")
		return
	}
	return list[rand.Intn(len(list))], nil
}

// initPool ?????????poolSize?????????bufConn
func (cp *connPool) initPool() {
	for _, addr := range cp.c.AddrList {
		cp.connCounter[addr] = 0
	}
	for i := 0; i < cp.c.PoolSize; i++ {
		if bufConn, err := cp.newBufConn(); err == nil {
			cp.putBufConn(bufConn)
		}
	}
}

// novalidUpstream check if there is no validUpstream
func (cp *connPool) novalidUpstream() bool {
	return len(cp.invalidUpstreams) == len(cp.c.AddrList)
}

//GetConn ???pool????????????BufConn
func (cp *connPool) getBufConn() (*LancerBufConn, error) {
	for {
		select {
		case bufConn := <-cp.validBufConnChan:
			if !bufConn.enabled {
				cp.putInvalidBufConn(bufConn)
				continue
			}
			return bufConn, nil
		case <-time.After(10 * time.Second):
			log.Warn("timeout when get conn from conn pool")
			continue
		}
	}
}

// setConn ???bufConn??????????????????Conn
func (cp *connPool) setConn(bufConn *LancerBufConn) (error) {
	if bufConn.conn != nil {
		if bufConn.enabled == false {
			cp.markUpstreamInvalid(bufConn.conn.RemoteAddr().String())
		}
		cp.connCounterDel(bufConn.conn.RemoteAddr().String())
		bufConn.conn.Close()
		bufConn.conn = nil
		bufConn.enabled = false
	}
	if conn, err := cp.newConn(); err == nil {
		bufConn.conn = conn
		bufConn.wr.Reset(conn)
		bufConn.ctime = time.Now()
		bufConn.enabled = true
		return nil
	} else {
		bufConn.enabled = false
		return err
	}
}

//putBufConn ???BufConn?????????pool???
func (cp *connPool) putBufConn(bufConn *LancerBufConn) {
	if bufConn.enabled == false {
		cp.putInvalidBufConn(bufConn)
		return
	}
	if bufConn.ctime.Add(time.Duration(cp.c.IdleTimeout)).Before(time.Now()) {
		bufConn.wr.Flush()
		cp.putInvalidBufConn(bufConn)
		return
	}
	cp.putValidBufConn(bufConn)
}

// putValidBufConn ??? bufConn???????????????pool???
func (cp *connPool) putValidBufConn(bufConn *LancerBufConn) {
	select {
	case cp.validBufConnChan <- bufConn:
		return
	default:
		log.Warn("BufConnChan full, discard???this shouldn't happen")
		return
	}
}

// putInvalidBufConn ???bufConn??????????????????pool???
func (cp *connPool) putInvalidBufConn(bufConn *LancerBufConn) {
	select {
	case cp.invalidBufConnChan <- bufConn:
		return
	default:
		log.Warn("invalidBufConnChan full, discard???this shouldn't happen")
		return
	}
}

// maintainBufConnPool ??????BufConnPool??????
func (cp *connPool) maintainBufConnPool() {
	for {
		select {
		case bufConn := <-cp.invalidBufConnChan:
			cp.setConn(bufConn)
			cp.putBufConn(bufConn)
		}
		time.Sleep(time.Second * 1)
	}
}

//markConnInvalid?????????????????????????????????upstreamserver??????????????????
func (cp *connPool) markUpstreamInvalid(addr string) (err error) {
	log.Error("mark upstream %s invalid", addr)
	cp.invalidUpstreamsLock.Lock()
	cp.invalidUpstreams[addr] = nil
	cp.invalidUpstreamsLock.Unlock()
	return
}

// markUpstreamValid ?????????addr??????????????????
func (cp *connPool) markUpstreamValid(addr string) (err error) {
	log.Info("%s is valid again", addr)
	cp.invalidUpstreamsLock.Lock()
	delete(cp.invalidUpstreams, addr)
	cp.invalidUpstreamsLock.Unlock()
	return
}

// connCounterAdd ?????????+1
func (cp *connPool) connCounterAdd(addr string) {
	cp.connCounterLock.Lock()
	defer cp.connCounterLock.Unlock()
	if _, ok := cp.connCounter[addr]; ok {
		cp.connCounter[addr] += 1
	} else {
		cp.connCounter[addr] = 1
	}
	return
}

//connCounterDel ?????????-1
func (cp *connPool) connCounterDel(addr string) {
	cp.connCounterLock.Lock()
	defer cp.connCounterLock.Unlock()
	if _, ok := cp.connCounter[addr]; ok {
		cp.connCounter[addr] -= 1
	}
}

// connPoolStatus ??????connPool??????
func (cp *connPool) connPoolStatus() interface{} {
	status := make(map[string]interface{})
	status["conn_num"] = cp.connCounter
	status["invalidUpstreams"] = cp.invalidUpstreams
	return status
}

// maintainUpstream ??????upstream???????????????
func (cp *connPool) maintainUpstream() {
	for {
		cp.invalidUpstreamsLock.RLock()
		tryAddrs := make([]string, 0, len(cp.invalidUpstreams))
		for k := range cp.invalidUpstreams {
			tryAddrs = append(tryAddrs, k)
		}
		cp.invalidUpstreamsLock.RUnlock()
		for _, addr := range tryAddrs {
			if conn, err := net.DialTimeout("tcp", addr, time.Duration(cp.c.DialTimeout)); err == nil && conn != nil {
				conn.Close()
				cp.markUpstreamValid(addr)
			}
		}
		time.Sleep(time.Second * 10)
	}
}

//ReleaseConnPool ??????????????????????????????
func (cp *connPool) ReleaseConnPool() {
	log.Info("Release Conn Pool")
	close(cp.validBufConnChan)
	close(cp.invalidBufConnChan)
	for conn := range cp.validBufConnChan {
		conn.enabled = false
		conn.wr.Flush()
		conn.conn.Close()
	}
}
