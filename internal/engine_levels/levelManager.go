package engine_levels

import (
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
)

type LevelManager interface {
	GetLevel() engine_shared.Renderer
}

func NewLevelManager(
	initialLevel engine_shared.Renderer,
	levelRenderer LevelRenderer,
	entityManager engine_entities.EntityManager) (LevelManager, error) {
	return &LevelManagerImpl{
		currentLevel: initialLevel,
	}, nil
}

type LevelManagerImpl struct {
	currentLevel engine_shared.Renderer
}

func (l *LevelManagerImpl) GetLevel() engine_shared.Renderer {
	return l.currentLevel
}
