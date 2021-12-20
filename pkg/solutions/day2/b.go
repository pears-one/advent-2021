package day2

import "github.com/evanfpearson/advent-2021/pkg/advent"

func B(input *advent.Input) (int, error) {
	depth := 0
	x := 0
	aim := 0
	for _, line := range *input {
		inst, err := parseInstruction(line)
		if err != nil {
			return 0, err
		}
		if inst.Direction == up {
			aim -= inst.Distance
		}
		if inst.Direction == down {
			aim += inst.Distance
		}
		if inst.Direction == forward {
			x += inst.Distance
			depth += aim * inst.Distance
		}
	}
	return depth*x, nil
}
