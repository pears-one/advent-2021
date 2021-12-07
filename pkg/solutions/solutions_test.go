package solutions

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"testing"
)

func TestSolutions(t *testing.T) {
	tests := []struct{
		day int
		expected int
		solution advent.Solution
	}{
		{
			1,
			7,
			SonarSweep,
		},
		{
			1,
			5,
			SonarSweepWindow,
		},
		{
			2,
			150,
			Dive,
		},
		{
			2,
			900,
			DiveWithAim,
		},
		{
			3,
			198,
			PowerConsumption,
		},
		{
			3,
			230,
			LifeSupportRating,
		},
		{
			4,
			4512,
			WinningBingoCardScore,
		},
		{
			4,
			1924,
			LosingBingoCardScore,
		},
	}
	for _, test := range tests {
		out, _ := test.solution(&testCases[test.day-1])
		if out != test.expected {
			t.Errorf("Day %d: expected %d, got %d", test.day, test.expected, out)
		}
	}
}
