package tileddecoder

import "fmt"

type Chunk struct {
	Data   []uint `json:"data"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

func (c *Chunk) String() string {
	return fmt.Sprintf("Chunk:\n"+
		"{\n"+
		"Data: %d\n"+
		"Height: %d\n"+
		"Width: %d\n"+
		"X: %d\n"+
		"Y: %d\n"+
		"}\n",
		c.Data,
		c.Height,
		c.Width,
		c.X,
		c.Y)
}
