package levels

import (
	"fmt"
	"strings"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_levels"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	"github.com/Ollie-Ave/Project_Kat_3/internal/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewLevelOne(
	levelLoader engine_levels.LevelLoader,
	levelRenderer engine_levels.LevelRenderer,
	entityManager engine_entities.EntityManager,
	entityFactory entities.EntityFactory) (*levelOne, error) {
	futureLevelFilePath := fmt.Sprintf("%s/%s/levelOne.json", engine_shared.AssetsPath, engine_shared.LevelPath)

	const defaultTimePeriod = engine_levels.FutureTimePeriod

	levelData, err := levelLoader.LoadLevelData(futureLevelFilePath, futureLevelFilePath, defaultTimePeriod)

	if err != nil {
		return nil, err
	}

	return &levelOne{
		levelData:     levelData,
		levelRenderer: levelRenderer,
		entityManager: entityManager,
		entityFactory: entityFactory,
	}, nil
}

type levelOne struct {
	levelRenderer engine_levels.LevelRenderer

	entityManager engine_entities.EntityManager

	levelData *engine_levels.LevelData

	entityFactory entities.EntityFactory
}

func (l *levelOne) Initialise() error {
	err := setLevelWidthOnCamera(l.levelData, l.entityManager)

	if err != nil {
		return err
	}

	levelCollider := l.setLevelCollider()

	err = l.spawnEntities(levelCollider)

	if err != nil {
		return err
	}

	return nil
}

func (l *levelOne) Update() {
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

func (l *levelOne) Render() {
	l.levelRenderer.Render(l.levelData)
}

func (l *levelOne) GetTimePeriod() string {
	if l.levelData.CurrentTimePeriod == l.levelData.PastPeriod {
		return engine_levels.PastTimePeriod
	} else {
		return engine_levels.FutureTimePeriod
	}
}

func setLevelWidthOnCamera(levelData *engine_levels.LevelData, entityManager engine_entities.EntityManager) error {
	camera, err := entityManager.GetCamera()

	if err != nil {
		return err
	}

	camera.SetLevelWidth(levelData.CurrentTimePeriod.TileWidth * levelData.CurrentTimePeriod.Width)

	return nil
}

func (l *levelOne) setLevelCollider() engine_shared.LevelCollider {
	timePeriod := l.GetTimePeriod()

	levelCollider := engine_levels.NewLevelCollider(l.levelData, timePeriod)

	l.entityManager.SetLevelCollider(levelCollider)

	return levelCollider.(engine_shared.LevelCollider)
}

func (l *levelOne) spawnEntities(levelCollider engine_shared.LevelCollider) error {

	for _, layer := range l.levelData.PastPeriod.Layers {
		if strings.Contains(layer.Name, engine_levels.SpawnerLayerIndicator) {
			err := l.spawnEntitiesForLayer(layer, levelCollider)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *levelOne) spawnEntitiesForLayer(layer *engine_levels.Layer, levelCollider engine_shared.LevelCollider) error {
	entitiesToSpawn := strings.Split(layer.Name, "_")

	if len(entitiesToSpawn) != 2 {
		return fmt.Errorf("spawner layer name must be in the format 'Spawner_EntityName'")
	}

	entityName := entitiesToSpawn[1]

	data, tileWidth, tileHeight := levelCollider.GetLayerCollisionData(layer.Name)

	for y, dataRow := range data {
		for x, dataPoint := range dataRow {
			if dataPoint {
				position := rl.NewVector2(float32((x*tileWidth)-tileWidth), float32((y*tileHeight)-tileHeight))

				entity, err := l.entityFactory.CreateEntityByType(entityName, position)

				if err != nil {
					return err
				}

				l.entityManager.SpawnEntity(entityName, entity)
			}

		}
	}

	return nil
}
