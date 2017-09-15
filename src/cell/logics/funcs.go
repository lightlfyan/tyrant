package logics

import (
	"math/rand"
	"strconv"
)

type FuncsType struct {
	l []string
}

func NewFuncsType() *FuncsType {
	max := len(FakeList)

	if max <= 0 {
		return &FuncsType{[]string{}}
	}

	idx := rand.Intn(max)
	return &FuncsType{FakeList[idx]}
}

func (f FuncsType) Label() string {
	return "funcs"
}

// rand.Intn(max - min) + min

func (f FuncsType) Random(arg ...int) string {
	if len(arg) > 0 {
		min := arg[0]
		max := arg[1]
		return strconv.Itoa(rand.Intn(max-min) + min + int(CellId))
	}
	return strconv.Itoa(rand.Intn(10000000000) + int(CellId))
}

func (f FuncsType) rand(i int) string {
	if len(f.l) < i || i < 0 {
		return "0"
	}
	return f.l[i-1]
}

func (f FuncsType) P1() string {
	return f.rand(1)
}
func (f FuncsType) P2() string {
	return f.rand(2)
}
func (f FuncsType) P3() string {
	return f.rand(3)
}
func (f FuncsType) P4() string {
	return f.rand(4)
}
func (f FuncsType) P5() string {
	return f.rand(5)
}
