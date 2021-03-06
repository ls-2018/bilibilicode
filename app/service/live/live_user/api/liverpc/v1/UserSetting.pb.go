// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/UserSetting.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UserSettingGetTagReq struct {
}

func (m *UserSettingGetTagReq) Reset()         { *m = UserSettingGetTagReq{} }
func (m *UserSettingGetTagReq) String() string { return proto.CompactTextString(m) }
func (*UserSettingGetTagReq) ProtoMessage()    {}
func (*UserSettingGetTagReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserSetting_6ecfe772383476d5, []int{0}
}
func (m *UserSettingGetTagReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserSettingGetTagReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserSettingGetTagReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserSettingGetTagReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingGetTagReq.Merge(dst, src)
}
func (m *UserSettingGetTagReq) XXX_Size() int {
	return m.Size()
}
func (m *UserSettingGetTagReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingGetTagReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingGetTagReq proto.InternalMessageInfo

type UserSettingGetTagResp struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	Data *UserSettingGetTagResp_Data `protobuf:"bytes,3,opt,name=data" json:"data"`
}

func (m *UserSettingGetTagResp) Reset()         { *m = UserSettingGetTagResp{} }
func (m *UserSettingGetTagResp) String() string { return proto.CompactTextString(m) }
func (*UserSettingGetTagResp) ProtoMessage()    {}
func (*UserSettingGetTagResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserSetting_6ecfe772383476d5, []int{1}
}
func (m *UserSettingGetTagResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserSettingGetTagResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserSettingGetTagResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserSettingGetTagResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingGetTagResp.Merge(dst, src)
}
func (m *UserSettingGetTagResp) XXX_Size() int {
	return m.Size()
}
func (m *UserSettingGetTagResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingGetTagResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingGetTagResp proto.InternalMessageInfo

func (m *UserSettingGetTagResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UserSettingGetTagResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UserSettingGetTagResp) GetData() *UserSettingGetTagResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type UserSettingGetTagResp_Tags struct {
	// tag id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// tag name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// ????????????
	Pic string `protobuf:"bytes,3,opt,name=pic,proto3" json:"pic"`
	// ????????????????????? 1??? 0???
	IsAdvice int64 `protobuf:"varint,4,opt,name=is_advice,json=isAdvice,proto3" json:"is_advice"`
	// ?????????id
	ParentId int64 `protobuf:"varint,5,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	// ????????????
	ParentName string `protobuf:"bytes,6,opt,name=parent_name,json=parentName,proto3" json:"parent_name"`
	// ??????id
	ActId int64 `protobuf:"varint,7,opt,name=act_id,json=actId,proto3" json:"act_id"`
}

func (m *UserSettingGetTagResp_Tags) Reset()         { *m = UserSettingGetTagResp_Tags{} }
func (m *UserSettingGetTagResp_Tags) String() string { return proto.CompactTextString(m) }
func (*UserSettingGetTagResp_Tags) ProtoMessage()    {}
func (*UserSettingGetTagResp_Tags) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserSetting_6ecfe772383476d5, []int{1, 0}
}
func (m *UserSettingGetTagResp_Tags) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserSettingGetTagResp_Tags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserSettingGetTagResp_Tags.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserSettingGetTagResp_Tags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingGetTagResp_Tags.Merge(dst, src)
}
func (m *UserSettingGetTagResp_Tags) XXX_Size() int {
	return m.Size()
}
func (m *UserSettingGetTagResp_Tags) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingGetTagResp_Tags.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingGetTagResp_Tags proto.InternalMessageInfo

func (m *UserSettingGetTagResp_Tags) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserSettingGetTagResp_Tags) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserSettingGetTagResp_Tags) GetPic() string {
	if m != nil {
		return m.Pic
	}
	return ""
}

func (m *UserSettingGetTagResp_Tags) GetIsAdvice() int64 {
	if m != nil {
		return m.IsAdvice
	}
	return 0
}

func (m *UserSettingGetTagResp_Tags) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *UserSettingGetTagResp_Tags) GetParentName() string {
	if m != nil {
		return m.ParentName
	}
	return ""
}

func (m *UserSettingGetTagResp_Tags) GetActId() int64 {
	if m != nil {
		return m.ActId
	}
	return 0
}

type UserSettingGetTagResp_OfflineTags struct {
	// ??????id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// ????????????
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
}

func (m *UserSettingGetTagResp_OfflineTags) Reset()         { *m = UserSettingGetTagResp_OfflineTags{} }
func (m *UserSettingGetTagResp_OfflineTags) String() string { return proto.CompactTextString(m) }
func (*UserSettingGetTagResp_OfflineTags) ProtoMessage()    {}
func (*UserSettingGetTagResp_OfflineTags) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserSetting_6ecfe772383476d5, []int{1, 1}
}
func (m *UserSettingGetTagResp_OfflineTags) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserSettingGetTagResp_OfflineTags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserSettingGetTagResp_OfflineTags.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserSettingGetTagResp_OfflineTags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingGetTagResp_OfflineTags.Merge(dst, src)
}
func (m *UserSettingGetTagResp_OfflineTags) XXX_Size() int {
	return m.Size()
}
func (m *UserSettingGetTagResp_OfflineTags) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingGetTagResp_OfflineTags.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingGetTagResp_OfflineTags proto.InternalMessageInfo

func (m *UserSettingGetTagResp_OfflineTags) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserSettingGetTagResp_OfflineTags) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserSettingGetTagResp_Data struct {
	// ??????????????????
	Tags []*UserSettingGetTagResp_Tags `protobuf:"bytes,1,rep,name=tags" json:"tags"`
	// ????????????????????????
	Offline []*UserSettingGetTagResp_OfflineTags `protobuf:"bytes,2,rep,name=offline" json:"offline"`
	// ??????id
	Uid int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid"`
	// ???????????? 1???????????? 0????????????????????????
	IsGray int64 `protobuf:"varint,4,opt,name=is_gray,json=isGray,proto3" json:"is_gray"`
}

func (m *UserSettingGetTagResp_Data) Reset()         { *m = UserSettingGetTagResp_Data{} }
func (m *UserSettingGetTagResp_Data) String() string { return proto.CompactTextString(m) }
func (*UserSettingGetTagResp_Data) ProtoMessage()    {}
func (*UserSettingGetTagResp_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_UserSetting_6ecfe772383476d5, []int{1, 2}
}
func (m *UserSettingGetTagResp_Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserSettingGetTagResp_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserSettingGetTagResp_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserSettingGetTagResp_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSettingGetTagResp_Data.Merge(dst, src)
}
func (m *UserSettingGetTagResp_Data) XXX_Size() int {
	return m.Size()
}
func (m *UserSettingGetTagResp_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSettingGetTagResp_Data.DiscardUnknown(m)
}

var xxx_messageInfo_UserSettingGetTagResp_Data proto.InternalMessageInfo

func (m *UserSettingGetTagResp_Data) GetTags() []*UserSettingGetTagResp_Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *UserSettingGetTagResp_Data) GetOffline() []*UserSettingGetTagResp_OfflineTags {
	if m != nil {
		return m.Offline
	}
	return nil
}

