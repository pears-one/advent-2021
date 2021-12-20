package day5

import "github.com/evanfpearson/advent-2021/pkg/advent"

func A(input *advent.Input) (int, error) {
	m, err := ParseVentMap(input, false)
	if err != nil {
		return 0, err
	}
	n := 0
	for _, strength := range m {
		if strength >= 2 {
			n++
		}
	}
	return n, nil
}
