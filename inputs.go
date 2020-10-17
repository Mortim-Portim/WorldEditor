package main

import (
	"marvin/GraphEng/GE"
	"math"

	"github.com/hajimehoshi/ebiten"
)

func mousebuttonleftPressed(w *Window) {
	dx, dy := ebiten.CursorPosition()
	wx, wy := w.wrld.GetTopLeft()
	x := int(math.Floor((float64(dx) - wx) / w.wrld.GetTileS()))
	y := int(math.Floor((float64(dy) - wy) / w.wrld.GetTileS()))

	if x >= 0 && x < w.wrld.TileMat.W() && y >= 0 && y < w.wrld.TileMat.H() {
		switch w.curType {
		case 0:
			if w.useSub {
				w.wrld.TileMat.Set(x, y, int16(w.selectedVar))
			} else {
				w.wrld.TileMat.Set(x, y, w.tilecollection[w.selectedVar].GetNum())
			}
		case 1:
			objID, _ := w.wrld.ObjMat.Get(x, y)
			if objID == 0 {
				structObj := GE.GetStructureObj(w.currentObject, float64(x)+w.wrld.TileMat.Focus().Min().X, float64(y)+w.wrld.TileMat.Focus().Min().Y)
				w.wrld.AddStructObj(structObj)
				w.wrld.UpdateObjMat()
			}
		case 2:
			lightID, _ := w.wrld.LIdxMat.Get(x, y)
			if lightID == -1 {
				light := GE.GetLightSource(&GE.Point{float64(x) + w.wrld.TileMat.Focus().Min().X, float64(y) + w.wrld.TileMat.Focus().Min().Y}, &GE.Vector{0, -1, 0}, 360, 400, 0.01, false)
				w.wrld.Lights = append(w.wrld.Lights, light)
				w.wrld.UpdateLIdxMat()
			}
		}
	}
}

func mousebuttonrightPressed(w *Window) {
	dx, dy := ebiten.CursorPosition()
	wx, wy := w.wrld.GetTopLeft()
	x := int(math.Floor((float64(dx) - wx) / w.wrld.GetTileS()))
	y := int(math.Floor((float64(dy) - wy) / w.wrld.GetTileS()))

	if x >= 0 && x < w.wrld.TileMat.W() && y >= 0 && y < w.wrld.TileMat.H() {
		switch w.curType {
		case 1:
			structureID, _ := w.wrld.ObjMat.Get(x, y)
			structureID--

			if structureID >= 0 {
				w.wrld.Objects[structureID] = w.wrld.Objects[len(w.wrld.Objects)-1]
				w.wrld.Objects = w.wrld.Objects[:len(w.wrld.Objects)-1]
				w.wrld.UpdateObjMat()
			}
		case 2:
			lightID, _ := w.wrld.LIdxMat.Get(x, y)
			if lightID >= 0 {
				w.wrld.Lights[lightID] = w.wrld.Lights[len(w.wrld.Lights)-1]
				w.wrld.Lights = w.wrld.Lights[:len(w.wrld.Lights)-1]
				w.wrld.UpdateLIdxMat()
			}
		}
	}
}
