package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"master/data"
	"master/protocol"
	"sort"
	"strconv"
	"strings"
)

type TaskProgress struct {
	Progress []*data.HeaderKV
	Clients  []*data.HeaderKV
}

func CreateTask(args *data.BoomArgs, uid string) (taskid int64, stat bool) {
	log.Println("CreateTask try get lock")

	taskid = 0
	stat = false

	data.TaskGlobalLock.Lock()
	defer data.TaskGlobalLock.Unlock()

	log.Println("CreateTask =============")

	workCount := data.WorkerCount

	if workCount <= 0 {
		log.Println("create fail: workcount == ", workCount)
		return
	}

	log.Println("call create s2c: ", args.Url)

	kvs := make([]*protocol.KV, 0, 10)
	for _, kv := range args.Headers {
		kvs = append(kvs, &protocol.KV{K: &kv.K, V: &kv.V})
	}
	num1 := args.Num / workCount
	conc1 := args.Conc / workCount

	newTaskId := data.GenTaskId()

	pb := &protocol.Task_Create_S2C{
		TaskId:      &newTaskId,
		Url:         &args.Url,
		Num:         &num1,
		Conc:        &conc1,
		Qps:         &args.Qps,
		Timeout:     &args.TimeOut,
		UserName:    &args.UserName,
		Password:    &args.Password,
		Body:        &args.Body,
		Accept:      &args.Accept,
		ContentType: &args.ContentType,
		Method:      &args.Method,
		KeepAlive:   &args.KeepAlive,
		ProxyAddr:   &args.ProxyAddr,
		Headers:     kvs,
		Host:        &args.Host,
		Uid:         &uid,
	}

	data.TaskQueue[newTaskId] = pb

	taskid = newTaskId
	stat = true
	return
}

func GetStatus() []byte {

	data.TaskGlobalLock.RLock()
	defer data.TaskGlobalLock.RUnlock()

	data.WorkerLock.RLock()
	defer data.WorkerLock.RUnlock()

	m1 := make(map[string]interface{}, data.MAX_WORKER)
	url, ok := data.TaskQueue[data.CurrTaskId]
	if ok {
		m1["url"] = url.Url
	} else {
		m1["url"] = "无"
	}

	m := make([]map[string]interface{}, data.MAX_WORKER)

	for clientId := 0; clientId < data.MAX_WORKER; clientId++ {
		worker, _ := data.Worker[int32(clientId)]
		if worker != nil {
			var addr string = worker.NetClient.Conn.RemoteAddr().String()
			p, _ := data.TaskProgress.Get(int32(clientId))
			m[clientId] = map[string]interface{}{
				addr: p,
			}
		} else {
			m[clientId] = nil
		}
	}

	m1["clients"] = m

	j, err := json.Marshal(m1)
	if err != nil {
		log.Println(err)
	}

	return j

}

func GetTaskQueueStatus() string {

	data.TaskGlobalLock.RLock()
	defer data.TaskGlobalLock.RUnlock()

	running := data.CurrTaskId
	if data.CurrTaskCounter <= 0 {
		running = 0
	}

	b := bytes.NewBufferString("")

	tmpl := `
		 <tr>
            <td>运行中:</td>
            <td>%d</td>
            <td></td>
        </tr>
        `
	tmpl = fmt.Sprintf(tmpl, running)
	b.WriteString(tmpl)

	keys := make([]int, 0, len(data.TaskQueue))

	for k, _ := range data.TaskQueue {
		keys = append(keys, int(k))
	}

	sort.Ints(keys)

	for _, id := range keys {
		v := data.TaskQueue[int64(id)]
		if v == nil {
			continue
		}

		if int64(id) == data.CurrTaskId && data.CellStopCounter > 0 {
			b.WriteString(fmt.Sprintf("<tr><td>%d</td><td>%s</td><td><span class='label label-danger'>停止中</span></td></tr>", id, *v.Url))
		} else {
			b.WriteString(fmt.Sprintf("<tr><td>%d</td><td>%s</td><td><a href='/deletetask?taskid=%d'><span class='label label-danger'>删除</span></a></td></tr>", id, *v.Url, id))
		}
	}
	return b.String()

	//return json.Marshal(m)
}

func GetReport(uid string) string {
	b := bytes.NewBufferString("")

	ks := make([]int, 0)
	dump := make(map[int]*data.TaskReport)

	data.TaskReportCache.ForeachRead(func(k int64, v1 interface{}) {
		v := v1.(*data.TaskReport)
		if v.Uid == uid || uid == "1000" {
			ks = append(ks, int(k))
			dump[int(k)] = v
		}
	})

	sort.Ints(ks)

	for _, key := range ks {
		report := dump[key]

		//"<a href='/chart?taskid=%d'>%s</a>"+

		rurl := strings.Split(report.Url, ";")[0]

		b.WriteString(fmt.Sprintf("<tr> <td colspan=22>"+
			"<a href='/deletereport?taskid=%d'>"+
			"<span class=\"label label-danger\">删除[%d]</span>"+
			"</a> "+
			"<span class=\"label label-info\">%s</span>"+
			"</td></tr>", report.TaskId, report.TaskId, rurl))

		fmtstr1 := fmt.Sprintf(`
			<tr>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%f</td>
			<td class=\"stats\">%f</td>
			<td class=\"stats\">%f</td>
			<td class=\"stats\">%f</td>

			<td class=\"stats\"><span class="label label-success">%.2f</span></td>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%d</td>
			<td class=\"stats\">%f</td>
			<td class=\"stats\">%f</td>
			<td class=\"stats\">%f</td>
			<td class=\"stats\">%f</td>
			<td class=\"stats\">%f</td>
			<td class=\"stats\">%f</td>
			<td class=\"stats\">%f</td>
			</tr>`,
			report.TaskId,
			report.Concurrency,
			report.Totalsecond,
			report.Fastest,
			report.Slowest,
			report.Avg,

			report.Rps,
			report.Sizet,
			report.Sizereq,
			report.C2,
			report.C3,
			report.C4,
			report.C5,
			report.CO,
			report.Err,
			report.P10,
			report.P25,
			report.P50,
			report.P75,
			report.P90,
			report.P95,
			report.P99,
		)
		b.WriteString(fmtstr1)
	}
	return b.String()
}

func GetChartData(TaskId string) string {
	Id, err := strconv.Atoi(TaskId)
	if err != nil {
		return ""
	}

	url, ok := data.TaskReportCache.Get(int64(Id))
	if !ok {
		return ""
	}

	b := bytes.NewBufferString("")

	var p10 float64 = 0
	var p25 float64 = 0
	var p50 float64 = 0
	var p75 float64 = 0
	var p90 float64 = 0
	var p95 float64 = 0
	var p99 float64 = 0

	reports := data.CellReports[int64(Id)]
	for _, report := range reports {
		p10 = *report.P10 * 1000
		p25 = *report.P25 * 1000
		p50 = *report.P50 * 1000
		p75 = *report.P75 * 1000
		p90 = *report.P90 * 1000
		p95 = *report.P95 * 1000
		p99 = *report.P99 * 1000
	}

	b.WriteString("fill([10, 25, 50, 75, 90, 95, 99],")
	b.WriteString(fmt.Sprintf("[%d, %d, %d, %d, %d, %d, %d],", int(p10), int(p25), int(p50), int(p75), int(p90), int(p95), int(p99)))
	b.WriteString(fmt.Sprintf("\"%s\"", url))
	b.WriteString(");")
	return b.String()
}
