package levels

import "github.com/Ollie-Ave/Project_Kat_3/internal/shared"

func NewLevelManager() (LevelManager, error) {
	levelOne, err := initLevelOne()

	if err != nil {
		return nil, err
	}

	return &LevelManagerImpl{
		currentLevel: levelOne,
	}, nil
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
