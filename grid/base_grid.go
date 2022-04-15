package grid

import (
	"MazeGo/sketch"
	"MazeGo/util"
	"fmt"
	"image"
	"image/color"
	"strings"

	"github.com/llgcode/draw2d/draw2dimg"
)

const (
	background = iota
	wall
)

type BaseGrid struct {
	RowLen int
	ColLen int
	Cells  [][]*Cell
}

func NewGrid(rowLen, colLen int) *BaseGrid {
	g := &BaseGrid{
		RowLen: rowLen,
		ColLen: colLen,
		Cells:  setCells(rowLen, colLen),
	}

	g.configureCells()

	return g

}

func (g *BaseGrid) Row() int {
	return g.RowLen
}

func (g *BaseGrid) Col() int {
	return g.ColLen
}

// func (g *BaseGrid) ContentOf(cell *Cell) string {
// 	return " "
// }

func (g *BaseGrid) DrawGrid(mode int, cell *Cell, img *image.RGBA, cellSize int, getBackGroundColor func(cell *Cell) color.Color) {
	x1 := cell.Column * cellSize
	y1 := cell.Row * cellSize
	x2 := (cell.Column + 1) * cellSize
	y2 := (cell.Row + 1) * cellSize

	//左上(0,0)から->　↓とやっていく
	//北と西はBorderだったらラインを引く
	//東と南はリンクしてなかったら線を引く
	if mode == background {
		color := getBackGroundColor(cell)
		if color != nil {
			sketch.DrawRectangle(x1, y1, x2, y2, img, color)
		}
	} else {
		if cell.Neighbors.North == nil {
			sketch.DrawLine(img, x1, y1, x2, y1, color.Black)
		}

		if cell.Neighbors.West == nil {
			sketch.DrawLine(img, x1, y1, x1, y2, color.Black)
		}

		if !cell.Linked(cell.Neighbors.East) {
			sketch.DrawLine(img, x2, y1, x2, y2, color.Black)
		}

		if !cell.Linked(cell.Neighbors.South) {
			sketch.DrawLine(img, x1, y2, x2, y2, color.Black)
		}
	}
}

func (g *BaseGrid) ToPng(cellSize int, getBackGroundColor func(cell *Cell) color.Color) error {

	imgHeight := cellSize * g.RowLen
	imgWidth := cellSize * g.ColLen

	img := sketch.NewImage(imgWidth+1, imgHeight+1)

	modes := []int{background, wall}

	for _, mode := range modes {
		for cell := range g.EachCell() {
			g.DrawGrid(mode, cell, img, cellSize, getBackGroundColor)
		}
	}

	draw2dimg.SaveToPngFile("maze.png", img)

	return nil
}

func (g *BaseGrid) ToStr(getCellContent func(cell *Cell) string) string {

	var builder strings.Builder

	builder.WriteString("+" + strings.Repeat("---+", g.ColLen) + "\n")

	top := "|"
	bottom := "+"

	getBoundary := func(cell, target *Cell, linkedStr, unlinkedStr string) string {
		if cell.Linked(target) {
			return linkedStr
		}

		return unlinkedStr
	}

	for i := 0; i < g.RowLen; i++ {

		var topBuilder strings.Builder
		var bottomBuilder strings.Builder

		topBuilder.WriteString(top)
		bottomBuilder.WriteString(bottom)

		// 横軸ごとに作っていく
		for j := 0; j < g.ColLen; j++ {

			target := g.GetCell(i, j)

			if target == nil {
				target = NewCell(-1, -1)
			}

			body := fmt.Sprintf("%3s", getCellContent(target))
			eastBoundary := getBoundary(target, target.Neighbors.East, " ", "|")

			topBuilder.WriteString(body)
			topBuilder.WriteString(eastBoundary)

			southBoundary := getBoundary(target, target.Neighbors.South, "   ", "---")
			corner := "+"

			bottomBuilder.WriteString(southBoundary)
			bottomBuilder.WriteString(corner)

		}

		builder.WriteString(topBuilder.String())
		builder.WriteString("\n")
		builder.WriteString(bottomBuilder.String())
		builder.WriteString("\n")

	}
	return builder.String()
}

func (g *BaseGrid) onBoard(r, c int) bool {
	return 0 <= r && r < g.RowLen && 0 <= c && c < g.ColLen
}

func (g *BaseGrid) GetCell(r, c int) *Cell {
	if !g.onBoard(r, c) {
		return nil
	}

	return g.Cells[r][c]
}

func (g *BaseGrid) EachCell() chan *Cell {
	c := make(chan *Cell)

	go func() {
		for i := 0; i < g.RowLen; i++ {
			for j := 0; j < g.ColLen; j++ {
				c <- g.Cells[i][j]
			}
		}

		close(c)
	}()

	return c
}

func (g *BaseGrid) configureCells() {
	for cell := range g.EachCell() {
		row, col := cell.Row, cell.Column

		cell.Neighbors.North = g.GetCell(row-1, col)
		cell.Neighbors.South = g.GetCell(row+1, col)
		cell.Neighbors.West = g.GetCell(row, col-1)
		cell.Neighbors.East = g.GetCell(row, col+1)

	}
}

func setCells(rowLen, colLen int) [][]*Cell {
	cells := make([][]*Cell, rowLen)

	for i := 0; i < rowLen; i++ {
		cells[i] = make([]*Cell, colLen)
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			cells[i][j] = NewCell(i, j)
		}
	}

	return cells
}

func (g *BaseGrid) GetRandomCell() *Cell {

	row := util.CreateRandNum(g.RowLen)
	col := util.CreateRandNum(g.ColLen)

	return g.Cells[row][col]
}

func (g *BaseGrid) Size() int {
	return g.RowLen * g.ColLen
}
