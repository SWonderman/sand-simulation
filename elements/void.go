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
