package levels

import "github.com/Ollie-Ave/Project_Kat_3/internal/shared"

func InitLevelManager() LevelManager {
	return &LevelManagerImpl{
		currentLevel: initLevelOne(),
	}
}

type LevelManager interface {
	GetLevel() shared.Renderer
}

type LevelManagerImpl struct {
	currentLevel shared.Renderer
}

func (l *LevelManagerImpl) GetLevel() shared.Renderer {
	return l.currentLevel
}
