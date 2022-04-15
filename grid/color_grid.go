package grid

import (
	"image/color"
	"math"
)

type ColorGrid struct {
	*BaseGrid
	distances *Distances
	maxDist   int
}

func NewColorGrid(rowLen, colLen int) *ColorGrid {
	return &ColorGrid{
		NewGrid(rowLen, colLen),
		nil,
		0,
	}
}

func (cg *ColorGrid) ContentOf(cell *Cell) string {
	return " "
}

func (cg *ColorGrid) String() string {
	return cg.ToStr(cg.ContentOf)
}

func (cg *ColorGrid) SetDistances(distances *Distances) {
	cg.distances = distances
	_, cg.maxDist = distances.Max()
}

//rootから一番遠い距離を基準としてintensityを決めていく
func (cg *ColorGrid) getBackGroundColor(cell *Cell) color.Color {
	distValue, exists := cg.distances.Cells[cell]
	if !exists {
		return nil
	}

	intensity := float64(cg.maxDist-distValue) / float64(cg.maxDist)
	dark := uint8(math.Round(255 * intensity))
	bright := uint8(128 + math.Round(127*intensity))

	return color.RGBA{
		dark,
		bright,
		dark,
		255,
	}

}

func (cg *ColorGrid) GeneratePng(cellSize int) error {
	return cg.ToPng(cellSize, cg.getBackGroundColor)
}
