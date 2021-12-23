package day15

import (
	"fmt"
	"github.com/evanfpearson/advent-2021/pkg/advent"
)

func B(input *advent.Input) (int, error) {
	g := NewBigCaveGraph(input, 5)
	dest := fmt.Sprintf("%d,%d", len(*input)*5-1, len((*input)[0])*5-1)
	return g.ShortestPath("0,0", dest), nil
}
