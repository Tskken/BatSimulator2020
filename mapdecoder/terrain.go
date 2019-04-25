package mapdecoder

import "fmt"

type Terrain struct {
	Name string `json:"name"`
	Tile int    `json:"tile"`
}

func (t *Terrain) String() string {
	return fmt.Sprintf("Terrain:\n"+
		"{\n"+
		"Name: %s\n"+
		"Tile: %d\n"+
		"}\n",
		t.Name,
		t.Tile)
}
