package engine_levels

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
)

func NewLevelCollider(levelData *LevelData, defaultTimePeriod string) engine_entities.EntityUpdater {
	pastCollisionData := make(map[string][][]bool)

	for _, layer := range levelData.PastPeriod.Layers {

		if layer.LayerType == engine_shared.TileLayer {
			pastCollisionData[layer.Name] = parseCollisionDataForLayer(layer, levelData.PastPeriod)
		}
	}

	futureCollisionData := make(map[string][][]bool)

	for _, layer := range levelData.FuturePeriod.Layers {

		if layer.LayerType == engine_shared.TileLayer {
			futureCollisionData[layer.Name] = parseCollisionDataForLayer(layer, levelData.FuturePeriod)
		}
	}

	return &levelColliderImpl{
		currentTimePeriod:   defaultTimePeriod,
		pastCollisionData:   pastCollisionData,
		futureCollisionData: futureCollisionData,
		tileWidth:           levelData.CurrentTimePeriod.TileWidth,
		tileHeight:          levelData.CurrentTimePeriod.TileHeight,
	}
}

func parseCollisionDataForLayer(layer *Layer, levelData *LevelTimePeriod) [][]bool {
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
	currentTimePeriod   string
	pastCollisionData   map[string][][]bool
	futureCollisionData map[string][][]bool

	tileWidth  int
	tileHeight int
}

func (l *levelColliderImpl) Update() error {
	return nil
}

func (l *levelColliderImpl) SetTimePeriod(timePeriod string) error {
	if timePeriod != PastTimePeriod && timePeriod != FutureTimePeriod {
		return fmt.Errorf("invalid time period: %s", timePeriod)
	}

	l.currentTimePeriod = timePeriod

	return nil
}

func (l *levelColliderImpl) GetLayerCollisionData(layerName string) ([][]bool, int, int) {
	var collisionData map[string][][]bool

	if l.currentTimePeriod == FutureTimePeriod {
		collisionData = l.futureCollisionData
	} else {
		collisionData = l.pastCollisionData
	}

	return collisionData[layerName], l.tileWidth, l.tileHeight
}
