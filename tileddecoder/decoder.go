package tileddecoder

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func LoadMap(rootPath, mapPath string) (*Map, error) {
	mapFile, err := os.Open(mapPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = mapFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	dec := json.NewDecoder(mapFile)

	m := new(Map)

	err = dec.Decode(m)
	if err != nil {
		return nil, err
	}

	for _, tileSet := range m.TileSets {

		tileSet.Image = filepath.Join(rootPath, tileSet.Image)
	}

	return m, nil
}
