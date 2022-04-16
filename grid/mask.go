package grid

import (
	"MazeGo/files"
	"MazeGo/util"
	"bufio"
	"image/color"
	"image/png"
	"os"
)

type Mask struct {
	rowLen  int
	colLen  int
	bitMask [][]bool
}

func NewMask(rowLen, colLen int) *Mask {
	bitmask := make([][]bool, rowLen)
	for i := 0; i < rowLen; i++ {
		bitmask[i] = make([]bool, colLen)
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			bitmask[i][j] = true
		}
	}

	return &Mask{rowLen: rowLen, colLen: colLen, bitMask: bitmask}
}

func getLines(fileName string) ([]string, error) {
	path := files.GetFilePath(fileName)
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil

}

func NewMaskWithTxt(fileName string) (*Mask, error) {

	lines, err := getLines(fileName)
	if err != nil {
		return nil, err
	}

	rowLen := len(lines)
	colLen := len(lines[0])
	mask := NewMask(rowLen, colLen)

	for i, line := range lines {
		for j, r := range line {
			if r == 'X' {
				mask.bitMask[i][j] = false
			} else {
				mask.bitMask[i][j] = true
			}
		}
	}

	return mask, nil
}

func NewMaskWithPng(fileName string) (*Mask, error) {
	path := files.GetFilePath(fileName)
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, err := png.Decode(bufio.NewReader(file))

	if err != nil {
		return nil, err
	}

	colLen := img.Bounds().Max.X
	rowLen := img.Bounds().Max.Y

	black := color.RGBA{0, 0, 0, 0}

	mask := NewMask(rowLen, colLen)
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			r, g, b, a := img.At(j, i).RGBA()
			c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

			if c.R == black.R && c.G == black.G && c.B == black.B {
				mask.bitMask[i][j] = false
			} else {
				mask.bitMask[i][j] = true
			}
		}
	}

	return mask, nil

}

func (m *Mask) isBound(row, col int) bool {
	return 0 <= row && row < m.rowLen && 0 <= col && col < m.colLen
}

func (m *Mask) TurnOff(row, col int) {
	if !m.isBound(row, col) {
		return
	}

	m.bitMask[row][col] = false
}

func (m *Mask) Count() int {
	count := 0
	for i := 0; i < m.rowLen; i++ {
		for j := 0; j < m.colLen; j++ {
			if m.bitMask[i][j] {
				count++
			}
		}
	}

	return count
}

func (m *Mask) RandomLocation() (int, int) {
	for {
		row, col := util.CreateRandNum(m.rowLen), util.CreateRandNum(m.colLen)

		if m.bitMask[row][col] {
			return row, col
		}
	}
}

func (m *Mask) Row() int {
	return m.rowLen
}

func (m *Mask) Col() int {
	return m.colLen
}
