package creator

import gr "MazeGo/grid"

var (
	emptyfilter = func(c *gr.Cell) bool {
		return len(c.Links) == 0
	}

	anyFilter = func(c *gr.Cell) bool {
		return len(c.Links) != 0
	}
)
