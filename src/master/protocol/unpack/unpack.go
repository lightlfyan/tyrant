package unpack

import (
	"log"
	"master/api"
	"master/protocol"

	"master/data"
	"master/function"
	"master/settings"

	"sync/atomic"

	"github.com/golang/protobuf/proto"
	"network"
)

func System_Heart_C2S(client *network.Client, context *network.Context) {
	obj := &protocol.System_Heart_C2S{}
	proto.Unmarshal(context.Msg.([]byte), obj)
	api.System_Heart_C2S_Handler(client, context, obj)
}

func System_Ping(client *network.Client, context *network.Context) {
	sendpb := &protocol.System_Pong{}
	function.NetworkManage.Send(settings.System_Pong, client, sendpb)
}

func System_Pong(client *network.Client, context *network.Context) {
}

func Cell_Login_C2S(client *network.Client, context *network.Context) {
	obj := &protocol.Cell_Login_C2S{}
	proto.Unmarshal(context.Msg.([]byte), obj)
	api.Cell_Login_C2S(client, context, obj)
}

func Cell_Fake_C2S(client *network.Client, context *network.Context) {
	atomic.AddInt32(&data.CellFakeLoadCounter, -1)
	log.Println("Cell_Fake_C2S")
}

func Task_Create_C2S(client *network.Client, context *network.Context) {
	obj := &protocol.Task_Create_C2S{}
	proto.Unmarshal(context.Msg.([]byte), obj)
	log.Println("Task_Create_C2S: ", *obj.TaskId)
}

func Task_Progress_C2S(client *network.Client, context *network.Context) {
	obj := &protocol.Task_Progress_C2S{}
	proto.Unmarshal(context.Msg.([]byte), obj)
	data.TaskProgress.Set(client.Uid, *obj.Progress)
}

func Task_Report_C2S(client *network.Client, context *network.Context) {
	obj := &protocol.Task_Report_C2S{}
	proto.Unmarshal(context.Msg.([]byte), obj)
	log.Println("Task_Report_C2S cellid:", *obj.Cellid)
	data.CellMailBox <- obj
}

func Cell_Stop_C2S(client *network.Client, context *network.Context) {
	api.Cell_Stop_C2S()
	log.Println("Cell_Stop_C2S: ", client.Uid)
}
