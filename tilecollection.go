package main

import (
	"math/rand"
)

type TileCollection struct {
	start, rang int
}

func (tc *TileCollection) randNum() (r int) {
	return rand.Intn(tc.rang) + tc.start
}
