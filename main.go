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
}


