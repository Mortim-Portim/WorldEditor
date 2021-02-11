package wrldedit

import (
	"image/color"

	"github.com/mortim-portim/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	ScreenWidth  = 1600
	ScreenHeight = 900
	tilewidth    = 100
	tileheight   = 100
	resourcefile = "./resource/tiles/"
)

type Window struct {
	wrld      *GE.WorldStructure
	objects   *Group
	pathlabel *GE.EditText

	frame, curType int

	//Tile
	useSub         bool
	selectedVar    int
	brushsize      int
	tilecollection []*TileCollection
	tilebuttons    *Group
	tilesubbuttons *Group
	importedTiles  []string

	//Object
	currentStructure *GE.Structure
	curretObject     *GE.StructureObj
	objectbuttons    *Group

	//Light
	lightbuttons *Group
}

func (w *Window) Update(screen *ebiten.Image) error {
	w.frame++

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mousebuttonleftPressed(w)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mousebuttonleftJustPressed(w)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		mousebuttonrightPressed(w)
	}

	keyPressed(w)

	_, y := ebiten.Wheel()

	if y < 0 {
		w.wrld.SetDisplayWH(w.wrld.TileMat.W()-1, w.wrld.TileMat.H()-1)
	}

	if y > 0 {
		w.wrld.SetDisplayWH(w.wrld.TileMat.W()-3, w.wrld.TileMat.H()-3)
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

	if g.frame%1800 == 0 {
		ExportWorld(g.pathlabel.GetText(), g)
	}
}

func (g *Window) draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0xA0, 0x00, 0xff})
	g.objects.Draw(screen)
	g.wrld.UpdateAllLightsIfNecassary()
	g.wrld.UpdateObjDrawables()
	g.wrld.Draw(screen)

	switch g.curType {
	case 0:
		g.tilebuttons.Draw(screen)
		g.tilesubbuttons.Draw(screen)
	case 1:
		g.objectbuttons.Draw(screen)
	}
}

func (g *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func GetWindow(wrld *GE.WorldStructure) (window *Window) {
	window = &Window{wrld: wrld, objects: GetGroup(), tilebuttons: GetGroup(), tilesubbuttons: GetGroup(), objectbuttons: GetGroup()}

	pathlabel := getPathLabel(1000, 120, 50, 25)
	window.pathlabel = pathlabel

	lightbar := getLightlevelScrollbar(1000, 50, 500, 30, window)
	importbutton := getImportButton(1000, 200, 50, "Import", window, pathlabel)
	exportbutton := getExportButton(1200, 200, 50, "Export", window, pathlabel)
	tilebutton := getTabButton(1000, 300, 50, 0, "Tile", window)
	objbutton := getTabButton(1200, 300, 50, 1, "Objects", window)
	lightbutton := getTabButton(1400, 300, 50, 2, "Light", window)
	window.objects.Add(lightbar, pathlabel, importbutton, exportbutton, tilebutton, objbutton, lightbutton)

	//autobutton := getAutocompleteButton(1000, 400, 50, window)
	brushscrollbar := getBrushScrollbar(1000, 400, 300, 30, window)
	window.tilebuttons.Add(brushscrollbar)

	ReadTilesFromFolder(resourcefile, wrld, window)
	readObjects("./resource/objects/", window)

	return
}
