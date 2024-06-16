package entities

import (
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewCamera() engine_entities.EntityUpdater {
	defaultOffset := rl.NewVector2(0, -300)
	defaultTarget := rl.NewVector2(0, 0)
	defaultRotation := float32(0)
	defaultZoom := float32(2)

	return &camera{
		camera:     rl.NewCamera2D(defaultOffset, defaultTarget, defaultRotation, defaultZoom),
		levelWidth: 0,
	}
}

type camera struct {
	camera rl.Camera2D

	levelWidth int
}

func (c *camera) Update() error {
	const cameraSpeed = 8

	if rl.IsKeyDown(rl.KeyD) {
		c.camera.Offset.X -= cameraSpeed

		maxTargetX := -(float32(c.levelWidth) * c.camera.Zoom) + shared.WindowWidth

		if c.camera.Offset.X <= float32(maxTargetX) {
			c.camera.Offset.X = float32(maxTargetX)
		}
	} else if rl.IsKeyDown(rl.KeyA) {
		c.camera.Offset.X += cameraSpeed

		if c.camera.Offset.X >= 0 {
			c.camera.Offset.X = 0
		}
	}

	return nil
}

func (c *camera) GetCamera() rl.Camera2D {
	return c.camera
}

func (c *camera) SetLevelWidth(width int) {
	c.levelWidth = width
}
