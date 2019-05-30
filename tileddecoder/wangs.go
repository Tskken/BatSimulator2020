package tileddecoder

import "fmt"

type WangSet struct {
	CornerColors []*WangColor `json:"cornercolors"`
	EdgeColors   []*WangColor `json:"edgecolors"`
	Name         string       `json:"name"`
	Tile         int          `json:"tile"`
	WangTiles    []*WangTile  `json:"wangtiles"`
}

func (w *WangSet) String() string {
	return fmt.Sprintf("WangSet\n"+
		"{\n"+
		"CornerColors: %v\n"+
		"EdgeColors: %v\n"+
		"Name: %s\n"+
		"Tile: %d\n"+
		"WangTiles: %v\n"+
		"}\n",
		w.CornerColors,
		w.EdgeColors,
		w.Name,
		w.Tile,
		w.WangTiles)
}

type WangColor struct {
	Color       string  `json:"color"`
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
	Tile        int     `json:"tile"`
}

func (w *WangColor) String() string {
	return fmt.Sprintf("WangColor:\n"+
		"{\n"+
		"Color: %s\n"+
		"Name: %s\n"+
		"Probability: %g\n"+
		"Tile: %d\n"+
		"}\n",
		w.Color,
		w.Name,
		w.Probability,
		w.Tile)
}

type WangTile struct {
	DFlip  bool    `json:"dflip"`
	HFlip  bool    `json:"hflip"`
	TileId int     `json:"tileid"`
	VFlip  bool    `json:"vflip"`
	WangId []uint8 `json:"wangid"`
}

func (w *WangTile) String() string {
	return fmt.Sprintf("WangTile:\n"+
		"{\n"+
		"DFlip: %t\n"+
		"HFlip: %t\n"+
		"TileId: %d\n"+
		"VFlip: %t\n"+
		"WangId: %d\n"+
		"}\n",
		w.DFlip,
		w.HFlip,
		w.TileId,
		w.VFlip,
		w.WangId)
}
