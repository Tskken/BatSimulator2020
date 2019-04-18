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

type Game struct {
	Window *pixelgl.Window
	Cfg    pixelgl.WindowConfig

	Bat       Bat
	Camera    Camera
	Animation Animation

	Objects ObjectCollection
	World World

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

	//pic, err := LoadPicture("./assets/textures/woodTexture.png")
	//if err != nil {
	//	panic(err)
	//}

	return &Game{
		Window: win,
		Cfg:    cfg,
		Bat: NewBat(win.Bounds().Center()),
		Camera: Camera{
			Position: win.Bounds().Center(),
			Bounds:   win.Bounds(),
		},
		Animation: NewAnimation(),
		Objects:map[bool][]Object{
			true: {
				//NewObject(32, 32, pixel.V(0, 0), pixel.NewSprite(pic, pixel.R(0,0,32,32))),
			},
			false: {
				//NewObject(32,32, pixel.V(32,32), pixel.NewSprite(pic, pixel.R(0,0,32,32))),
			},
		},
		World:NewWorld(),
		FPSTimer: time.Tick(time.Second),
		Last:     time.Now(),
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

		g.Objects.Draw(g.Window, Scale)

		//imd := imdraw.New(nil)
		//imd.Color = pixel.RGB(0, 1, 0)
		//imd.Push(g.Objects[CollideTrue][0][0].Sprite.Frame().Min.Scaled(WorldScale))
		//imd.Push(g.Objects[CollideTrue][0][0].Sprite.Frame().Min.Scaled(WorldScale))
		//imd.Push(pixel.V(g.Objects[CollideTrue][0][0].Sprite.Frame().Min.Scaled(WorldScale).X, g.Objects[CollideTrue][0][0].Sprite.Frame().Max.Scaled(WorldScale).Y))
		//imd.Push(pixel.V(g.Objects[CollideTrue][0][0].Sprite.Frame().Max.Scaled(WorldScale).X, g.Objects[CollideTrue][0][0].Sprite.Frame().Min.Scaled(WorldScale).Y))
		//imd.Rectangle(0)
		//
		//imd.Draw(g.Window)

		//imd := imdraw.New(nil)
		//
		//var i int
		//var j int
		//
		//for i = 0; i < int(SpriteWidthCount); i++ {
		//	if i % 2 != 0 {
		//		j = 1
		//	} else {
		//		j = 0
		//	}
		//	for;j < int(SpriteHeightCount); j+=2{
		//		imd.Color = pixel.RGBA{G:1,A:0.1}
		//		imd.Push(g.World.Objects[CollideFalse][i][j].Bounds.Min.Scaled(WorldScale))
		//		imd.Push(g.World.Objects[CollideFalse][i][j].Bounds.Min.Scaled(WorldScale))
		//		imd.Push(pixel.V(g.World.Objects[CollideFalse][i][j].Bounds.Min.Scaled(WorldScale).X, g.World.Objects[CollideFalse][i][j].Bounds.Max.Scaled(WorldScale).Y))
		//		imd.Push(pixel.V(g.World.Objects[CollideFalse][i][j].Bounds.Max.Scaled(WorldScale).X, g.World.Objects[CollideFalse][i][j].Bounds.Min.Scaled(WorldScale).Y))
		//		imd.Rectangle(0)
		//	}
		//}
		//
		//imd.Draw(g.Window)

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
