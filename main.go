package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/mortim-portim/WorldEditor/mathwe"
)

type window struct {
	zoom   float64
	matrix *mathwe.Matrix
	images []ebiten.Image
}

const screenwidth = 960
const screenheight = 520
const width = 10
const height = 5
const pixelwidth = 16

func (w *window) Update(screen *ebiten.Image) error {
	//input
	_, dy := ebiten.Wheel()
	w.zoom += dy / 5

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		ix, iy := ebiten.CursorPosition()
		x := math.Floor(float64(ix) / w.zoom / pixelwidth)
		y := math.Floor(float64(iy) / w.zoom / pixelwidth)

		w.matrix.Set(int(x), int(y), 1, 0)
	}

	//drawing
	screen.Fill(color.NRGBA{0x00, 0xA0, 0x31, 0xff})

	for x := 0; x < width; x++ {
		if float64(x*16)*w.zoom > float64(screenwidth)/3.0*2.0 {
			continue
		}

		for y := 0; y < height; y++ {
			if float64(y*16)*w.zoom > float64(screenheight)/3.0*2.0 {
				continue
			}

			img := w.images[w.matrix.Get(x, y, 0)]
			op := ebiten.DrawImageOptions{}

			op.GeoM.Scale(w.zoom, w.zoom)
			op.GeoM.Translate(float64(x*16)*w.zoom, float64(y*16)*w.zoom)

			screen.DrawImage(&img, &op)
		}
	}
	return nil
}

func (w *window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

//Main funciton
func main() {
	path := []string{
		"./resource/16.png",
	}

	images := []ebiten.Image{}

	for p := range path {
		image, _, _ := ebitenutil.NewImageFromFile(path[p], ebiten.FilterDefault)
		images = append(images, *image)
	}

	matrix := mathwe.Matrix{X: width, Y: height, Z: 1}
	matrix.Init(0)

	window := &window{zoom: 4, matrix: &matrix, images: images}

	ebiten.SetWindowSize(screenwidth, screenheight)
	if err := ebiten.RunGame(window); err != nil {
		panic(err)
	}
}
