package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

func GetFrame(rect GE.Rectangle, thickness float64, alpha uint8) (frame *GE.ImageObj) {
	x, y, w, h := rect.Min().X, rect.Min().Y, rect.Bounds().X, rect.Bounds().Y
	frame = &GE.ImageObj{X: x, Y: y, W: w, H: h}

	frameImg, _ := ebiten.NewImage(int(w), int(h), ebiten.FilterDefault)
	frameImg.Fill(&color.RGBA{0, 0, 0, 0})

	left := GE.GetLineOfPoints(0, 0, 0, h, thickness)
	top := GE.GetLineOfPoints(0, 0, w, 0, thickness)
	right := GE.GetLineOfPoints(w, 0, w, h, thickness)
	bottom := GE.GetLineOfPoints(0, h, w, h, thickness)
	col := &color.RGBA{0, 0, 0, alpha}

	left.Fill(frameImg, col)
	top.Fill(frameImg, col)
	right.Fill(frameImg, col)
	bottom.Fill(frameImg, col)
	oImg := image.Image(frameImg)
	frame.OriginalImg = &oImg
	frame.ScaleOriginal(w, h)
	frame.ScaleToOriginalSize()
	return
}
