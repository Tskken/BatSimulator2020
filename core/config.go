package core

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	ConfigRelPath = "./core/config.json"
)

var (
	GlobalConfig = &Config{}

	DefaultConfig = &Config{
		WindowConfig: &WindowConfig{
			Title: "Bat Simulator 2020??? Kappa...",
			WindowWidth:  1026,
			WindowHeight: 768,
			VSync: true,
			Undecorated: false,
			Resizable: false,
		},
		BatConfig: &BatConfig{
			BatScale:         2,
			BatSpeed:         500,
			BatAnimationRate: time.Second / 7,
		},
		WorldConfig: &WorldConfig{
			WorldScale: 2.5,
		},
	}
)

type Config struct {
	*WindowConfig `json:"window_config"`
	*BatConfig    `json:"bat_config"`
	*WorldConfig  `json:"world_config"`
}

type WindowConfig struct {
	Title string `json:"title"`
	WindowWidth  float64 `json:"window_width"`
	WindowHeight float64 `json:"window_height"`
	VSync bool `json:"v_sync"`
	Undecorated bool `json:"undecorated"`
	Resizable bool `json:"resizable"`
}

type BatConfig struct {
	BatScale         float64       `json:"bat_scale"`
	BatSpeed         float64       `json:"bat_speed"`
	BatAnimationRate time.Duration `json:"bat_animation_rate"`
}

type WorldConfig struct {
	WorldScale float64 `json:"world_scale"`
}

func LoadConfigs() {
	path, err := filepath.Abs(ConfigRelPath)
	if err != nil {
		log.Println(err)
		GlobalConfig = DefaultConfig
		return
	}

	file, err := os.Open(path)
	if err != nil {
		GlobalConfig = DefaultConfig
		return
	}

	dec := json.NewDecoder(file)

	err = dec.Decode(GlobalConfig)
	if err != nil {
		log.Println(err)
		GlobalConfig = DefaultConfig
	}
}

func SaveToConfig() {
	path, err := filepath.Abs(ConfigRelPath)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(file)

	err = enc.Encode(GlobalConfig)
	if err != nil {
		panic(err)
	}
}
