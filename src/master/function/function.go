package function

import (
	"github.com/golang/protobuf/proto"
	"network"
)

type ProtocolManager interface {
	Start()
	Get(uint32) func(*network.Client, *network.Context)
}

type NetworkManager interface {
	Start()
	Send(method uint32, client *network.Client, pb proto.Message)
}

var ProtocolManage ProtocolManager
var NetworkManage NetworkManager
