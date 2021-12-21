package day14

import "github.com/evanfpearson/advent-2021/pkg/advent"

func A(input *advent.Input) (int, error) {
	template, rules := parse(input)
	for i := 0; i < 40; i++ {
		template = makeInsertions(template, rules)
	}
	return getScore(template), nil
}
