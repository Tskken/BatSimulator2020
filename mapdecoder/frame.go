package mapdecoder

import "fmt"

type Frame struct {
	Duration int `json:"duration"`
	TileId   int `json:"tileid"`
}

func (f *Frame) String() string {
	return fmt.Sprintf("Frame:\n"+
		"{\n"+
		"Duration: %d\n"+
		"TileId: %d\n"+
		"}\n",
		f.Duration,
		f.TileId)
}
