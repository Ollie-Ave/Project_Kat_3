package engine_shared

import rl "github.com/gen2brain/raylib-go/raylib"

type CameraPosessor interface {
	GetCamera() rl.Camera2D

	SetLevelWidth(width int)
}

type Renderable interface {
	Render()
}

type LevelCollider interface {
	GetLayerCollisionData(layerName string) ([][]bool, int, int)
}
