package api

import (
	"master/protocol"
	"master/protocol/pack"
	"time"

	"network"
)

func System_Heart_C2S_Handler(client *network.Client, context *network.Context, pb *protocol.System_Heart_C2S) {
	pack.System_Heart_S2C(client, time.Now().UnixNano())
}
