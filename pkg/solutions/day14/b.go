package day14

import "github.com/evanfpearson/advent-2021/pkg/advent"

func B(input *advent.Input) (int, error) {
	pd := parse(input)
	for i := 0; i < 40; i++ {
		pd.makeInsertions()
	}
	return pd.getScore(), nil
}
