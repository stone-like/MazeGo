package grid

import (
	"image/color"
	"strconv"
)

type DistanceGrid struct {
	*BaseGrid
	distances *Distances
}

func NewDistanceGrid(rowLen, colLen int) *DistanceGrid {
	grid := NewGrid(rowLen, colLen)
	grid.SetCells(CreateCells(rowLen, colLen))
	grid.ConfigureCells()
	return &DistanceGrid{
		grid,
		nil,
	}
}

func (dg *DistanceGrid) ContentOf(cell *Cell) string {

	if dg.distances == nil {
		return " "
	}

	value, exists := dg.distances.Cells[cell]

	if !exists {
		return " "
	}

	return strconv.Itoa(value)
}

func (dg *DistanceGrid) String() string {
	return dg.ToStr(dg.ContentOf)
}

func (dg *DistanceGrid) SetDistances(distances *Distances) {
	dg.distances = distances
}

func (dg *DistanceGrid) getBackGroundColor(cell *Cell) color.Color {
	return nil
}

func (dg *DistanceGrid) GeneratePng(cellSize int) error {
	return dg.ToPng(cellSize, dg.getBackGroundColor)
}

func (dg *DistanceGrid) GetRandomCell() *Cell {
	return dg.CreateRandomCell()
}

func (dg *DistanceGrid) Size() int {
	return dg.GetSize()
}
