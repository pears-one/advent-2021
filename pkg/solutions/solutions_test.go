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
		{
			5,
			5,
			DangerZones,
		},
		{
			5,
			12,
			DangerZonesWithDiagonals,
		},
		{
			6,
			5934,
			LanternfishPopulation80Days,
		},

		{
			6,
			26984457539,
			LanternfishPopulation256Days,
		},
	}
	for _, test := range tests {
		out, _ := test.solution(&testCases[test.day-1])
		if out != test.expected {
			t.Errorf("Day %d: expected %d, got %d", test.day, test.expected, out)
		}
	}
}
