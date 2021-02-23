package wrldedit

import (
	"bufio"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/Compression"
	"github.com/mortim-portim/GraphEng/GE"
)

func ImportWorld(input string, window *Window) {
	data, err := ioutil.ReadFile("./resource/maps/" + input + "/map.txt")
	if err != nil {
		return
	}

	bs := Compression.DecompressAll(data, []int{8, 8, 2, 15})
	tilMat := GE.GetMatrix(0, 0, 0)
	err = tilMat.Decompress(bs[5])
	if err != nil {
		return
	}
	window.wrld.TileMat = tilMat

	window.wrld.BytesToObjects(bs[6])
	window.wrld.BytesToLights(bs[7])

	regMat := GE.GetMatrix(0, 0, 0)
	if len(bs) >= 9 {
		regMat.Decompress(bs[8])
	}
	window.wrld.RegionMat = regMat

	window.wrld.SetMiddle(int(Compression.BytesToInt64(bs[0])), int(Compression.BytesToInt64(bs[1])), false)

	//Region
	window.wrld.Region = make([]*Region, 0)
	region, _ := os.Open("./resource/maps/" + input + "/region.txt")
	scanner := bufio.NewScanner(region)

	for scanner.Scan() {
		line := scanner.Text()
		rgb := strings.Split(line, "-")

		r, _ := strconv.Atoi(rgb[0])
		g, _ := strconv.Atoi(rgb[1])
		b, _ := strconv.Atoi(rgb[2])

		region := GetRegion(color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		AddRegion(region, window)
	}
}

func AddRegion(region *Region, window *Window) {
	i := len(window.wrld.Region)
	regbtn := GE.GetImageButton(region.color.Img, float64(1000+(i%8)*70), 630+(math.Ceil(float64(i/8)))*70, 64, 64)
	regbtn.Data = i + 1
	regbtn.RegisterOnLeftEvent(func(btn *GE.Button) {
		window.selectedVar = btn.Data.(int)
	})
	window.wrld.Region = append(window.wrld.Region, region)
	window.tabview.Screens.Members[2].(*Group).Members[0].(*GE.ScrollPanel).Add(regbtn)
}

func ExportWorld(input string, window *Window) {
	if input == "" {
		input = "AutoSave"
	}

	folder := "./resource/maps/" + input + "/"
	os.Mkdir(folder, 0755)
	os.Mkdir(folder+"tile/", 0755)
	window.wrld.Save(folder + "map.txt")

	//Tile
	index, _ := os.Create(folder + "tile/#index.txt")

	light := &GE.ImageObj{X: 0, Y: 0, W: 16, H: 16}
	dark := &GE.ImageObj{X: 0, Y: 16, W: 16, H: 16}
	for i, tile := range window.wrld.Tiles {
		name := strconv.Itoa(i) + ".png"
		file, _ := os.Create(folder + "/tile/" + name)
		img := ebiten.NewImage(16, 32)

		tile.Draw(img, light, 255)
		tile.Draw(img, dark, 0)

		png.Encode(file, img)

		index.WriteString(name + "\n")
		file.Close()
	}

	//Region
	regionfile, _ := os.Create(folder + "region.txt")

	for _, region := range window.wrld.Region {
		r, g, b, _ := region.color.Img.At(0, 0).RGBA()
		r >>= 8
		g >>= 8
		b >>= 8

		regionfile.WriteString(strconv.Itoa(int(r)) + "-" + strconv.Itoa(int(g)) + "-" + strconv.Itoa(int(b)) + "\n")
	}
}
