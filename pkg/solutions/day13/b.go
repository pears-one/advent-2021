package day13

import "github.com/evanfpearson/advent-2021/pkg/advent"

func B(input *advent.Input) (int, error) {
	o := parseOrigami(input)
	for n := range o.folds {
		o.Fold(n)
	}
	o.Print()
	return o.dots.Cardinality(), nil
}
