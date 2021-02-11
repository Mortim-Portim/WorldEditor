package wrldedit

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

func GetWorldStructure(x, y, w, h float64, wTiles, hTiles, screenWT, screenHT int) *WorldStructure {
	geWrldStrctr := GE.GetWorldStructure(x, y, w, h, wTiles, hTiles, screenWT, screenHT)
	drawer := &GE.ImageObj{W: geWrldStrctr.GetTileS(), H: geWrldStrctr.GetTileS()}
	return &WorldStructure{geWrldStrctr, drawer, make([]Region, 0), false}
}

type WorldStructure struct {
	*GE.WorldStructure

	drawer     *GE.ImageObj
	Region     []Region
	drawRegion bool
}

func (wrld *WorldStructure) Draw(screen *ebiten.Image) {
	wrld.Draw(screen)

	if wrld.drawRegion {
		for y := 0; y < wrld.RegionMat.H(); y++ {
			for x := 0; x < wrld.RegionMat.W(); x++ {
				tile_idx, err := wrld.RegionMat.Get(x, y)
				if err == nil {
					if int(tile_idx) >= 0 && int(tile_idx) < len(wrld.Region) {
						xStart, yStart := wrld.GetTopLeft()
						middleDx, middleDy := wrld.GetMiddleDelta()
						wrld.drawer.X, wrld.drawer.Y = float64(x)*wrld.GetTileS()+xStart+float64(middleDx), float64(y)*wrld.GetTileS()+yStart+float64(middleDy)
						//wrld.Region[idx].Draw(screen, p.drawer, 0)
					}
				}
			}
		}
	}
}
