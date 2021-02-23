package wrldedit

import (
	"image/color"
	"strconv"

	"github.com/mortim-portim/GraphEng/GE"
)

func getTabButton(x, y, w, h float64, id int, name string, window *Window) (btn *GE.Button) {
	btn = GE.GetSizedTextButton(name, GE.StandardFont, x, y, w, h, color.Black, color.White)
	btn.Data = id
	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		window.curType = btn.Data.(int)
	})
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

func getToggleFrameButton(x, y, h float64, window *Window) (btn *GE.Button) {
	btn = GE.GetTextButton("Grid", "", GE.StandardFont, x, y, h, color.Black, color.White)
	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			return
		}

		if window.wrld.HasFrame() {
			window.wrld.GetFrame(0, 255, 1)
		} else {
			window.wrld.GetFrame(1, 255, 1)
		}
	})

	return
}

func getCrtNwRegionButton(x, y, h float64, window *Window, rgb *GE.EditText) (btn *GE.Button) {
	btn = GE.GetTextButton("New", "", GE.StandardFont, x, y, h, color.Black, color.White)
	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			return
		}
		col, err := strconv.ParseInt(rgb.GetText(), 16, 64)

		if err != nil {
			return
		}

		r := (col & 0xFF0000) >> 16
		g := (col & 0xFF00) >> 8
		b := (col & 0xFF)

		region := GetRegion(color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		AddRegion(region, window)
	})
	return
}

func getChangeSizeButton(x, y, h float64, window *Window, width, height *GE.EditText) (btn *GE.Button) {
	btn = GE.GetTextButton("Change", "", GE.StandardFont, x, y, h, color.Black, color.White)
	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if !btn.LPressed {
			return
		}

		width, err0 := strconv.Atoi(width.GetText())
		height, err1 := strconv.Atoi(height.GetText())

		if err0 != nil || err1 != nil {
			return
		}

		window.wrld.ScaleTo(width, height)
	})

	return
}

func getRegionAlphaScrollbar(x, y, w, h float64, window *Window) (scrollbar *GE.ScrollBar) {
	scrollbar = GE.GetStandardScrollbar(x, y, w, h, 0, 100, 0, GE.StandardFont)
	scrollbar.HideValue()
	scrollbar.RegisterOnChange(func(sb *GE.ScrollBar) {
		window.wrld.Regionalpha = float64(sb.Current()) / 100
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
	scrollbar.HideValue()
	scrollbar.RegisterOnChange(func(sb *GE.ScrollBar) {
		window.wrld.SetLightLevel(int16(scrollbar.Current()))
	})

	return
}
