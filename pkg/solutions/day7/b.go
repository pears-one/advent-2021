package day7

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
)

func B(input *advent.Input) (int, error) {
	var cp CrabPositions
	cp, err := utils.ParseIntList((*input)[0], ",")
	if err != nil {
		return 0, err
	}
	cp.Sort()
	s := cp[0]
	pu := cp.FuelUsageToPosition(s, TriangularUsage)
	for {
		nu := cp.FuelUsageToPosition(s, TriangularUsage)
		if nu > pu {
			return pu, nil
		}
		pu = nu
		s++
	}
}
