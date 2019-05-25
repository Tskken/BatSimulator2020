package main

import (
	"BatSimulator2020/mapdecoder"
	"github.com/Tskken/QuadGo"
	"github.com/faiface/pixel"
	"image/color"
)

const (
	TileSize   = 16
	WorldScale = 2.5
)

type World struct {
	BackgroundColor color.RGBA
	Layers          []*Layer
	Type            string
}

func NewWorld() *World {
	m := GetDecodedMap()

	w := &World{
		Layers: make([]*Layer, len(m.Layers), len(m.Layers)),
		Type:   m.Type,
	}

	if m.BackgroundColor != "" {
		w.BackgroundColor = StringToColor(m.BackgroundColor[1:])
	}

	for i := range m.Layers {
		w.Layers[i] = NewLayer(m, i)
	}

	return w
}

type Layer struct {
	Name    string
	Sprites []*Sprite
	Type    string

	Changed bool
	Batch   *pixel.Batch
}

func NewLayer(m *mapdecoder.Map, index int) *Layer {
	l := &Layer{
		Type:    m.Layers[index].Type,
		Name:    m.Layers[index].Name,
		Changed: true,
	}

	if l.Type != "tilelayer" {
		if len(m.Layers[index].Objects) != 0 {
			for _, obj := range m.Layers[index].Objects {
				o := NewObject(obj, m)
				QGo.Insert(QuadGo.NewBounds(o.Rect.Min.X, o.Rect.Min.Y, o.Rect.Max.X, o.Rect.Max.Y), o)
			}
		}
	} else {
		sprites, spritesheet := LoadSpiteMap(m)

		if len(m.Layers[index].Data) != 0 {
			l.Sprites = GenerateMap(m.Layers[index].Data, sprites, m)
		}

		l.Batch = pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)
	}

	return l
}

func (w *World) Draw(target pixel.Target) {
	for _, layer := range w.Layers {
		if layer.Type == "tilelayer" {
			if layer.Sprites != nil {

				if layer.Changed {
					layer.Batch.Clear()

					for _, sprite := range layer.Sprites {
						if sprite.Visible {
							sprite.Sprite.Draw(layer.Batch, sprite.Matrix)
						}
					}

					layer.Changed = false
				}

				layer.Batch.Draw(target)
			}
		}
	}
}
