package solutions

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"strconv"
	"strings"
)

// Day One

func SonarSweep(input *advent.Input) (int, error) {
	numIncreases := 0
	depths, err := input.ToInt()
	if err != nil {
		return 0, err
	}
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			numIncreases++
		}
	}
	return numIncreases, nil
}

func SonarSweepWindow(input *advent.Input) (int, error) {
	numIncreases := 0
	windowSize := 3
	depths, err := input.ToInt()
	if err != nil {
		return 0, err
	}
	for i := 1; i < len(depths)-windowSize+1; i++ {
		a := depths[i:i+windowSize]
		b := depths[i-1:i+windowSize-1]
		if sum(a) > sum(b) {
			numIncreases++
		}
	}
	return numIncreases, nil
}

// Day Two

type Instruction struct {
	Direction string
	Distance int
}

func parse(instruction string) (Instruction, error) {
	s := strings.SplitN(instruction, " ", 2)
	dist, err := strconv.Atoi(s[1])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{
		Direction: s[0],
		Distance:  dist,
	}, nil
}

func Dive(input *advent.Input) (int, error) {
	depth := 0
	x := 0
	for _, instruction := range *input {
		inst, err := parse(instruction)
		if err != nil {
			return 0, err
		}
		if inst.Direction == "up" {
			depth -= inst.Distance
		}
		if inst.Direction == "down" {
			depth += inst.Distance
		}
		if inst.Direction == "forward" {
			x += inst.Distance
		}
	}
	return depth*x, nil
}

func DiveWithAim(input *advent.Input) (int, error) {
	depth := 0
	x := 0
	aim := 0
	for _, instruction := range *input {
		inst, err := parse(instruction)
		if err != nil {
			return 0, err
		}
		if inst.Direction == "up" {
			aim -= inst.Distance
		}
		if inst.Direction == "down" {
			aim += inst.Distance
		}
		if inst.Direction == "forward" {
			x += inst.Distance
			depth += aim * inst.Distance
		}
	}
	return depth*x, nil
}