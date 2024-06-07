package levels

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_levels"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
)

func NewLevelOne(
	levelLoader engine_levels.LevelLoader,
	levelRenderer engine_levels.LevelRenderer,
	entityManager engine_entities.EntityManager) (*levelOne, error) {
	levelOneFilePath := fmt.Sprintf("%s/%s/testLevel.json", engine_shared.AssetsPath, engine_shared.LevelPath)

	levelData, err := levelLoader.LoadLevelData(&levelOneFilePath)

	if err != nil {
		return nil, err
	}

	camera, err := entityManager.GetCamera()

	if err != nil {
		return nil, err
	}

	camera.SetLevelWidth(levelData.TileWidth * levelData.Width)

	return &levelOne{
		levelData:     levelData,
		levelRenderer: levelRenderer,
		entityManager: entityManager,
	}, nil
}

type levelOne struct {
	levelRenderer engine_levels.LevelRenderer

	entityManager engine_entities.EntityManager

	levelData *engine_levels.LevelData
}

func (l *levelOne) Render() {
	l.levelRenderer.Render(l.levelData)
}
