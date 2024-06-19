package main

import (
	"fmt"
	"os"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_levels"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_shared"
	"github.com/Ollie-Ave/Project_Kat_3/internal/entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/levels"
	"github.com/Ollie-Ave/Project_Kat_3/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	setupWindow()

	entityManager := engine_entities.NewEntityManager()
	entityManager.SpawnEntity(shared.CameraEntityName, entities.NewCamera())

	playerStartingPosition := rl.NewVector2(200, 150)
	player, err := entities.NewPlayer(
		playerStartingPosition,
		entityManager,
		engine_entities.NewPhysicsHandler(entityManager),
		engine_entities.NewAnimationHandler(entities.PlayerIdleAnimation),
	)

	if err != nil {
		panic(err)
	}

	entityManager.SpawnEntity(shared.PlayerEntityName, player)

	levelRenderer := engine_levels.NewLevelRenderer(entityManager)
	levelLoader := engine_levels.NewLevelLoader()
	levelOne, err := levels.NewLevelOne(levelLoader, levelRenderer, entityManager)

	if err != nil {
		panic(err)
	}

	levelManager := engine_levels.NewLevelManager(levelOne, levelRenderer, entityManager)

	for !rl.WindowShouldClose() {
		update(levelManager, entityManager)
	}

	rl.CloseWindow()
}

func setupWindow() {
	rl.InitWindow(shared.WindowWidth, shared.WindowHeight, shared.WindowTitle)

	rl.SetTargetFPS(shared.WindowTargetFPS)
	rl.SetExitKey(shared.WindowExitKey)

	rl.SetConfigFlags(rl.FlagVsyncHint)
}

func update(levelManager engine_levels.LevelManager, entityManager engine_entities.EntityManager) {
	rl.BeginDrawing()

	err := beginCameraMode2D(entityManager)

	if err != nil {
		panic(err)
	}

	rl.ClearBackground(shared.WindowBackgroundColor)

	levelManager.
		GetLevel().
		Render()

	for _, entity := range entityManager.GetEntities() {

		entity.Update()

		if entity, ok := entity.(engine_shared.Renderable); ok {
			entity.Render()
		}
	}

	rl.EndMode2D()

	renderFPS()

	rl.EndDrawing()

	handleDebugMode()
}

func beginCameraMode2D(entityManager engine_entities.EntityManager) error {
	camera, err := entityManager.GetCamera()

	if err != nil {
		return err
	}

	rl.BeginMode2D(camera.GetCamera())

	return nil
}

func renderFPS() {
	fpsText := fmt.Sprintf("FPS: %d", rl.GetFPS())

	rl.DrawText(fpsText, 10, 10, 20, rl.White)
}

func handleDebugMode() {
	if rl.IsKeyReleased(rl.KeyF3) {
		newDebugState := "true"

		if os.Getenv(engine_shared.DebugModeEnvironmentVariable) == "true" {
			newDebugState = "false"
		}

		os.Setenv(engine_shared.DebugModeEnvironmentVariable, newDebugState)
	}
}
