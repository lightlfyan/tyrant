package pack

import (
	"github.com/golang/protobuf/proto"

	"master/function"
	"master/protocol"
	"master/settings"

	"network"
)

func System_Heart_S2C(client *network.Client, time int64) {
	obj := &protocol.System_Heart_S2C{
		ServerTime: proto.Int64(time),
	}
	function.NetworkManage.Send(settings.System_Heart_S2C, client, obj)
}

func Cell_Fake_S2C(client *network.Client, filename string) {
	obj := &protocol.Cell_Fake_S2C{
		Filename: &filename,
	}
	function.NetworkManage.Send(settings.Cell_Fake_S2C, client, obj)
}

func Cell_Stop_S2C(client *network.Client) {
	obj := &protocol.Cell_Stop_S2C{}
	function.NetworkManage.Send(settings.Cell_Stop_S2C, client, obj)
}
