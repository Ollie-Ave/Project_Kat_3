package levels

import (
	"encoding/json"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	assetsPath  = "/home/oliver/Code/Project_Kat_3/assets"
	levelPath   = "levels/export"
	tileSrcPath = "tileSets/src"
)

func initLevelOne() (*levelOne, error) {
	levelOneFilePath := fmt.Sprintf("%s/%s/testLevel.json", assetsPath, levelPath)

	levelData, err := loadLevelData(&levelOneFilePath)

	if err != nil {
		return nil, err
	}

	return &levelOne{
		levelData:     levelData,
		levelRenderer: newLevelRenderer(),
	}, nil
}

func loadLevelData(filePath *string) (*levelData, error) {
	var levelData *levelData

	levelFile, err := os.ReadFile(*filePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(levelFile, &levelData)

	if err != nil {
		return nil, err
	}

	for _, tileSet := range levelData.TileSets {
		tileTexurePath := fmt.Sprintf("%s/%s/%s", assetsPath, tileSrcPath, tileSet.Image)

		_, err := os.Stat(tileTexurePath)

		if err != nil {
			return nil, err
		}

		tileSet.Texture = rl.LoadTexture(tileTexurePath)
	}

	for _, layer := range levelData.Layers {
		imagePath := fmt.Sprintf("%s/%s/%s", assetsPath, levelPath, layer.ImagePath)

		_, err := os.Stat(imagePath)

		if err != nil {
			return nil, err
		}

		layer.ImageTexture = rl.LoadTexture(imagePath)
	}

	return levelData, nil
}

type levelOne struct {
	levelRenderer levelRenderer

	levelData *levelData
}

func (l *levelOne) Render() {
	l.levelRenderer.Render(l.levelData)
}
