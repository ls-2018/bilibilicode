package service

import (
	"context"
	"encoding/csv"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"go-common/app/admin/main/answer/model"
	"go-common/library/ecode"
	"go-common/library/log"
)

// QuestionList .
func (s *Service) QuestionList(c context.Context, arg *model.ArgQue) (res *model.QuestionPage, err error) {
	res = &model.QuestionPage{}
	if res.Total, err = s.dao.QuestionCount(c, arg); err != nil {
		return
	}
	res.Items = []*model.QuestionDB{}
	if res.Total > 0 {
		if res.Items, err = s.dao.QuestionList(c, arg); err != nil {
			return
		}
	}
	return
}

// UpdateStatus update question state
func (s *Service) UpdateStatus(c context.Context, qid int64, state int8, operator string) (err error) {
	var (
		r int64
		q *model.Question
	)
	if q, err = s.dao.QueByID(c, qid); err != nil {
		log.Error("dao QueByID(%d) error(%v)", qid, err)
		return
	}
	if q == nil || q.State == state {
		return
	}
	if r, err = s.dao.UpdateStatus(c, state, qid, operator); err != nil || r != 1 {
		return
	}
	return
}

// BatchUpdateState bacth update question state.
func (s *Service) BatchUpdateState(c context.Context, qids []int64, state int8, operator string) (err error) {
	for _, id := range qids {
		s.UpdateStatus(c, id, state, operator)
	}
	return
}

// Types question type
func (s *Service) Types(c context.Context) (res []*model.TypeInfo, err error) {
	return s.dao.Types(c)
}

// ReadCsv read csv file
func (s *Service) ReadCsv(f multipart.File, h *multipart.FileHeader) (rs [][]string, err error) {
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error("upload question ReadCsv error(%v)", err)
			break
		}
		if len(record) == model.ArgsCount {
			rs = append(rs, record)
		}
	}
	return
}

// UploadQsts upload questions
func (s *Service) UploadQsts(c context.Context, f multipart.File, h *multipart.FileHeader, operator string) (msg string, err error) {
	defer f.Close()
	if h != nil && !strings.HasSuffix(h.Filename, ".csv") {
		msg = "not csv file."
		return
	}
	sz, ok := f.(model.Sizer)
	if !ok {
		msg = "get file size faild."
		return
	}
	size := sz.Size()
	if size > model.FileMaxSize {
		msg = "file size more than 2M."
		return
	}
	rs, err := s.ReadCsv(f, h)
	log.Info("file %s, len(%d)", h.Filename, len(rs))
	if len(rs) == 0 || len(rs) > model.MaxCount {
		msg = "file size count is 0 or more than " + strconv.FormatInt(model.MaxCount, 10)
		return
	}
	for _, r := range rs {
		typeID, err := strconv.ParseInt(r[0], 10, 8)
		if err != nil {
			log.Error("strconv.ParseInt(%+v) err(%v)", r[0], err)
		}
		if err == nil {
			q := &model.QuestionDB{
				TypeID:   int8(typeID),
				Question: r[1],
				Ans1:     r[2],
				Ans2:     r[3],
				Ans3:     r[4],
				Ans4:     r[5],
				Operator: operator,
			}
			if err = s.QuestionAdd(c, q); err != nil {
				log.Error("s.QuestionAdd(%+v) error(%v)", q, err)
			}
		}
	}
	return
}

// QuestionAdd add register question
func (s *Service) QuestionAdd(c context.Context, q *model.QuestionDB) (err error) {
	if len(q.Question) < model.MinQuestion || len(q.Question) > model.MaxQuestion {
		err = ecode.QuestionStrNotAllow
		return
	}
	if len(q.Ans1) < model.MinAns || len(q.Ans1) > model.MaxAns ||
		len(q.Ans2) < model.MinAns || len(q.Ans2) > model.MaxAns ||
		len(q.Ans3) < model.MinAns || len(q.Ans3) > model.MaxAns ||
		len(q.Ans4) < model.MinAns || len(q.Ans4) > model.MaxAns {
		err = ecode.QuestionAnsNotAllow
		return
	}
	if q.Tips != "" && (len(q.Tips) < model.MinTips || len(q.Tips) > model.MaxTips) {
		err = ecode.QuestionTipsNotAllow
		return
	}
	if q.TypeID <= 0 {
		err = ecode.QuestionTypeNotAllow
		return
	}
	// only sourport text question
	q.MediaType = model.TextMediaType
	q.State = model.PassCheck
	q.Ctime = time.Now()
	if _, err = s.dao.QuestionAdd(c, q); err != nil {
		return
	}
	qid := q.ID
	s.eventChan.Save(func() {
		s.CreateBFSImg(context.Background(), []int64{qid})
	})
	return
}

func (s *Service) loadtypes(c context.Context) (t map[int64]*model.TypeInfo, err error) {
	var tys []*model.TypeInfo
	tys, err = s.dao.Types(c)
	if err != nil {
		log.Error("s.dao.Types error(%v)", err)
		return
	}
	t = make(map[int64]*model.TypeInfo)
	for _, v := range tys {
		if v.Parentid == 0 && t[v.ID] == nil {
			t[v.ID] = &model.TypeInfo{ID: v.ID, Name: v.Name, Subs: []*model.SubType{}}
		} else if t[v.Parentid] != nil {
			t[v.Parentid].Subs = append(t[v.Parentid].Subs, &model.SubType{ID: v.ID, Name: v.Name, LabelName: v.LabelName})
		}
	}
	return
}

