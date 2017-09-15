package data

import (
	"log"
	"master/protocol"
	"sync"
	"sync/atomic"

	"time"

	"master/function"
	"master/settings"

	"master/protocol/pack"

	"sort"

	"network"
)

const (
	MAX_WORKER = 10
	Token      = "tyrant-token"
	Uid        = "tyrant-uid"
)

var taskId int64

var CurrTaskId int64
var CurrTaskCounter int32
var CellFakeLoadCounter int32
var CellStopCounter int32

var Worker map[int32]*worker
var WorkerCount int32 = 0
var WorkerLock *sync.RWMutex

var CellReports map[int64][]*protocol.Task_Report_C2S
var CellMailBox (chan *protocol.Task_Report_C2S)

var TaskProgress *TaskProgressType

var TaskGlobalLock *sync.RWMutex

var TaskQueue map[int64]*protocol.Task_Create_S2C
var FakeQueue chan string

var UploadLock *sync.Mutex

func init() {

	UploadLock = &sync.Mutex{}

	FakeQueue = make(chan string, 100)

	TaskProgress = &TaskProgressType{}
	TaskProgress.Init()

	CurrTaskId = 1
	CellFakeLoadCounter = 0
	CellStopCounter = 0

	WorkerLock = &sync.RWMutex{}

	TaskGlobalLock = &sync.RWMutex{}

	Worker = make(map[int32]*worker, MAX_WORKER)
	var i int32
	for i = 0; i < MAX_WORKER; i++ {
		Worker[i] = nil
	}

	CellReports = make(map[int64][]*protocol.Task_Report_C2S)
	CellMailBox = make(chan *protocol.Task_Report_C2S, 1000)

	CurrTaskCounter = 0
	TaskQueue = make(map[int64]*protocol.Task_Create_S2C, 100)

	boltInit()
}

func Start() {
	go ReceiveReport()
}

func GenTaskId() int64 {
	return atomic.AddInt64(&taskId, 1)
}

func ResetTaskProgress() {
	for i := 0; i < MAX_WORKER; i++ {
		TaskProgress.Set(int32(i), 0)
	}
}

func AddWork(client *network.Client) (id int32, result bool) {
	WorkerLock.Lock()
	defer WorkerLock.Unlock()

	id = -1
	result = false
	for i := 0; i < MAX_WORKER; i++ {
		c, ok := Worker[int32(i)]
		if ok && c == nil {
			id = int32(i)
			break
		}
	}

	if id == -1 {
		return
	}

	client.Uid = id

	// 同一台机器只能有一个cell登录
	for _, v := range Worker {
		if v != nil && v.NetClient.Conn.RemoteAddr().String() == client.Conn.RemoteAddr().String() {
			return
		}
	}

	Worker[id] = &worker{NetClient: client}
	WorkerCount += 1

	result = true
	return
}

func DelWork(id int32) {
	WorkerLock.Lock()
	defer WorkerLock.Unlock()

	_, ok := Worker[id]
	if ok {
		Worker[id] = nil
		WorkerCount -= 1
	}
}

func ReceiveReport() {

	for {
		// 更新文件
		if atomic.LoadInt32(&CellFakeLoadCounter) > 1 {
			time.Sleep(time.Second)
			log.Println("CellFakeLoadCounter: ", CellFakeLoadCounter)
			continue
		}

		// 停止任务
		if atomic.LoadInt32(&CellStopCounter) > 1 {
			time.Sleep(time.Second)
			log.Println("CellStopCounter: ", CellStopCounter)
			continue
		}

		if atomic.LoadInt32(&CurrTaskCounter) <= 0 {

			// 没有任务的时候 尝试更新文件或者 发送新任务
			select {
			case fileName := <-FakeQueue:
				cellUpdateFake(fileName)
			default:
				fetchTask()
			}

		} else {
			select {
			case r := <-CellMailBox:
				TaskProgress.Set(*r.Cellid, 100)
				TaskGlobalLock.Lock()

				list, ok := CellReports[*r.Taskid]
				if !ok {
					list = make([]*protocol.Task_Report_C2S, 0, 10)
				}

				list = append(list, r)
				CellReports[*r.Taskid] = list

				atomic.AddInt32(&CurrTaskCounter, -1)
				if CurrTaskCounter <= 0 {
					whenTaskDone()
				}

				TaskGlobalLock.Unlock()
			default:
			}
		}
	}
}

