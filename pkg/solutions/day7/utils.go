package day7

import (
	"github.com/evanfpearson/advent-2021/pkg/utils"
	"sort"
)

type CrabPositions []int

func (c *CrabPositions) Sort() {
	sort.Ints(*c)
}

func (c *CrabPositions) FindOptimalPosition() int {
	c.Sort()
	return (*c)[len(*c)/2]
}

func LinearUsage(diff int) int {
	return utils.Abs(diff)
}

func TriangularUsage(diff int) int {
	return utils.Abs(diff) * (utils.Abs(diff) + 1) / 2
}

func (c *CrabPositions) FuelUsageToPosition(p int, usage func(int) int) int {
	s := 0
	for _, cp := range *c {
		s += usage(cp - p)
	}
	return s
}