// QuestionEdit .
func (s *Service) QuestionEdit(c context.Context, arg *model.QuestionDB) (aff int64, err error) {
	if aff, err = s.dao.QuestionEdit(c, arg); err != nil {
		return
	}
	s.eventChan.Save(func() {
		s.CreateBFSImg(context.Background(), []int64{arg.ID})
	})
	return
}

// LoadTypes .
func (s *Service) LoadTypes(c context.Context) (err error) {
	var allType = []*model.TypeInfo{
		{ID: 1, Parentid: 0, Name: "??????"},
		{ID: 2, Parentid: 0, Name: "??????"},
		{ID: 3, Parentid: 0, Name: "??????"},
		{ID: 4, Parentid: 0, Name: "??????"},
		{ID: 5, Parentid: 0, Name: "??????"},
		{ID: 6, Parentid: 0, Name: "????????????"},
		{ID: 7, Parentid: 0, Name: "??????"},
		{ID: 8, Parentid: 1, Name: "????????????", LabelName: "??????"},
		{ID: 9, Parentid: 1, Name: "????????????", LabelName: "??????"},
		{ID: 12, Parentid: 1, Name: "???????????? ", LabelName: "??????"},
		{ID: 13, Parentid: 1, Name: "???????????? ", LabelName: "??????"},
		{ID: 14, Parentid: 1, Name: "???????????? ", LabelName: "??????"},
		{ID: 15, Parentid: 2, Name: "????????? ", LabelName: "??????"},
		{ID: 16, Parentid: 2, Name: "?????? ", LabelName: "??????"},
		{ID: 17, Parentid: 2, Name: "????????? ", LabelName: "??????"},
		{ID: 18, Parentid: 3, Name: "?????? ", LabelName: "??????"},
		{ID: 19, Parentid: 3, Name: "?????? ", LabelName: "??????"},
		{ID: 20, Parentid: 3, Name: "?????? ", LabelName: "??????"},
		{ID: 21, Parentid: 3, Name: "?????? ", LabelName: "??????"},
		{ID: 22, Parentid: 3, Name: "?????? ", LabelName: "??????"},
		{ID: 23, Parentid: 3, Name: "?????? ", LabelName: "??????"},
		{ID: 24, Parentid: 3, Name: "?????? ", LabelName: "??????"},
		{ID: 25, Parentid: 3, Name: "?????? ", LabelName: "??????"},
		{ID: 26, Parentid: 3, Name: "???????????? ", LabelName: "??????"},
		{ID: 27, Parentid: 4, Name: "???????????? ", LabelName: "??????"},
		{ID: 28, Parentid: 4, Name: "???????????? ", LabelName: "??????"},
		{ID: 29, Parentid: 5, Name: "ACG?????? ", LabelName: "??????"},
		{ID: 30, Parentid: 5, Name: "??????????????? ", LabelName: "??????"},
		{ID: 31, Parentid: 5, Name: "?????? ", LabelName: "??????"},
		{ID: 32, Parentid: 6, Name: "?????? ", LabelName: "????????????"},
		{ID: 33, Parentid: 6, Name: "?????? ", LabelName: "????????????"},
		{ID: 34, Parentid: 6, Name: "?????? ", LabelName: "????????????"},
		{ID: 35, Parentid: 7, Name: "?????? ", LabelName: "??????"},
		{ID: 36, Parentid: 0, Name: "?????????", LabelName: "?????????"},
	}
	for _, v := range allType {
		if _, err := s.dao.TypeSave(context.Background(), v); err != nil {
			log.Error("s.dao.TypeSave(%+v) err(%v)", v, err)
		}
	}
	return
}

// LoadImg .
func (s *Service) LoadImg(c context.Context) (err error) {
	qss, err := s.dao.AllQS(c)
	if err != nil {
		log.Error("s.dao.AllQS() err(%v)", err)
	}
	for _, qs := range qss {
		lastID := qs.ID
		if err = s.eventChan.Save(func() {
			s.CreateBFSImg(context.Background(), []int64{lastID})
		}); err != nil {
			log.Error("s.CreateBFSImg(%d) err(%v)", lastID, err)
		}
	}
	return
}

// QueHistory .
func (s *Service) QueHistory(c context.Context, arg *model.ArgHistory) (res *model.HistoryPage, err error) {
	res = &model.HistoryPage{}
	if res.Total, err = s.dao.HistoryCount(c, arg); err != nil {
		return
	}
	res.Items = []*model.AnswerHistoryDB{}
	if res.Total > 0 {
		if res.Items, err = s.dao.QueHistory(c, arg); err != nil {
			return
		}
	}
	return
}

// History .
func (s *Service) History(c context.Context, arg *model.ArgHistory) (res *model.HistoryPage, err error) {
	if arg.Pn <= 0 || arg.Ps <= 0 {
		arg.Pn, arg.Ps = 1, 1000
	}
	res = &model.HistoryPage{}
	if res.Items, err = s.dao.HistoryES(c, arg); err != nil {
		return
	}
	res.Total = int64(len(res.Items))
	return
}
