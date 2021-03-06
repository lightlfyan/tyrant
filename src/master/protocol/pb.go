// Code generated by protoc-gen-go.
// source: pb.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	Msg
	System_Msg_S2C
	System_Heart_C2S
	System_Heart_S2C
	System_Message_S2C
	System_Ping
	System_Pong
	Cell_Login_C2S
	Cell_Login_S2C
	Cell_Logout_C2S
	Cell_Fake_S2C
	Cell_Fake_C2S
	Cell_Stop_S2C
	Cell_Stop_C2S
	KV
	Task_Create_S2C
	Task_Create_C2S
	Task_Progress_S2C
	Task_Progress_C2S
	Task_Report_C2S
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Msg struct {
	Cmd              *string `protobuf:"bytes,1,req,name=cmd" json:"cmd,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Msg) GetCmd() string {
	if m != nil && m.Cmd != nil {
		return *m.Cmd
	}
	return ""
}

// --------SYSTEM------------------------------------
// 10000
type System_Msg_S2C struct {
	Code             *int32 `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *System_Msg_S2C) Reset()                    { *m = System_Msg_S2C{} }
func (m *System_Msg_S2C) String() string            { return proto.CompactTextString(m) }
func (*System_Msg_S2C) ProtoMessage()               {}
func (*System_Msg_S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *System_Msg_S2C) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

