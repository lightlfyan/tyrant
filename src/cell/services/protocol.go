package services

import (
	"cell/function"
	"cell/unpack"
	"master/settings"
)

type protobuf struct {
	proto2fun map[uint32]function.HandlefunType
}

func (p *protobuf) Start() {
	p.proto2fun[settings.System_Heart_S2C] = unpack.System_Heart_S2C
	p.proto2fun[settings.System_Ping] = unpack.System_Ping
	p.proto2fun[settings.System_Pong] = unpack.System_Pong
	p.proto2fun[settings.Cell_Login_S2C] = unpack.Cell_Login_S2C

	p.proto2fun[settings.Task_Create_S2C] = unpack.Task_Create_S2C
	p.proto2fun[settings.Task_Progress_S2C] = unpack.Task_Progress_S2C

	p.proto2fun[settings.Cell_Fake_S2C] = unpack.Cell_Fake_S2C
	p.proto2fun[settings.Cell_Stop_S2C] = unpack.Cell_Stop_S2C
}

func (p *protobuf) Get(method uint32) function.HandlefunType {
	return p.proto2fun[method]
}

func init() {
	function.ProtocolManage = &protobuf{
		proto2fun: make(map[uint32]function.HandlefunType),
	}
}
