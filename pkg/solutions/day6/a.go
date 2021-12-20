package day6

import "github.com/evanfpearson/advent-2021/pkg/advent"

func A(input *advent.Input) (int, error) {
	population, err := parsePopulation(input)
	if err != nil {
		return 0, err
	}
	population.After(80)
	return population.Size(), nil
}
