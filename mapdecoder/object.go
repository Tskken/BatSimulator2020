package mapdecoder

import "fmt"

type Object struct {
	Ellipse    bool        `json:"ellipse"`
	GId        int         `json:"gid"`
	Height     float64     `json:"height"`
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Point      bool        `json:"point"`
	Polygon    []*Pixel    `json:"polygon"`
	Polyline   []*Pixel    `json:"polyline"`
	Properties []*Property `json:"properties"`
	Rotation   float64     `json:"rotation"`
	Template   string      `json:"template"`
	Text       *Text       `json:"text"`
	Type       string      `json:"type"`
	Visible    bool        `json:"visible"`
	Width      float64     `json:"width"`
	X          float64     `json:"x"`
	Y          float64     `json:"y"`
}

func (o *Object) String() string {
	return fmt.Sprintf("Object:\n"+
		"{\n"+
		"Ellipse: %t\n"+
		"GId: %d\n"+
		"Height: %g\n"+
		"Id: %d\n"+
		"Name: %s\n"+
		"Point: %t\n"+
		"Polygon: %v\n"+
		"Polyline: %v\n"+
		"Properties: %v\n"+
		"Rotation: %g\n"+
		"Template: %s\n"+
		"Text: %v\n"+
		"Type: %s\n"+
		"Visible: %t\n"+
		"X: %g\n"+
		"Y: %g\n"+
		"}\n",
		o.Ellipse,
		o.GId,
		o.Height,
		o.Id,
		o.Name,
		o.Point,
		o.Polygon,
		o.Polyline,
		o.Properties,
		o.Rotation,
		o.Template,
		o.Text,
		o.Type,
		o.Visible,
		o.X,
		o.Y)
}

type ObjectTemplate struct {
	Type    string   `json:"type"`
	TileSet *TileSet `json:"tileset"`
	Object  Object   `json:"object"`
}

func (o *ObjectTemplate) String() string {
	return fmt.Sprintf("ObjectTemplate:\n"+
		"{\n"+
		"	Type: %s/n"+
		"	TileSet: %v\n"+
		"	Object: %v"+
		"}\n",
		o.Type,
		o.TileSet,
		o.Object)
}
