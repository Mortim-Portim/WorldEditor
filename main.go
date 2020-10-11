package main

import (
	"image/color"
	"log"

	"marvin/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1600
	screenHeight = 900
	tilewidth    = 100
	tileheight   = 100
	resourcefile = "./resource/tiles/"
)

func main() {
	GE.Init("./resource/VT323.ttf")

	tileMat := GE.GetMatrix(tilewidth, tileheight, 0)
	lightidxMat := GE.GetMatrix(tilewidth, tileheight, 255)

	wrld := GE.GetWorldStructure(20, 40, 900, 800, 18, 16)
	wrld.TileMat = tileMat
	wrld.CurrentLightMat = lightidxMat
	wrld.UpdateObjMat()
	wrld.SetMiddle(0, 0)

	rect, _ := ebiten.NewImage(16, 32, ebiten.FilterDefault)
	rect.Fill(color.Black)
	wrld.AddTile(&GE.Tile{Img: GE.CreateDayNightImg(rect, 16, 16, 1, 1, 0), Name: "black"})

	window := &window{wrld: wrld, tileMat: tileMat}

	autobutton := getAutocompleteButton(1000, 50, 60, &color.RGBA{255, 0, 0, 255}, &color.RGBA{0, 0, 255, 255}, wrld, window)
	window.buttons = append(window.buttons, autobutton)

	fillbutton := getFillButton(1200, 50, 60, &color.RGBA{255, 0, 0, 255}, &color.RGBA{0, 0, 255, 255}, window)
	window.buttons = append(window.buttons, fillbutton)

	readTileCollection(resourcefile, wrld, window)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GameEngine Test")
	if err := ebiten.RunGame(window); err != nil {
		log.Fatal(err)
	}
}
