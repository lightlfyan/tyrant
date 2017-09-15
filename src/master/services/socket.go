package services

import (
	"log"
	"time"

	"master/data"
	"master/function"
	"master/protocol"
	"master/settings"

	"github.com/golang/protobuf/proto"
	"network"
)

type socket struct {
	server network.SocketServer
}

func AcceptFun(client *network.Client) {
	log.Println("accept", client.Conn.RemoteAddr().String())
}

func KickFun(client *network.Client) {
	log.Println(client.Uid, " disconnect")
	data.DelWork(client.Uid)
	log.Println("kick", client.Conn.RemoteAddr().String())
}

func HandlerFun(client *network.Client, context network.Context) {
	if client.Uid == -1 && context.Method != settings.Cell_Login_C2S {
		log.Println("not register")
		return
	}
	unpackFunc := function.ProtocolManage.Get(context.Method)
	unpackFunc(client, &context)
}

func (n *socket) Send(method uint32, client *network.Client, pb proto.Message) {
	sendData, _ := proto.Marshal(pb)
	client.SenderBox <- network.Context{
		Cmd:    "send",
		Method: method,
		Msg:    sendData,
	}
}

func (n *socket) Start() {
	n.server = network.SocketServer{
		Net:  "tcp4",
		Addr: ":9999",

		SenderBoxQueueSize:   2,
		ReceiverBoxQueueSize: 2,
		MainBoxQueueSize:     64,

		ReadDeadLine:  60 * time.Second,
		WriteDeadLine: 60 * time.Second,

		Kick:    KickFun,
		Handler: HandlerFun,
		Accept:  AcceptFun,
	}
	go n.PingPong()
	n.server.Start()
}

func (n *socket) PingPong() {
	t := time.NewTicker(time.Second * 10)
	for {
		<-t.C
		data.WorkerLock.RLock()
		for _, worker := range data.Worker {
			if worker != nil {
				pb := &protocol.System_Ping{}
				n.Send(settings.System_Ping, worker.NetClient, pb)
			}
		}
		data.WorkerLock.RUnlock()
	}
}

func init() {
	function.NetworkManage = &socket{}
}
