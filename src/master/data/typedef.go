package data

import "network"

type worker struct {
	NetClient *network.Client
}
type HeaderKV struct {
	K string
	V string
}
type BoomArgs struct {
	Url         string
	Num         int32
	Qps         int32
	Conc        int32
	TimeOut     int32 // seconds
	UserName    string
	Password    string
	Body        string
	Accept      string
	ContentType string
	Method      string
	Headers     []HeaderKV
	KeepAlive   bool
	ProxyAddr   string
	Host        string
}

type TaskReport struct {
	TaskId int64
	Uid    string
	Url    string
	//Number      int
	Concurrency int64
	Rps         float64

	Totalsecond float64
	Sizet       int64
	Sizereq     int64
	Fastest     float64
	Slowest     float64
	Avg         float64

	C2  int64
	C3  int64
	C4  int64
	C5  int64
	CO  int64
	Err int64

	P10 float64
	P25 float64
	P50 float64
	P75 float64
	P90 float64
	P95 float64
	P99 float64
}
