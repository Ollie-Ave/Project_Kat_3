package engine_levels

import (
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
)

func NewLevelCollider(levelData *LevelData) engine_entities.EntityUpdatable {
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

type levelCollider struct {
	layerCollisionData map[string][][]bool
}

func (l *levelCollider) Update() {
}
