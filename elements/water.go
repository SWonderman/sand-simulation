package elements

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WaterElement struct {
	cell   *Cell
	family ElementFamily
	color  rl.Color
}

func (waterElement *WaterElement) GetCell() *Cell {
	return waterElement.cell
}

func (waterElement *WaterElement) SetCell(cell *Cell) {
	waterElement.cell = cell
}

func (waterElement *WaterElement) GetFamily() ElementFamily {
	return waterElement.family
}

func (waterElement *WaterElement) Update(matrix *Grid) {
	// TODO: impl!
	return
}

func (waterElement *WaterElement) GetColor() rl.Color {
	return waterElement.color
}

type WaterFamily struct {
	elementType ElementType
	spread      int
	colors      map[int]rl.Color
}

func NewWaterFamily() *WaterFamily {
	return &WaterFamily{
		elementType: Void,
		spread:      2,
		colors: map[int]rl.Color{
			1: {155, 206, 235, 255},
			2: {145, 196, 225, 255},
			3: {135, 186, 215, 255},
			4: {125, 176, 205, 255},
		},
	}
}

func (waterFamily *WaterFamily) GetType() ElementType {
	return waterFamily.elementType
}

func (waterFamily *WaterFamily) GetSpread() int {
	return waterFamily.spread
}

func (waterFamily *WaterFamily) GetColors() map[int]rl.Color {
	return waterFamily.colors
}

func (waterFamily *WaterFamily) SelectRandomColor() rl.Color {
	return waterFamily.colors[rand.Intn(len(waterFamily.colors)-1)+1]
}

func (waterFamily *WaterFamily) CreateElement(cell *Cell) Element {
	return &WaterElement{
		cell:   cell,
		family: waterFamily,
		color:  waterFamily.SelectRandomColor(),
	}
}

func (waterFamily *WaterFamily) CreateElements(grid *Grid, cell *Cell) []Element {
	createdElements := []Element{}
	spread := waterFamily.GetSpread()

	for i := -spread; i <= spread; i++ {
		for j := -spread; j <= spread; j++ {
			if rand.Intn(10) >= 3 {
				continue
			}

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

			createdElements = append(createdElements, waterFamily.CreateElement(&Cell{Row: newRow, Column: newColumn}))
		}
	}

	return createdElements
}
