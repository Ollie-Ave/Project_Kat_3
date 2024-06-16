package engine_levels

import (
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
)

func NewLevelCollider(levelData *LevelData) engine_entities.EntityUpdater {
	layerCollisionData := make(map[string][][]bool)

	for _, layer := range levelData.Layers {

		if layer.LayerType == engine_shared.TileLayer {
			layerCollisionData[layer.Name] = parseCollisionDataForLayer(layer, levelData)
		}
	}

	return &levelColliderImpl{
		layerCollisionData: layerCollisionData,
		tileWidth:          levelData.TileWidth,
		tileHeight:         levelData.TileHeight,
	}
}

func parseCollisionDataForLayer(layer *Layer, levelData *LevelData) [][]bool {
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

type levelColliderImpl struct {
	layerCollisionData map[string][][]bool

	tileWidth  int
	tileHeight int
}

func (l *levelColliderImpl) Update() error {
	return nil
}

func (l *levelColliderImpl) GetLayerCollisionData(layerName string) ([][]bool, int, int) {
	return l.layerCollisionData[layerName], l.tileWidth, l.tileHeight
}
