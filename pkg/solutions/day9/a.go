package day9

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
)

func A(input *advent.Input) (int, error) {
	s := 0
	hm := parseHeightMap(input)
	for row := 0; row < hm.Height(); row++ {
		for col := 0; col < hm.Width(); col++ {
			pt := utils.Point{X: col, Y: row}
			if hm.IsLowPoint(pt) {
				s += hm.RiskLevel(pt)
			}
		}
	}
	return s, nil
}
