package creator

import (
	gr "MazeGo/grid"
	"MazeGo/util"
)

func GetRandCell(m map[*gr.Cell]struct{}) *gr.Cell {
	targetIndex := util.CreateRandNum(len(m))

	index := 0
	for cell, _ := range m {
		if targetIndex == index {
			return cell
		}

		index++
	}

	return nil

}

func Include(unvisited map[*gr.Cell]struct{}, cell *gr.Cell) bool {
	_, exists := unvisited[cell]

	return exists
}

func OnWilsons(grid gr.Grid) gr.Grid {
	unvisited := make(map[*gr.Cell]struct{})

	//スタート前に一つランダムにvisitedとする

	fillUnvisited := func() {
		index := 0
		targetIndex := util.CreateRandNum(grid.Size())
		for cell := range grid.EachCell() {

			if targetIndex != index {
				unvisited[cell] = struct{}{}
			}

			index++
		}
	}

	fillUnvisited()

	createPath := func(path []*gr.Cell, start *gr.Cell) []*gr.Cell {

		cell := start
		//pathを使ってvisitedにたどり着くまでやる
		for Include(unvisited, cell) {
			cell = cell.GetRandomNeighbor()
			position := cell.IndexOf(path)

			if position != -1 {
				//ループが検出されたら,ループの始点をpathの最後にしてcontinue
				path = path[:position+1]
				continue
			}
			path = append(path, cell)

		}

		return path
	}

	for len(unvisited) != 0 {
		cell := GetRandCell(unvisited)
		path := createPath([]*gr.Cell{cell}, cell)

		//visitedまでたどり着いたPathに下記を行う
		//createPathでpathの最後尾はvisitedになっているはずなのでlen(path)-2までがunvisitedのPathということになる

		for i := 0; i <= len(path)-2; i++ {
			path[i].Link(path[i+1], true)
			delete(unvisited, path[i])
		}

	}

	return grid

}
