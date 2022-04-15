package creator

import (
	gr "MazeGo/grid"
	"MazeGo/util"
)

func OnSideWinder(grid gr.Grid) gr.Grid {

	shouldCloseout := func(cell *gr.Cell) bool {
		if cell.Neighbors.East == nil {
			return true
		}

		index := util.CreateRandNum(2)
		return cell.Neighbors.North != nil && index == 0
	}

	for i := 0; i < grid.Row(); i++ {

		//runは->方向に堆積していく
		var run []*gr.Cell
		for j := 0; j < grid.Col(); j++ {

			cell := grid.GetCell(i, j)
			run = append(run, cell)

			if !shouldCloseout(cell) {
				cell.Link(cell.Neighbors.East, true)
				continue
			}

			index := util.CreateRandNum(len(run))
			member := run[index]

			if member.Neighbors.North != nil {
				member.Link(member.Neighbors.North, true)
			}

			run = []*gr.Cell{}

		}
	}
	return grid
}
