package main

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day1"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day10"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day11"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day12"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day13"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day2"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day3"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day4"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day5"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day6"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day7"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day8"
	"github.com/evanfpearson/advent-2021/pkg/solutions/day9"
)

func main() {
	runner := advent.NewRunner()
	runner.Run(1, day1.A)
	runner.Run(1, day1.B)
	runner.Run(2, day2.A)
	runner.Run(2, day2.B)
	runner.Run(3, day3.A)
	runner.Run(3, day3.B)
	runner.Run(4, day4.A)
	runner.Run(4, day4.B)
	runner.Run(5, day5.A)
	runner.Run(5, day5.B)
	runner.Run(6, day6.A)
	runner.Run(6, day6.B)
	runner.Run(7, day7.A)
	runner.Run(7, day7.B)
	runner.Run(8, day8.A)
	runner.Run(8, day8.B)
	runner.Run(9, day9.A)
	runner.Run(9, day9.B)
	runner.Run(10, day10.A)
	runner.Run(10, day10.B)
	runner.Run(11, day11.A)
	runner.Run(11, day11.B)
	runner.Run(12, day12.A)
	runner.Run(12, day12.B)
	runner.Run(13, day13.A)
	runner.Run(13, day13.B)
}
