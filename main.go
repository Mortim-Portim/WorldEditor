package main

import (
	"log"

	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/WorldEditor/wrldedit"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	GE.Init("./resource/VT323.ttf")

	wrld := GE.GetWorldStructure(20, 40, 900, 800, 100, 100, 900, 800)
	wrld.SetMiddle(0, 0, true)
	wrld.SetDisplayWH(18, 16)
	wrld.SetLightStats(0, 255, 0)

	window := wrldedit.GetWindow(wrld)

	ebiten.SetWindowSize(wrldedit.ScreenWidth, wrldedit.ScreenHeight)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("WorldEditor")
	ebiten.SetMaxTPS(30)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetRunnableOnUnfocused(true)
	if err := ebiten.RunGame(window); err != nil {
		log.Fatal(err)
	}
}
