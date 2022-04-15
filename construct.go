package main

import (
	"MazeGo/creator"
	"MazeGo/grid"
)

//Grid
const (
	Simple    = "simple"
	Distances = "distances"
	Color     = "color"
)

//Creator
const (
	BinaryTree = "binaryTree"
	SideWinder = "sidewinder"
)

func CreateGrid(gridType string, rowLen, colLen int) grid.Grid {
	switch gridType {
	case Simple:
		return grid.NewSimpleGrid(rowLen, colLen)
	case Distances:
		return grid.NewDistanceGrid(rowLen, colLen)
	case Color:
		return grid.NewColorGrid(rowLen, colLen)
	default:
		return grid.NewSimpleGrid(rowLen, colLen)
	}
}

func GetCreator(creatorType string) creator.CreatFunc {
	switch creatorType {
	case BinaryTree:
		return creator.OnBinaryTree
	case SideWinder:
		return creator.OnSideWinder
	default:
		return creator.OnBinaryTree
	}
}

func Construct(rowLen, colLen int, gridType, creatorType string) grid.Grid {
	grid := CreateGrid(gridType, rowLen, colLen)
	creator := GetCreator(creatorType)

	creator(grid)

	return grid
}
