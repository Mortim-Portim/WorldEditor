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
	wrld      *WorldStructure
	objects   *Group
	pathlabel *GE.EditText

	frame, curType int
	selectedVar    int
	brushsize      int

	//Tile
	useSub         bool
	tilecollection []*TileCollection
	tilebuttons    *Group
	tilesubbuttons *Group
	importedTiles  []string

	//Object
	currentStructure *GE.Structure
	curretObject     *GE.StructureObj
	objectbuttons    *GE.ScrollPanel

	//Region
	regionbuttons   *Group
	regselectbutton *Group
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

	w.update()
	w.draw(screen)

	return nil
}

func (w *Window) update() {
	w.wrld.Update(w.frame)
	w.objects.Update(w.frame)

	switch w.curType {
	case 0:
		w.tilebuttons.Update(w.frame)
		w.tilesubbuttons.Update(w.frame)
	case 1:
		w.objectbuttons.Update(w.frame)
	case 2:
		w.regionbuttons.Update(w.frame)
		w.regselectbutton.Update(w.frame)
	}

	if w.frame%1800 == 0 {
		ExportWorld(w.pathlabel.GetText(), w)
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
	case 2:
		g.regionbuttons.Draw(screen)
		g.regselectbutton.Draw(screen)
	}
}

func (g *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func GetWindow(wrld *WorldStructure) (window *Window) {
	window = &Window{wrld: wrld, objects: GetGroup(), tilebuttons: GetGroup(), tilesubbuttons: GetGroup(), regionbuttons: GetGroup(), regselectbutton: GetGroup()}

	pathlabel := GE.GetEditText("Path", 1000, 120, 50, 25, GE.StandardFont, color.Black, color.White)
	window.pathlabel = pathlabel

	lightbar := getLightlevelScrollbar(1000, 50, 500, 30, window)
	importbutton := getImportButton(1000, 200, 50, "Import", window, pathlabel)
	exportbutton := getExportButton(1200, 200, 50, "Export", window, pathlabel)
	tilebutton := getTabButton(1000, 300, 50, 0, "Tile", window)
	objbutton := getTabButton(1200, 300, 50, 1, "Objects", window)
	lightbutton := getTabButton(1400, 300, 50, 2, "Region", window)
	window.objects.Add(lightbar, pathlabel, importbutton, exportbutton, tilebutton, objbutton, lightbutton)

	brushlabel := GE.GetTextImage("Brush:", 1000, 400, 30, GE.StandardFont, color.Black, color.Transparent)
	brushscrollbar := getBrushScrollbar(1200, 400, 300, 30, window)
	window.tilebuttons.Add(brushlabel, brushscrollbar)

	alphalabel := GE.GetTextImage("Overlay:", 1000, 500, 30, GE.StandardFont, color.Black, color.Transparent)
	regalphbar := getRegionAlphaScrollbar(1200, 500, 300, 30, window)
	colorlabel := GE.GetEditText("Col", 1200, 560, 50, 6, GE.StandardFont, color.Black, color.White)
	crtnwregionbutton := getCrtNwRegionButton(1000, 560, 50, window, colorlabel)
	window.regionbuttons.Add(brushlabel, brushscrollbar, alphalabel, regalphbar, colorlabel, crtnwregionbutton)

	ReadTilesFromFolder(resourcefile, wrld, window)
	ReadObjects("./resource/objects/", window)

	return
}
