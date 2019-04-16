package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	_ "image/png"
	"os"
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

	Frames   int
	FPSTimer <-chan time.Time

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

	path, err := filepath.Abs("./assets/sprites/bat-sprite.png")
	if err != nil {
		panic(err)
	}

	// Load sprite sheet from file
	spritesheet, err := LoadPicture(path)
	if err != nil {
		panic(err)
	}

	var sprites []*pixel.Sprite

	// Save sprites from sprite sheet to array
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			sprites = append(sprites, pixel.NewSprite(spritesheet, pixel.R(x, y, x+32, y+32)))
		}
	}

	return &Game{
		Window: win,
		Cfg:    cfg,
		Bat: Bat{
			Sprite:   sprites[0],
			HitBox:   pixel.R(0, 0, 32, 32),
			Position: win.Bounds().Center(),
			Speed:    500.0,
		},
		Camera: Camera{
			Position: win.Bounds().Center(),
			Bounds:   win.Bounds(),
			Speed:    500.0,
		},
		Animation: Animation{
			SpriteMap: map[Action][]*pixel.Sprite{
				Up: {
					sprites[5],
					sprites[9],
					sprites[13],
				},
				Down: {
					sprites[7],
					sprites[11],
					sprites[15],
				},
				Left: {
					sprites[4],
					sprites[8],
					sprites[12],
				},
				Right: {
					sprites[6],
					sprites[10],
					sprites[14],
				},
				Idle: {
					sprites[7],
					sprites[11],
					sprites[15],
				},
			},
			Action:         Idle,
			AnimationTimer: time.Tick(AnimationTime),
		},
		FPSTimer: time.Tick(time.Second),
		Last:     time.Now(),
	}, nil
}

func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
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
		}

		// Set camera position
		g.Camera.Matrix = pixel.IM.Moved(g.Window.Bounds().Center().Sub(g.Camera.Position))
		g.Window.SetMatrix(g.Camera.Matrix)

		// Clear window
		g.Window.Clear(colornames.Skyblue)

		// Draw bat on screen
		g.Bat.Sprite.Draw(g.Window, pixel.IM.Scaled(pixel.ZV, 4).Moved(g.Bat.Position))

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
