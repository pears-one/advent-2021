package day10

import "github.com/evanfpearson/advent-2021/pkg/advent"

func B(input *advent.Input) (int, error) {
	linter := NewLinter(closeByOpens)
	var scores []int
	for _, line := range *input {
		if completed, err := linter.Autocomplete(line); err == nil {
			score := 0
			for _, bracket := range completed[len(line):] {
				score *= 5
				score += scoreByBracketB[bracket]
			}
			scores = append(scores, score)
		}
	}
	return median(scores)
}
