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
	return &DistanceGrid{
		NewGrid(rowLen, colLen),
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
