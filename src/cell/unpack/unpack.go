package unpack

import (
	"cell/api"
	"cell/function"
	"cell/logics"
	"cell/protocol"
	"log"

	"master/settings"

	"github.com/golang/protobuf/proto"
	"network"
)

func System_Ping(context *network.Context) {
	//log.Println("receive ping")
	sendpb := &protocol.System_Pong{}
	function.NetworkManage.Send(settings.System_Pong, sendpb)
}

func System_Pong(context *network.Context) {
	//log.Println("recive pong")
}

func System_Heart_S2C(context *network.Context) {
	pb := &protocol.System_Heart_S2C{}
	proto.Unmarshal(context.Msg.([]byte), pb)

	sendpb := &protocol.System_Heart_C2S{}
	function.NetworkManage.Send(settings.System_Heart_C2S, sendpb)
}

func Cell_Login_S2C(context *network.Context) {
	pb := &protocol.Cell_Login_S2C{}
	proto.Unmarshal(context.Msg.([]byte), pb)
	api.Cell_Login_S2C(pb)

}

func Task_Create_S2C(context *network.Context) {
	pb := &protocol.Task_Create_S2C{}
	proto.Unmarshal(context.Msg.([]byte), pb)
	api.Task_Create_C2S(pb)
}

func Task_Progress_S2C(context *network.Context) {
	if logics.CurrBoom == nil {
		return
	}

	pb := &protocol.Task_Progress_S2C{}
	proto.Unmarshal(context.Msg.([]byte), pb)

	var p int32 = 100
	if logics.CurrBoom != nil {
		p = logics.CurrBoom.Progress()
	}

	sendpb := &protocol.Task_Progress_C2S{
		TaskId:   &logics.TaskId,
		Progress: &p,
	}

	function.NetworkManage.Send(settings.Task_Progress_C2S, sendpb)
}

func Cell_Fake_S2C(context *network.Context) {
	pb := &protocol.Cell_Fake_S2C{}
	proto.Unmarshal(context.Msg.([]byte), pb)

	api.DownloadFakeData(*pb.Filename)
	api.Loadfakedata(*pb.Filename)

	sendpb := &protocol.Cell_Fake_C2S{}
	function.NetworkManage.Send(settings.Cell_Fake_C2S, sendpb)
}

func Cell_Stop_S2C(context *network.Context) {
	if logics.CurrBoom != nil {
		logics.CurrBoom.Stop()
	}

	log.Println("Cell_Stop_S2C: allstop")

	sendpb := &protocol.Cell_Stop_C2S{}
	function.NetworkManage.Send(settings.Cell_Stop_C2S, sendpb)
}
