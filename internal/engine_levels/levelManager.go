package engine_levels

import (
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
)

type LevelManager interface {
	GetLevel() LevelHandler
}

func NewLevelManager(
	initialLevel LevelHandler,
	levelRenderer LevelRenderer,
	entityManager engine_entities.EntityManager) LevelManager {

	return &LevelManagerImpl{
		currentLevel: initialLevel,
	}
}

type LevelManagerImpl struct {
	currentLevel LevelHandler
}

func (l *LevelManagerImpl) GetLevel() LevelHandler {
	return l.currentLevel
}
