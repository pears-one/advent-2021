package day11

import "github.com/evanfpearson/advent-2021/pkg/advent"

func B(input *advent.Input) (int, error) {
	grid := parseOctopusGrid(input)
	i := 0
	for {
		i++
		grid.Step()
		if grid.isSynchronised() {
			return i, nil
		}
	}
}
