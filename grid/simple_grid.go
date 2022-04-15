package grid

import "image/color"

type SimpleGrid struct {
	*BaseGrid
}

func NewSimpleGrid(rowLen, colLen int) *SimpleGrid {
	return &SimpleGrid{
		NewGrid(rowLen, colLen),
	}
}

func (sg *SimpleGrid) ContentOf(cell *Cell) string {
	return " "
}

func (sg *SimpleGrid) String() string {
	return sg.ToStr(sg.ContentOf)
}

func (sg *SimpleGrid) getBackGroundColor(cell *Cell) color.Color {
	return nil
}

func (sg *SimpleGrid) GeneratePng(cellSize int) error {
	return sg.ToPng(cellSize, sg.getBackGroundColor)
}
