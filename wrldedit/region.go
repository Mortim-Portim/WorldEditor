package wrldedit

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

func GetRegion(color color.Color) *Region {
	img, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	img.Fill(color)
	imgobj := &GE.ImageObj{Img: img}
	return &Region{imgobj}
}

type Region struct {
	color *GE.ImageObj
}

func (r *Region) Draw(screen *ebiten.Image, drawer *GE.ImageObj, alpha float64) {
	drawer.CopyXYWHTo(r.color)
	r.color.DrawImageObjAlpha(screen, alpha)
}
