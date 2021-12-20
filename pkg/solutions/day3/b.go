package day3

import "github.com/evanfpearson/advent-2021/pkg/advent"

func B(input *advent.Input) (int, error) {
	ogr, err := getOxygenGeneratorRating(*input)
	if err != nil {
		return 0, err
	}
	csr, err := getCO2ScrubberRating(*input)
	if err != nil {
		return 0, err
	}
	return ogr * csr, nil
}
