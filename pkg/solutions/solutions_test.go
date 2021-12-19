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
		{
			7,
			37,
			FuelMinUsageA,
		},
		{
			7,
			168,
			FuelMinUsageB,
		},
		{
			8,
			26,
			EasySevenSegments,
		},
		{
			8,
			61229,
			HardSevenSegments,
		},
		{
			9,
			15,
			LowPointRiskLevels,
		},
		{
			9,
			1134,
			FindBasins,
		},
		{
			10,
			26397,
			LintBrackets,
		},
		{
			10,
			288957,
			AutocompleteBrackets,
		},
		{
			11,
			1656,
			OctopusFlashes,
		},
		{
			11,
			195,
			OctopusSynchronise,
		},
		{
			12,
			226,
			PassagePathingA,
		},
		{
			12,
			3509,
			PassagePathingB,
		},
		{
			13,
			17,
			FirstFold,
		},
		{
			13,
			16,
			AllFolds,
		},
	}
	for _, test := range tests {
		out, _ := test.solution(&testCases[test.day-1])
		if out != test.expected {
			t.Errorf("Day %d: expected %d, got %d", test.day, test.expected, out)
		}
	}
}
