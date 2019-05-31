package core

import (
	"BatSimulator2020/tileddecoder"
	"encoding/hex"
	"errors"
	"github.com/faiface/pixel"
	"image"
	"image/color"
	"math"
	"os"
	"path/filepath"
)

const (
	FlippedHorizontallyFlag uint = 0x80000000
	FlippedVerticallyFlag   uint = 0x40000000
	FlippedDiagonallyFlag   uint = 0x20000000
)

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func StringToColor(c string) color.RGBA {
	b, err := hex.DecodeString(c)
	if err != nil {
		panic(err)
	}

	switch len(b) {
	case 4:
		return color.RGBA{
			R: b[1],
			G: b[2],
			B: b[3],
			A: 255,
		}
	case 5:
		return color.RGBA{
			R: b[1],
			G: b[2],
			B: b[3],
			A: b[0],
		}
	default:
		panic(errors.New("improper format of hex color in NewWorld()"))
		return color.RGBA{}
	}
}

func TileRotation(gid uint) (uint, pixel.Matrix) {
	matrix := pixel.IM

	if gid&FlippedDiagonallyFlag == FlippedDiagonallyFlag {
		matrix = matrix.Rotated(pixel.ZV, 90*math.Pi/180).ScaledXY(pixel.ZV, pixel.V(1, -1))
	}

	if gid&FlippedHorizontallyFlag == FlippedHorizontallyFlag {
		matrix = matrix.ScaledXY(pixel.ZV, pixel.V(-1, 1))
	}

	if gid&FlippedVerticallyFlag == FlippedVerticallyFlag {
		matrix = matrix.ScaledXY(pixel.ZV, pixel.V(1, -1))
	}

	// Clear the flags
	gid &= ^(FlippedHorizontallyFlag |
		FlippedVerticallyFlag |
		FlippedDiagonallyFlag)

	return gid, matrix
}

func LoadSpiteMap(m *tileddecoder.Map) ([]*pixel.Sprite, pixel.Picture) {
	spritesheet, err := LoadPicture(m.TileSets[0].Image)
	if err != nil {
		panic(err)
	}

	var sprites []*pixel.Sprite

	// Save sprites from sprite sheet to array
	for y := spritesheet.Bounds().Max.Y - TileSize; y >= spritesheet.Bounds().Min.Y; y -= TileSize {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += TileSize {
			sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(x, y, x+TileSize, y+TileSize)))
		}
	}

	return sprites, spritesheet
}

func GetDecodedMap() *tileddecoder.Map {
	mapConfigPath, err := filepath.Abs("./assets/maps/cave_map_v1.json")
	if err != nil {
		panic(err)
	}

	tileSetRootPath, err := filepath.Abs("../BatSimulator2020/assets/tilesets")
	if err != nil {
		panic(err)
	}

	m, err := tileddecoder.LoadMap(tileSetRootPath, mapConfigPath)
	if err != nil {
		panic(err)
	}

	return m
}

func GenerateMap(tileIds []uint, tileSprites []*pixel.Sprite, m *tileddecoder.Map) (mapSprites []*Sprite) {
	mapSprites = make([]*Sprite, 0, m.Height*m.Width)

	var tileIndex uint = 0

	for y := m.Height - 1; y >= 0; y-- {
		for x := 0; x < m.Width; x++ {
			gid, matrix := TileRotation(tileIds[tileIndex])

			// Resolve the tile
			for i := len(m.TileSets) - 1; i >= 0; i-- {
				tileset := m.TileSets[i]
				if uint(tileset.FirstGId) <= gid {
					mapSprites = append(mapSprites, &Sprite{
						Sprite:  tileSprites[gid-uint(m.TileSets[i].FirstGId)],
						Matrix:  matrix.Scaled(pixel.ZV, GlobalConfig.WorldScale).Moved(pixel.V(float64(x)*TileSize*GlobalConfig.WorldScale, float64(y)*TileSize*GlobalConfig.WorldScale)),
						Visible: true,
					})
					break
				}
			}
			tileIndex++
		}
	}
	return
}
