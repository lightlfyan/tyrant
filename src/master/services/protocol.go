package services

import (
	"master/function"
	"master/protocol/unpack"
	"master/settings"

	"network"
)

type protobuf struct {
	p map[uint32]func(*network.Client, *network.Context)
}

func (p *protobuf) Start() {
	p.p[settings.System_Heart_C2S] = unpack.System_Heart_C2S
	p.p[settings.System_Ping] = unpack.System_Ping
	p.p[settings.System_Pong] = unpack.System_Pong
	p.p[settings.Cell_Login_C2S] = unpack.Cell_Login_C2S

	p.p[settings.Cell_Fake_C2S] = unpack.Cell_Fake_C2S

	p.p[settings.Task_Create_C2S] = unpack.Task_Create_C2S
	p.p[settings.Task_Progress_C2S] = unpack.Task_Progress_C2S
	p.p[settings.Task_Report_C2S] = unpack.Task_Report_C2S
	p.p[settings.Cell_Stop_C2S] = unpack.Cell_Stop_C2S
}

func (p *protobuf) Get(method uint32) func(*network.Client, *network.Context) {
	return p.p[method]
}

func init() {
	function.ProtocolManage = &protobuf{
		p: make(map[uint32]func(*network.Client, *network.Context)),
	}
}
