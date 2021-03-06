package grid

import "MazeGo/util"

type Neighbors struct {
	North *Cell
	South *Cell
	East  *Cell
	West  *Cell
}

type Cell struct {
	Row       int
	Column    int
	Links     map[*Cell]bool
	Neighbors *Neighbors
}

func NewCell(row, column int) *Cell {
	return &Cell{
		Row:       row,
		Column:    column,
		Links:     make(map[*Cell]bool),
		Neighbors: &Neighbors{},
	}
}

func (c *Cell) Distances() *Distances {
	//dist.Cellsがvisitedの役割も担う
	dist := NewDistances(c)
	queue := []*Cell{}

	queue = append(queue, c)
	for len(queue) != 0 {
		target := queue[0]
		queue = queue[1:]

		for _, cell := range target.GetLinks() {
			if _, exists := dist.Cells[cell]; exists {
				continue
			}

			dist.Cells[cell] = dist.Cells[target] + 1
			queue = append(queue, cell)

		}

	}

	return dist
}

//Linkは迷宮の壁に穴をあけることに相当する
func (c *Cell) Link(cell *Cell, biDir bool) *Cell {
	c.Links[cell] = true

	if biDir {
		cell.Link(c, false)
	}
	return c
}

func (c *Cell) UnLink(cell *Cell, biDir bool) *Cell {
	delete(c.Links, cell)
	if biDir {
		cell.UnLink(c, false)
	}

	return c
}

func (c *Cell) GetLinks() []*Cell {
	temp := make([]*Cell, len(c.Links))

	index := 0
	for key := range c.Links {
		temp[index] = key
		index++
	}

	return temp
}

func (c *Cell) Linked(cell *Cell) bool {
	_, exists := c.Links[cell]

	return exists
}

func (n *Neighbors) ToList(fn func(c *Cell) bool) []*Cell {
	var temp []*Cell
	if n.North != nil && fn(n.North) {
		temp = append(temp, n.North)
	}
	if n.South != nil && fn(n.South) {
		temp = append(temp, n.South)
	}
	if n.East != nil && fn(n.East) {
		temp = append(temp, n.East)
	}
	if n.West != nil && fn(n.West) {
		temp = append(temp, n.West)
	}
	return temp
}

func (c *Cell) GetNeighbors() []*Cell {
	return c.Neighbors.ToList(func(c *Cell) bool {
		return true
	})
}

func (c *Cell) GetFilteredNeighbors(fn func(c *Cell) bool) []*Cell {
	return c.Neighbors.ToList(fn)
}

func (c *Cell) GetRandomNeighbor() *Cell {

	list := c.GetNeighbors()
	index := util.CreateRandNum(len(list))
	return list[index]
}

func (c *Cell) IndexOf(cells []*Cell) int {
	for ind, cell := range cells {
		if c == cell {
			return ind
		}
	}

	return -1
}
