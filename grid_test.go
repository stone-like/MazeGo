package main

import (
	gr "MazeGo/grid"
	"fmt"
	"testing"
)

func TestGrid(t *testing.T) {

	grid := Construct(25, 25, Color, SideWinder)

	start := grid.GetCell(grid.Row()/2, grid.Col()/2)
	distances := start.Distances()

	cg := grid.(*gr.ColorGrid)

	cg.SetDistances(distances)

	fmt.Println(grid)

	cg.GeneratePng(100)

	// err := grid.ToPng(100)
	// fmt.Println(grid)
	// fmt.Println(err)
}
