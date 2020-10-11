package main

import (
	"image/color"
	"marvin/GraphEng/GE"
)

func getAutocompleteButton(x, y, h float64, textCol, backCol color.Color, wrld *GE.WorldStructure, window *window) (btn *GE.Button) {
	btn = GE.GetTextButton("Connect", "Edasdit", GE.StandardFont, x, y, h, textCol, backCol)

	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed == false {
			return
		}

		for x := 0; x < tilewidth; x++ {
			for y := 0; y < tileheight; y++ {
				tilename := wrld.Tiles[window.tileMat.GetAbs(x, y)].Name
				var tilecollection TileCollection

				for _, antilecollection := range window.tilecollection {
					if antilecollection.GetString() == tilename {
						tilecollection = antilecollection
						break
					}
				}

				if tilecollection == nil {
					continue
				}

				switch tc := tilecollection.(type) {
				case *ConnectedTC:
					n := getOnPos(x, y-1, tc.start, tc.rang, window.tileMat)
					w := getOnPos(x-1, y, tc.start, tc.rang, window.tileMat)
					s := getOnPos(x, y+1, tc.start, tc.rang, window.tileMat)
					e := getOnPos(x+1, y, tc.start, tc.rang, window.tileMat)

					window.tileMat.SetAbs(x, y, int16(tc.GetIndex(n, w, s, e)))
				}
			}
		}
	})

	return
}

func getOnPos(x, y, s, r int, tilemat *GE.Matrix) int {
	if x < 0 || x >= tilewidth || y < 0 || y >= tileheight {
		return 0
	}

	if tilemat.GetAbs(x, y) >= int16(s) && tilemat.GetAbs(x, y) < int16(s+r) {
		return 0
	}

	return 1
}

func getFillButton(x, y, h float64, textCol, backCol color.Color, window *window) (btn *GE.Button) {
	btn = GE.GetTextButton("Fill", "Fill", GE.StandardFont, x, y, h, textCol, backCol)

	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed == false {
			return
		}

		for x := 0; x < window.tileMat.WAbs(); x++ {
			for y := 0; y < window.tileMat.HAbs(); y++ {
				window.tileMat.SetAbs(x, y, window.tilecollection[window.curImg].GetNum())
			}
		}
	})
	return
}
