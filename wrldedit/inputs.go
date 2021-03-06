package wrldedit

import (
	"math"
	"strconv"

	"github.com/mortim-portim/GraphEng/GE"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//Todo: once per Tile
//Bug: Object Matrix upadate after moving

func mousebuttonleftPressed(w *Window) {
	dx, dy := ebiten.CursorPosition()
	wx, wy := w.wrld.GetTopLeft()
	x := (float64(dx) - wx) / w.wrld.GetTileS()
	y := (float64(dy) - wy) / w.wrld.GetTileS()
	ix, iy := int(x), int(y)

	if x >= 0 && ix < w.wrld.TileMat.W() && y >= 0 && iy < w.wrld.TileMat.H() {
		switch w.tabview.CurrentTab {
		case 0:
			if w.useSub {
				w.wrld.TileMat.Set(ix, iy, int64(w.selectedVar))
			} else {
				brushsize := w.brushsize

				w.wrld.TileMat.Fill(ix-brushsize, iy-brushsize, ix+brushsize, iy+brushsize, int64(w.tilecollection[w.selectedVar].GetStart()))
				for dx := ix - brushsize - 1; dx <= ix+brushsize+1; dx++ {
					for dy := iy - brushsize - 1; dy <= iy+brushsize+1; dy++ {
						connectTiles(dx, dy, w)
					}
				}
			}
		case 2:
			w.wrld.RegionMat.Fill(ix-w.brushsize, iy-w.brushsize, ix+w.brushsize, iy+w.brushsize, int64(w.selectedVar))
		}

		/*lightID, _ := w.wrld.LIdxMat.Get(ix, iy)
		if lightID == -1 {
			w.wrld.AddLights(GE.GetLightSource(&GE.Point{float64(x) + w.wrld.TileMat.Focus().Min().X, float64(y) + w.wrld.TileMat.Focus().Min().Y}, &GE.Vector{0, -1, 0}, 360, 400, 0.01, false))
		}*/
	}
}

func connectTiles(x, y int, window *Window) {
	tileID, _ := window.wrld.TileMat.Get(x, y)
	tileName := window.wrld.Tiles[tileID].Name
	tcID, _ := strconv.Atoi(tileName)
	tc := window.tilecollection[tcID]

	surrTileName := make(map[string]int)
	surrtile := make([]int, 0)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			index, _ := window.wrld.TileMat.Get(dx+x, dy+y)
			surrname := window.wrld.Tiles[int(index)].Name
			surid, _ := strconv.Atoi(surrname)

			if tc.GetSame(window.tilecollection[surid].name) {
				surrtile = append(surrtile, 0)
			} else {
				surrtile = append(surrtile, 1)
				surrTileName[surrname]++
			}
		}
	}

	surrTileName["Default"] = 0
	most := "Default"
	for tilenam, value := range surrTileName {
		if value > surrTileName[most] {
			most = tilenam
		}
	}

	if most != "Default" {
		id, _ := strconv.Atoi(most)
		most = window.tilecollection[id].name
	}

	window.wrld.TileMat.Set(x, y, int64(tc.GetIndex(surrtile[3], surrtile[1], surrtile[4], surrtile[6], surrtile[0], surrtile[5], surrtile[2], surrtile[7], most)))
}

func mousebuttonleftJustPressed(w *Window) {
	dx, dy := ebiten.CursorPosition()
	wx, wy := w.wrld.GetTopLeft()
	x := (float64(dx) - wx) / w.wrld.GetTileS()
	y := (float64(dy) - wy) / w.wrld.GetTileS()
	ix, iy := int(x), int(y)
	tx, ty := w.wrld.TileMat.Focus().Min().X, w.wrld.TileMat.Focus().Min().Y

	if x >= 0 && ix < w.wrld.TileMat.W() && y >= 0 && iy < w.wrld.TileMat.H() {
		switch w.tabview.CurrentTab {
		case 1:
			rx, ry := math.Floor(x+tx), math.Floor(y+ty)
			structObj := GE.GetStructureObj(w.wrld.Structures[w.curType], rx, ry)
			w.curretObject = structObj
			w.wrld.AddStructObj(structObj)
			w.wrld.UpdateObjMat()
		}
	}
}

func mousebuttonrightPressed(w *Window) {
	dx, dy := ebiten.CursorPosition()
	wx, wy := w.wrld.GetTopLeft()
	x := (float64(dx) - wx) / w.wrld.GetTileS()
	y := (float64(dy) - wy) / w.wrld.GetTileS()
	ix, iy := int(x), int(y)

	if x >= 0 && ix < w.wrld.TileMat.W() && y >= 0 && iy < w.wrld.TileMat.H() {
		switch w.tabview.CurrentTab {
		case 1:
			structureID, _ := findObject(w.wrld, x+w.wrld.TileMat.Focus().Min().X, y+w.wrld.TileMat.Focus().Min().Y)

			if structureID == -1 {
				break
			}

			if structureID >= 0 {
				w.wrld.Objects[structureID] = w.wrld.Objects[len(w.wrld.Objects)-1]
				w.wrld.Objects = w.wrld.Objects[:len(w.wrld.Objects)-1]
				w.wrld.UpdateObjMat()
			}
		case 2:
			/*lightID, _ := w.wrld.LIdxMat.Get(ix, iy)
			if lightID >= 0 {
				w.wrld.Lights[lightID] = w.wrld.Lights[len(w.wrld.Lights)-1]
				w.wrld.Lights = w.wrld.Lights[:len(w.wrld.Lights)-1]
				w.wrld.RemoveLight(int(lightID))
			}*/
		}
	}
}

func keyPressed(w *Window) {
	if w.curType == 1 && w.curretObject != nil {
		objctX := w.curretObject.Hitbox.Min().X
		objctY := w.curretObject.Hitbox.Min().Y
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			w.curretObject.SetToXY(objctX, objctY-0.1)
			w.wrld.UpdateObjMat()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			w.curretObject.SetToXY(objctX, objctY+0.1)
			w.wrld.UpdateObjMat()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			w.curretObject.SetToXY(objctX+0.1, objctY)
			w.wrld.UpdateObjMat()
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			w.curretObject.SetToXY(objctX-0.1, objctY)
			w.wrld.UpdateObjMat()
		}
	}
}

func findObject(w *WorldStructure, x, y float64) (idx int, obj *GE.StructureObj) {
	for i, obj := range w.Objects {
		point := GE.Point{x, y}
		if point.InBounds(obj.Drawbox) {
			return i, obj
		}
	}
	return -1, nil
}
