package solutions

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"testing"
)

func TestSolutions(t *testing.T) {
	tests := []struct{
		input *advent.Input
		expected int
		solution advent.Solution
	}{
		{
			&advent.Input{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"},
			7,
			SonarSweep,
		},
		{
			&advent.Input{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"},
			5,
			SonarSweepWindow,
		},
		{
			&advent.Input{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			150,
			Dive,
		},
		{
			&advent.Input{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			900,
			DiveWithAim,
		},
	}
	for _, test := range tests {
		out, _ := test.solution(test.input)
		if out != test.expected {
			t.Errorf("expected %d, got %d", test.expected, out)
		}
	}
}
