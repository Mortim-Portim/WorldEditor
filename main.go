package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"marvin/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/WorldEditor/util"
)

const (
	screenWidth  = 1600
	screenHeight = 900
)

type window struct {
	btn       *GE.Button
	scrollbar *GE.ScrollBar
	label     *GE.EditText
	pictures  *GE.Button
	wrld      *GE.WorldStructure
	tileMat   *GE.Matrix

	frame, curImg  int
	tilecollection []*TileCollection
	imgButtons     []*GE.Button
}

func (g *window) Update(screen *ebiten.Image) error {
	screen.Fill(color.RGBA{0x00, 0xA0, 0x00, 0xff})
	g.frame++

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		dx, dy := ebiten.CursorPosition()
		wx, wy := g.wrld.GetTopLeft()
		x := int(math.Floor((float64(dx) - wx) / 50.0))
		y := int(math.Floor((float64(dy) - wy) / 50.0))

		if x >= 0 && x < g.tileMat.W() && y >= 0 && y < g.tileMat.H() {
			g.tileMat.Set(x, y, int16(g.tilecollection[g.curImg].randNum()))
		}
	}

	g.btn.Update(g.frame)
	g.scrollbar.Update(g.frame)
	g.label.Update(g.frame)

	for _, bt := range g.imgButtons {
		bt.Update(g.frame)
	}

	g.btn.Draw(screen)
	g.scrollbar.Draw(screen)
	g.label.Draw(screen)

	for _, bt := range g.imgButtons {
		bt.Draw(screen)
	}

	g.wrld.TileMat = g.tileMat
	g.wrld.DrawBack(screen)
	g.wrld.DrawFront(screen)
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

	tileMat := GE.GetMatrix(10, 10, 0)
	lightMat := GE.GetMatrix(10, 10, 255)
	objMat := GE.GetMatrix(0, 0, 0)

	wrld := GE.GetWorldStructure(0, 50, 500, 500, tileMat.W(), tileMat.H())
	wrld.TileMat = tileMat
	wrld.LightMat = lightMat
	wrld.ObjMat = objMat
	wrld.GetFrame(2, 90)

	rect, _ := ebiten.NewImage(16, 32, ebiten.FilterDefault)
	rect.Fill(color.Black)
	wrld.AddTile(&GE.Tile{GE.CreateDayNightImg(rect, 16, 16, 1, 1, 0), "black"})

	w := &window{btn: btn, scrollbar: scrollbar, label: label, wrld: wrld, tileMat: tileMat}

	label.RegisterOnChange(func(t *GE.EditText) {
		imgs, _ := GE.ReadTiles(label.GetText())

		if len(w.tilecollection) == 0 {
			w.tilecollection = append(w.tilecollection, &TileCollection{1, len(imgs)})
		} else {
			w.tilecollection = append(w.tilecollection, &TileCollection{w.tilecollection[len(w.tilecollection)].start + w.tilecollection[len(w.tilecollection)].rang, len(imgs)})
		}

		for _, img := range imgs {
			wrld.AddTile(img)
		}

		btimg, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
		imgs[0].Img.Draw(btimg, 1)

		button := GE.GetImageButton(btimg, float64(1000+(len(w.imgButtons)%5)*64), 500+(math.Ceil(float64(len(w.imgButtons)/5)))*64, 64, 64)
		button.RegisterOnLeftEvent(func(b *GE.Button) {
			w.curImg = util.IndexOf(w.imgButtons, b)
			fmt.Print(w.curImg)
		})
		w.imgButtons = append(w.imgButtons, button)
	})

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GameEngine Test")
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}
}
