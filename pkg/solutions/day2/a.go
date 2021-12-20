package day2

import "github.com/evanfpearson/advent-2021/pkg/advent"

func A(input *advent.Input) (int, error) {
	depth := 0
	x := 0
	for _, instruction := range *input {
		inst, err := parseInstruction(instruction)
		if err != nil {
			return 0, err
		}
		if inst.Direction == up {
			depth -= inst.Distance
		}
		if inst.Direction == down {
			depth += inst.Distance
		}
		if inst.Direction == forward {
			x += inst.Distance
		}
	}
	return depth*x, nil
}
