package entities

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	"github.com/Ollie-Ave/Project_Kat_3/internal/shared"
)

func NewEntityManager() engine_entities.EntityManager {
	return &EntityManagerImpl{
		entities:           make(map[string]engine_entities.EntityUpdatable),
		duplicateEntityIds: make(map[string]int),
	}
}

type EntityManagerImpl struct {
	entities           map[string]engine_entities.EntityUpdatable
	duplicateEntityIds map[string]int
}

func (e *EntityManagerImpl) SpawnEntity(id string, entity engine_entities.EntityUpdatable) {
	if e.entities[id] != nil {
		e.duplicateEntityIds[id]++

		id = fmt.Sprintf("%s-%d", id, e.duplicateEntityIds[id])
	}

	e.entities[id] = entity
}

func (e *EntityManagerImpl) GetEntities() map[string]engine_entities.EntityUpdatable {
	return e.entities
}

func (e *EntityManagerImpl) GetEntityById(id string) (engine_entities.EntityUpdatable, error) {
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
