package wrldedit

import (
	"math/rand"

	"github.com/mortim-portim/GraphEng/GE"
)

type TileCollection struct {
	name        string
	same        []string
	start, rang int
	subbuttons  []*GE.Button
	index       map[uint8]map[string][]int64
}

func (tc *TileCollection) GetString() string {
	return tc.name
}

func (tc *TileCollection) GetStart() int {
	return tc.start
}

func (tc *TileCollection) GetRange() int {
	return tc.rang
}

func (tc *TileCollection) GetSubButtons() []*GE.Button {
	return tc.subbuttons
}

func (tc *TileCollection) GetIndex(u, l, d, r, ul, ur, dl, dr int, tile string) int64 {
	surround := uint8(1*u + 2*l + 4*d + 8*r + 16*ul + 32*ur + 64*dl + 128*dr)
	redSurround := uint8(1*u + 2*l + 4*d + 8*r)

	surrtiles, avab := tc.index[surround]

	if !avab {
		surrtiles, avab = tc.index[redSurround]
	}

	if !avab {
		surrtiles = tc.index[0]
	}

	cnttile, avab := surrtiles[tile]

	if !avab {
		cnttile, avab = surrtiles["Default"]
	}

	if !avab {
		return int64(tc.start)
	}

	return cnttile[rand.Intn(len(cnttile))]
}

func (tc *TileCollection) GetSame(name string) bool {
	if name == tc.name {
		return true
	}

	for _, nam := range tc.same {
		if name == nam {
			return true
		}
	}

	return false
}
