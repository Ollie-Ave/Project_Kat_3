package engine_levels

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewLevelLoader() LevelLoader {
	return &levelLoaderImpl{}
}

type LevelLoader interface {
	LoadLevelData(filePath *string) (*LevelData, error)
}

type levelLoaderImpl struct {
}

func (l *levelLoaderImpl) LoadLevelData(filePath *string) (*LevelData, error) {

	var levelData *LevelData

	levelFile, err := os.ReadFile(*filePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(levelFile, &levelData)

	if err != nil {
		return nil, err
	}

	for _, tileSet := range levelData.TileSets {
		tileTexurePath := fmt.Sprintf("%s/%s/%s", engine_shared.AssetsPath, engine_shared.TileSrcPath, tileSet.Image)

		_, err := os.Stat(tileTexurePath)

		if err != nil {
			return nil, err
		}

		tileSet.Texture = rl.LoadTexture(tileTexurePath)
	}

	for _, layer := range levelData.Layers {
		imagePath := fmt.Sprintf("%s/%s/%s", engine_shared.AssetsPath, engine_shared.LevelPath, layer.ImagePath)

		_, err := os.Stat(imagePath)

		if err != nil {
			return nil, err
		}

		layer.ImageTexture = rl.LoadTexture(imagePath)
	}

	return levelData, nil

}
