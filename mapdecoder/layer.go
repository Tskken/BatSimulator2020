package mapdecoder

import "fmt"

type Layer struct {
	Chunks           []*Chunk    `json:"chunks"`
	Compression      string      `json:"compression"`
	Data             []uint      `json:"data"`
	DrawOrder        string      `json:"draworder"`
	Encoding         string      `json:"encoding"`
	Height           int         `json:"height"`
	Id               int         `json:"id"`
	Image            string      `json:"image"`
	Layers           []*Layer    `json:"layers"`
	Name             string      `json:"name"`
	Objects          []*Object   `json:"objects"`
	OffSetx          float64     `json:"offsetx"`
	OffSety          float64     `json:"offsety"`
	Opacity          float64     `json:"opacity"`
	Properties       []*Property `json:"properties"`
	TransparentColor string      `json:"transparentcolor"`
	Type             string      `json:"type"`
	Visible          bool        `json:"visible"`
	Width            int         `json:"width"`
	X                int         `json:"x"`
	Y                int         `json:"y"`
}

func (l *Layer) String() string {
	return fmt.Sprintf("Layer:\n"+
		"{\n"+
		"Chunks: %v\n"+
		"Compression %s\n"+
		"Data: %d\n"+
		"DrawOrder: %s\n"+
		"Encoding: %s\n"+
		"Height: %d\n"+
		"Id: %d\n"+
		"Image: %v\n"+
		"Layers: %v\n"+
		"Name: %s\n"+
		"Objects: %v\n"+
		"Offsetx: %g\n"+
		"Offsety: %g\n"+
		"Opacity: %g\n"+
		"Properties: %v\n"+
		"TransparentColor: %s\n"+
		"Type: %s\n"+
		"Visible: %t\n"+
		"Width: %d\n"+
		"X: %d\n"+
		"Y: %d\n"+
		"}\n",
		l.Chunks,
		l.Compression,
		l.Data,
		l.DrawOrder,
		l.Encoding,
		l.Height,
		l.Id,
		l.Image,
		l.Layers,
		l.Name,
		l.Objects,
		l.OffSetx,
		l.OffSety,
		l.Opacity,
		l.Properties,
		l.TransparentColor,
		l.Type,
		l.Visible,
		l.Width,
		l.X,
		l.Y)
}
