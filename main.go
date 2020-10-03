package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/mortim-portim/WorldEditor/math"
)

type Window struct {
	matrix *math.Matrix
	images []ebiten.Image
}

func (w *Window) Update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	return nil
}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

//Main funciton
func main() {
	window := &Window{}

	path := []string{"./resource/16.png"}

	for p := range path {
		image, _, _ := ebitenutil.NewImageFromFile(path[p], ebiten.FilterDefault)
		window.images = append(window.images, *image)
	}

	if err := ebiten.RunGame(window); err != nil {
		panic(err)
	}
}
