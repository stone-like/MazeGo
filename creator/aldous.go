package creator

import (
	gr "MazeGo/grid"
)

func OnAldousBroder(grid gr.Grid) gr.Grid {
	cell := grid.GetRandomCell()
	unvisitedNum := grid.Size() - 1

	for unvisitedNum > 0 {
		neighbor := cell.GetRandomNeighbor()

		//一回訪れたところもまた訪れることを許容している
		//一回目に訪れたときのみLink
		if len(neighbor.Links) == 0 {
			cell.Link(neighbor, true)
			unvisitedNum--
		}

		cell = neighbor
	}

	return grid

}
