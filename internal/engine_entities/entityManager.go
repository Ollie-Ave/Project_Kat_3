package engine_entities

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	"github.com/Ollie-Ave/Project_Kat_3/internal/shared"
)

type EntityManager interface {
	SpawnEntity(string, EntityUpdatable)

	SetLevelCollider(EntityUpdatable)

	GetEntities() map[string]EntityUpdatable

	GetEntityById(string) (EntityUpdatable, error)

	GetCamera() (engine_shared.CameraPosessor, error)
}

func NewEntityManager() EntityManager {
	return &EntityManagerImpl{
		entities:           make(map[string]EntityUpdatable),
		duplicateEntityIds: make(map[string]int),
	}
}

type EntityManagerImpl struct {
	entities           map[string]EntityUpdatable
	duplicateEntityIds map[string]int
}

func (e *EntityManagerImpl) SpawnEntity(id string, entity EntityUpdatable) {
	if e.entities[id] != nil {
		e.duplicateEntityIds[id]++

		id = fmt.Sprintf("%s-%d", id, e.duplicateEntityIds[id])
	}

	e.entities[id] = entity
}

func (e *EntityManagerImpl) GetEntities() map[string]EntityUpdatable {
	return e.entities
}

func (e *EntityManagerImpl) GetEntityById(id string) (EntityUpdatable, error) {
	entity := e.entities[id]

	if entity == nil {
		return nil, fmt.Errorf("entity with id %s not found", id)
	}

	return entity, nil
}

func (e *EntityManagerImpl) GetCamera() (engine_shared.CameraPosessor, error) {
	camera, err := e.GetEntityById(shared.CameraEntityName)

	if err != nil {
		return nil, err
	}

	return camera.(engine_shared.CameraPosessor), nil
}

func (e *EntityManagerImpl) SetLevelCollider(entity EntityUpdatable) {
	e.entities[engine_shared.LevelColliderEntityName] = entity
}
