package main

import (
	"math"
	"strconv"

	"github.com/mortim-portim/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
)

//Todo: once per Tile

func mousebuttonleftPressed(w *Window) {
	dx, dy := ebiten.CursorPosition()
	wx, wy := w.wrld.GetTopLeft()
	x := int(math.Floor((float64(dx) - wx) / w.wrld.GetTileS()))
	y := int(math.Floor((float64(dy) - wy) / w.wrld.GetTileS()))

	if x >= 0 && x < w.wrld.TileMat.W() && y >= 0 && y < w.wrld.TileMat.H() {
		switch w.curType {
		case 0:
			if w.useSub {
				w.wrld.TileMat.Set(x, y, int64(w.selectedVar))
			} else {
				brushsize := w.brushsize

				w.wrld.TileMat.Fill(x-brushsize, y-brushsize, x+brushsize, y+brushsize, int64(w.tilecollection[w.selectedVar].GetStart()))
				for dx := x - brushsize - 1; dx <= x+brushsize+1; dx++ {
					for dy := y - brushsize - 1; dy <= y+brushsize+1; dy++ {
						connectTiles(dx, dy, w)
					}
				}
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
				w.wrld.AddLights(GE.GetLightSource(&GE.Point{float64(x) + w.wrld.TileMat.Focus().Min().X, float64(y) + w.wrld.TileMat.Focus().Min().Y}, &GE.Vector{0, -1, 0}, 360, 400, 0.01, false))
			}
		}
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

			if surrname == tileName {
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

func getOnPos(x, y, s, r int, tilemat *GE.Matrix) int {
	if x < 0 || x >= tilewidth || y < 0 || y >= tileheight {
		return 0
	}

	tileID, _ := tilemat.Get(x, y)
	if tileID >= int64(s) && tileID < int64(s+r) {
		return 0
	}

	return 1
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
				w.wrld.RemoveLight(int(lightID))
			}
		}
	}
}
