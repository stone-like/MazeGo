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
	Mask      = "mask"
)

//Creator
const (
	BinaryTree     = "binaryTree"
	SideWinder     = "sidewinder"
	Aldrous        = "aldrous"
	Wilsons        = "wilsons"
	HuntAndKill    = "huntandkill"
	RecurBackTrack = "recurbacktrack"
)

type Options struct {
	mask *grid.Mask
}

type Option func(*Options)

func MaskOp(mask *grid.Mask) Option {
	return func(args *Options) {
		args.mask = mask
	}
}

func NewOptions(setters ...Option) *Options {
	option := &Options{}

	for _, setter := range setters {
		setter(option)
	}

	return option
}

func CreateGrid(gridType string, rowLen, colLen int, option *Options) grid.Grid {
	switch gridType {
	case Simple:
		return grid.NewSimpleGrid(rowLen, colLen)
	case Distances:
		return grid.NewDistanceGrid(rowLen, colLen)
	case Color:
		return grid.NewColorGrid(rowLen, colLen)
	case Mask:
		if option.mask == nil {
			return grid.NewSimpleGrid(rowLen, colLen)
		}
		return grid.NewMaskGrid(rowLen, colLen, option.mask)
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
	case Aldrous:
		return creator.OnAldousBroder
	case Wilsons:
		return creator.OnWilsons
	case HuntAndKill:
		return creator.OnHuntAndKill
	case RecurBackTrack:
		return creator.OnRecurBackTrack
	default:
		return creator.OnBinaryTree
	}
}

func Construct(rowLen, colLen int, gridType, creatorType string, option *Options) grid.Grid {
	grid := CreateGrid(gridType, rowLen, colLen, option)
	creator := GetCreator(creatorType)

	creator(grid)

	return grid
}
