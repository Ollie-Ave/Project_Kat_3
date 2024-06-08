package engine_levels

import (
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
)

type LevelManager interface {
	GetLevel() engine_shared.Renderable
}

func NewLevelManager(
	initialLevel engine_shared.Renderable,
	levelRenderer LevelRenderer,
	entityManager engine_entities.EntityManager) LevelManager {

	return &LevelManagerImpl{
		currentLevel: initialLevel,
	}
}

type LevelManagerImpl struct {
	currentLevel engine_shared.Renderable
}

func (l *LevelManagerImpl) GetLevel() engine_shared.Renderable {
	return l.currentLevel
}
