package main

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/engine_levels"
	"github.com/Ollie-Ave/Project_Kat_3/internal/entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/levels"
	"github.com/Ollie-Ave/Project_Kat_3/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	setupWindow()

	entityManager := entities.NewEntityManager()
	entityManager.SpawnEntity(shared.CameraEntityName, entities.NewCamera())

	levelRenderer := engine_levels.NewLevelRenderer(entityManager)
	levelLoader := engine_levels.NewLevelLoader()
	levelOne, err := levels.NewLevelOne(levelLoader, levelRenderer, entityManager)

	if err != nil {
		panic(err)
	}

	levelManager, err := engine_levels.NewLevelManager(levelOne, levelRenderer, entityManager)

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting Game")

	for !rl.WindowShouldClose() {
		update(levelManager, entityManager)
	}

	rl.CloseWindow()
}

func setupWindow() {
	rl.InitWindow(shared.WindowWidth, shared.WindowHeight, shared.WindowTitle)

	rl.SetTargetFPS(shared.WindowTargetFPS)
	rl.SetExitKey(shared.WindowExitKey)
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
	}

	rl.EndMode2D()

	renderFPS()

	rl.EndDrawing()

	renderFPS()
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
