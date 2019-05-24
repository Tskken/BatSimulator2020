package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/png"
	"time"
)

const (
	WindowWidth  = 1024
	WindowHeight = 768
)

type Game struct {
	Window *pixelgl.Window
	Cfg    *pixelgl.WindowConfig

	Bat       *Bat
	Camera    *Camera
	Animation *Animation
	World     *World

	FPSTimer  *FPSTimer
	DeltaTime *DeltaTime
}

type FPSTimer struct {
	Frames   int
	FPSTimer <-chan time.Time
}

type DeltaTime struct {
	Last time.Time
	DT   float64
}

func NewGame() (*Game, error) {
	cfg := pixelgl.WindowConfig{
		Title:  "Bat Simulator 2020???...",
		Bounds: pixel.R(0, 0, WindowWidth, WindowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return nil, err
	}

	return &Game{
		Window: win,
		Cfg:    &cfg,
		Bat:    NewBat(win.Bounds().Center()),
		Camera: &Camera{
			Position: win.Bounds().Center(),
			Bounds:   win.Bounds(),
		},
		Animation: NewAnimation(),
		World:     LoadMap(),
		FPSTimer: &FPSTimer{
			Frames:   0,
			FPSTimer: time.Tick(time.Second),
		},
		DeltaTime: &DeltaTime{
			Last: time.Now(),
		},
	}, nil
}

func (g *Game) MainGameLoop() {
	for !g.Window.Closed() {
		g.DeltaTime.UpdateDT()

		// Handle possible actions
		g.ActionHandler()

		// Animation update counter
		if time.Since(g.Animation.AnimationTimer) >= AnimationTime || g.Bat.Sprite == nil {
			g.Bat.Sprite = g.Animation.SpriteMap[g.Animation.Action][g.Animation.Index]
			g.Animation.AnimationTimer = time.Now()
		}

		g.Draw()

		g.FPSTimer.FPSCounter(g)
	}
}

func (f *FPSTimer) FPSCounter(g *Game) {
	// Get FPS counter
	select {
	case <-f.FPSTimer:
		g.Window.SetTitle(fmt.Sprintf("%s | FPS: %d", g.Cfg.Title, f.Frames))
		f.Frames = 0
	default:
		f.Frames++
	}
}

func (g *Game) Draw() {
	g.Camera.UpdateCamera(g.Bat.HitBox.Center(), g.DeltaTime.DT)
	g.Window.SetMatrix(g.Camera.Matrix)

	imd := imdraw.New(nil)
	for _, obj := range g.World.Layers[1].Objects {
		imd.Color = pixel.RGB(0, 1, 0)
		imd.Push(obj.Rect.Min, obj.Rect.Max)
		imd.Rectangle(0)
	}

	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(g.Bat.HitBox.Min, g.Bat.HitBox.Max)
	imd.Rectangle(0)

	// Clear window
	g.Window.Clear(colornames.Black)
	g.World.Draw(g.Window)

	imd.Draw(g.Window)

	g.Bat.Draw(g.Window)

	// Update window
	g.Window.Update()
}

func (d *DeltaTime) UpdateDT() {
	d.DT = time.Since(d.Last).Seconds()
	d.Last = time.Now()
}
