package main

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/entities"
	"github.com/Ollie-Ave/Project_Kat_3/internal/levels"
	"github.com/Ollie-Ave/Project_Kat_3/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	setupWindow()

	levelManager, err := levels.NewLevelManager()

	entityManager := entities.NewEntityManager()
	entityManager.SpawnEntity(shared.CameraEntityName, entities.NewCamera())

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

func update(levelManager levels.LevelManager, entityManager entities.EntityManager) {
	rl.BeginDrawing()

	err := beginCameraMode2D(entityManager)

	if err != nil {
		panic(err)
	}

	rl.ClearBackground(shared.WindowBackgroundColor)

	renderFPS()

	levelManager.
		GetLevel().
		Render()

	for _, entity := range entityManager.GetEntities() {
		entity.Update()
	}

	rl.EndMode2D()

	rl.EndDrawing()

	renderFPS()
}

func beginCameraMode2D(entityManager entities.EntityManager) error {
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
