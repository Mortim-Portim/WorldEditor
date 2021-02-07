package main

import (
	"image/color"

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

		ImportWorld(input.GetText(), window)
	})
	return
}

func getExportButton(x, y, h float64, name string, window *Window, input *GE.EditText) (btn *GE.Button) {
	btn = GE.GetTextButton(name, "", GE.StandardFont, x, y, h, color.Black, color.White)
	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			return
		}

		ExportWorld(input.GetText(), window)
	})
	return
}
