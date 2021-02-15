package wrldedit

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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

func (wrld *WorldStructure) Update(frame int) {
	x, y := ebiten.CursorPosition()
	_, scroll := ebiten.Wheel()
	hasFocus := int(wrld.X) <= x && x < int(wrld.X+wrld.W) && int(wrld.Y) <= y && y < int(wrld.Y+wrld.H)

	if hasFocus {
		shiftpressed := ebiten.IsKeyPressed(ebiten.KeyShift)

		if (ebiten.IsKeyPressed(ebiten.KeyA) && !shiftpressed) || (inpututil.IsKeyJustPressed(ebiten.KeyA) && shiftpressed) {
			wrld.Move(-1, 0, true, false)
		}

		if (ebiten.IsKeyPressed(ebiten.KeyD) && !shiftpressed) || (inpututil.IsKeyJustPressed(ebiten.KeyD) && shiftpressed) {
			wrld.Move(1, 0, true, false)
		}

		if (ebiten.IsKeyPressed(ebiten.KeyW) && !shiftpressed) || (inpututil.IsKeyJustPressed(ebiten.KeyW) && shiftpressed) {
			wrld.Move(0, -1, true, false)
		}

		if (ebiten.IsKeyPressed(ebiten.KeyS) && !shiftpressed) || (inpututil.IsKeyJustPressed(ebiten.KeyS) && shiftpressed) {
			wrld.Move(0, 1, true, false)
		}

		if scroll < 0 {
			wrld.SetDisplayWH(wrld.TileMat.W()-1, wrld.TileMat.H()-1)
		}

		if scroll > 0 {
			wrld.SetDisplayWH(wrld.TileMat.W()-3, wrld.TileMat.H()-3)
		}
	}
}
