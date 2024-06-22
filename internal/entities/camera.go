package entities

import (
	"fmt"
	"os"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	"github.com/Ollie-Ave/Project_Kat_3/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	cameraMovementOffset = 100
	maxCameraSpeed       = playerSpeed
	cameraSpeed          = 50
	cameraDrag           = 25
)

func NewCamera(xInitialPosition float32, entityManager engine_entities.EntityManager) engine_entities.EntityUpdater {
	defaultOffset := rl.NewVector2(float32(shared.WindowWidth)/2, -float32(shared.WindowHeight)/2)
	defaultTarget := rl.NewVector2(xInitialPosition, -50)
	defaultRotation := float32(0)
	defaultZoom := float32(2)

	c := &camera{
		camera:        rl.NewCamera2D(defaultOffset, defaultTarget, defaultRotation, defaultZoom),
		levelWidth:    0,
		entityManager: entityManager,
		velocity:      rl.NewVector2(0, 0),
	}

	c.camera.Target = c.getNewCameraTarget(false)

	return c
}

type camera struct {
	camera rl.Camera2D

	levelWidth int
	velocity   rl.Vector2

	entityManager engine_entities.EntityManager
}

func (c *camera) Update() error {

	c.velocity = applyCameraDrag(c.velocity)

	playerHitbox, err := c.getPlayerHitbox()

	if err != nil {
		return err
	}

	c.velocity = c.getNewCameraVelocity(playerHitbox)

	c.camera.Target = c.getNewCameraTarget(true)

	return nil
}

func (c *camera) GetCamera() rl.Camera2D {
	return c.camera
}

func (c *camera) SetLevelWidth(width int) {
	c.levelWidth = width
}

func (c *camera) getPlayerHitbox() (*rl.Rectangle, error) {
	playerEntity, err := c.entityManager.GetEntityById(shared.PlayerEntityName)

	if err != nil {
		return nil, err
	}

	playerCollider, isCollider := playerEntity.(engine_entities.Collider)

	if !isCollider {
		return nil, fmt.Errorf("found and entity with id %s but it does not implement the Collider interface", shared.PlayerEntityName)
	}

	playerHitbox := playerCollider.GetHitbox()

	return &playerHitbox, nil
}

func clampValue(value, max, min float32) float32 {
	if value > max {
		return max
	} else if value < min {
		return min
	}

	return value
}

func (c *camera) getNewCameraVelocity(playerHitbox *rl.Rectangle) rl.Vector2 {
	velocity := c.velocity

	playerCenterPoint := playerHitbox.X + (playerHitbox.Width / 2)

	windowMidpoint := c.camera.Target.X

	if playerCenterPoint > float32(windowMidpoint+cameraMovementOffset) {
		velocity.X = clampValue(velocity.X+cameraSpeed, maxCameraSpeed, -maxCameraSpeed)
	} else if playerCenterPoint < float32(windowMidpoint-cameraMovementOffset) {
		velocity.X = clampValue(velocity.X-cameraSpeed, maxCameraSpeed, -maxCameraSpeed)
	}

	if os.Getenv(engine_shared.DebugModeEnvironmentVariable) == "true" {
		rl.DrawLine(int32(playerCenterPoint), 0, int32(playerCenterPoint), shared.WindowHeight, rl.Red)

		rl.DrawLine(int32(windowMidpoint), 0, int32(windowMidpoint), shared.WindowHeight, rl.Red)
		rl.DrawLine(int32(windowMidpoint-cameraMovementOffset), 0, int32(windowMidpoint-cameraMovementOffset), shared.WindowHeight, rl.Green)
		rl.DrawLine(int32(windowMidpoint+cameraMovementOffset), 0, int32(windowMidpoint+cameraMovementOffset), shared.WindowHeight, rl.Green)
	}

	return velocity
}

func (c *camera) getNewCameraTarget(checkVelocity bool) rl.Vector2 {
	currentTarget := c.camera.Target

	deltaTime := rl.GetFrameTime()

	currentTarget.X += c.velocity.X * deltaTime

	cameraLeftSide := c.camera.Target.X - (c.camera.Offset.X / 2)
	cameraRightSide := c.camera.Target.X + (c.camera.Offset.X / 2)

	maxTargetX := float32(c.levelWidth)

	const extraPixelsToHackFixRenderingArtifact = 5

	if cameraLeftSide <= extraPixelsToHackFixRenderingArtifact && (c.velocity.X < 0 || !checkVelocity) {
		currentTarget.X = 0 + (c.camera.Offset.X / 2) + 1
	} else if cameraRightSide >= maxTargetX-extraPixelsToHackFixRenderingArtifact && (c.velocity.X > 0 || !checkVelocity) {
		currentTarget.X = maxTargetX - (c.camera.Offset.X / 2) - 1
	}

	return currentTarget
}

func applyCameraDrag(currentVelocity rl.Vector2) rl.Vector2 {
	if currentVelocity.X > 0 {
		currentVelocity.X = clampValue(currentVelocity.X-cameraDrag, maxCameraSpeed, 0)
	} else if currentVelocity.X < 0 {
		currentVelocity.X = clampValue(currentVelocity.X+cameraDrag, 0, -maxCameraSpeed)
	}

	return currentVelocity
}