func RemoveCurrTaskNoLock() {
	ResetTaskProgress()

	delete(CellReports, CurrTaskId)

	if _, ok := TaskQueue[CurrTaskId]; ok {
		delete(TaskQueue, CurrTaskId)
		CurrTaskId += 1
	}

	CurrTaskCounter = 0
}

func whenTaskDone() {
	AddReport(genReport())

	delete(CellReports, CurrTaskId)
	delete(TaskQueue, CurrTaskId)

	CurrTaskId += 1
}

func genReport() *TaskReport {
	task := TaskQueue[CurrTaskId]

	reports := CellReports[CurrTaskId]

	var totalsecond float64 = 0
	var rps float64 = 0
	var sizet int64 = 0
	var sizereq int64 = 0
	var fastest float64 = 10000000000
	var slowest float64 = 0
	var avg float64 = 0
	var c2 int64 = 0
	var c3 int64 = 0
	var c4 int64 = 0
	var c5 int64 = 0
	var co int64 = 0
	var conc int64 = 0
	var errorcode int64 = 0

	var p10 float64 = 0
	var p25 float64 = 0
	var p50 float64 = 0
	var p75 float64 = 0
	var p90 float64 = 0
	var p95 float64 = 0
	var p99 float64 = 0

	for _, report := range reports {
		if report != nil {

			rps += *report.Rps
			sizet += *report.Sizetotal

			avg = (avg + *report.Average) / 2
			conc += *report.Conc

			if *report.Totalsecond > totalsecond {
				totalsecond = *report.Totalsecond
			}

			sizereq = *report.Sizereq

			if *report.Fastest < fastest {
				fastest = *report.Fastest
			}

			if *report.Slowest > slowest {
				slowest = *report.Slowest
			}

			c2 += *report.Code2
			c3 += *report.Code3
			c4 += *report.Code4
			c5 += *report.Code5
			co += *report.CodeOther
			errorcode += *report.Error

			p10 = *report.P10
			p25 = *report.P25
			p50 = *report.P50
			p75 = *report.P75
			p90 = *report.P90
			p95 = *report.P95
			p99 = *report.P99
		}
	}

	return &TaskReport{
		TaskId:      CurrTaskId,
		Uid:         *task.Uid,
		Url:         *task.Url,
		Concurrency: conc,
		Rps:         rps,
		Totalsecond: totalsecond,
		Sizet:       sizet,
		Sizereq:     sizereq,
		Fastest:     fastest,
		Slowest:     slowest,
		Avg:         avg,
		C2:          c2,
		C3:          c3,
		C4:          c4,
		C5:          c5,
		CO:          co,
		Err:         errorcode,

		P10: p10,
		P25: p25,
		P50: p50,
		P75: p75,
		P90: p90,
		P95: p95,
		P99: p99,
	}
}

func fetchTask() {

	TaskGlobalLock.Lock()
	defer TaskGlobalLock.Unlock()

	WorkerLock.RLock()
	defer WorkerLock.RUnlock()

	if len(TaskQueue) <= 0 {
		return
	}

	keys := make([]int, 0, len(TaskQueue))
	for k := range TaskQueue {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	var pb *protocol.Task_Create_S2C = nil

	for _, k := range keys {
		CurrTaskId = int64(k)
		pb = TaskQueue[CurrTaskId]
		if pb != nil {
			break
		} else {
			delete(TaskQueue, CurrTaskId)
		}
	}

	if pb == nil {
		CurrTaskId = 0
		return
	}

	ResetTaskProgress()
	CurrTaskCounter = WorkerCount
	for _, worker := range Worker {
		if worker != nil {
			function.NetworkManage.Send(settings.Task_Create_S2C, worker.NetClient, pb)
		}
	}
}

func cellUpdateFake(filename string) {
	WorkerLock.RLock()
	WorkerLock.RUnlock()

	CellFakeLoadCounter = WorkerCount
	for _, worker := range Worker {
		if worker != nil {
			pack.Cell_Fake_S2C(worker.NetClient, filename)
		}
	}
}
