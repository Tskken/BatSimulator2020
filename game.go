package main

import (
	"BatSimulator2020/mapdecoder"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/png"
	"path/filepath"
	"time"
)

const (
	WindowWidth  = 1024
	WindowHeight = 768
)

type Game struct {
	Window *pixelgl.Window
	Cfg    pixelgl.WindowConfig

	Bat       Bat
	Camera    Camera
	Animation Animation
	World     World

	Frames   int
	FPSTimer <-chan time.Time

	Last time.Time
	DT   float64
}

/*
	TODO: Config setup
		- Add load config from file support
		- Streamline code and comment
*/
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

	mapPath, err := filepath.Abs("./assets/maps/cave_map_v1.json")
	if err != nil {
		panic(err)
	}

	rootPath, err := filepath.Abs("../BatSimulator2020/assets/tilesets")
	if err != nil {
		panic(err)
	}

	m, err := mapdecoder.LoadMap(rootPath, mapPath)
	if err != nil {
		panic(err)
	}

	w, err := LoadMap(m)
	if err != nil {
		panic(err)
	}

	return &Game{
		Window: win,
		Cfg:    cfg,
		Bat:    NewBat(win.Bounds().Center()),
		Camera: Camera{
			Position: win.Bounds().Center(),
			Bounds:   win.Bounds(),
		},
		Animation: NewAnimation(),
		World:     *w,
		FPSTimer:  time.Tick(time.Second),
		Last:      time.Now(),
	}, nil
}

func (g *Game) MainGameLoop() {
	for !g.Window.Closed() {
		g.UpdateDT()

		// Handle possible actions
		g.ActionHandler()

		// Animation update counter
		select {
		case <-g.Animation.AnimationTimer:
			g.Bat.Sprite = g.Animation.SpriteMap[g.Animation.Action][g.Animation.Index]
		default:
			if g.Bat.Sprite == nil {
				g.Bat.Sprite = g.Animation.SpriteMap[Idle][0]
			}
		}

		// Clear window
		g.Window.Clear(colornames.Gray)

		g.Camera.UpdateCamera(g.Bat.Position, g.DT)

		g.Window.SetMatrix(g.Camera.Matrix)

		g.World.Draw(g.Window)

		g.Bat.Draw(g.Window)

		// Update window
		g.Window.Update()

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

func (g *Game) UpdateDT() {
	g.DT = time.Since(g.Last).Seconds()
	g.Last = time.Now()
}
