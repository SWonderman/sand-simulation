package main

import (
	"fmt"
	"sw/sandy/elements"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const ROWS int32 = 80
	const COLUMNS int32 = 160
	const CELL_SIZE = 10

	const WIN_WIDTH int32 = COLUMNS * CELL_SIZE
	const WIN_HEIGHT int32 = ROWS * CELL_SIZE

	rl.InitWindow(WIN_WIDTH, WIN_HEIGHT, "Sandy")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	voidFamily := elements.NewVoidFamily()
	sandFamily := elements.NewSandFamily()
	stoneFamily := elements.NewStoneFamily()
	waterFamily := elements.NewWaterFamily()

	var selectedFamily elements.ElementFamily
	selectedFamily = stoneFamily
	currentlySeletedFamily := "Stone"

	matrix := elements.NewGrid(int(COLUMNS), int(ROWS), voidFamily)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyQ) {
			matrix.Clear()
		}

		if rl.IsKeyDown(rl.KeyOne) {
			selectedFamily = stoneFamily
			currentlySeletedFamily = "Stone"
		} else if rl.IsKeyDown(rl.KeyTwo) {
			selectedFamily = sandFamily
			currentlySeletedFamily = "Sand"
		} else if rl.IsKeyDown(rl.KeyThree) {
			selectedFamily = voidFamily
			currentlySeletedFamily = "Void"
		} else if rl.IsKeyDown(rl.KeyFour) {
			selectedFamily = waterFamily
			currentlySeletedFamily = "Water"
		}

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			matrix.FillCells(rl.GetMousePosition(), CELL_SIZE, selectedFamily)
		}

		// Check and apply rules
		for r := ROWS - 1; r >= 0; r-- {
			for c := COLUMNS - 1; c >= 0; c-- {
				currentCell := matrix.GetElement(&elements.Cell{Row: int(r), Column: int(c)})
				if currentCell != nil {
					currentCell.Update(matrix)
				}
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		for r := range ROWS {
			for c := range COLUMNS {
				currentElement := matrix.GetElement(&elements.Cell{Row: int(r), Column: int(c)})
				color := currentElement.GetColor()
				rl.DrawRectangle(c*CELL_SIZE, r*CELL_SIZE, CELL_SIZE, CELL_SIZE, color)
			}
		}

		rl.DrawText(fmt.Sprintf("Selected: %s", currentlySeletedFamily), 10, 10, 15, rl.RayWhite)

		rl.EndDrawing()
	}
}
