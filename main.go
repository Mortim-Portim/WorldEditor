package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"marvin/GraphEng/GE"
	"marvin/GraphEng/GE/WObjs"

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
	wrld             *GE.WorldPainter
	idxMat, layerMat *GE.Matrix

	frame      int
	imgButtons []*GE.Button
}

func (g *window) Update(screen *ebiten.Image) error {
	screen.Fill(color.RGBA{0x00, 0xA0, 0x00, 0xff})
	g.frame++

	g.scrollbar.Update(g.frame)
	g.label.Update(g.frame)

	g.btn.Draw(screen)
	g.scrollbar.Draw(screen)
	g.label.Draw(screen)

	for _, bt := range g.imgButtons {
		bt.Draw(screen)
	}

	g.wrld.Paint(screen, g.idxMat, g.layerMat, 0)
	return nil
}

func (g *window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	GE.Init("./resource/VT323.ttf")

	mat := &GE.Matrix{X: 3, Y: 3, Z: 3}
	mat.InitIdx()

	btn := GE.GetTextButton("Edit", "Edasdit", GE.StandardFont, 1000, 50, 60, &color.RGBA{255, 0, 0, 255}, &color.RGBA{0, 0, 255, 255})

	scrollbar := GE.GetStandardScrollbar(1000, 200, 300, 60, 0, 3, 0, GE.StandardFont)
	label := GE.GetEditText("Path", 1000, 400, 60, 20, GE.StandardFont, color.RGBA{255, 120, 20, 255}, GE.EditText_Selected_Col)

	fmt.Println(label.ImageObj.Img)

	wmatI := &GE.Matrix{X: 10, Y: 9, Z: 1}
	wmatI.Init(0)
	wmatL := &GE.Matrix{X: 10, Y: 9, Z: 1}
	wmatL.Init(0)

	wrld := GE.GetWorldPainter(0, 50, 500, 500, wmatI.X, wmatI.Y)
	wrld.GetFrame(2, 90)

	rect, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	rect.Fill(color.Black)
	wrld.AddTile(&WObjs.Tile{rect})

	w := &window{btn: btn, scrollbar: scrollbar, label: label, wrld: wrld, idxMat: wmatI, layerMat: wmatL}

	label.RegisterOnChange(func(t *GE.EditText) {
		img := GE.LoadEbitenImg(label.GetText())
		wrld.AddTile(&WObjs.Tile{img})
		button := GE.GetImageButton(img, float64(1000+(len(w.imgButtons)%5)*64), 500+(math.Ceil(float64(len(w.imgButtons)/5)))*64, 64, 64)
		w.imgButtons = append(w.imgButtons, button)
	})

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GameEngine Test")
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}
}
