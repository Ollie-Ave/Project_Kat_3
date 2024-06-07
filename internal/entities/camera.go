package entities

import rl "github.com/gen2brain/raylib-go/raylib"

func NewCamera() *Camera {
	defaultOffset := rl.NewVector2(0, 0)
	defaultTarget := rl.NewVector2(0, 0)
	defaultRotation := float32(0)
	defaultZoom := float32(1.45)

	return &Camera{
		camera: rl.NewCamera2D(defaultOffset, defaultTarget, defaultRotation, defaultZoom),
	}
}

type Camera struct {
	camera rl.Camera2D
}

func (c *Camera) Update() {
	if rl.IsKeyDown(rl.KeyA) {
		c.camera.Offset.X += 2
	} else if rl.IsKeyDown(rl.KeyD) {
		c.camera.Offset.X -= 2
	}
}

func (c *Camera) GetCamera() rl.Camera2D {
	return c.camera
}
