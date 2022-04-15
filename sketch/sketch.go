package sketch

import (
	"image"
	"image/color"
	"image/draw"
)

func NewImage(w, h int) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			m.Set(x, y, color.White)
		}
	}
	return m
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//ブレゼンハムのアルゴリズム
func DrawLine(img *image.RGBA, x1 int, y1 int,
	x2 int, y2 int, linecolor color.Color) {
	var step int = 0
	var dx int = AbsInt(x2 - x1)
	var dy int = AbsInt(y2 - y1)

	if dx > dy {
		if x1 > x2 {
			step = 0
			if y1 > y2 {
				step = 1
			} else {
				step = -1
			}
			x1, x2 = x2, x1 // swap
			y1 = y2
		} else {
			if y1 < y2 {
				step = 1
			} else {
				step = -1
			}
		}
		img.Set(x1, y1, linecolor)
		var s int = dx >> 1
		x1++
		for x1 <= x2 {
			s -= dy
			if s < 0 {
				s += dx
				y1 += step
			}
			img.Set(x1, y1, linecolor)
			x1++
		}
	} else {
		if y1 > y2 {
			if x1 > x2 {
				step = 1
			} else {
				step = -1
			}
			y1, y2 = y2, y1 // swap
			x1 = x2
		} else {
			if x1 < x2 {
				step = 1
			} else {
				step = -1
			}
		}
		img.Set(x1, y1, linecolor)
		var s int = dy >> 1
		y1++
		for y1 <= y2 {
			s -= dx
			if s < 0 {
				s += dy
				x1 += step
			}
			img.Set(x1, y1, linecolor)
			y1++
		}
	}
}

func DrawRectangle(x1 int, y1 int, x2 int, y2 int, img draw.Image, color color.Color) {
	var xmin, xmax, ymin, ymax int

	if x1 < x2 {
		xmin, xmax = x1, x2
	} else {
		xmin, xmax = x2, x1
	}

	if y1 < y2 {
		ymin, ymax = y1, y2
	} else {
		ymin, ymax = y2, y1
	}

	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			img.Set(x, y, color)
		}
	}
}
