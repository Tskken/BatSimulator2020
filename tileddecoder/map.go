package tileddecoder

import (
	"encoding/json"
	"fmt"
)

type Map struct {
	BackgroundColor string       `json:"backgroundcolor"`
	Height          int          `json:"height"`
	HexSideLength   int          `json:"hexsidelength"`
	Infinite        bool         `json:"infinite"`
	Layers          []*Layer     `json:"layers"`
	NextLayerId     int          `json:"nextlayerid"`
	NextObjectId    int          `json:"nextobjectid"`
	Orientation     string       `json:"orientation"`
	Properties      []*Property  `json:"properties"`
	RenderOrder     string       `json:"renderorder"`
	StaggerAxis     string       `json:"staggeraxis"`
	StaggerIndex    string       `json:"staggerindex"`
	TiledVersion    string       `json:"tiledversion"`
	TileHeight      int          `json:"tileheight"`
	TileSets        []*TileSet   `json:"tilesets"`
	TileWidth       int          `json:"tilewidth"`
	Type            string       `json:"type"`
	Version         *json.Number `json:"version"`
	Width           int          `json:"width"`
}

func (m *Map) String() string {
	return fmt.Sprintf("Map:\n"+
		"{\n"+
		"BackgroundColor: %s\n"+
		"Height %d\n"+
		"HexSideLength: %d\n"+
		"Infinite: %t\n"+
		"Layers: %v\n"+
		"NextLayerId: %d\n"+
		"NextObjectId: %d\n"+
		"Orientation: %s\n"+
		"Properties: %v\n"+
		"RenderOrder: %s\n"+
		"StaggerAxis: %s\n"+
		"StaggerIndex: %s\n"+
		"TiledVersion: %s\n"+
		"TileHeight: %d\n"+
		"TileSets: %v\n"+
		"TileWidth: %d\n"+
		"Type: %s\n"+
		"Version: %v\n"+
		"Width: %d\n"+
		"}\n",
		m.BackgroundColor,
		m.Height,
		m.HexSideLength,
		m.Infinite,
		m.Layers,
		m.NextLayerId,
		m.NextObjectId,
		m.Orientation,
		m.Properties,
		m.RenderOrder,
		m.StaggerAxis,
		m.StaggerIndex,
		m.TiledVersion,
		m.TileHeight,
		m.TileSets,
		m.TileWidth,
		m.Type,
		m.Version,
		m.Width)
}
