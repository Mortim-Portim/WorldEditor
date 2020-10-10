package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"math"
	"strings"

	"marvin/GraphEng/GE"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1600
	screenHeight = 900
	tilewidth    = 10
	tileheight   = 10
	resourcefile = "./resource/tiles/"
)

type window struct {
	btn       *GE.Button
	scrollbar *GE.ScrollBar
	label     *GE.EditText
	pictures  *GE.Button
	wrld      *GE.WorldStructure
	tileMat   *GE.Matrix

	frame, curImg  int
	useSub         bool
	tilecollection []TileCollection
	imgButtons     []*GE.Button
	subButtons     []*GE.Button
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
			if g.useSub {
				g.tileMat.Set(x, y, int16(g.curImg))
			} else {
				g.tileMat.Set(x, y, int16(g.tilecollection[g.curImg].GetNum()))
			}
		}
	}

	g.btn.Update(g.frame)
	g.scrollbar.Update(g.frame)
	g.label.Update(g.frame)

	for _, bt := range g.imgButtons {
		bt.Update(g.frame)
	}

	for _, bt := range g.subButtons {
		bt.Update(g.frame)
	}

	g.btn.Draw(screen)
	g.scrollbar.Draw(screen)
	g.label.Draw(screen)

	for _, bt := range g.imgButtons {
		bt.Draw(screen)
	}

	for _, bt := range g.subButtons {
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

	btn := GE.GetTextButton("Autocomplete", "Edasdit", GE.StandardFont, 1000, 50, 60, &color.RGBA{255, 0, 0, 255}, &color.RGBA{0, 0, 255, 255})
	scrollbar := GE.GetStandardScrollbar(1000, 200, 300, 60, 0, 3, 0, GE.StandardFont)
	label := GE.GetEditText("Path", 1000, 400, 60, 25, GE.StandardFont, color.RGBA{255, 120, 20, 255}, GE.EditText_Selected_Col)

	tileMat := GE.GetMatrix(tilewidth, tileheight, 0)
	lightMat := GE.GetMatrix(tilewidth, tileheight, 255)
	objMat := GE.GetMatrix(0, 0, 0)

	wrld := GE.GetWorldStructure(0, 50, 500, 500, tileMat.W(), tileMat.H())
	wrld.TileMat = tileMat
	wrld.LightMat = lightMat
	wrld.ObjMat = objMat
	wrld.GetFrame(2, 90)

	rect, _ := ebiten.NewImage(16, 32, ebiten.FilterDefault)
	rect.Fill(color.Black)
	wrld.AddTile(&GE.Tile{GE.CreateDayNightImg(rect, 16, 16, 1, 1, 0), "black"})

	world := &window{btn: btn, scrollbar: scrollbar, label: label, wrld: wrld, tileMat: tileMat}

	btn.RegisterOnLeftEvent(func(btn *GE.Button) {
		if btn.LPressed == false {
			return
		}

		for x := 0; x < tilewidth; x++ {
			for y := 0; y < tilewidth; y++ {
				tilename := wrld.Tiles[wrld.TileMat.GetAbs(x, y)].Name
				var tilecollection TileCollection

				for _, antilecollection := range world.tilecollection {
					if antilecollection.GetString() == tilename {
						tilecollection = antilecollection
						break
					}
				}

				if tilecollection == nil {
					continue
				}

				switch tc := tilecollection.(type) {
				case *ConnectedTC:
					n := getOnPos(x, y-1, tc.start, tc.rang, world.tileMat)
					w := getOnPos(x-1, y, tc.start, tc.rang, world.tileMat)
					s := getOnPos(x, y+1, tc.start, tc.rang, world.tileMat)
					e := getOnPos(x+1, y, tc.start, tc.rang, world.tileMat)

					tileMat.Set(x, y, int16(tc.GetIndex(n, w, s, e)))
				}
			}
		}
	})

	files, err := ioutil.ReadDir(resourcefile)

	if err != nil {
		fmt.Println("Resource filepath is false")
		return
	}

	for _, file := range files {
		imgs, _ := GE.ReadTiles(resourcefile + file.Name() + "/")

		var tilecollection TileCollection
		var lastnum int

		if len(world.tilecollection) == 0 {
			lastnum = 1
		} else {
			lastnum = world.tilecollection[len(world.tilecollection)-1].GetLast()
		}

		switch strings.Split(file.Name(), "-")[1] {
		case "r":
			tilecollection = &RandomTC{file.Name(), lastnum, len(imgs)}
		case "c":
			tilecollection = &ConnectedTC{file.Name(), lastnum, len(imgs)}
		default:
			fmt.Println(file.Name() + " does not have a correct ending")
			continue
		}

		for _, tile := range imgs {
			tile.Name = file.Name()
		}

		world.tilecollection = append(world.tilecollection, tilecollection)

		for _, img := range imgs {
			wrld.AddTile(img)
		}

		btimg, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
		imgs[0].Img.Draw(btimg, 1)

		button := GE.GetImageButton(btimg, float64(1000+(len(world.imgButtons)%9)*64), 500+(math.Ceil(float64(len(world.imgButtons)/9)))*64, 64, 64)
		button.Data = len(world.imgButtons)
		button.RegisterOnLeftEvent(func(b *GE.Button) {
			if !b.LPressed {
				return
			}
			world.curImg = b.Data.(int)
			world.useSub = false
			setSubBtn(world, wrld, world.tilecollection[world.curImg])
		})
		world.imgButtons = append(world.imgButtons, button)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("GameEngine Test")
	if err := ebiten.RunGame(world); err != nil {
		log.Fatal(err)
	}
}

func setSubBtn(window *window, wrld *GE.WorldStructure, tc TileCollection) {
	window.subButtons = nil

	for i := tc.GetStart(); i < tc.GetLast(); i++ {
		btimg, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
		wrld.Tiles[i].Img.Draw(btimg, 1)

		//btimg.DrawImage(wrld.Tiles[i].Img.Day.Img, &ebiten.DrawImageOptions{})

		button := GE.GetImageButton(btimg, float64(1000+(len(window.subButtons)%9)*64), 700+(math.Ceil(float64(len(window.subButtons)/9)))*64, 64, 64)
		button.Data = i
		button.RegisterOnLeftEvent(func(b *GE.Button) {
			if !b.LPressed {
				return
			}

			window.curImg = b.Data.(int)
			window.useSub = true
		})
		window.subButtons = append(window.subButtons, button)
	}
}

func getOnPos(x, y, s, r int, tilemat *GE.Matrix) int {
	if x < 0 || x >= tilewidth || y < 0 || y >= tileheight {
		return 0
	}

	if tilemat.Get(x, y) >= int16(s) && tilemat.Get(x, y) < int16(s+r) {
		return 0
	}

	return 1
}
