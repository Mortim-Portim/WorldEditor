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

	tabview *GE.TabView

	//Tile
	useSub         bool
	tilecollection []*TileCollection

	//Object
	curretObject *GE.StructureObj
}

func (w *Window) Update() error {
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

	w.wrld.Update(w.frame)
	w.objects.Update(w.frame)
	w.tabview.Update(w.frame)
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	w.objects.Draw(screen)
	w.wrld.UpdateAllLightsIfNecassary()
	w.wrld.UpdateObjDrawables()
	w.wrld.Draw(screen)
	w.tabview.Draw(screen)
}

func (g *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func GetWindow() (window *Window) {
	window = &Window{}
	window.wrld = GetWorldStructure(20, 40, 900, 800, 100, 100, 900, 800, window)

	pathlabel := GE.GetEditText("Path", 1000, 120, 30, 25, GE.StandardFont, color.Black, color.White)
	window.pathlabel = pathlabel

	lightbar := getLightlevelScrollbar(1000, 50, 500, 30, window)
	importbutton := getImportButton(1000, 200, 30, "Import", window, pathlabel)
	exportbutton := getExportButton(1200, 200, 30, "Export", window, pathlabel)
	gridbutton := getToggleFrameButton(1400, 200, 30, window)
	window.objects = GetGroup(lightbar, pathlabel, importbutton, exportbutton, gridbutton)

	brushlabel := GE.GetTextImage("Brush:", 1000, 400, 30, GE.StandardFont, color.White, color.Transparent)
	brushscrollbar := getBrushScrollbar(1200, 400, 300, 30, window)
	tilesubscrollpanel := GE.GetScrollPanel(1000, 700, 600, 290)
	tilebtns, tilecollection := ReadTilesFromFolder(resourcefile, window.wrld, window)
	window.tilecollection = tilecollection
	tilegroup := GetGroup(tilesubscrollpanel, tilebtns, brushlabel, brushscrollbar)

	objctbuttons := ReadObjects("./resource/objects/", window)

	regionscrollpanel := GE.GetScrollPanel(1000, 630, 600, 260)
	alphalabel := GE.GetTextImage("Overlay:", 1000, 500, 30, GE.StandardFont, color.White, color.Transparent)
	regalphbar := getRegionAlphaScrollbar(1200, 500, 300, 30, window)
	colorlabel := GE.GetEditText("Col", 1200, 560, 50, 6, GE.StandardFont, color.Black, color.White)
	crtnwregionbutton := getCrtNwRegionButton(1000, 560, 50, window, colorlabel)
	regionbuttons := GetGroup(regionscrollpanel, brushlabel, brushscrollbar, alphalabel, regalphbar, colorlabel, crtnwregionbutton)

	widthlabel := GE.GetEditText("Width", 1000, 400, 30, 6, GE.StandardFont, color.White, color.White)
	heigthlabel := GE.GetEditText("Heigth", 1100, 400, 30, 6, GE.StandardFont, color.White, color.White)
	changesizebutton := getChangeSizeButton(1200, 400, 30, window, widthlabel, heigthlabel)
	mapbuttons := GetGroup(widthlabel, heigthlabel, changesizebutton)

	//Tile:0, Object:1, Region:2, Map:3
	window.tabview = GE.GetTabView(&GE.TabViewParams{X: 1000, Y: 300, W: 50, H: 30, TabH: 30, Nms: []string{"Tile", "Objct", "Region", "Map"}, Scrs: []GE.UpdateAble{tilegroup, objctbuttons, regionbuttons, mapbuttons}})

	return
}
