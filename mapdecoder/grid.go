package mapdecoder

import "fmt"

type Grid struct {
	Orientation string `json:"orientation"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid\n"+
		"{\n"+
		"Orientation: %s\n"+
		"Width: %d\n"+
		"Height: %d\n"+
		"}\n",
		g.Orientation,
		g.Width,
		g.Height)
}
