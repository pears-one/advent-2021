package day12

import "github.com/evanfpearson/advent-2021/pkg/advent"

func B(input *advent.Input) (int, error) {
	cg := parseCaveGraph(input)
	paths := cg.Walk([]Path{{start}}, PartBIsValid)
	return len(paths), nil
}
