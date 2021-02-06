package main

import (
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/Compression"
	"github.com/mortim-portim/GraphEng/GE"
)

//Fix Fillbutton to use random texture

func getFillButton(x, y, h float64, window *Window) (btn *GE.Button) {
	btn = GE.GetTextButton("Fill", "Fill", GE.StandardFont, x, y, h, color.Black, color.White)

	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed == false {
			return
		}

		for x := 0; x < window.wrld.TileMat.WAbs(); x++ {
			for y := 0; y < window.wrld.TileMat.HAbs(); y++ {
				window.wrld.TileMat.SetAbs(x, y, int64(window.tilecollection[window.selectedVar].GetStart()))
			}
		}
	})
	return
}

func getBrushScrollbar(x, y, w, h float64, window *Window) (scrollbar *GE.ScrollBar) {
	scrollbar = GE.GetStandardScrollbar(x, y, w, h, 0, 10, 0, GE.StandardFont)
	scrollbar.RegisterOnChange(func(sb *GE.ScrollBar) {
		window.brushsize = sb.Current()
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

		folder := "./resource/maps/" + input.GetText() + "/"
		os.Mkdir(folder, 0755)
		os.Mkdir(folder+"tile/", 0755)
		window.wrld.Save(folder + "map.txt")

		index, _ := os.Create(folder + "/tile/#index.txt")

		light := &GE.ImageObj{X: 0, Y: 0, W: 16, H: 16}
		dark := &GE.ImageObj{X: 0, Y: 16, W: 16, H: 16}
		for i, tile := range window.wrld.Tiles {
			name := strconv.Itoa(i) + ".png"
			file, _ := os.Create(folder + "/tile/" + name)
			img, _ := ebiten.NewImage(16, 32, ebiten.FilterDefault)

			tile.Draw(img, light, 255)
			tile.Draw(img, dark, 0)

			png.Encode(file, img)

			index.WriteString(name + "\n")
			file.Close()
		}
	})
	return
}
