package core

import "github.com/faiface/pixel"

type Sprite struct {
	Sprite  *pixel.Sprite
	Matrix  pixel.Matrix
	Visible bool
}

func (s *Sprite) Update(sprite *pixel.Sprite, vec pixel.Vec) {
	if sprite != nil {
		s.Sprite = sprite
	}
	s.Matrix = s.Matrix.Moved(vec)
}

func (s *Sprite) Draw(target pixel.Target) {
	s.Sprite.Draw(target, s.Matrix)
}
