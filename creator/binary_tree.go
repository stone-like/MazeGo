package creator

import (
	gr "MazeGo/grid"
	"MazeGo/util"
)

func OnBinaryTree(grid gr.Grid) gr.Grid {

	for cell := range grid.EachCell() {
		var neighbors []*gr.Cell

		if cell.Neighbors.North != nil {
			neighbors = append(neighbors, cell.Neighbors.North)
		}

		if cell.Neighbors.East != nil {
			neighbors = append(neighbors, cell.Neighbors.East)
		}

		if len(neighbors) == 0 {
			continue
		}

		index := util.CreateRandNum(len(neighbors))

		neighbor := neighbors[index]

		cell.Link(neighbor, true)
	}
	return grid
}
