package elements

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	Row    int
	Column int
}

type Grid struct {
	Width      int
	Height     int
	Matrix     [][]Element
	BaseFamily ElementFamily
}

func NewGrid(width int, height int, elementFamily ElementFamily) *Grid {
	matrix := make([][]Element, height)
	for m := range matrix {
		matrix[m] = make([]Element, width)
	}

	for r := range height {
		for c := range width {
			matrix[r][c] = elementFamily.CreateElement(&Cell{Row: r, Column: c})
		}
	}

	return &Grid{
		Width:      width,
		Height:     height,
		Matrix:     matrix,
		BaseFamily: elementFamily,
	}
}

func (grid *Grid) Clear() {
	for r := range grid.Height {
		for c := range grid.Width {
			grid.Matrix[r][c] = grid.BaseFamily.CreateElement(&Cell{Row: r, Column: c})
		}
	}
}

func (grid *Grid) convertMousePositionToGrid(mousePosition rl.Vector2, cellSize int32) *Cell {
	row := int(mousePosition.Y / float32(cellSize))
	column := int(mousePosition.X / float32(cellSize))

	// snap the coordinates to the grid size to prevent index out of range errors
	if row < 0 {
		row = 0
	} else if column < 0 {
		column = 0
	} else if row >= grid.Height-1 {
		row = grid.Height - 1
	} else if column >= grid.Width-1 {
		column = grid.Width - 1
	}

	return &Cell{
		Row:    row,
		Column: column,
	}
}

func (grid *Grid) isCellWithinBounds(cell *Cell) bool {
	return cell.Row >= 0 && cell.Row < grid.Height && cell.Column >= 0 && cell.Column < grid.Width
}

func (grid *Grid) GetElement(cell *Cell) Element {
	if grid.isCellWithinBounds(cell) {
		return grid.Matrix[cell.Row][cell.Column]
	}

	return nil
}

func (grid *Grid) SetElement(element Element) {
	cell := element.GetCell()
	if grid.isCellWithinBounds(cell) {
		grid.Matrix[cell.Row][cell.Column] = element
	}
}

func (grid *Grid) SwapCells(fromCell *Cell, toCell *Cell) {
	grid.Matrix[fromCell.Row][fromCell.Column], grid.Matrix[toCell.Row][toCell.Column] = grid.Matrix[toCell.Row][toCell.Column], grid.Matrix[fromCell.Row][fromCell.Column]
}

func (grid *Grid) SwapElements(fromElement Element, toElement Element) {
	fromCell := fromElement.GetCell()
	toCell := toElement.GetCell()

	grid.Matrix[fromCell.Row][fromCell.Column], grid.Matrix[toCell.Row][toCell.Column] = grid.Matrix[toCell.Row][toCell.Column], grid.Matrix[fromCell.Row][fromCell.Column]

	fromElement.SetCell(toCell)
	toElement.SetCell(fromCell)
}

func (grid *Grid) FillCells(currentMousePosition rl.Vector2, cellSize int32, elementFamily ElementFamily) {
	mousePos := grid.convertMousePositionToGrid(currentMousePosition, cellSize)
	spread := elementFamily.GetSpread()

	for i := -spread; i <= spread; i++ {
		for j := -spread; j <= spread; j++ {
			if rand.Intn(10) >= 3 {
				continue
			}

			newRow := mousePos.Row + i
			newColumn := mousePos.Column + j

			if newRow <= 0 {
				newRow = 0
			} else if newRow >= grid.Height-1 {
				newRow = grid.Height - 1
			}

			if newColumn <= 0 {
				newColumn = 0
			} else if newColumn >= grid.Width-1 {
				newColumn = grid.Width - 1
			}

			grid.SetElement(elementFamily.CreateElement(&Cell{Row: newRow, Column: newColumn}))
		}
	}
}
