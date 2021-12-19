package main

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/solutions"
)

func main() {
	runner := advent.NewRunner()
	runner.Run(1, solutions.SonarSweep)
	runner.Run(1, solutions.SonarSweepWindow)
	runner.Run(2, solutions.Dive)
	runner.Run(2, solutions.DiveWithAim)
	runner.Run(3, solutions.PowerConsumption)
	runner.Run(3, solutions.LifeSupportRating)
	runner.Run(4, solutions.WinningBingoCardScore)
	runner.Run(4, solutions.LosingBingoCardScore)
	runner.Run(5, solutions.DangerZones)
	runner.Run(5, solutions.DangerZonesWithDiagonals)
	runner.Run(6, solutions.LanternfishPopulation80Days)
	runner.Run(6, solutions.LanternfishPopulation256Days)
	runner.Run(7, solutions.FuelMinUsageA)
	runner.Run(7, solutions.FuelMinUsageB)
	runner.Run(8, solutions.EasySevenSegments)
	runner.Run(8, solutions.HardSevenSegments)
	runner.Run(9, solutions.LowPointRiskLevels)
	runner.Run(9, solutions.FindBasins)
	runner.Run(10, solutions.LintBrackets)
	runner.Run(10, solutions.AutocompleteBrackets)
	runner.Run(11, solutions.OctopusFlashes)
	runner.Run(11, solutions.OctopusSynchronise)
	runner.Run(12, solutions.PassagePathingA)
	runner.Run(12, solutions.PassagePathingB)
	runner.Run(13, solutions.FirstFold)
}


