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

	levelCollider := newLevelCollider(levelData)
	entityManager.SetLevelCollider(levelCollider)

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

func newLevelCollider(levelData *engine_levels.LevelData) engine_entities.EntityUpdatable {
	layerCollisionData := make(map[string][][]bool)

	for _, layer := range levelData.Layers {

		if layer.LayerType == engine_shared.TileLayer {
			layerCollisionData[layer.Name] = parseCollisionDataForLayer(layer, levelData)
		}
	}

	return &levelCollider{
		layerCollisionData: layerCollisionData,
	}
}

func parseCollisionDataForLayer(layer *engine_levels.Layer, levelData *engine_levels.LevelData) [][]bool {
	collisionData := assign2DArrayBuffer[bool](levelData.Width, levelData.Height)

	for y := 0; y < levelData.Height; y++ {
		for x := 0; x < levelData.Width; x++ {
			layerDataIndex := y*levelData.Width + x

			collisionData[y][x] = layer.Data[layerDataIndex] != 0
		}
	}

	return collisionData
}

func assign2DArrayBuffer[T any](rows, cols int) [][]T {
	buffer := make([][]T, cols)

	for i := range buffer {
		buffer[i] = make([]T, rows)
	}

	return buffer
}

type levelCollider struct {
	layerCollisionData map[string][][]bool
}

func (l *levelCollider) Update() {
}
