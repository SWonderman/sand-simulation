package elements

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Element interface {
	GetCell() *Cell
	SetCell(cell *Cell)
	GetFamily() ElementFamily
	Update(matrix *Grid)
	GetColor() rl.Color
}

type ElementType int

const (
	Void  ElementType = iota
	Sand  ElementType = iota
	Stone ElementType = iota
	Water ElementType = iota
)

type ElementFamily interface {
	GetType() ElementType
	GetSpread() int
	GetColors() map[int]rl.Color
	GetName() string
	SelectRandomColor() rl.Color
	CreateElement(cell *Cell) Element
	CreateElements(grid *Grid, cell *Cell) []Element
	// IsMovable() bool // TODO: How about adding this?
}
