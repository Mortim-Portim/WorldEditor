package wrldedit

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"

	"github.com/mortim-portim/GraphEng/GE"
)

//Bug: Deleting Tile Folder
//Bug: Changing Tile Folder

func ReadTilesFromFolder(path string, ws *GE.WorldStructure, window *Window) {
	files, err := ioutil.ReadDir(path)
	Check(err, "Resource filepath is false")

	indexs, index := scanfile(path + "index.txt")

	for _, mis := range detectMissing(files, indexs) {
		index.WriteString(mis + "\n")
		indexs = append(indexs, mis)
	}

	tilebutton := make([]UpdateAble, 0)
	lastnum := 0

	for i, str := range indexs {
		folder, _ := ioutil.ReadDir(path + str + "/")
		tiles, cltindex := scanfile(path + str + "/index.txt")
		tiledata := make([]InputParam, len(tiles))

		for y, n := range tiles {
			tiledata[y] = ReadTileInfo(n)
		}

		for _, mistxt := range detectMissingIP(folder, tiledata) {
			cltindex.WriteString(mistxt + "\n")
			tiles = append(tiles, mistxt)
		}

		subbtn := make([]UpdateAble, 0)
		index := make(map[uint8]map[string][]int64)

		for k, tile := range tiledata {
			registerDirection(tile.GetStringElse("Direction", ""), tile.GetStringElse("Tile", "Default"), int64(k+lastnum), index)

			for m := 0; true; k++ {
				direction, avab := tile.GetString("Direction" + strconv.Itoa(m))

				if !avab {
					break
				}

				registerDirection(direction, tile.GetStringElse("Tile"+strconv.Itoa(k), "Default"), int64(k+lastnum), index)
			}

			rotation := tile.GetFloat64Else("Rotation", 0)
			name, _ := tile.GetString("Name")

			fmt.Println(name)
			img, err := GE.LoadDayNightImg(path+str+"/"+name, 0, 0, 0, 0, rotation)
			Check(err, "Hi")
			img.ScaleToOriginalSize()
			ws.AddTile(&GE.Tile{img, strconv.Itoa(i)})

			button := GE.GetImageButton(img.GetDay(), float64(1000+(k%8)*70), 700+(math.Ceil(float64(k/8)))*70, 64, 64)
			button.Img.Angle = rotation
			button.Data = k + lastnum
			button.RegisterOnLeftEvent(func(b *GE.Button) {
				if !b.LPressed {
					return
				}

				window.selectedVar = b.Data.(int)
				window.useSub = true
			})

			subbtn = append(subbtn, button)
		}

		mbutton := GE.GetImageButton(subbtn[0].(*GE.Button).Img.Img, float64(1000+(i%8)*70), 500+(math.Ceil(float64(i/8)))*70, 64, 64)
		mbutton.Data = i
		mbutton.RegisterOnLeftEvent(func(b *GE.Button) {
			if !b.LPressed {
				return
			}
			window.selectedVar = b.Data.(int)
			window.useSub = false
			window.tilesubbuttons = window.tilecollection[b.Data.(int)].GetSubButtons()
		})

		tilebutton = append(tilebutton, mbutton)

		var tc *TileCollection

		subgroup := &Group{subbtn}

		tc = &TileCollection{str, lastnum, len(tiles), subgroup, index}

		window.tilecollection = append(window.tilecollection, tc)

		lastnum += len(tiles)
	}

	window.tilebuttons.Add(tilebutton...)
}

func Check(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func scanfile(path string) (line []string, file *os.File) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 666)
	Check(err, "Failed creating file: "+path)
	scanner := bufio.NewScanner(file)
	line = make([]string, 0)

	for scanner.Scan() {
		line = append(line, scanner.Text())
	}

	return
}

func detectMissing(directory []os.FileInfo, line []string) (missing []string) {
	missing = make([]string, 0)
	for _, fi := range directory {
		name := fi.Name()
		exist := false

		if fi.Name() == "index.txt" || fi.Name() == "info.txt" {
			continue
		}

		for _, nam := range line {
			if name == nam {
				exist = true
				break
			}
		}

		if !exist {
			missing = append(missing, name)
		}
	}

	return
}

func detectMissingIP(directory []os.FileInfo, ip []InputParam) (missing []string) {
	line := make([]string, len(ip))

	for k, param := range ip {
		line[k], _ = param.GetString("Name")
	}

	missing = detectMissing(directory, line)
	return
}

func registerDirection(direction, cnttile string, i int64, index map[uint8]map[string][]int64) {
	spldirection := []rune(direction)
	num := uint8(0)

	for _, r := range spldirection {
		switch r {
		case 'U':
			num++
		case 'L':
			num += 2
		case 'D':
			num += 4
		case 'R':
			num += 8
		case 'Q':
			num += 16
		case 'K':
			num += 32
		case 'N':
			num += 64
		case 'M':
			num += 128
		case 'F':
			num += 240
		}
	}

	if index[num] == nil {
		index[num] = make(map[string][]int64)
	}

	if index[num][cnttile] == nil {
		index[num][cnttile] = make([]int64, 0)
	}

	index[num][cnttile] = append(index[num][cnttile], i)
}

func readObjects(path string, window *Window) {
	objects, _ := GE.ReadStructures(path)

	for i, object := range objects {
		btnImg := object.NUA.GetDay()
		button := GE.GetImageButton(btnImg, float64(1000+(i%8)*70), 500+(math.Ceil(float64(i/8)))*70, 64, 64)

		button.Data = object
		button.RegisterOnLeftEvent(func(btn *GE.Button) {
			window.currentStructure = button.Data.(*GE.Structure)
		})

		window.objectbuttons.Add(button)

		window.wrld.AddStruct(object)
	}
}
