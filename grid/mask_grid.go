package grid

import "image/color"

type MaskGrid struct {
	*BaseGrid
	mask *Mask
}

func NewMaskGrid(rowLen, colLen int, mask *Mask) *MaskGrid {

	grid := NewGrid(rowLen, colLen)
	SetCellsWithMask(grid, mask)
	grid.ConfigureCells()
	return &MaskGrid{
		grid,
		mask,
	}
}

func SetCellsWithMask(g *BaseGrid, mask *Mask) {
	cells := make([][]*Cell, g.RowLen)

	for i := 0; i < g.RowLen; i++ {
		cells[i] = make([]*Cell, g.ColLen)
	}

	for i := 0; i < g.RowLen; i++ {
		for j := 0; j < g.ColLen; j++ {

			if mask.bitMask[i][j] {
				cells[i][j] = NewCell(i, j)
			}
		}
	}

	g.SetCells(cells)
}

func (mg *MaskGrid) GetRandomCell() *Cell {
	row, col := mg.mask.RandomLocation()
	return mg.GetCell(row, col)
}

func (mg *MaskGrid) Size() int {
	return mg.mask.Count()
}

func (mg *MaskGrid) ContentOf(cell *Cell) string {
	return " "
}

func (mg *MaskGrid) String() string {
	return mg.ToStr(mg.ContentOf)
}

func (mg *MaskGrid) getBackGroundColor(cell *Cell) color.Color {
	return nil
}

func (mg *MaskGrid) GeneratePng(cellSize int) error {
	return mg.ToPng(cellSize, mg.getBackGroundColor)
}
