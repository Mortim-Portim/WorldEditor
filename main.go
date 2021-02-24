package main

import (
	"log"

	"github.com/mortim-portim/GraphEng/GE"
	"github.com/mortim-portim/WorldEditor/wrldedit"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	GE.Init("./resource/VT323.ttf", 60)
	window := wrldedit.GetWindow()

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
