package grid

import "image/color"

type SimpleGrid struct {
	*BaseGrid
}

func NewSimpleGrid(rowLen, colLen int) *SimpleGrid {
	grid := NewGrid(rowLen, colLen)
	grid.SetCells(CreateCells(rowLen, colLen))
	grid.ConfigureCells()
	return &SimpleGrid{
		grid,
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

func (sg *SimpleGrid) GetRandomCell() *Cell {
	return sg.CreateRandomCell()
}

func (sg *SimpleGrid) Size() int {
	return sg.GetSize()
}
