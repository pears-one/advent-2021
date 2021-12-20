package day11

import "github.com/evanfpearson/advent-2021/pkg/advent"

func A(input *advent.Input) (int, error) {
	grid := parseOctopusGrid(input)
	for i := 0; i < 100; i++ {
		grid.Step()
	}
	return grid.totalFlashCount, nil
}
