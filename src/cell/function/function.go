package function

import (
	"github.com/golang/protobuf/proto"
	"network"
)

type HandlefunType func(*network.Context)

type ICellProtocolManager interface {
	Start()
	Get(uint32) HandlefunType
}

type ICellNetworkManager interface {
	Start(addr *string)
	Send(method uint32, pb proto.Message)
}

var ProtocolManage ICellProtocolManager
var NetworkManage ICellNetworkManager
