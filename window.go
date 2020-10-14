package main

import (
	"image/color"
	"marvin/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Window struct {
	wrld    *GE.WorldStructure
	objects *GE.Group

	frame, curType int

	//Tile
	useSub         bool
	selectedVar    int
	tilecollection []TileCollection
	tilebuttons    *GE.Group
	tilesubbuttons *GE.Group

	//Object
	currentObject *GE.Structure
	objectbuttons *GE.Group
}

func (w *Window) Update(screen *ebiten.Image) error {
	w.frame++

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mousebuttonleftPressed(w)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mousebuttonleftJustPressed(w)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		mousebuttonrightJustPressed(w)
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		w.wrld.Move(-1, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		w.wrld.Move(1, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		w.wrld.Move(0, -1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		w.wrld.Move(0, 1)
	}

	_, y := ebiten.Wheel()

	if y < 0 {
		w.wrld.SetDisplayWH(w.wrld.TileMat.W()+1, w.wrld.TileMat.H()+1)
	}

	if y > 0 {
		w.wrld.SetDisplayWH(w.wrld.TileMat.W()-1, w.wrld.TileMat.H()-1)
	}

	w.update()
	w.draw(screen)

	return nil
}

func (g *Window) update() {
	g.objects.Update(g.frame)

	switch g.curType {
	case 0:
		g.tilebuttons.Update(g.frame)
		g.tilesubbuttons.Update(g.frame)
	case 1:
		g.objectbuttons.Update(g.frame)
	}
}

func (g *Window) draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0xA0, 0x00, 0xff})
	g.objects.Draw(screen)
	g.wrld.DrawLights(false)
	g.wrld.DrawBack(screen)

	switch g.curType {
	case 0:
		g.tilebuttons.Draw(screen)
		g.tilesubbuttons.Draw(screen)
	case 1:
		g.objectbuttons.Draw(screen)
	}
}

func (g *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func getWindow(wrld *GE.WorldStructure) (window *Window) {
	window = &Window{wrld: wrld, objects: GE.GetGroup(), tilebuttons: GE.GetGroup(), tilesubbuttons: GE.GetGroup(), objectbuttons: GE.GetGroup()}

	lightbar := getLightlevelScrollbar(1000, 50, 500, 30, window)
	tilebutton := getTabButton(1000, 300, 50, 0, "Tile", window)
	objbutton := getTabButton(1200, 300, 50, 1, "Objects", window)
	window.objects.Add(lightbar, tilebutton, objbutton)

	autobutton := getAutocompleteButton(1000, 400, 50, window)
	fillbutton := getFillButton(1300, 400, 50, window)
	window.tilebuttons.Add(autobutton, fillbutton)
	return
}
