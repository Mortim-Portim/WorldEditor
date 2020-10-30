package main

import (
	"image/color"
	"io/ioutil"
	"marvin/GraphEng/Compression"
	"marvin/GraphEng/GE"
	"os"
)

func getAutocompleteButton(x, y, h float64, window *Window) (btn *GE.Button) {
	btn = GE.GetTextButton("Connect", "Edasdit", GE.StandardFont, x, y, h, color.Black, color.White)

	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed == false {
			return
		}

		for x := 0; x < tilewidth; x++ {
			for y := 0; y < tileheight; y++ {
				tileID, _ := window.wrld.TileMat.GetAbs(x, y)
				tilename := window.wrld.Tiles[tileID].Name
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
					n := getOnPos(x, y-1, tc.start, tc.rang, window.wrld.TileMat)
					w := getOnPos(x-1, y, tc.start, tc.rang, window.wrld.TileMat)
					s := getOnPos(x, y+1, tc.start, tc.rang, window.wrld.TileMat)
					e := getOnPos(x+1, y, tc.start, tc.rang, window.wrld.TileMat)
					nw := getOnPos(x-1, y-1, tc.start, tc.rang, window.wrld.TileMat)
					ne := getOnPos(x+1, y, tc.start, tc.rang, window.wrld.TileMat)
					sw := getOnPos(x-1, y+1, tc.start, tc.rang, window.wrld.TileMat)
					se := getOnPos(x+1, y+1, tc.start, tc.rang, window.wrld.TileMat)

					window.wrld.TileMat.SetAbs(x, y, int16(tc.GetIndex(n, w, s, e, ne, nw, sw, se)))
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

	tileID, _ := tilemat.GetAbs(x, y)
	if tileID >= int16(s) && tileID < int16(s+r) {
		return 0
	}

	return 1
}

func getFillButton(x, y, h float64, window *Window) (btn *GE.Button) {
	btn = GE.GetTextButton("Fill", "Fill", GE.StandardFont, x, y, h, color.Black, color.White)

	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed == false {
			return
		}

		for x := 0; x < window.wrld.TileMat.WAbs(); x++ {
			for y := 0; y < window.wrld.TileMat.HAbs(); y++ {
				window.wrld.TileMat.SetAbs(x, y, window.tilecollection[window.selectedVar].GetNum())
			}
		}
	})
	return
}

func getLightlevelScrollbar(x, y, w, h float64, window *Window) (scrollbar *GE.ScrollBar) {
	scrollbar = GE.GetStandardScrollbar(x, y, w, h, 0, 255, 255, GE.StandardFont)
	scrollbar.RegisterOnChange(func(sb *GE.ScrollBar) {
		window.wrld.SetLightLevel(int16(sb.Current()))
	})

	return
}

func getTabButton(x, y, h float64, id int, name string, window *Window) (btn *GE.Button) {
	btn = GE.GetTextButton(name, "", GE.StandardFont, x, y, h, color.Black, color.White)
	btn.Data = id
	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		window.curType = btn.Data.(int)
	})
	return
}

func getPathLabel(x, y, h float64, max int) (lbl *GE.EditText) {
	lbl = GE.GetEditText("Path", x, y, h, max, GE.StandardFont, color.Black, color.White)
	return
}

func getImportButton(x, y, h float64, name string, window *Window, input *GE.EditText) (btn *GE.Button) {
	btn = GE.GetTextButton(name, "", GE.StandardFont, x, y, h, color.Black, color.White)
	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			return
		}

		data, err1 := ioutil.ReadFile("./resource/maps/" + input.GetText() + ".map")
		if err1 != nil {
			return
		}

		bs := Compression.DecompressAll(data, []int{8, 8, 2, 2, 8})
		tilMat := GE.GetMatrix(0, 0, 0)
		err2 := tilMat.Decompress(bs[5])
		if err2 != nil {
			return
		}

		window.wrld.TileMat = tilMat

		window.wrld.Objects = nil
		window.wrld.BytesToObjects(bs[6])
		window.wrld.UpdateObjMat()

		window.wrld.Lights = nil
		window.wrld.BytesToLights(bs[7])
		window.wrld.UpdateLIdxMat()

		window.wrld.SetMiddle(int(Compression.BytesToInt64(bs[0])), int(Compression.BytesToInt64(bs[1])), false)
		window.wrld.SetLightStats(Compression.BytesToInt16(bs[2]), Compression.BytesToInt16(bs[3]), Compression.BytesToFloat64(bs[4]))
	})
	return
}

func getExportButton(x, y, h float64, name string, window *Window, input *GE.EditText) (btn *GE.Button) {
	btn = GE.GetTextButton(name, "", GE.StandardFont, x, y, h, color.Black, color.White)
	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			return
		}

		window.wrld.Save("./resource/maps/" + input.GetText() + ".map")

		file, _ := os.Create("./resource/maps/" + input.GetText() + ".index")

		for _, line := range window.importedTiles {
			file.WriteString(line + "\n")
		}
	})
	return
}
