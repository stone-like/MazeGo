package main

import (
	gr "MazeGo/grid"
	"fmt"
	"testing"
)

func TestGrid(t *testing.T) {
	mask, _ := gr.NewMaskWithPng("maze_text.png")
	option := NewOptions(MaskOp(mask))

	grid := Construct(mask.Row(), mask.Col(), Mask, RecurBackTrack, option)

	fmt.Println(grid)
	grid.GeneratePng(5)

	// option := NewOptions()

	// grid := Construct(25, 25, Color, RecurBackTrack, option)

	// start := grid.GetCell(grid.Row()/2, grid.Col()/2)
	// distances := start.Distances()

	// cg := grid.(*gr.ColorGrid)

	// cg.SetDistances(distances)

	// fmt.Println(grid)

	// cg.GeneratePng(100)

}
