package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello World"); err != nil {
		panic(err)
	}
}

var square *ebiten.Image

func update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DebugPrint(screen, "Hello World")

	if square == nil {
		square, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)
	}

	square.Fill(color.White)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(64, 64)
	screen.DrawImage(square, opts)

	return nil
}
