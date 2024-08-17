package elements

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StoneElement struct {
	cell   *Cell
	family ElementFamily
	color  rl.Color
}

func (stoneElement *StoneElement) GetCell() *Cell {
	return stoneElement.cell
}

func (stoneElement *StoneElement) SetCell(cell *Cell) {
	stoneElement.cell = cell
}

func (stoneElement *StoneElement) GetFamily() ElementFamily {
	return stoneElement.family
}

func (stoneElement *StoneElement) GetColor() rl.Color {
	return stoneElement.color
}

func (stoneElement *StoneElement) Update(matrix *Grid) {
	return
}

type StoneFamily struct {
	elementType ElementType
	spread      int
	colors      map[int]rl.Color
}

func NewStoneFamily() *StoneFamily {
	return &StoneFamily{
		elementType: Stone,
		spread:      1,
		colors: map[int]rl.Color{
			1: {120, 120, 120, 255},
			2: {130, 130, 130, 255},
			3: {140, 140, 140, 255},
			4: {150, 150, 150, 255},
		},
	}
}

func (stoneFamily *StoneFamily) GetType() ElementType {
	return stoneFamily.elementType
}

func (stoneFamily *StoneFamily) GetSpread() int {
	return stoneFamily.spread
}

func (stoneFamily *StoneFamily) GetColors() map[int]rl.Color {
	return stoneFamily.colors
}

func (stoneFamily *StoneFamily) GetName() string {
	return "Stone"
}

func (stoneFamily *StoneFamily) SelectRandomColor() rl.Color {
	return stoneFamily.colors[rand.Intn(len(stoneFamily.colors)-1)+1]
}

func (stoneFamily *StoneFamily) CreateElement(cell *Cell) Element {
	return &StoneElement{
		cell:   cell,
		family: stoneFamily,
		color:  stoneFamily.SelectRandomColor(),
	}
}

func (stoneFamily *StoneFamily) CreateElements(grid *Grid, cell *Cell) []Element {
	createdElements := []Element{}
	spread := stoneFamily.GetSpread()

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

			createdElements = append(createdElements, stoneFamily.CreateElement(&Cell{Row: newRow, Column: newColumn}))
		}
	}

	return createdElements
}
