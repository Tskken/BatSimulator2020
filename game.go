package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/png"
	"time"
)

const (
	WindowWidth  = 1024
	WindowHeight = 768
)

type State uint8

const (
	Running State = iota
	Paused
	Stopped
)

type Action uint8

const (
	Up Action = iota
	Down
	Left
	Right
	Idle
)

type Game struct {
	Window *pixelgl.Window
	Cfg    *pixelgl.WindowConfig

	Bat    *Bat
	Camera *Camera
	World  *World

	State State

	Frames   int
	FPSTimer <-chan time.Time

	Last time.Time
	DT   time.Duration
}

func NewGame() (*Game, error) {
	//w, h := pixelgl.PrimaryMonitor().Size()

	cfg := pixelgl.WindowConfig{
		Title:  "Bat Simulator 2020???...",
		Bounds: pixel.R(0, 0, WindowWidth, WindowHeight),
		//Undecorated:true,
		Resizable:true,
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
		World:    NewWorld(),
		State:    Running,
		Frames:   0,
		FPSTimer: time.Tick(time.Second),
		Last:     time.Now(),
	}, nil
}

func (g *Game) MainGameLoop() {
	for !g.Window.Closed() {
		g.DT = time.Since(g.Last)
		g.Last = time.Now()

		g.Update()

		g.Draw()

		// Get FPS counter
		select {
		case <-g.FPSTimer:
			g.Window.SetTitle(fmt.Sprintf("%s | FPS: %d", g.Cfg.Title, g.Frames))
			g.Frames = 0
		default:
			g.Frames++
		}
	}
}

func (g *Game) HandleInputs() (pixel.Vec, Action) {
	if g.Window.JustPressed(pixelgl.KeyP) {
		if g.State == Paused {
			g.State = Running
		} else {
			g.State = Paused
		}
	}

	if g.Window.JustPressed(pixelgl.KeyEscape) {
		g.Window.SetClosed(true)
	}

	vec := pixel.ZV
	act := Idle

	if g.State != Paused {
		if g.Window.Pressed(pixelgl.KeyW) {
			v := pixel.ZV.Add(pixel.V(0, g.Bat.Speed*g.DT.Seconds()))
			if !g.Bat.CollisionCheck(v) {
				vec = vec.Add(v)
				act = Up
			}
		} else if g.Window.Pressed(pixelgl.KeyS) {
			v := pixel.ZV.Sub(pixel.V(0, g.Bat.Speed*g.DT.Seconds()))
			if !g.Bat.CollisionCheck(v) {
				vec = vec.Add(v)
				act = Down
			}
		}

		if g.Window.Pressed(pixelgl.KeyA) {
			v := pixel.ZV.Sub(pixel.V(g.Bat.Speed*g.DT.Seconds(), 0))
			if !g.Bat.CollisionCheck(v) {
				vec = vec.Add(v)
				act = Left
			}
		} else if g.Window.Pressed(pixelgl.KeyD) {
			v := pixel.ZV.Add(pixel.V(g.Bat.Speed*g.DT.Seconds(), 0))
			if !g.Bat.CollisionCheck(v) {
				vec = vec.Add(v)
				act = Right
			}
		}
	}

	return vec, act
}

func (g *Game) Update() {

	vec, act := g.HandleInputs()

	if g.State != Paused {
		g.Bat.Update(vec, act, g.DT)
	}
}

func (g *Game) Draw() {
	g.Camera.UpdateCamera(g.Bat.HitBox.Center(), g.DT.Seconds())
	g.Window.SetMatrix(g.Camera.Matrix)

	// Clear window
	g.Window.Clear(colornames.Black)

	g.World.Draw(g.Window)

	g.Bat.Draw(g.Window)

	// Update window
	g.Window.Update()
}
