package day7

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
)

func A(input *advent.Input) (int, error) {
	var cp CrabPositions
	cp, err := utils.ParseIntList((*input)[0], ",")
	if err != nil {
		return 0, err
	}
	return cp.FuelUsageToPosition(cp.FindOptimalPosition(), LinearUsage), nil
}