// 10001
type System_Heart_C2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *System_Heart_C2S) Reset()                    { *m = System_Heart_C2S{} }
func (m *System_Heart_C2S) String() string            { return proto.CompactTextString(m) }
func (*System_Heart_C2S) ProtoMessage()               {}
func (*System_Heart_C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 10002
type System_Heart_S2C struct {
	ServerTime       *int64 `protobuf:"varint,1,req,name=server_time,json=serverTime" json:"server_time,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *System_Heart_S2C) Reset()                    { *m = System_Heart_S2C{} }
func (m *System_Heart_S2C) String() string            { return proto.CompactTextString(m) }
func (*System_Heart_S2C) ProtoMessage()               {}
func (*System_Heart_S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *System_Heart_S2C) GetServerTime() int64 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

// 10003
type System_Message_S2C struct {
	Msg              *Msg   `protobuf:"bytes,1,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *System_Message_S2C) Reset()                    { *m = System_Message_S2C{} }
func (m *System_Message_S2C) String() string            { return proto.CompactTextString(m) }
func (*System_Message_S2C) ProtoMessage()               {}
func (*System_Message_S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *System_Message_S2C) GetMsg() *Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

// 10004
type System_Ping struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *System_Ping) Reset()                    { *m = System_Ping{} }
func (m *System_Ping) String() string            { return proto.CompactTextString(m) }
func (*System_Ping) ProtoMessage()               {}
func (*System_Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// 10005
type System_Pong struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *System_Pong) Reset()                    { *m = System_Pong{} }
func (m *System_Pong) String() string            { return proto.CompactTextString(m) }
func (*System_Pong) ProtoMessage()               {}
func (*System_Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// --------Cell--------------------------------------
// 10100
type Cell_Login_C2S struct {
	Id               *int32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Cell_Login_C2S) Reset()                    { *m = Cell_Login_C2S{} }
func (m *Cell_Login_C2S) String() string            { return proto.CompactTextString(m) }
func (*Cell_Login_C2S) ProtoMessage()               {}
func (*Cell_Login_C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Cell_Login_C2S) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

// 10101
type Cell_Login_S2C struct {
	Id               *int32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Result           *bool  `protobuf:"varint,2,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Cell_Login_S2C) Reset()                    { *m = Cell_Login_S2C{} }
func (m *Cell_Login_S2C) String() string            { return proto.CompactTextString(m) }
func (*Cell_Login_S2C) ProtoMessage()               {}
func (*Cell_Login_S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Cell_Login_S2C) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Cell_Login_S2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

// 1010
type Cell_Logout_C2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Cell_Logout_C2S) Reset()                    { *m = Cell_Logout_C2S{} }
func (m *Cell_Logout_C2S) String() string            { return proto.CompactTextString(m) }
func (*Cell_Logout_C2S) ProtoMessage()               {}
func (*Cell_Logout_C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// 10103
type Cell_Fake_S2C struct {
	Filename         *string `protobuf:"bytes,1,req,name=filename" json:"filename,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Cell_Fake_S2C) Reset()                    { *m = Cell_Fake_S2C{} }
func (m *Cell_Fake_S2C) String() string            { return proto.CompactTextString(m) }
func (*Cell_Fake_S2C) ProtoMessage()               {}
func (*Cell_Fake_S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Cell_Fake_S2C) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

// 10104
type Cell_Fake_C2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Cell_Fake_C2S) Reset()                    { *m = Cell_Fake_C2S{} }
func (m *Cell_Fake_C2S) String() string            { return proto.CompactTextString(m) }
func (*Cell_Fake_C2S) ProtoMessage()               {}
func (*Cell_Fake_C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type Cell_Stop_S2C struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Cell_Stop_S2C) Reset()                    { *m = Cell_Stop_S2C{} }
func (m *Cell_Stop_S2C) String() string            { return proto.CompactTextString(m) }
func (*Cell_Stop_S2C) ProtoMessage()               {}
func (*Cell_Stop_S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type Cell_Stop_C2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Cell_Stop_C2S) Reset()                    { *m = Cell_Stop_C2S{} }
func (m *Cell_Stop_C2S) String() string            { return proto.CompactTextString(m) }
func (*Cell_Stop_C2S) ProtoMessage()               {}
func (*Cell_Stop_C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type KV struct {
	K                *string `protobuf:"bytes,1,req,name=k" json:"k,omitempty"`
	V                *string `protobuf:"bytes,2,req,name=v" json:"v,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *KV) Reset()                    { *m = KV{} }
func (m *KV) String() string            { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()               {}
func (*KV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *KV) GetK() string {
	if m != nil && m.K != nil {
		return *m.K
	}
	return ""
}

func (m *KV) GetV() string {
	if m != nil && m.V != nil {
		return *m.V
	}
	return ""
}

type Task_Create_S2C struct {
	TaskId           *int64   `protobuf:"varint,1,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	Url              *string  `protobuf:"bytes,2,req,name=url" json:"url,omitempty"`
	Num              *int32   `protobuf:"varint,3,req,name=num" json:"num,omitempty"`
	Conc             *int32   `protobuf:"varint,4,req,name=conc" json:"conc,omitempty"`
	Qps              *int32   `protobuf:"varint,5,req,name=qps" json:"qps,omitempty"`
	Timeout          *int32   `protobuf:"varint,6,req,name=timeout" json:"timeout,omitempty"`
	UserName         *string  `protobuf:"bytes,7,req,name=UserName,json=userName" json:"UserName,omitempty"`
	Password         *string  `protobuf:"bytes,8,req,name=Password,json=password" json:"Password,omitempty"`
	Body             *string  `protobuf:"bytes,9,req,name=Body,json=body" json:"Body,omitempty"`
	Accept           *string  `protobuf:"bytes,10,req,name=Accept,json=accept" json:"Accept,omitempty"`
	ContentType      *string  `protobuf:"bytes,11,req,name=contentType" json:"contentType,omitempty"`
	Method           *string  `protobuf:"bytes,12,req,name=Method,json=method" json:"Method,omitempty"`
	ProxyAddr        *string  `protobuf:"bytes,13,req,name=ProxyAddr,json=proxyAddr" json:"ProxyAddr,omitempty"`
	KeepAlive        *bool    `protobuf:"varint,14,req,name=KeepAlive,json=keepAlive" json:"KeepAlive,omitempty"`
	Headers          []*KV    `protobuf:"bytes,15,rep,name=Headers,json=headers" json:"Headers,omitempty"`
	Host             *string  `protobuf:"bytes,16,req,name=Host,json=host" json:"Host,omitempty"`
	Urls             []string `protobuf:"bytes,17,rep,name=urls" json:"urls,omitempty"`
	Uid              *string  `protobuf:"bytes,18,req,name=uid" json:"uid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Task_Create_S2C) Reset()                    { *m = Task_Create_S2C{} }
func (m *Task_Create_S2C) String() string            { return proto.CompactTextString(m) }
func (*Task_Create_S2C) ProtoMessage()               {}
func (*Task_Create_S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Task_Create_S2C) GetTaskId() int64 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

func (m *Task_Create_S2C) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Task_Create_S2C) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *Task_Create_S2C) GetConc() int32 {
	if m != nil && m.Conc != nil {
		return *m.Conc
	}
	return 0
}

func (m *Task_Create_S2C) GetQps() int32 {
	if m != nil && m.Qps != nil {
		return *m.Qps
	}
	return 0
}

func (m *Task_Create_S2C) GetTimeout() int32 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *Task_Create_S2C) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *Task_Create_S2C) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *Task_Create_S2C) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *Task_Create_S2C) GetAccept() string {
	if m != nil && m.Accept != nil {
		return *m.Accept
	}
	return ""
}

func (m *Task_Create_S2C) GetContentType() string {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return ""
}

func (m *Task_Create_S2C) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *Task_Create_S2C) GetProxyAddr() string {
	if m != nil && m.ProxyAddr != nil {
		return *m.ProxyAddr
	}
	return ""
}

func (m *Task_Create_S2C) GetKeepAlive() bool {
	if m != nil && m.KeepAlive != nil {
		return *m.KeepAlive
	}
	return false
}

func (m *Task_Create_S2C) GetHeaders() []*KV {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Task_Create_S2C) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *Task_Create_S2C) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *Task_Create_S2C) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

type Task_Create_C2S struct {
	TaskId           *int64 `protobuf:"varint,1,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Task_Create_C2S) Reset()                    { *m = Task_Create_C2S{} }
func (m *Task_Create_C2S) String() string            { return proto.CompactTextString(m) }
func (*Task_Create_C2S) ProtoMessage()               {}
func (*Task_Create_C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Task_Create_C2S) GetTaskId() int64 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

type Task_Progress_S2C struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Task_Progress_S2C) Reset()                    { *m = Task_Progress_S2C{} }
func (m *Task_Progress_S2C) String() string            { return proto.CompactTextString(m) }
func (*Task_Progress_S2C) ProtoMessage()               {}
func (*Task_Progress_S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type Task_Progress_C2S struct {
	TaskId           *int64 `protobuf:"varint,1,req,name=task_id,json=taskId" json:"task_id,omitempty"`
	Progress         *int32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Task_Progress_C2S) Reset()                    { *m = Task_Progress_C2S{} }
func (m *Task_Progress_C2S) String() string            { return proto.CompactTextString(m) }
func (*Task_Progress_C2S) ProtoMessage()               {}
func (*Task_Progress_C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *Task_Progress_C2S) GetTaskId() int64 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

func (m *Task_Progress_C2S) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type Task_Report_C2S struct {
	Totalsecond      *float64 `protobuf:"fixed64,1,req,name=totalsecond" json:"totalsecond,omitempty"`
	Fastest          *float64 `protobuf:"fixed64,2,req,name=fastest" json:"fastest,omitempty"`
	Slowest          *float64 `protobuf:"fixed64,3,req,name=slowest" json:"slowest,omitempty"`
	Average          *float64 `protobuf:"fixed64,4,req,name=average" json:"average,omitempty"`
	Rps              *float64 `protobuf:"fixed64,5,req,name=rps" json:"rps,omitempty"`
	Sizetotal        *int64   `protobuf:"varint,6,req,name=sizetotal" json:"sizetotal,omitempty"`
	Sizereq          *int64   `protobuf:"varint,7,req,name=sizereq" json:"sizereq,omitempty"`
	Taskid           *int64   `protobuf:"varint,8,req,name=taskid" json:"taskid,omitempty"`
	Code2            *int64   `protobuf:"varint,9,req,name=code2" json:"code2,omitempty"`
	Code3            *int64   `protobuf:"varint,10,req,name=code3" json:"code3,omitempty"`
	Code4            *int64   `protobuf:"varint,11,req,name=code4" json:"code4,omitempty"`
	Code5            *int64   `protobuf:"varint,12,req,name=code5" json:"code5,omitempty"`
	CodeOther        *int64   `protobuf:"varint,13,req,name=codeOther" json:"codeOther,omitempty"`
	P10              *float64 `protobuf:"fixed64,14,req,name=p10" json:"p10,omitempty"`
	P25              *float64 `protobuf:"fixed64,15,req,name=p25" json:"p25,omitempty"`
	P50              *float64 `protobuf:"fixed64,16,req,name=p50" json:"p50,omitempty"`
	P75              *float64 `protobuf:"fixed64,17,req,name=p75" json:"p75,omitempty"`
	P90              *float64 `protobuf:"fixed64,18,req,name=p90" json:"p90,omitempty"`
	P95              *float64 `protobuf:"fixed64,19,req,name=p95" json:"p95,omitempty"`
	P99              *float64 `protobuf:"fixed64,20,req,name=p99" json:"p99,omitempty"`
	Conc             *int64   `protobuf:"varint,21,req,name=conc" json:"conc,omitempty"`
	Error            *int64   `protobuf:"varint,22,req,name=error" json:"error,omitempty"`
	Cellid           *int32   `protobuf:"varint,23,req,name=cellid" json:"cellid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Task_Report_C2S) Reset()                    { *m = Task_Report_C2S{} }
func (m *Task_Report_C2S) String() string            { return proto.CompactTextString(m) }
func (*Task_Report_C2S) ProtoMessage()               {}
func (*Task_Report_C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *Task_Report_C2S) GetTotalsecond() float64 {
	if m != nil && m.Totalsecond != nil {
		return *m.Totalsecond
	}
	return 0
}

func (m *Task_Report_C2S) GetFastest() float64 {
	if m != nil && m.Fastest != nil {
		return *m.Fastest
	}
	return 0
}

func (m *Task_Report_C2S) GetSlowest() float64 {
	if m != nil && m.Slowest != nil {
		return *m.Slowest
	}
	return 0
}

func (m *Task_Report_C2S) GetAverage() float64 {
	if m != nil && m.Average != nil {
		return *m.Average
	}
	return 0
}

func (m *Task_Report_C2S) GetRps() float64 {
	if m != nil && m.Rps != nil {
		return *m.Rps
	}
	return 0
}

func (m *Task_Report_C2S) GetSizetotal() int64 {
	if m != nil && m.Sizetotal != nil {
		return *m.Sizetotal
	}
	return 0
}

func (m *Task_Report_C2S) GetSizereq() int64 {
	if m != nil && m.Sizereq != nil {
		return *m.Sizereq
	}
	return 0
}

func (m *Task_Report_C2S) GetTaskid() int64 {
	if m != nil && m.Taskid != nil {
		return *m.Taskid
	}
	return 0
}

func (m *Task_Report_C2S) GetCode2() int64 {
	if m != nil && m.Code2 != nil {
		return *m.Code2
	}
	return 0
}

func (m *Task_Report_C2S) GetCode3() int64 {
	if m != nil && m.Code3 != nil {
		return *m.Code3
	}
	return 0
}

func (m *Task_Report_C2S) GetCode4() int64 {
	if m != nil && m.Code4 != nil {
		return *m.Code4
	}
	return 0
}

func (m *Task_Report_C2S) GetCode5() int64 {
	if m != nil && m.Code5 != nil {
		return *m.Code5
	}
	return 0
}

func (m *Task_Report_C2S) GetCodeOther() int64 {
	if m != nil && m.CodeOther != nil {
		return *m.CodeOther
	}
	return 0
}

func (m *Task_Report_C2S) GetP10() float64 {
	if m != nil && m.P10 != nil {
		return *m.P10
	}
	return 0
}

func (m *Task_Report_C2S) GetP25() float64 {
	if m != nil && m.P25 != nil {
		return *m.P25
	}
	return 0
}

func (m *Task_Report_C2S) GetP50() float64 {
	if m != nil && m.P50 != nil {
		return *m.P50
	}
	return 0
}

func (m *Task_Report_C2S) GetP75() float64 {
	if m != nil && m.P75 != nil {
		return *m.P75
	}
	return 0
}

func (m *Task_Report_C2S) GetP90() float64 {
	if m != nil && m.P90 != nil {
		return *m.P90
	}
	return 0
}

func (m *Task_Report_C2S) GetP95() float64 {
	if m != nil && m.P95 != nil {
		return *m.P95
	}
	return 0
}

func (m *Task_Report_C2S) GetP99() float64 {
	if m != nil && m.P99 != nil {
		return *m.P99
	}
	return 0
}

func (m *Task_Report_C2S) GetConc() int64 {
	if m != nil && m.Conc != nil {
		return *m.Conc
	}
	return 0
}

func (m *Task_Report_C2S) GetError() int64 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return 0
}

func (m *Task_Report_C2S) GetCellid() int32 {
	if m != nil && m.Cellid != nil {
		return *m.Cellid
	}
	return 0
}

func init() {
	proto.RegisterType((*Msg)(nil), "protocol.Msg")
	proto.RegisterType((*System_Msg_S2C)(nil), "protocol.System_Msg_S2C")
	proto.RegisterType((*System_Heart_C2S)(nil), "protocol.System_Heart_C2S")
	proto.RegisterType((*System_Heart_S2C)(nil), "protocol.System_Heart_S2C")
	proto.RegisterType((*System_Message_S2C)(nil), "protocol.System_Message_S2C")
	proto.RegisterType((*System_Ping)(nil), "protocol.System_Ping")
	proto.RegisterType((*System_Pong)(nil), "protocol.System_Pong")
	proto.RegisterType((*Cell_Login_C2S)(nil), "protocol.Cell_Login_C2S")
	proto.RegisterType((*Cell_Login_S2C)(nil), "protocol.Cell_Login_S2C")
	proto.RegisterType((*Cell_Logout_C2S)(nil), "protocol.Cell_Logout_C2S")
	proto.RegisterType((*Cell_Fake_S2C)(nil), "protocol.Cell_Fake_S2C")
	proto.RegisterType((*Cell_Fake_C2S)(nil), "protocol.Cell_Fake_C2S")
	proto.RegisterType((*Cell_Stop_S2C)(nil), "protocol.Cell_Stop_S2C")
	proto.RegisterType((*Cell_Stop_C2S)(nil), "protocol.Cell_Stop_C2S")
	proto.RegisterType((*KV)(nil), "protocol.KV")
	proto.RegisterType((*Task_Create_S2C)(nil), "protocol.Task_Create_S2C")
	proto.RegisterType((*Task_Create_C2S)(nil), "protocol.Task_Create_C2S")
	proto.RegisterType((*Task_Progress_S2C)(nil), "protocol.Task_Progress_S2C")
	proto.RegisterType((*Task_Progress_C2S)(nil), "protocol.Task_Progress_C2S")
	proto.RegisterType((*Task_Report_C2S)(nil), "protocol.Task_Report_C2S")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x94, 0xe1, 0x4e, 0x23, 0x37,
	0x10, 0xc7, 0xb5, 0x59, 0x42, 0x12, 0x07, 0x08, 0xec, 0xd1, 0xc3, 0xaa, 0x2a, 0x35, 0x5a, 0x55,
	0x15, 0x6a, 0x25, 0x44, 0xc3, 0xad, 0xae, 0xf9, 0x48, 0x91, 0x2a, 0x2a, 0x4a, 0x8b, 0x0c, 0xbd,
	0xaf, 0xd1, 0xde, 0x7a, 0x6e, 0x59, 0x65, 0xb3, 0x5e, 0x6c, 0x27, 0xd7, 0xf0, 0x26, 0x7d, 0xa5,
	0x3e, 0x55, 0x35, 0x63, 0x1b, 0x08, 0x52, 0xef, 0x53, 0xfc, 0xff, 0xcd, 0xcc, 0xda, 0xe3, 0xf9,
	0x3b, 0xac, 0xdf, 0x7e, 0x3c, 0x69, 0xb5, 0xb2, 0x2a, 0xe9, 0xd3, 0x4f, 0xa1, 0xea, 0xf4, 0x88,
	0xc5, 0xd7, 0xa6, 0x4c, 0xf6, 0x59, 0x5c, 0x2c, 0x24, 0x8f, 0xc6, 0x9d, 0xe3, 0x81, 0xc0, 0x65,
	0xfa, 0x1d, 0xdb, 0xbb, 0x5d, 0x1b, 0x0b, 0x8b, 0xd9, 0xb5, 0x29, 0x67, 0xb7, 0x93, 0x8b, 0x24,
	0x61, 0x5b, 0x85, 0x92, 0x40, 0x49, 0x5d, 0x41, 0xeb, 0x34, 0x61, 0xfb, 0x3e, 0xeb, 0x12, 0x72,
	0x6d, 0x67, 0x17, 0x93, 0xdb, 0xf4, 0xec, 0x15, 0xc3, 0xda, 0x6f, 0xd9, 0xd0, 0x80, 0x5e, 0x81,
	0x9e, 0xd9, 0x6a, 0xe1, 0x3e, 0x11, 0x0b, 0xe6, 0xd0, 0x5d, 0xb5, 0x80, 0x34, 0x63, 0x49, 0xd8,
	0x0e, 0x8c, 0xc9, 0x4b, 0xf0, 0x65, 0xf1, 0xc2, 0x94, 0x94, 0x3e, 0x9c, 0xec, 0x9e, 0x84, 0x53,
	0x9f, 0x5c, 0x9b, 0x52, 0x60, 0x24, 0xdd, 0x65, 0x43, 0x5f, 0x76, 0x53, 0x35, 0x1b, 0x52, 0x35,
	0x65, 0x3a, 0x66, 0x7b, 0x17, 0x50, 0xd7, 0xb3, 0xdf, 0x55, 0x59, 0x35, 0x78, 0xb6, 0x64, 0x8f,
	0x75, 0x2a, 0xe9, 0x3b, 0xe8, 0x54, 0x32, 0xfd, 0x79, 0x23, 0x03, 0xb7, 0x7c, 0x95, 0x91, 0xbc,
	0x65, 0xdb, 0x1a, 0xcc, 0xb2, 0xb6, 0xbc, 0x33, 0xee, 0x1c, 0xf7, 0x85, 0x57, 0xe9, 0x01, 0x1b,
	0x85, 0x4a, 0xb5, 0x74, 0x8d, 0xff, 0xc8, 0x76, 0x09, 0xfd, 0x9a, 0xcf, 0xdd, 0xf1, 0xbf, 0x66,
	0xfd, 0x4f, 0x55, 0x0d, 0x4d, 0xee, 0x5b, 0x1e, 0x88, 0x27, 0x9d, 0x8e, 0x5e, 0x26, 0x63, 0x75,
	0x00, 0xb7, 0x56, 0xb5, 0x58, 0xbd, 0x09, 0x30, 0x63, 0xcc, 0x3a, 0x57, 0x1f, 0x92, 0x1d, 0x16,
	0xcd, 0xfd, 0xd7, 0xa2, 0x39, 0xaa, 0x15, 0x9d, 0x6c, 0x20, 0xa2, 0x55, 0xfa, 0x6f, 0xcc, 0x46,
	0x77, 0xb9, 0x99, 0xcf, 0x2e, 0x34, 0xe4, 0xd6, 0x1d, 0xe2, 0x88, 0xf5, 0x2c, 0x22, 0xdf, 0x55,
	0x2c, 0xb6, 0x51, 0xfe, 0x26, 0x71, 0xe6, 0x4b, 0x5d, 0xfb, 0x62, 0x5c, 0x22, 0x69, 0x96, 0x0b,
	0x1e, 0x53, 0xf3, 0xb8, 0x74, 0x33, 0x6f, 0x0a, 0xbe, 0x15, 0x66, 0xde, 0x14, 0x98, 0xf5, 0xd0,
	0x1a, 0xde, 0x75, 0x59, 0x0f, 0xad, 0x49, 0x38, 0xeb, 0xe1, 0x58, 0xd5, 0xd2, 0xf2, 0x6d, 0xa2,
	0x41, 0xe2, 0x0d, 0xfc, 0x65, 0x40, 0xff, 0x81, 0x37, 0xd0, 0x73, 0x37, 0xb0, 0xf4, 0x1a, 0x63,
	0x37, 0xb9, 0x31, 0x9f, 0x95, 0x96, 0xbc, 0xef, 0x62, 0xad, 0xd7, 0xb8, 0xef, 0x2f, 0x4a, 0xae,
	0xf9, 0x80, 0xf8, 0xd6, 0x47, 0x25, 0xd7, 0x38, 0x89, 0xf3, 0xa2, 0x80, 0xd6, 0x72, 0x46, 0x74,
	0x3b, 0x27, 0x95, 0x8c, 0xd9, 0xb0, 0x50, 0x8d, 0x85, 0xc6, 0xde, 0xad, 0x5b, 0xe0, 0x43, 0x0a,
	0xbe, 0x44, 0x58, 0x79, 0x0d, 0xf6, 0x5e, 0x49, 0xbe, 0xe3, 0x2a, 0x17, 0xa4, 0x92, 0x6f, 0xd8,
	0xe0, 0x46, 0xab, 0xbf, 0xd7, 0xe7, 0x52, 0x6a, 0xbe, 0x4b, 0xa1, 0x41, 0x1b, 0x00, 0x46, 0xaf,
	0x00, 0xda, 0xf3, 0xba, 0x5a, 0x01, 0xdf, 0xa3, 0xe1, 0x0f, 0xe6, 0x01, 0x24, 0xdf, 0xb3, 0xde,
	0x25, 0xe4, 0x12, 0xb4, 0xe1, 0xa3, 0x71, 0x7c, 0x3c, 0x9c, 0xec, 0x3c, 0xdb, 0xf3, 0xea, 0x83,
	0xe8, 0xdd, 0xbb, 0x20, 0x76, 0x72, 0xa9, 0x8c, 0xe5, 0xfb, 0xae, 0x93, 0x7b, 0x65, 0x2c, 0xb2,
	0xa5, 0xae, 0x0d, 0x3f, 0x18, 0xc7, 0xc8, 0x70, 0x4d, 0xd3, 0xa8, 0x24, 0x4f, 0xfc, 0x34, 0x2a,
	0x99, 0xfe, 0xb0, 0x39, 0x4b, 0xb4, 0xef, 0xff, 0xcd, 0x32, 0x7d, 0xc3, 0x0e, 0x28, 0xf7, 0x46,
	0xab, 0x52, 0x83, 0x31, 0x64, 0xa0, 0xcb, 0xd7, 0xf0, 0x4b, 0x9f, 0xc0, 0x71, 0xb4, 0x3e, 0x91,
	0x3c, 0xd1, 0x15, 0x4f, 0x3a, 0xfd, 0x67, 0xcb, 0x9f, 0x45, 0x40, 0xab, 0xdc, 0x33, 0xc7, 0x6b,
	0xb7, 0xca, 0xe6, 0xb5, 0x81, 0x42, 0x35, 0xee, 0x63, 0x91, 0x78, 0x89, 0xd0, 0x16, 0x9f, 0x72,
	0x63, 0xc1, 0xb8, 0xb7, 0x13, 0x89, 0x20, 0x31, 0x62, 0x6a, 0xf5, 0x19, 0x23, 0xb1, 0x8b, 0x78,
	0x89, 0x91, 0x7c, 0x05, 0x3a, 0x2f, 0x81, 0x3c, 0x17, 0x89, 0x20, 0xf1, 0x82, 0xb4, 0xb7, 0x5d,
	0x24, 0x70, 0x89, 0x03, 0x32, 0xd5, 0x23, 0xd0, 0x96, 0x64, 0xbc, 0x58, 0x3c, 0x03, 0xda, 0xa3,
	0x7a, 0x04, 0x0d, 0x0f, 0xe4, 0xbc, 0x58, 0x04, 0x89, 0x76, 0xc0, 0x9e, 0x2b, 0x67, 0x3b, 0x7f,
	0x03, 0x95, 0x4c, 0x0e, 0x59, 0x17, 0xff, 0xd4, 0x26, 0xe4, 0xba, 0x58, 0x38, 0x11, 0xe8, 0x19,
	0xb9, 0xce, 0xd3, 0xb3, 0x40, 0xdf, 0x91, 0xdd, 0x3c, 0x7d, 0x17, 0x68, 0x46, 0x3e, 0xf3, 0x34,
	0xc3, 0x73, 0xe2, 0xe2, 0x4f, 0x7b, 0x0f, 0xce, 0x66, 0xb1, 0x78, 0x06, 0xd8, 0x57, 0xfb, 0xd3,
	0x29, 0x19, 0x2c, 0x12, 0xb8, 0x24, 0x32, 0xc9, 0xf8, 0xc8, 0x93, 0x49, 0x46, 0x24, 0x3b, 0x25,
	0x0f, 0x21, 0xc9, 0x5c, 0xce, 0xfb, 0x8c, 0x1f, 0x78, 0xf2, 0xde, 0xe5, 0x4c, 0x4f, 0xc9, 0x40,
	0x48, 0xa6, 0x2e, 0x67, 0x9a, 0xf1, 0x37, 0x81, 0xf8, 0x9c, 0x29, 0x3f, 0x0c, 0x64, 0xfa, 0xf4,
	0xc0, 0xbf, 0xa2, 0x63, 0xb9, 0x07, 0x7e, 0xc8, 0xba, 0xa0, 0xb5, 0xd2, 0xfc, 0xad, 0xeb, 0x82,
	0x04, 0xde, 0x5a, 0x01, 0x75, 0x5d, 0x49, 0x7e, 0x44, 0xee, 0xf0, 0xea, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2a, 0x6d, 0xe0, 0x66, 0x56, 0x06, 0x00, 0x00,
}
