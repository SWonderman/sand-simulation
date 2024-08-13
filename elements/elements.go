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
