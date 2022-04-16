package creator

import (
	gr "MazeGo/grid"
	"MazeGo/util"
)

func OnHuntAndKill(grid gr.Grid) gr.Grid {
	current := grid.GetRandomCell()

	searchValidCell := func() (*gr.Cell, *gr.Cell) {

		//linkがanyのやつだけをfilter->visitedのやつだけをfilter
		for cell := range grid.EachCell() {
			visited_neighbors := cell.GetFilteredNeighbors(anyFilter)
			//cellがunvisitedかつ、visited_neighborが一人でもいれば
			if len(cell.Links) == 0 && len(visited_neighbors) != 0 {
				index := util.CreateRandNum(len(visited_neighbors))
				neighbor := visited_neighbors[index]
				return cell, neighbor
			}
		}

		return nil, nil
	}

	for current != nil {

		//linkがemptyのやつだけをfilter->unvisitedのやつだけをfilter
		unvisited_neighbors := current.GetFilteredNeighbors(emptyfilter)

		if len(unvisited_neighbors) != 0 {

			index := util.CreateRandNum(len(unvisited_neighbors))
			neighbor := unvisited_neighbors[index]
			current.Link(neighbor, true)
			current = neighbor
			continue
		}

		newCurrent, neighbor := searchValidCell()

		if newCurrent != nil {
			newCurrent.Link(neighbor, true)
		}

		current = newCurrent

	}

	return grid
}
