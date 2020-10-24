package main

import (
	"fmt"
	"io/ioutil"
	"marvin/GraphEng/GE"
	"math"
	"strings"
)

func readTileCollection(path string, window *Window) {
	files, err := ioutil.ReadDir(path)
	window.importedTiles = make([]string, 0)

	if err != nil {
		fmt.Println("Resource filepath is false")
		return
	}

	for i, file := range files {
		subimg, _ := ioutil.ReadDir(resourcefile + file.Name() + "/")

		for _, name := range subimg {
			split := strings.Split(name.Name(), ".")
			if split[1] == "png" {
				window.importedTiles = append(window.importedTiles, split[0])
			}
		}

		tiles, _ := GE.ReadTiles(resourcefile + file.Name() + "/")
		subbuttons := GE.GetGroup()

		var lastnum int

		if len(window.tilecollection) == 0 {
			lastnum = 0
		} else {
			lastnum = window.tilecollection[len(window.tilecollection)-1].GetLast()
		}

		for k, tile := range tiles {
			tile.Name = file.Name()
			window.wrld.AddTile(tile)

			btimg := tile.Img.GetDay()
			button := GE.GetImageButton(btimg, float64(1000+(k%8)*70), 700+(math.Ceil(float64(k/8)))*70, 64, 64)
			button.Data = k + lastnum
			button.RegisterOnLeftEvent(func(b *GE.Button) {
				if !b.LPressed {
					return
				}

				window.selectedVar = b.Data.(int)
				window.useSub = true

				fmt.Printf("%v", window.selectedVar)
			})

			subbuttons.Add(button)
		}

		var tilecollection TileCollection

		switch strings.Split(file.Name(), "-")[1] {
		case "r":
			tilecollection = &RandomTC{DefaultTC{file.Name(), lastnum, len(tiles), subbuttons}}
		case "c":
			tilecollection = &ConnectedTC{DefaultTC{file.Name(), lastnum, len(tiles), subbuttons}}
		default:
			fmt.Println(file.Name() + " does not have a correct ending")
			continue
		}

		window.tilecollection = append(window.tilecollection, tilecollection)

		btimg := tiles[0].Img.GetDay()

		button := GE.GetImageButton(btimg, float64(1000+(i%8)*70), 500+(math.Ceil(float64(i/8)))*70, 64, 64)
		button.Data = i
		button.RegisterOnLeftEvent(func(b *GE.Button) {
			if !b.LPressed {
				return
			}
			window.selectedVar = b.Data.(int)
			window.useSub = false
			window.tilesubbuttons = window.tilecollection[b.Data.(int)].GetSubButtons()
		})

		window.tilebuttons.Add(button)
	}
}

func readObjects(path string, window *Window) {
	objects, _ := GE.ReadStructures(path)

	for i, object := range objects {
		btnImg := object.NUA.GetDay()
		button := GE.GetImageButton(btnImg, float64(1000+(i%8)*70), 500+(math.Ceil(float64(i/8)))*70, 64, 64)

		button.Data = object
		button.RegisterOnLeftEvent(func(btn *GE.Button) {
			window.currentObject = button.Data.(*GE.Structure)
		})

		window.objectbuttons.Add(button)

		window.wrld.AddStruct(object)
	}
}
