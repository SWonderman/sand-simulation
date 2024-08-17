package elements

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type WaterElement struct {
	cell         *Cell
	family       ElementFamily
	color        rl.Color
    disperseRate int
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
	cell := waterElement.GetCell()

	elementBelow := matrix.GetElement(&Cell{Row: cell.Row + 1, Column: cell.Column})
	elementDiagonallyRight := matrix.GetElement(&Cell{Row: cell.Row + 1, Column: cell.Column + 1})
	elementDiagonallyLeft := matrix.GetElement(&Cell{Row: cell.Row + 1, Column: cell.Column - 1})
	elementRight := matrix.GetElement(&Cell{Row: cell.Row, Column: cell.Column + 1})
	elementLeft := matrix.GetElement(&Cell{Row: cell.Row, Column: cell.Column - 1})

	if elementBelow != nil && elementBelow.GetFamily().GetType() == Void {
		matrix.SwapElements(waterElement, elementBelow)
	} else if elementDiagonallyLeft != nil && elementDiagonallyLeft.GetFamily().GetType() == Void {
        waterElement.applyDispersion(matrix, elementDiagonallyLeft, &rl.Vector2{X: -1, Y: 1})
	} else if elementDiagonallyRight != nil && elementDiagonallyRight.GetFamily().GetType() == Void {
        waterElement.applyDispersion(matrix, elementDiagonallyRight, &rl.Vector2{X: 1, Y: 1})
    } else {
        decider := rand.Intn(2)

        if elementLeft != nil && elementLeft.GetFamily().GetType() == Void && decider == 0 {
            waterElement.applyDispersion(matrix, elementLeft, &rl.Vector2{X: -1, Y: 0})
        } else if elementRight != nil && elementRight.GetFamily().GetType() == Void && decider == 1 {
            waterElement.applyDispersion(matrix, elementRight, &rl.Vector2{X: 1, Y: 0})
        }
    }
}

func (waterElement *WaterElement) GetColor() rl.Color {
	return waterElement.color
}

func (waterElement *WaterElement) applyDispersion(matrix *Grid, element Element, movement *rl.Vector2) {
    currentCell := element.GetCell()
    currentRow := currentCell.Row
    currentColumn := currentCell.Column
    for _ = range waterElement.disperseRate {
        adjacentElement := matrix.GetElement(&Cell{Row: currentRow, Column: currentColumn})
        if adjacentElement != nil && adjacentElement.GetFamily().GetType() == Void {
            matrix.SwapElements(waterElement, adjacentElement)
            currentRow += int(movement.Y)
            currentColumn += int(movement.X)
        } else {
            return
        }
    }
}

type WaterFamily struct {
	elementType ElementType
	spread      int
	colors      map[int]rl.Color
}

func NewWaterFamily() *WaterFamily {
	return &WaterFamily{
		elementType: Water,
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

func (waterFamily *WaterFamily) GetName() string { 
	return "Water"
}

func (waterFamily *WaterFamily) SelectRandomColor() rl.Color {
	return waterFamily.colors[rand.Intn(len(waterFamily.colors)-1)+1]
}

func (waterFamily *WaterFamily) CreateElement(cell *Cell) Element {
	return &WaterElement{
		cell:   cell,
		family: waterFamily,
		color:  waterFamily.SelectRandomColor(),
        disperseRate: 3,
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
