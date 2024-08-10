package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func canFall(matrix *[][]byte, row int, column int) bool {
	if (*matrix)[row][column] == 0 {
		return false
	}

	if row+1 >= len((*matrix)) {
		return false
	}

	cellBelow := (*matrix)[row+1][column]
	if cellBelow == 1 {
		return false
	}

	return true
}

func canRollLeft(matrix *[][]byte, row int, column int) bool {
	if (*matrix)[row][column] == 0 {
		return false
	}

	bottomRow := row + 1
	leftColumn := column - 1

	if bottomRow >= len((*matrix)) {
		return false
	}

	if leftColumn < 0 {
		return false
	}

	leftDiagonalCellValue := (*matrix)[bottomRow][leftColumn]
	if leftDiagonalCellValue == 1 {
		return false
	}

	return true
}

func canRollRight(matrix *[][]byte, row int, column int) bool {
	if (*matrix)[row][column] == 0 {
		return false
	}

	bottomRow := row + 1
	rightColumn := column + 1

	if bottomRow >= len((*matrix)) {
		return false
	}

	if rightColumn >= len((*matrix)[0]) {
		return false
	}

	rightDiagonalCellValue := (*matrix)[bottomRow][rightColumn]
	if rightDiagonalCellValue == 1 {
		return false
	}

	return true
}

func moveDown(matrix *[][]byte, row int, column int) {
	(*matrix)[row][column] = 0
	(*matrix)[row+1][column] = 1
}

func main() {
	const ROWS int32 = 40
	const COLUMNS int32 = 40
	const CELL_SIZE = 15

	const WIN_WIDTH int32 = COLUMNS * CELL_SIZE
	const WIN_HEIGHT int32 = ROWS * CELL_SIZE

	sandColors := map[int]rl.Color{
		0: {237, 201, 175, 255},
		1: {220, 189, 152, 255},
		2: {210, 178, 140, 255},
		3: {194, 165, 127, 255},
		4: {172, 147, 106, 255},
		5: {157, 130, 89, 255},
	}

	rl.InitWindow(WIN_WIDTH, WIN_HEIGHT, "Sandy")
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	gameMatrix := createEmptyMatrix(ROWS, COLUMNS)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyQ) {
			clearMatrix(&gameMatrix)
		}

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			row, column := convertMousePositionToGrid(&gameMatrix, rl.GetMousePosition(), CELL_SIZE)
			gameMatrix[row][column] = 1
		}

		// Check and apply rules
		for r := ROWS - 1; r >= 0; r-- {
			for c := COLUMNS - 1; c >= 0; c-- {
				// Gravity
				if canFall(&gameMatrix, int(r), int(c)) {
					gameMatrix[r][c], gameMatrix[r+1][c] = gameMatrix[r+1][c], gameMatrix[r][c]
				} else if canRollLeft(&gameMatrix, int(r), int(c)) {
					gameMatrix[r][c], gameMatrix[r+1][c-1] = gameMatrix[r+1][c-1], gameMatrix[r][c]
				} else if canRollRight(&gameMatrix, int(r), int(c)) {
					gameMatrix[r][c], gameMatrix[r+1][c+1] = gameMatrix[r+1][c+1], gameMatrix[r][c]
				}
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		for r := range ROWS {
			for c := range COLUMNS {
				cellValue := gameMatrix[r][c]
				color := rl.Black

				if cellValue == 1 {
					color = sandColors[1]
				}

				// Draw cells
				rl.DrawRectangle(c*CELL_SIZE, r*CELL_SIZE, CELL_SIZE, CELL_SIZE, color)

				// Draw grid
				rl.DrawLine(c*CELL_SIZE, 0, c*CELL_SIZE, WIN_HEIGHT, rl.NewColor(130, 130, 130, 100))
				rl.DrawLine(0, r*CELL_SIZE, WIN_WIDTH, r*CELL_SIZE, rl.NewColor(130, 130, 130, 100))
			}
		}

		rl.EndDrawing()
	}
}
