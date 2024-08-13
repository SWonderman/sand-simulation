package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

type Cell struct {
	Row    int
	Column int
}

type ElementType int

const (
	Void ElementType = iota
	Sand ElementType = iota
)

type ElementFamily interface {
	GetType() ElementType
	GetSpread() int
	GetColors() map[int]rl.Color
	SelectRandomColor() rl.Color
	CreateElement(cell *Cell) Element
	// IsMovable() bool // TODO: How about adding this?
}

type VoidFamily struct {
	elementType ElementType
	spread      int
	colors      map[int]rl.Color
}

func NewVoidFamily() *VoidFamily {
	return &VoidFamily{
		elementType: Void,
		spread:      1,
		colors: map[int]rl.Color{
			1: {20, 20, 20, 255},
		},
	}
}

func (voidFamily *VoidFamily) GetType() ElementType {
	return voidFamily.elementType
}

func (voidFamily *VoidFamily) GetSpread() int {
	return voidFamily.spread
}

func (voidFamily *VoidFamily) GetColors() map[int]rl.Color {
	return voidFamily.colors
}

func (voidFamily *VoidFamily) SelectRandomColor() rl.Color {
	return voidFamily.colors[1]
}

func (voidFamily *VoidFamily) CreateElement(cell *Cell) Element {
	return &VoidElement{
		cell:   cell,
		family: voidFamily,
		color:  voidFamily.SelectRandomColor(),
	}
}

type SandFamily struct {
	elementType ElementType
	spread      int
	colors      map[int]rl.Color
}

func NewSandFamily() *SandFamily {
	return &SandFamily{
		elementType: Sand,
		spread:      2,
		colors: map[int]rl.Color{
			1: {237, 201, 175, 255},
			2: {220, 189, 152, 255},
			3: {210, 178, 140, 255},
			4: {194, 165, 127, 255},
		},
	}
}

func (sandFamily *SandFamily) GetType() ElementType {
	return sandFamily.elementType
}

func (sandFamily *SandFamily) GetSpread() int {
	return sandFamily.spread
}

func (sandFamily *SandFamily) GetColors() map[int]rl.Color {
	return sandFamily.colors
}

func (sandFamily *SandFamily) SelectRandomColor() rl.Color {
	return sandFamily.colors[rand.Intn(len(sandFamily.colors)-1)+1]
}

func (sandFamily *SandFamily) CreateElement(cell *Cell) Element {
	return &SandElement{
		cell:   cell,
		family: sandFamily,
		color:  sandFamily.SelectRandomColor(),
	}
}

type Element interface {
	GetCell() *Cell
	SetCell(cell *Cell)
	GetFamily() ElementFamily
	Update(matrix *Grid)
	GetColor() rl.Color
}

type VoidElement struct {
	cell   *Cell
	family ElementFamily
	color  rl.Color
}

func (voidElement *VoidElement) GetCell() *Cell {
	return voidElement.cell
}

func (voidElement *VoidElement) SetCell(cell *Cell) {
	voidElement.cell = cell
}

func (voidElement *VoidElement) GetFamily() ElementFamily {
	return voidElement.family
}

func (voidElement *VoidElement) Update(matrix *Grid) {
	return
}

func (voidElement *VoidElement) GetColor() rl.Color {
	return voidElement.color
}

type SandElement struct {
	cell   *Cell
	family ElementFamily
	color  rl.Color
}

func (sandElement *SandElement) GetCell() *Cell {
	return sandElement.cell
}

func (sandElement *SandElement) SetCell(cell *Cell) {
	sandElement.cell = cell
}

func (sandElement *SandElement) GetFamily() ElementFamily {
	return sandElement.family
}

func (sandElement *SandElement) Update(matrix *Grid) {
	cell := sandElement.GetCell()
	if matrix.GetElement(cell) == nil {
		return
	}

	elementBelow := matrix.GetElement(&Cell{Row: cell.Row + 1, Column: cell.Column})
	if elementBelow == nil {
		return
	}

	elementDiagonallyRight := matrix.GetElement(&Cell{Row: cell.Row + 1, Column: cell.Column + 1})
	elementDiagonallyLeft := matrix.GetElement(&Cell{Row: cell.Row + 1, Column: cell.Column - 1})

	// Sand is falling!
	if elementBelow.GetFamily().GetType() == Void {
		matrix.SwapElements(sandElement, elementBelow)
	} else if elementDiagonallyRight != nil && elementDiagonallyRight.GetFamily().GetType() == Void {
		matrix.SwapElements(sandElement, elementDiagonallyRight)
	} else if elementDiagonallyLeft != nil && elementDiagonallyLeft.GetFamily().GetType() == Void {
		matrix.SwapElements(sandElement, elementDiagonallyLeft)
	}
}

func (sandElement *SandElement) GetColor() rl.Color {
	return sandElement.color
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

func main() {
	const ROWS int32 = 80
	const COLUMNS int32 = 160
	const CELL_SIZE = 10

	const WIN_WIDTH int32 = COLUMNS * CELL_SIZE
	const WIN_HEIGHT int32 = ROWS * CELL_SIZE

	rl.InitWindow(WIN_WIDTH, WIN_HEIGHT, "Sandy")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	voidFamily := NewVoidFamily()
	sandFamily := NewSandFamily()

	selectedFamily := sandFamily

	matrix := NewGrid(int(COLUMNS), int(ROWS), voidFamily)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyQ) {
			matrix.Clear()
		}

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			matrix.FillCells(rl.GetMousePosition(), CELL_SIZE, selectedFamily)
		}

		// Check and apply rules
		for r := ROWS - 1; r >= 0; r-- {
			for c := COLUMNS - 1; c >= 0; c-- {
				currentCell := matrix.GetElement(&Cell{Row: int(r), Column: int(c)})
				if currentCell != nil {
					currentCell.Update(matrix)
				}
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		for r := range ROWS {
			for c := range COLUMNS {
				currentElement := matrix.GetElement(&Cell{Row: int(r), Column: int(c)})
				color := currentElement.GetColor()
				rl.DrawRectangle(c*CELL_SIZE, r*CELL_SIZE, CELL_SIZE, CELL_SIZE, color)
			}
		}

		rl.EndDrawing()
	}
}
