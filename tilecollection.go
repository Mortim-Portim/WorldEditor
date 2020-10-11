package main

import (
	"math/rand"
)

type TileCollection interface {
	GetString() string
	GetNum() int16
	GetLast() int
	GetStart() int
	GetRange() int
}

type RandomTC struct {
	name        string
	start, rang int
}

func (tc *RandomTC) GetString() string {
	return tc.name
}

func (tc *RandomTC) GetNum() int16 {
	return int16(rand.Intn(tc.rang) + tc.start)
}

func (tc *RandomTC) GetLast() int {
	return tc.start + tc.rang
}

func (tc *RandomTC) GetStart() int {
	return tc.start
}

func (tc *RandomTC) GetRange() int {
	return tc.rang
}

type ConnectedTC struct {
	name        string
	start, rang int
}

func (tc *ConnectedTC) GetString() string {
	return tc.name
}

func (tc *ConnectedTC) GetNum() int16 {
	return int16(tc.start)
}

func (tc *ConnectedTC) GetLast() int {
	return tc.start + tc.rang
}

func (tc *ConnectedTC) GetIndex(n, w, s, e int) int {
	return n + 2*w + 4*s + 8*e + tc.start
}

func (tc *ConnectedTC) GetStart() int {
	return tc.start
}

func (tc *ConnectedTC) GetRange() int {
	return tc.rang
}
