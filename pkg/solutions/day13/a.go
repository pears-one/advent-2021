package day13

import "github.com/evanfpearson/advent-2021/pkg/advent"

func A(input *advent.Input) (int, error) {
	o := parseOrigami(input)
	o.Fold(0)
	return o.dots.Cardinality(), nil
}
