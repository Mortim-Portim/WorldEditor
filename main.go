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

	wrld := GE.GetWorldStructure(20, 40, 900, 800, 100, 100)
	wrld.SetMiddle(0, 0)
	wrld.SetDisplayWH(18, 16)
	wrld.SetLightStats(0, 255, 0)

	rect, _ := ebiten.NewImage(16, 32, ebiten.FilterDefault)
	rect.Fill(color.Black)
	wrld.AddTile(&GE.Tile{Img: GE.CreateDayNightImg(rect, 16, 16, 1, 1, 0), Name: "black"})

	window := getWindow(wrld)

	readTileCollection(resourcefile, window)
	readObjects("./resource/objects/", window)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("WorldEditor")
	if err := ebiten.RunGame(window); err != nil {
		log.Fatal(err)
	}
}
