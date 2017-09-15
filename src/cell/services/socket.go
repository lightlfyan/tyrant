package services

import (
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"network"

	"cell/api"
	"cell/function"
	"cell/logics"
)

type socket struct {
	client network.SocketClient
}

func HandlerFun(context network.Context) {
	function.ProtocolManage.Get(context.Method)(&context)
}

func WhenConnDone() {
	api.Login(logics.CellId)
}

func WhenConnExit() {
	log.Println("conn exit")
}

func (n *socket) Send(method uint32, pb proto.Message) {
	sendData, _ := proto.Marshal(pb)
	ctx := network.Context{
		Cmd:    "send",
		Method: method,
		Msg:    sendData,
	}
	n.client.Send(ctx)
}

func (n *socket) Start(addr *string) {
	n.client.Addr = *addr
	n.client.Start()
}

func init() {
	client := network.SocketClient{
		Net:                 "tcp4",
		Addr:                "127.0.0.1:9999",
		HandlerBoxQueueSize: 64,
		ReadDeadLine:        60 * time.Second,
		WriteDeadLine:       60 * time.Second,
		Handler:             HandlerFun,
		WhenConnDone:        WhenConnDone,
		WhenConnExit:        WhenConnExit,
	}
	function.NetworkManage = &socket{client: client}
}
