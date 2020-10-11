package main

import (
	"fmt"
	"io/ioutil"
	"marvin/GraphEng/GE"
	"math"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

func readTileCollection(path string, wrld *GE.WorldStructure, window *window) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("Resource filepath is false")
		return
	}

	for i, file := range files {
		imgs, _ := GE.ReadTiles(resourcefile + file.Name() + "/")

		var tilecollection TileCollection
		var lastnum int

		if len(window.tilecollection) == 0 {
			lastnum = 1
		} else {
			lastnum = window.tilecollection[len(window.tilecollection)-1].GetLast()
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

		window.tilecollection = append(window.tilecollection, tilecollection)

		for _, img := range imgs {
			wrld.AddTile(img)
		}

		btimg, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
		imgs[0].Img.Draw(btimg, 1)

		button := GE.GetImageButton(btimg, float64(1000+(i%9)*64), 500+(math.Ceil(float64(i/9)))*64, 64, 64)
		button.Data = i
		button.RegisterOnLeftEvent(func(b *GE.Button) {
			if !b.LPressed {
				return
			}
			window.curImg = b.Data.(int)
			window.useSub = false
			setSubBtn(window, wrld, window.tilecollection[window.curImg])
		})
		window.buttons = append(window.buttons, button)
	}
}

func setSubBtn(window *window, wrld *GE.WorldStructure, tc TileCollection) {
	window.subButtons = nil

	for i := tc.GetStart(); i < tc.GetLast(); i++ {
		btimg := wrld.Tiles[i].Img.GetDay()
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
