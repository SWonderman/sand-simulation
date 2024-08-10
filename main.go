package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	Row    int
	Column int
}

func createEmptyMatrix(rows int32, columns int32) [][]byte {
	matrix := make([][]byte, rows)
	for m := range matrix {
		matrix[m] = make([]byte, columns)
	}

	for r := range rows {
		for c := range columns {
			matrix[r][c] = 0
		}
	}

	return matrix
}

func clearMatrix(matrix *[][]byte) {
	for r := range len((*matrix)) {
		for c := range len((*matrix)[0]) {
			(*matrix)[r][c] = 0
		}
	}
}

func convertMousePositionToGrid(matrix *[][]byte, mousePosition rl.Vector2, cellSize int32) (int, int) {
	row := int(mousePosition.Y / float32(cellSize))
	column := int(mousePosition.X / float32(cellSize))

	// snap the coordinates to the grid size to prevent index out of range errors
	if row < 0 {
		row = 0
	} else if column < 0 {
		column = 0
	} else if row >= len((*matrix))-1 {
		row = len((*matrix)) - 1
	} else if column >= len((*matrix)[0])-1 {
		column = len((*matrix)[0]) - 1
	}

	return row, column
}

func canFall(matrix *[][]byte, cell *Cell) bool {
	if cell.Row+1 >= len((*matrix)) {
		return false
	}

	cellBelow := (*matrix)[cell.Row+1][cell.Column]
	if cellBelow >= 1 {
		return false
	}

	return true
}

func canRollLeft(matrix *[][]byte, cell *Cell) bool {
	bottomRow := cell.Row + 1
	leftColumn := cell.Column - 1

	if bottomRow >= len((*matrix)) {
		return false
	}

	if leftColumn < 0 {
		return false
	}

	leftDiagonalCellValue := (*matrix)[bottomRow][leftColumn]
	if leftDiagonalCellValue >= 1 {
		return false
	}

	return true
}

func canRollRight(matrix *[][]byte, cell *Cell) bool {
	bottomRow := cell.Row + 1
	rightColumn := cell.Column + 1

	if bottomRow >= len((*matrix)) {
		return false
	}

	if rightColumn >= len((*matrix)[0]) {
		return false
	}

	rightDiagonalCellValue := (*matrix)[bottomRow][rightColumn]
	if rightDiagonalCellValue >= 1 {
		return false
	}

	return true
}

func applyGravity(matrix *[][]byte, cell *Cell) {
	if (*matrix)[cell.Row][cell.Column] == 0 {
		return
	}

	if canFall(matrix, cell) {
		(*matrix)[cell.Row][cell.Column], (*matrix)[cell.Row+1][cell.Column] = (*matrix)[cell.Row+1][cell.Column], (*matrix)[cell.Row][cell.Column]
	} else if canRollLeft(matrix, cell) {
		(*matrix)[cell.Row][cell.Column], (*matrix)[cell.Row+1][cell.Column-1] = (*matrix)[cell.Row+1][cell.Column-1], (*matrix)[cell.Row][cell.Column]
	} else if canRollRight(matrix, cell) {
		(*matrix)[cell.Row][cell.Column], (*matrix)[cell.Row+1][cell.Column+1] = (*matrix)[cell.Row+1][cell.Column+1], (*matrix)[cell.Row][cell.Column]
	}
}

func fillCells(matrix *[][]byte, mousePos *Cell, spread int, sandColorsCount int) {
	for i := -spread; i <= spread; i++ {
		for j := -spread; j <= spread; j++ {
			if rand.Intn(10) >= 3 {
				continue
			}

			newRow := mousePos.Row + i
			newColumn := mousePos.Column + j

			if newRow <= 0 {
				newRow = 0
			} else if newRow >= len((*matrix))-1 {
				newRow = len((*matrix)) - 1
			}

			if newColumn <= 0 {
				newColumn = 0
			} else if newColumn >= len((*matrix)[0])-1 {
				newColumn = len((*matrix)[0]) - 1
			}

			(*matrix)[newRow][newColumn] = byte(rand.Intn(sandColorsCount-1)) + 1
		}
	}
}

func main() {
	const ROWS int32 = 80
	const COLUMNS int32 = 160
	const CELL_SIZE = 10

	const WIN_WIDTH int32 = COLUMNS * CELL_SIZE
	const WIN_HEIGHT int32 = ROWS * CELL_SIZE

	const SPREAD_SIZE int = 5
	const SPREAD int = int(SPREAD_SIZE / 2)

	sandColors := map[int]rl.Color{
		1: {237, 201, 175, 255},
		2: {220, 189, 152, 255},
		3: {210, 178, 140, 255},
		4: {194, 165, 127, 255},
	}

	rl.InitWindow(WIN_WIDTH, WIN_HEIGHT, "Sandy")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	gameMatrix := createEmptyMatrix(ROWS, COLUMNS)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyQ) {
			clearMatrix(&gameMatrix)
		}

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			row, column := convertMousePositionToGrid(&gameMatrix, rl.GetMousePosition(), CELL_SIZE)
			fillCells(&gameMatrix, &Cell{Row: row, Column: column}, SPREAD, len(sandColors))
		}

		// Check and apply rules
		for r := ROWS - 1; r >= 0; r-- {
			for c := COLUMNS - 1; c >= 0; c-- {
				applyGravity(&gameMatrix, &Cell{Row: int(r), Column: int(c)})
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		for r := range ROWS {
			for c := range COLUMNS {
				cellValue := gameMatrix[r][c]
				color := rl.Black

				if cellValue >= 1 {
					color = sandColors[int(cellValue)]
				}

				rl.DrawRectangle(c*CELL_SIZE, r*CELL_SIZE, CELL_SIZE, CELL_SIZE, color)
			}
		}

		rl.EndDrawing()
	}
}