func (m *UserSettingGetTagResp_Data) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserSettingGetTagResp_Data) GetIsGray() int64 {
	if m != nil {
		return m.IsGray
	}
	return 0
}

func init() {
	proto.RegisterType((*UserSettingGetTagReq)(nil), "live_user.v1.UserSettingGetTagReq")
	proto.RegisterType((*UserSettingGetTagResp)(nil), "live_user.v1.UserSettingGetTagResp")
	proto.RegisterType((*UserSettingGetTagResp_Tags)(nil), "live_user.v1.UserSettingGetTagResp.Tags")
	proto.RegisterType((*UserSettingGetTagResp_OfflineTags)(nil), "live_user.v1.UserSettingGetTagResp.OfflineTags")
	proto.RegisterType((*UserSettingGetTagResp_Data)(nil), "live_user.v1.UserSettingGetTagResp.Data")
}
func (m *UserSettingGetTagReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserSettingGetTagReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserSettingGetTagResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserSettingGetTagResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.Data != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(m.Data.Size()))
		n1, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *UserSettingGetTagResp_Tags) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserSettingGetTagResp_Tags) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(m.Id))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Pic) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(len(m.Pic)))
		i += copy(dAtA[i:], m.Pic)
	}
	if m.IsAdvice != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(m.IsAdvice))
	}
	if m.ParentId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(m.ParentId))
	}
	if len(m.ParentName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(len(m.ParentName)))
		i += copy(dAtA[i:], m.ParentName)
	}
	if m.ActId != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(m.ActId))
	}
	return i, nil
}

func (m *UserSettingGetTagResp_OfflineTags) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserSettingGetTagResp_OfflineTags) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(m.Id))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *UserSettingGetTagResp_Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserSettingGetTagResp_Data) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for _, msg := range m.Tags {
			dAtA[i] = 0xa
			i++
			i = encodeVarintUserSetting(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Offline) > 0 {
		for _, msg := range m.Offline {
			dAtA[i] = 0x12
			i++
			i = encodeVarintUserSetting(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Uid != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(m.Uid))
	}
	if m.IsGray != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintUserSetting(dAtA, i, uint64(m.IsGray))
	}
	return i, nil
}

