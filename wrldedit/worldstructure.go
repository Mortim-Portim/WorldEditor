package wrldedit

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

func GetWorldStructure(x, y, w, h float64, wTiles, hTiles, screenWT, screenHT int) *WorldStructure {
	geWrldStrctr := GE.GetWorldStructure(x, y, w, h, wTiles, hTiles, screenWT, screenHT)
	geWrldStrctr.SetMiddle(0, 0, true)
	geWrldStrctr.SetDisplayWH(18, 16)
	geWrldStrctr.SetLightStats(0, 255, 0)

	return &WorldStructure{geWrldStrctr, make([]*Region, 0), 0}
}

type WorldStructure struct {
	*GE.WorldStructure

	Region      []*Region
	Regionalpha float64
}

func (wrld *WorldStructure) Draw(screen *ebiten.Image) {
	wrld.WorldStructure.Draw(screen)

	for y := 0; y < wrld.RegionMat.H(); y++ {
		for x := 0; x < wrld.RegionMat.W(); x++ {
			tile_idx, err := wrld.RegionMat.Get(x, y)
			tile_idx--
			if err == nil {
				if int(tile_idx) >= 0 && int(tile_idx) < len(wrld.Region) {
					xStart, yStart := wrld.GetTopLeft()
					middleDx, middleDy := wrld.GetMiddleDelta()
					drawer := wrld.GetDrawer()
					drawer.X, drawer.Y = float64(x)*wrld.GetTileS()+xStart+float64(middleDx), float64(y)*wrld.GetTileS()+yStart+float64(middleDy)
					wrld.Region[tile_idx].Draw(screen, drawer, wrld.Regionalpha)
				}
			}
		}
	}
}
