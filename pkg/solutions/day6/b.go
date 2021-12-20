package day6

import "github.com/evanfpearson/advent-2021/pkg/advent"

func B(input *advent.Input) (int, error) {
	population, err := parsePopulation(input)
	if err != nil {
		return 0, err
	}
	population.After(256)
	return population.Size(), nil
}
