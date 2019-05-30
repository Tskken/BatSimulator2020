package tileddecoder

import "fmt"

type TileSet struct {
	Columns          int         `json:"columns"`
	FirstGId         int         `json:"firstgid"`
	Grid             *Grid       `json:"grid"`
	Image            string      `json:"image"`
	ImageWidth       int         `json:"imagewidth"`
	ImageHeight      int         `json:"imageheight"`
	Margin           int         `json:"margin"`
	Name             string      `json:"name"`
	Properties       []*Property `json:"properties"`
	Spacing          int         `json:"spacing"`
	Terrains         []*Terrain  `json:"terrains"`
	TileCount        int         `json:"tilecount"`
	TileHeight       int         `json:"tileheight"`
	TileOffSet       *Pixel      `json:"tileoffset"`
	Tiles            []*Tile     `json:"tiles"`
	TileWidth        int         `json:"tilewidth"`
	TransparentColor string      `json:"transparentcolor"`
	Type             string      `json:"type"`
	WangSets         []*WangSet  `json:"wang_sets"`
}

func (t *TileSet) String() string {
	return fmt.Sprintf("TileSet:\n"+
		"{\n"+
		"Columns: %d\n"+
		"FirstGId: %d\n"+
		"Grid: %v\n"+
		"Image: %s\n"+
		"ImageWidth: %d\n"+
		"ImageHeight: %d\n"+
		"Margin: %d\n"+
		"Name: %s\n"+
		"Properties: %v\n"+
		"Spacing: %d\n"+
		"Terrains: %v"+
		"TileCount: %d\n"+
		"TileHeight: %d\n"+
		"TileOffSet: %v\n"+
		"Tiles: %v\n"+
		"TileWidth: %d\n"+
		"TransparentColor: %s\n"+
		"Type: %s\n"+
		"WangSets: %v\n"+
		"}\n",
		t.Columns,
		t.FirstGId,
		t.Grid,
		t.Image,
		t.ImageWidth,
		t.ImageHeight,
		t.Margin,
		t.Name,
		t.Properties,
		t.Spacing,
		t.Terrains,
		t.TileCount,
		t.TileHeight,
		t.TileOffSet,
		t.Tiles,
		t.TileWidth,
		t.TransparentColor,
		t.Type,
		t.WangSets)
}

type Tile struct {
	Animation   []*Frame    `json:"animation"`
	Id          int         `json:"id"`
	Image       string      `json:"image"`
	ImageHeight int         `json:"imageheight"`
	ImageWidth  int         `json:"imagewidth"`
	ObjectGroup *Layer      `json:"object_group"`
	Properties  []*Property `json:"properties"`
	Terrain     []int       `json:"terrain"`
	Type        string      `json:"type"`
}

func (t *Tile) String() string {
	return fmt.Sprintf("Tile:\n"+
		"{\n"+
		"Animation: %v\n"+
		"Id: %d\n"+
		"Image: %s\n"+
		"ImageHeight: %d\n"+
		"ImageWidth: %d\n"+
		"ObjectGroup: %v\n"+
		"Properties: %v\n"+
		"Terrain: %d\n"+
		"Type: %s\n"+
		"}\n",
		t.Animation,
		t.Id,
		t.Image,
		t.ImageHeight,
		t.ImageWidth,
		t.ObjectGroup,
		t.Properties,
		t.Terrain,
		t.Type)
}
