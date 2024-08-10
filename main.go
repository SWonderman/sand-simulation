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
	} else if row >= len((*matrix)) {
		row = len((*matrix)) - 1
	} else if column >= len((*matrix)[0]) {
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

func main() {
	const ROWS int32 = 40
	const COLUMNS int32 = 80
	const CELL_SIZE = 10

	const WIN_WIDTH int32 = COLUMNS * CELL_SIZE
	const WIN_HEIGHT int32 = ROWS * CELL_SIZE

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
			gameMatrix[row][column] = byte(rand.Intn(len(sandColors)-1)) + 1
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
