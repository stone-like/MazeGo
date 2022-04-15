package grid

type Distances struct {
	Root  *Cell
	Cells map[*Cell]int //*Cell / distance
}

func NewDistances(c *Cell) *Distances {

	cells := make(map[*Cell]int)
	cells[c] = 0
	return &Distances{
		Root:  c,
		Cells: cells,
	}
}

func (d *Distances) ToList() []*Cell {
	temp := make([]*Cell, len(d.Cells))

	index := 0
	for cell, _ := range d.Cells {
		temp[index] = cell
		index++
	}

	return temp
}

func (d *Distances) ShortestPath(goal *Cell) *Distances {
	current := goal

	breadcrumbs := NewDistances(d.Root)
	breadcrumbs.Cells[current] = d.Cells[current]

	for current != d.Root {

		for _, link := range current.GetLinks() {
			//リンクしているところからルートに近づく数字のlinkを取ればOK
			if d.Cells[link] < d.Cells[current] {
				breadcrumbs.Cells[link] = d.Cells[link]
				current = link
				break
			}
		}
	}

	return breadcrumbs

}

//rootからのMaxCell,MaxDistを返す
func (d *Distances) Max() (*Cell, int) {
	max_dist := 0
	max_cell := d.Root

	for cell, dist := range d.Cells {
		if dist > max_dist {
			max_cell = cell
			max_dist = dist
		}
	}

	return max_cell, max_dist
}

func (d *Distances) GetLongestPathDist() *Distances {
	//ルートから一番遠いところを新しいスタートにする
	newStart, _ := d.Max()

	newDist := newStart.Distances()

	//新しいスタートから一番遠いところを得る
	goal, _ := newDist.Max()

	//新しいスタートから一番ところまでのPathを得る
	return newDist.ShortestPath(goal)

}
