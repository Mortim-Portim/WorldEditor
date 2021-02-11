package wrldedit

import (
	"image/png"
	"io/ioutil"
	"os"
	"strconv"

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

	window.wrld.SetMiddle(int(Compression.BytesToInt64(bs[0])), int(Compression.BytesToInt64(bs[1])), false)
	window.wrld.SetLightStats(Compression.BytesToInt16(bs[2]), Compression.BytesToInt16(bs[3]), Compression.BytesToFloat64(bs[4]))
}

func ExportWorld(input string, window *Window) {
	if input == "" {
		input = "AutoSave"
	}

	folder := "./resource/maps/" + input + "/"
	os.Mkdir(folder, 0755)
	os.Mkdir(folder+"tile/", 0755)
	window.wrld.Save(folder + "map.txt")

	index, _ := os.Create(folder + "/tile/#index.txt")

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
}
