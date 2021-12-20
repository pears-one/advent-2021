package day9

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"sort"
)

func B(input *advent.Input) (int, error) {
	bf := NewBasinFinder(input)
	bf.FindAll()
	basinSizes := make([]int, len(bf.basins))
	for i, basin := range bf.basins {
		basinSizes[i] = basin.Cardinality()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2], nil
}
