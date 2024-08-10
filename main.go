package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	const ROWS int32 = 40
	const COLUMNS int32 = 40
	const CELL_SIZE = 15

	rl.InitWindow(COLUMNS*CELL_SIZE, ROWS*CELL_SIZE, "Sandy")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
