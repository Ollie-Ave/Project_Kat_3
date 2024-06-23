package entities

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityFactory interface {
	CreateEntityByType(entityType string, position rl.Vector2) (engine_entities.EntityUpdater, error)
}

type EntityFactoryImpl struct {
	entityManager engine_entities.EntityManager
}

func NewEntityFactory(entityManager engine_entities.EntityManager) EntityFactory {
	return &EntityFactoryImpl{
		entityManager: entityManager,
	}
}

func (e *EntityFactoryImpl) CreateEntityByType(entityType string, position rl.Vector2) (engine_entities.EntityUpdater, error) {
	switch entityType {
	case shared.PlayerEntityName:
		return NewPlayer(
			position,
			e.entityManager,
			engine_entities.NewPhysicsHandler(e.entityManager),
			engine_entities.NewAnimationHandler(PlayerIdleAnimation),
		)
	default:
		return nil, fmt.Errorf("could not find entity with type '%s'", entityType)
	}

}
