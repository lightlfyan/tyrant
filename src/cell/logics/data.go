package logics

var CellId int32
var TaskId int64

var CurrBoom *Boomer

var FakeList [][]string

func init() {
	FakeList = make([][]string, 0)
}

func NewTask(boom *Boomer, id int64) {
	if CurrBoom != nil {
		CurrBoom.StopAsync()
	}

	CurrBoom = boom
	TaskId = id
}
