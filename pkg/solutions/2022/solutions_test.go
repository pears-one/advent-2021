package solutions

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/solutions/2022/day3"
	"testing"
)

func TestSolutions(t *testing.T) {
	tests := []struct {
		day      int
		expected int
		solution advent.Solution
	}{
		{
			3,
			157,
			day3.A,
		},
	}
	for _, test := range tests {
		out, _ := test.solution(&testCases[test.day-1])
		if out != test.expected {
			t.Errorf("Day %d: expected %d, got %d", test.day, test.expected, out)
		}
	}
}
