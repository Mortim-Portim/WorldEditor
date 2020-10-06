package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"marvin/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1600
	screenHeight = 900
)

type window struct {
	btn              *GE.Button
	scrollbar        *GE.ScrollBar
	label            *GE.EditText
	pictures         *GE.Button
	wrld             *GE.WorldStructure
	idxMat, layerMat *GE.Matrix

	frame      int
	curImg     int
	imgButtons []*GE.Button
}

func (g *window) Update(screen *ebiten.Image) error {
	screen.Fill(color.RGBA{0x00, 0xA0, 0x00, 0xff})
	g.frame++

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		dx, dy := ebiten.CursorPosition()

		x := int16(math.Floor((float64(dx) - g.wrld.GetTopLeft().X) / 50.0))
		y := int16(math.Floor((float64(dy) - g.wrld.GetTopLeft().Y) / 50.0))

		if x >= 0 && x < g.wrld.IdxMat.W() && y >= 0 && y < g.wrld.IdxMat.H() {
			g.idxMat.Set(int16(x), int16(y), 1)
			g.layerMat.Set(x, y, int16(g.scrollbar.Current()))
		}
	}

	g.scrollbar.Update(g.frame)
	g.label.Update(g.frame)

	g.btn.Draw(screen)
	g.scrollbar.Draw(screen)
	g.label.Draw(screen)

	for _, bt := range g.imgButtons {
		bt.Draw(screen)
	}

	g.wrld.IdxMat = g.idxMat
	g.wrld.LayerMat = g.layerMat
	g.wrld.Draw(screen)
	return nil
}

func (g *window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	GE.Init("./resource/VT323.ttf")

	btn := GE.GetTextButton("Edit", "Edasdit", GE.StandardFont, 1000, 50, 60, &color.RGBA{255, 0, 0, 255}, &color.RGBA{0, 0, 255, 255})

	scrollbar := GE.GetStandardScrollbar(1000, 200, 300, 60, 0, 3, 0, GE.StandardFont)
	label := GE.GetEditText("Path", 1000, 400, 60, 25, GE.StandardFont, color.RGBA{255, 120, 20, 255}, GE.EditText_Selected_Col)

	wmatI := GE.GetMatrix(10, 10, 0)
	wmatL := GE.GetMatrix(10, 10, 0)

	wrld := GE.GetWorldStructure(0, 50, 500, 500, wmatI.W(), wmatI.H())
	wrld.GetFrame(2, 90)

	rect, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	rect.Fill(color.Black)
	wrld.AddTile(&GE.Tile{rect})

	w := &window{btn: btn, scrollbar: scrollbar, label: label, wrld: wrld, idxMat: wmatI, layerMat: wmatL}

	label.RegisterOnChange(func(t *GE.EditText) {
		img, _ := GE.LoadEbitenImg(label.GetText())
		wrld.AddTile(&GE.Tile{img})
		button := GE.GetImageButton(img, float64(1000+(len(w.imgButtons)%5)*64), 500+(math.Ceil(float64(len(w.imgButtons)/5)))*64, 64, 64)
		button.RegisterOnLeftEvent(func(b *GE.Button) {
			w.curImg = len(w.imgButtons)
			fmt.Println(len(w.imgButtons))
		})
		w.imgButtons = append(w.imgButtons, button)
	})

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GameEngine Test")
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}
}
