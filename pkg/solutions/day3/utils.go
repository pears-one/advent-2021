package day3

import (
	"errors"
	"fmt"
	"math"
)

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

func Btoi(binary string) int {
	var i float64
	l := len(binary)
	for j := 1; j <= l; j++ {
		if string(binary[l-j]) == "1" {
			i += math.Pow(2, float64(j-1))
		}
	}
	return int(i)
}

func mostCommonBit(numbers []string, bitPosition int) (string, error) {
	var count [2]int
	for _, bits := range numbers {
		if string(bits[bitPosition]) == "1" {
			count[1]++
		} else if string(bits[bitPosition]) == "0" {
			count[0]++
		} else {
			return "", errors.New("not a binary string")
		}
	}
	if count[0] > count[1] {
		return "0", nil
	}
	return "1", nil
}

func leastCommonBit(numbers []string, bitPosition int) (string, error) {
	mcb, err := mostCommonBit(numbers, bitPosition)
	if err != nil {
		return "", err
	}
	if mcb == "1" {
		return "0", nil
	}
	return "1", nil
}

func getOxygenGeneratorRating(diagnosticReport []string) (int, error) {
	l := len(diagnosticReport)
	for bitPos := 0; l > 1; bitPos++ {
		var keep []string
		mcb, err := mostCommonBit(diagnosticReport, bitPos)
		if err != nil {
			return 0, err
		}
		for _, bits := range diagnosticReport {
			if string(bits[bitPos]) == mcb {
				keep = append(keep, bits)
			}
		}
		l = len(keep)
		diagnosticReport = keep
	}
	return Btoi(diagnosticReport[0]), nil
}

func getCO2ScrubberRating(diagnosticReport []string) (int, error) {
	l := len(diagnosticReport)
	for bitPos := 0; l > 1; bitPos++ {
		var keep []string
		lcb, err := leastCommonBit(diagnosticReport, bitPos)
		if err != nil {
			return 0, err
		}
		for _, bits := range diagnosticReport {
			if string(bits[bitPos]) == lcb {
				keep = append(keep, bits)
			}
		}
		l = len(keep)
		diagnosticReport = keep
	}
	return Btoi(diagnosticReport[0]), nil
}
