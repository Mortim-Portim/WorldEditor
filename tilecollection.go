package main

import (
	"marvin/GraphEng/GE"
	"math/rand"
)

type TileCollection interface {
	GetString() string
	GetNum() int16
	GetLast() int
	GetStart() int
	GetRange() int
	GetSubButtons() *GE.Group
}

type DefaultTC struct {
	name        string
	start, rang int
	subbuttons  *GE.Group
}

func (tc *DefaultTC) GetString() string {
	return tc.name
}

func (tc *DefaultTC) GetLast() int {
	return tc.start + tc.rang
}

func (tc *DefaultTC) GetStart() int {
	return tc.start
}

func (tc *DefaultTC) GetRange() int {
	return tc.rang
}

func (tc *DefaultTC) GetSubButtons() *GE.Group {
	return tc.subbuttons
}

type RandomTC struct {
	DefaultTC
}

func (tc *RandomTC) GetNum() int16 {
	return int16(rand.Intn(tc.rang) + tc.start)
}

type ConnectedTC struct {
	DefaultTC
	index []uint8
}

func (tc *ConnectedTC) GetNum() int16 {
	return int16(tc.start)
}

func (tc *ConnectedTC) GetIndex(n, w, s, e, nw, ne, sw, se int) int {
	idx, subIdx := 0, 0
	surround := uint8(1*n + 2*w + 4*s + 8*e + 16*nw + 32*ne + 64*sw + 128*se)
	redSurround := uint8(1*n + 2*w + 4*s + 8*e)

	for i, num := range tc.index {
		if num == surround {
			idx = i
		}

		if num == redSurround {
			subIdx = i
		}
	}

	if idx == 0 {
		idx = subIdx
	}

	//fmt.Printf("%v %v %v %v %v %v %v %v -> %v %v %v\n", n, w, s, e, nw, ne, sw, se, surround, redSurround, idx)

	return idx + tc.start
}
