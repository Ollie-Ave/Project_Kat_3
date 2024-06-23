package levels

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_levels"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewTestLevel(
	levelLoader engine_levels.LevelLoader,
	levelRenderer engine_levels.LevelRenderer,
	entityManager engine_entities.EntityManager) (*testLevel, error) {
	futureLevelFilePath := fmt.Sprintf("%s/%s/testLevel.json", engine_shared.AssetsPath, engine_shared.LevelPath)
	pastLevelFilePath := fmt.Sprintf("%s/%s/testLevel_Past.json", engine_shared.AssetsPath, engine_shared.LevelPath)

	const defaultTimePeriod = engine_levels.FutureTimePeriod

	levelData, err := levelLoader.LoadLevelData(futureLevelFilePath, pastLevelFilePath, defaultTimePeriod)

	if err != nil {
		return nil, err
	}

	camera, err := entityManager.GetCamera()

	if err != nil {
		return nil, err
	}

	camera.SetLevelWidth(levelData.CurrentTimePeriod.TileWidth * levelData.CurrentTimePeriod.Width)

	levelCollider := engine_levels.NewLevelCollider(levelData, defaultTimePeriod)
	entityManager.SetLevelCollider(levelCollider)

	return &testLevel{
		levelData:     levelData,
		levelRenderer: levelRenderer,
		entityManager: entityManager,
	}, nil
}

type testLevel struct {
	levelRenderer engine_levels.LevelRenderer

	entityManager engine_entities.EntityManager

	levelData *engine_levels.LevelData
}

func (l *testLevel) Update() {
	if rl.IsKeyPressed(rl.KeyR) {
		if l.levelData.CurrentTimePeriod == l.levelData.PastPeriod {
			l.levelData.CurrentTimePeriod = l.levelData.FuturePeriod
		} else {
			l.levelData.CurrentTimePeriod = l.levelData.PastPeriod
		}

		l.entityManager.
			GetLevelCollider().
			SetTimePeriod(l.GetTimePeriod())
	}
}

func (l *testLevel) Render() {
	l.levelRenderer.Render(l.levelData)
}

func (l *testLevel) GetTimePeriod() string {
	if l.levelData.CurrentTimePeriod == l.levelData.PastPeriod {
		return engine_levels.PastTimePeriod
	} else {
		return engine_levels.FutureTimePeriod
	}
}
