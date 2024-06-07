package engine_entities

import "github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"

type EntityManager interface {
	SpawnEntity(string, EntityUpdatable)

	GetEntities() map[string]EntityUpdatable

	GetEntityById(string) (EntityUpdatable, error)

	GetCamera() (engine_shared.CameraPosessor, error)
}
