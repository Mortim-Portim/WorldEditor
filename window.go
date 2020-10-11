package main

import (
	"image/color"
	"marvin/GraphEng/GE"
	"math"

	"github.com/hajimehoshi/ebiten"
)

type window struct {
	wrld    *GE.WorldStructure
	tileMat *GE.Matrix
	buttons []*GE.Button

	frame, curImg  int
	useSub         bool
	tilecollection []TileCollection
	subButtons     []*GE.Button
}

func (g *window) Update(screen *ebiten.Image) error {
	g.frame++

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		dx, dy := ebiten.CursorPosition()
		wx, wy := g.wrld.GetTopLeft()
		x := int(math.Floor((float64(dx) - wx) / (g.wrld.W / g.wrld.TileMat.Focus().Bounds().X)))
		y := int(math.Floor((float64(dy) - wy) / (g.wrld.H / g.wrld.TileMat.Focus().Bounds().Y)))

		if x >= 0 && x < g.tileMat.W() && y >= 0 && y < g.tileMat.H() {
			if g.useSub {
				g.tileMat.Set(x, y, int16(g.curImg))
			} else {
				g.tileMat.Set(x, y, g.tilecollection[g.curImg].GetNum())
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.wrld.Move(-1, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.wrld.Move(1, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.wrld.Move(0, -1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.wrld.Move(0, 1)
	}

	g.wrld.TileMat = g.tileMat

	g.update()
	g.draw(screen)

	return nil
}

func (g *window) update() {
	for _, bt := range g.buttons {
		bt.Update(g.frame)
	}

	for _, bt := range g.subButtons {
		bt.Update(g.frame)
	}
}

func (g *window) draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0xA0, 0x00, 0xff})

	for _, bt := range g.buttons {
		bt.Draw(screen)
	}

	for _, bt := range g.subButtons {
		bt.Draw(screen)
	}

	g.wrld.DrawBack(screen)
	//g.wrld.DrawFront(screen)
}

func (g *window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
