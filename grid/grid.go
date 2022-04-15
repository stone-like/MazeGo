package grid

type Grid interface {
	ContentOf(cell *Cell) string
	String() string
	GeneratePng(cellSize int) error
	EachCell() chan *Cell
	GetRandomCell() *Cell
	GetCell(r int, c int) *Cell
	Size() int
	Row() int
	Col() int
}
