package main

import (
	"github.com/dhconnelly/rtreego"
	"github.com/faiface/pixel"
	"image"
	"os"
)

var (
	RectZV = pixel.R(0,0,0,0)
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

func ToRect(rct pixel.Rect) *rtreego.Rect {
	rect, err := rtreego.NewRect(rtreego.Point{rct.Min.X, rct.Min.Y}, []float64{rct.W(), rct.H()})
	if err != nil {
		panic(err)
	}

	return rect
}
