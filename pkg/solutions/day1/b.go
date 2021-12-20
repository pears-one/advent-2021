package day1

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
)

func B(input *advent.Input) (int, error) {
	numIncreases := 0
	windowSize := 3
	depths, err := input.ToInt()
	if err != nil {
		return 0, err
	}
	for i := 1; i < len(depths)-windowSize+1; i++ {
		a := depths[i : i+windowSize]
		b := depths[i-1 : i+windowSize-1]
		if utils.Sum(a) > utils.Sum(b) {
			numIncreases++
		}
	}
	return numIncreases, nil
}
