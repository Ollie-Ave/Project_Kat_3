package main

import (
	"fmt"

	"github.com/Ollie-Ave/Project_Kat_3/internal/constants"
	"github.com/Ollie-Ave/Project_Kat_3/internal/levels"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	setupWindow()

	levelManager, err := levels.InitLevelManager()

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting Game")

	for !rl.WindowShouldClose() {
		update(levelManager)
	}

	rl.CloseWindow()
}

func setupWindow() {
	rl.InitWindow(constants.WindowWidth, constants.WindowHeight, constants.WindowTitle)

	rl.SetTargetFPS(constants.WindowTargetFPS)
	rl.SetExitKey(constants.WindowExitKey)
}

func update(levelManager levels.LevelManager) {
	rl.BeginDrawing()

	rl.ClearBackground(constants.WindowBackgroundColor)

	renderFPS()

	levelManager.
		GetLevel().
		Render()

	rl.EndDrawing()

	renderFPS()
}

func renderFPS() {
	fpsText := fmt.Sprintf("FPS: %d", rl.GetFPS())

	rl.DrawText(fpsText, 10, 10, 20, rl.White)
}
