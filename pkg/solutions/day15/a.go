package day15

import (
	"fmt"
	"github.com/evanfpearson/advent-2021/pkg/advent"
)

func A(input *advent.Input) (int, error) {
	g := NewCaveGraph(input)
	dest := fmt.Sprintf("%d,%d", len(*input)-1, len((*input)[len(*input)-1])-1)
	return g.ShortestPath("0,0", dest), nil
}
