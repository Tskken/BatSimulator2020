package main

import (
	"github.com/faiface/pixel"
	"time"
)

const (
	AnimationTime = time.Second / 8
)

type Animation struct {
	SpriteMap map[Action][]*pixel.Sprite
	Index     int
	Action    Action

	AnimationTimer <-chan time.Time
}

func (a *Animation) Update(action Action) {
	if a.Action == action {
		a.Index++
		if a.Index >= len(a.SpriteMap[a.Action]) {
			a.Index = 0
		}
	} else {
		a.Action = action
	}
}
