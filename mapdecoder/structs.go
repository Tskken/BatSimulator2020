package mapdecoder

import "fmt"

type Property struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func (p *Property) String() string {
	return fmt.Sprintf("Property:\n"+
		"{\n"+
		"Name: %s\n"+
		"Type: %s\n"+
		"Value: %v\n"+
		"}\n",
		p.Name,
		p.Type,
		p.Value)
}

type Text struct {
	Text string `json:"text"`
	Wrap bool   `json:"wrap"`
}

func (t *Text) String() string {
	return fmt.Sprintf("Text:\n"+
		"{\n"+
		"Text: %s\n"+
		"Wrap: %t\n"+
		"}\n",
		t.Text,
		t.Wrap)
}

type Pixel struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (p *Pixel) String() string {
	return fmt.Sprintf("Pixel:\n"+
		"{\n"+
		"X: %f\n"+
		"Y: %f\n"+
		"}\n",
		p.X,
		p.Y)
}
