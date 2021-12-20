package day3

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"math"
)

func A(input *advent.Input) (int, error) {
	numBits := len((*input)[0])
	max := int(math.Pow(2, float64(numBits)) - 1)
	gamma, err := findGamma(*input)
	if err != nil {
		return 0, err
	}
	gammaInt := Btoi(gamma)
	return gammaInt * (max - gammaInt), nil
}
