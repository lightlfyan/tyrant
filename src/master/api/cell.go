package api

import (
	"log"
	"master/data"
	"master/function"
	"master/protocol"
	"master/settings"

	"master/protocol/pack"
	"sync/atomic"

	"network"
)

func Cell_Login_C2S(client *network.Client, context *network.Context, pb *protocol.Cell_Login_C2S) {
	workerId, result := data.AddWork(client)
	log.Println("login", workerId)

	sendpb := &protocol.Cell_Login_S2C{
		Id:     &workerId,
		Result: &result,
	}

	function.NetworkManage.Send(settings.Cell_Login_S2C, client, sendpb)
}

func StopCurrent() {
	data.WorkerLock.RLock()
	data.WorkerLock.RUnlock()

	atomic.AddInt32(&data.CellStopCounter, data.WorkerCount)
	for _, worker := range data.Worker {
		if worker != nil {
			pack.Cell_Stop_S2C(worker.NetClient)
		}
	}

}

func Cell_Stop_C2S() {
	data.TaskGlobalLock.Lock()
	defer data.TaskGlobalLock.Unlock()

	data.CellStopCounter -= 1

	if data.CellStopCounter <= 0 {
		data.RemoveCurrTaskNoLock()
	}
}
