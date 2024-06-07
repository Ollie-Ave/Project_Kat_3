package shared

import rl "github.com/gen2brain/raylib-go/raylib"

type CameraPosessor interface {
	GetCamera() rl.Camera2D
}

type Renderer interface {
	Render()
}
