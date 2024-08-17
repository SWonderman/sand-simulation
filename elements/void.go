package elements

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func (voidFamily *VoidFamily) GetName() string {
	return "Void"
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

func (voidFamily *VoidFamily) CreateElements(grid *Grid, cell *Cell) []Element {
	createdElements := []Element{}
	spread := voidFamily.GetSpread()

	for i := -spread; i <= spread; i++ {
		for j := -spread; j <= spread; j++ {
			newRow := cell.Row + i
			newColumn := cell.Column + j

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

			createdElements = append(createdElements, voidFamily.CreateElement(&Cell{Row: newRow, Column: newColumn}))
		}
	}

	return createdElements
}
