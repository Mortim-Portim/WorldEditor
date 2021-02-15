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
	data, err1 := ioutil.ReadFile("./resource/maps/" + input + "/map.txt")
	if err1 != nil {
		return
	}

	bs := Compression.DecompressAll(data, []int{8, 8, 2, 2, 8})
	tilMat := GE.GetMatrix(0, 0, 0)
	err2 := tilMat.Decompress(bs[5])
	if err2 != nil {
		return
	}

	window.wrld.TileMat = tilMat

	window.wrld.Objects = nil
	window.wrld.BytesToObjects(bs[6])
	window.wrld.UpdateObjMat()

	window.wrld.Lights = nil
	window.wrld.BytesToLights(bs[7])
	window.wrld.UpdateLIdxMat()

	regMat := GE.GetMatrix(0, 0, 0)
	if len(bs) >= 9 {
		regMat.Decompress(bs[8])
	}
	window.wrld.RegionMat = regMat

	window.wrld.SetMiddle(int(Compression.BytesToInt64(bs[0])), int(Compression.BytesToInt64(bs[1])), false)
	window.wrld.SetLightStats(Compression.BytesToInt16(bs[2]), Compression.BytesToInt16(bs[3]), Compression.BytesToFloat64(bs[4]))

	//Region
	window.regselectbutton = GetGroup()
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
	window.regselectbutton.Add(regbtn)
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
		img, _ := ebiten.NewImage(16, 32, ebiten.FilterDefault)

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
