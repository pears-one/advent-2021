package solutions

import (
	"errors"
	"fmt"
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"math"
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

// Day 3

func findGamma(bitList []string) (string, error) {
	numBits := len(bitList[0])
	bitCounts := make([]int, numBits)
	reportLength := len(bitList)
	for _, bits := range bitList {
		for i, bit := range bits {
			if string(bit) == "1" {
				bitCounts[i]++
			} else if string(bit) != "0" {
				return "", errors.New(fmt.Sprintf("input not binary string: %s", bits))
			}
		}
	}
	gamma := ""
	for _, c := range bitCounts {
		if c > reportLength/2 {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}
	return gamma, nil
}

func PowerConsumption(diagnosticReport *advent.Input) (int, error) {
	numBits := len((*diagnosticReport)[0])
	max := int(math.Pow(2, float64(numBits)) - 1)
	gamma, err := findGamma(*diagnosticReport)
	if err != nil {
		return 0, err
	}
	gammaInt := Btoi(gamma)
	return gammaInt * (max - gammaInt), nil
}

func LifeSupportRating(diagnosticReport *advent.Input) (int, error) {
	ogr, err := getOxygenGeneratorRating(*diagnosticReport)
	if err != nil {
		return 0, err
	}
	csr, err := getCO2ScrubberRating(*diagnosticReport)
	if err != nil {
		return 0, err
	}
	return ogr * csr, nil
}