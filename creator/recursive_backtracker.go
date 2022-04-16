package creator

import (
	gr "MazeGo/grid"
	"MazeGo/util"
)

func OnRecurBackTrack(grid gr.Grid) gr.Grid {

	first := grid.GetRandomCell()
	stack := []*gr.Cell{}
	stack = append(stack, first)

	for len(stack) != 0 {
		current := stack[len(stack)-1]

		//unvisitedなneighborをゲット
		neighbors := current.GetFilteredNeighbors(emptyfilter)

		if len(neighbors) == 0 {
			stack = stack[:len(stack)-1]
			continue
		}

		index := util.CreateRandNum(len(neighbors))
		neighbor := neighbors[index]
		current.Link(neighbor, true)
		stack = append(stack, neighbor)
	}

	return grid
}
