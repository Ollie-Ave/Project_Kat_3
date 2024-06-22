package engine_entities

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	"github.com/Ollie-Ave/Project_Kat_3/internal/shared"
)

type EntityManager interface {
	SpawnEntity(string, EntityUpdater)

	SetLevelCollider(EntityUpdater)

	GetLevelCollider() engine_shared.LevelCollider

	GetEntities() map[string]EntityUpdater

	GetEntityById(string) (EntityUpdater, error)

	GetCamera() (engine_shared.CameraPosessor, error)
}

func NewEntityManager() EntityManager {
	return &entityManagerImpl{
		entities:           make(map[string]EntityUpdater),
		duplicateEntityIds: make(map[string]int),
	}
}

type entityManagerImpl struct {
	entities           map[string]EntityUpdater
	duplicateEntityIds map[string]int
}

func (e *entityManagerImpl) SpawnEntity(id string, entity EntityUpdater) {
	if e.entities[id] != nil {
		e.duplicateEntityIds[id]++

		id = fmt.Sprintf("%s-%d", id, e.duplicateEntityIds[id])
	}

	e.entities[id] = entity
}

func (e *entityManagerImpl) GetEntities() map[string]EntityUpdater {
	return e.entities
}

func (e *entityManagerImpl) GetEntityById(id string) (EntityUpdater, error) {
	entity := e.entities[id]

	if entity == nil {
		return nil, fmt.Errorf("entity with id %s not found", id)
	}

	return entity, nil
}

func (e *entityManagerImpl) GetCamera() (engine_shared.CameraPosessor, error) {
	camera, err := e.GetEntityById(shared.CameraEntityName)

	if err != nil {
		return nil, err
	}

	return camera.(engine_shared.CameraPosessor), nil
}

func (e *entityManagerImpl) SetLevelCollider(entity EntityUpdater) {
	e.entities[engine_shared.LevelColliderEntityName] = entity
}

func (e *entityManagerImpl) GetLevelCollider() engine_shared.LevelCollider {
	return e.entities[engine_shared.LevelColliderEntityName].(engine_shared.LevelCollider)
}
