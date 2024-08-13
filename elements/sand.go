package elements

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
