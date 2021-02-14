package wrldedit

import (
	"image/color"
	"math"
	"strconv"

	"github.com/mortim-portim/GraphEng/GE"
)

func getTabButton(x, y, h float64, id int, name string, window *Window) (btn *GE.Button) {
	btn = GE.GetTextButton(name, "", GE.StandardFont, x, y, h, color.Black, color.White)
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
		i := len(window.wrld.Region)
		regbtn := GE.GetImageButton(region.color.Img, float64(1000+(i%8)*70), 550+(math.Ceil(float64(i/8)))*70, 64, 64)
		regbtn.Data = i + 1
		regbtn.RegisterOnLeftEvent(func(btn *GE.Button) {
			window.selectedVar = btn.Data.(int)
		})
		window.wrld.Region = append(window.wrld.Region, region)
		window.regionbuttons.Add(regbtn)
	})
	return
}

func getRegionAlphaScrollbar(x, y, w, h float64, window *Window) (scrollbar *GE.ScrollBar) {
	scrollbar = GE.GetStandardScrollbar(x, y, w, h, 0, 100, 0, GE.StandardFont)
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
	scrollbar.RegisterOnChange(func(sb *GE.ScrollBar) {
		window.wrld.SetLightLevel(int16(sb.Current()))
	})

	return
}