func encodeVarintUserSetting(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserSettingGetTagReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UserSettingGetTagResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovUserSetting(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovUserSetting(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovUserSetting(uint64(l))
	}
	return n
}

func (m *UserSettingGetTagResp_Tags) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUserSetting(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUserSetting(uint64(l))
	}
	l = len(m.Pic)
	if l > 0 {
		n += 1 + l + sovUserSetting(uint64(l))
	}
	if m.IsAdvice != 0 {
		n += 1 + sovUserSetting(uint64(m.IsAdvice))
	}
	if m.ParentId != 0 {
		n += 1 + sovUserSetting(uint64(m.ParentId))
	}
	l = len(m.ParentName)
	if l > 0 {
		n += 1 + l + sovUserSetting(uint64(l))
	}
	if m.ActId != 0 {
		n += 1 + sovUserSetting(uint64(m.ActId))
	}
	return n
}

func (m *UserSettingGetTagResp_OfflineTags) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovUserSetting(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovUserSetting(uint64(l))
	}
	return n
}

func (m *UserSettingGetTagResp_Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for _, e := range m.Tags {
			l = e.Size()
			n += 1 + l + sovUserSetting(uint64(l))
		}
	}
	if len(m.Offline) > 0 {
		for _, e := range m.Offline {
			l = e.Size()
			n += 1 + l + sovUserSetting(uint64(l))
		}
	}
	if m.Uid != 0 {
		n += 1 + sovUserSetting(uint64(m.Uid))
	}
	if m.IsGray != 0 {
		n += 1 + sovUserSetting(uint64(m.IsGray))
	}
	return n
}

