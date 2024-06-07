package engine_shared

import rl "github.com/gen2brain/raylib-go/raylib"

type CameraPosessor interface {
	GetCamera() rl.Camera2D

	SetLevelWidth(width int)
}

type Renderer interface {
	Render()
}
