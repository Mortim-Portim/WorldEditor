package main

import (
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

	wrld := GE.GetWorldStructure(20, 40, 900, 800, 100, 100, 900, 800)
	wrld.SetMiddle(0, 0, true)
	wrld.SetDisplayWH(18, 16)
	wrld.SetLightStats(0, 255, 0)

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