func sovUserSetting(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUserSetting(x uint64) (n int) {
	return sovUserSetting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserSettingGetTagReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserSetting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserSettingGetTagReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserSettingGetTagReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipUserSetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUserSetting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserSettingGetTagResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserSetting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UserSettingGetTagResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserSettingGetTagResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUserSetting
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUserSetting
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &UserSettingGetTagResp_Data{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserSetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUserSetting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserSettingGetTagResp_Tags) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserSetting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Tags: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tags: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUserSetting
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUserSetting
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsAdvice", wireType)
			}
			m.IsAdvice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsAdvice |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentId", wireType)
			}
			m.ParentId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParentId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUserSetting
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActId", wireType)
			}
			m.ActId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUserSetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUserSetting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserSettingGetTagResp_OfflineTags) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserSetting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OfflineTags: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OfflineTags: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUserSetting
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserSetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUserSetting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UserSettingGetTagResp_Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserSetting
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUserSetting
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, &UserSettingGetTagResp_Tags{})
			if err := m.Tags[len(m.Tags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUserSetting
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Offline = append(m.Offline, &UserSettingGetTagResp_OfflineTags{})
			if err := m.Offline[len(m.Offline)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsGray", wireType)
			}
			m.IsGray = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsGray |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUserSetting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUserSetting
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUserSetting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserSetting
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUserSetting
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthUserSetting
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUserSetting
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUserSetting(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUserSetting = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserSetting   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v1/UserSetting.proto", fileDescriptor_UserSetting_6ecfe772383476d5) }

var fileDescriptor_UserSetting_6ecfe772383476d5 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xde, 0x64, 0x7f, 0x76, 0xa2, 0x08, 0x43, 0x2d, 0x31, 0x94, 0x64, 0x5d, 0x3d, 0x2c, 0x82,
	0x59, 0xb7, 0xfe, 0x05, 0xae, 0x62, 0xe9, 0x45, 0x61, 0xac, 0x1e, 0xbc, 0x84, 0xd9, 0xcc, 0xec,
	0x38, 0xd0, 0xdd, 0xa4, 0x99, 0x49, 0xa0, 0x77, 0xff, 0x00, 0xff, 0x2c, 0x4f, 0xd2, 0xa3, 0xa7,
	0x20, 0xbb, 0x07, 0x21, 0x7f, 0x85, 0xcc, 0x4b, 0x6a, 0x73, 0x28, 0xb8, 0x78, 0xf9, 0x78, 0xef,
	0xcb, 0xbc, 0xef, 0x7d, 0x33, 0x1f, 0x41, 0x87, 0xc5, 0x7c, 0xf6, 0x51, 0xf1, 0xec, 0x03, 0xd7,
	0x5a, 0x6e, 0x44, 0x98, 0x66, 0x89, 0x4e, 0xf0, 0xbd, 0x0b, 0x59, 0xf0, 0x28, 0x57, 0x3c, 0x0b,
	0x8b, 0xb9, 0xf7, 0x5c, 0x48, 0xfd, 0x25, 0x5f, 0x86, 0x71, 0xb2, 0x9e, 0x89, 0x44, 0x24, 0x33,
	0x38, 0xb4, 0xcc, 0x57, 0xd0, 0x41, 0x03, 0x55, 0x3d, 0x3c, 0x39, 0x42, 0x87, 0x2d, 0xc5, 0x53,
	0xae, 0xcf, 0xa9, 0x20, 0xfc, 0x72, 0xf2, 0xa3, 0x8f, 0x1e, 0xde, 0xf1, 0x41, 0xa5, 0xf8, 0x18,
	0xf5, 0xe2, 0x84, 0x71, 0xd7, 0x1a, 0x5b, 0xd3, 0xee, 0x62, 0x54, 0x95, 0x01, 0xf4, 0x04, 0x10,
	0x3f, 0x42, 0xdd, 0xb5, 0x12, 0xae, 0x3d, 0xb6, 0xa6, 0x07, 0x8b, 0x61, 0x55, 0x06, 0xa6, 0x25,
	0x06, 0xf0, 0x5b, 0xd4, 0x63, 0x54, 0x53, 0xb7, 0x3b, 0xb6, 0xa6, 0xce, 0xc9, 0x34, 0x6c, 0xdb,
	0x0e, 0xef, 0xdc, 0x15, 0xbe, 0xa1, 0x9a, 0xd6, 0x2b, 0xcc, 0x24, 0x01, 0xf4, 0xbe, 0xda, 0xa8,
	0x77, 0x4e, 0x85, 0xc2, 0x47, 0xc8, 0x96, 0xac, 0xf1, 0x31, 0xa8, 0xca, 0xc0, 0x96, 0x8c, 0xd8,
	0x92, 0x19, 0x87, 0x1b, 0xba, 0xe6, 0x8d, 0x09, 0x18, 0x37, 0x3d, 0x01, 0x34, 0x0e, 0x53, 0x19,
	0x83, 0x8b, 0xc6, 0x61, 0x2a, 0x63, 0x62, 0x00, 0x3f, 0x43, 0x07, 0x52, 0x45, 0x94, 0x15, 0x32,
	0xe6, 0x6e, 0x0f, 0x74, 0xef, 0x57, 0x65, 0x70, 0x4b, 0x92, 0x91, 0x54, 0xaf, 0xa0, 0x32, 0x67,
	0x53, 0x9a, 0xf1, 0x8d, 0x8e, 0x24, 0x73, 0xfb, 0xb7, 0x67, 0xff, 0x92, 0x64, 0x54, 0x97, 0x67,
	0x0c, 0xbf, 0x40, 0x4e, 0x43, 0x83, 0xaf, 0x01, 0xac, 0x7e, 0x50, 0x95, 0x41, 0x9b, 0x26, 0xa8,
	0x6e, 0xde, 0x19, 0x93, 0x8f, 0xd1, 0x80, 0xc6, 0x20, 0x3d, 0x04, 0x69, 0x54, 0x95, 0x41, 0xc3,
	0x90, 0x3e, 0x8d, 0xf5, 0x19, 0xf3, 0x5e, 0x23, 0xe7, 0xfd, 0x6a, 0x75, 0x21, 0x37, 0xfc, 0xff,
	0x1f, 0xc3, 0xfb, 0x6d, 0xa1, 0x9e, 0x79, 0x64, 0x13, 0x8e, 0xa6, 0x42, 0xb9, 0xd6, 0xb8, 0xbb,
	0x6f, 0x38, 0x66, 0x6d, 0x2d, 0x68, 0x26, 0x09, 0x20, 0xfe, 0x84, 0x86, 0x49, 0xed, 0xca, 0xb5,
	0x41, 0x6a, 0xb6, 0x8f, 0x54, 0xeb, 0x22, 0x0b, 0xa7, 0x2a, 0x83, 0x1b, 0x0d, 0x72, 0x53, 0x98,
	0xd4, 0x72, 0xc9, 0x20, 0xb5, 0x6e, 0x9d, 0x5a, 0x2e, 0x19, 0x31, 0x80, 0x9f, 0xa2, 0xa1, 0x54,
	0x91, 0xc8, 0xe8, 0x55, 0x93, 0x19, 0x28, 0x34, 0x14, 0x19, 0x48, 0x75, 0x9a, 0xd1, 0xab, 0x13,
	0x8a, 0x9c, 0xd6, 0x6e, 0x4c, 0xd0, 0x50, 0x70, 0x1d, 0x69, 0x2a, 0xf0, 0xe4, 0x9f, 0x0e, 0x2f,
	0xbd, 0x27, 0x7b, 0xdc, 0x62, 0x71, 0xfc, 0x7d, 0xeb, 0x5b, 0xd7, 0x5b, 0xdf, 0xfa, 0xb5, 0xf5,
	0xad, 0x6f, 0x3b, 0xbf, 0x73, 0xbd, 0xf3, 0x3b, 0x3f, 0x77, 0x7e, 0xe7, 0xb3, 0x5d, 0xcc, 0x97,
	0x03, 0xf8, 0xe1, 0x5e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xa7, 0xd5, 0xa9, 0xc5, 0x03,
	0x00, 0x00,
}
