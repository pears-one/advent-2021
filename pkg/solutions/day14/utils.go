package day14

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"math"
	"regexp"
	"strings"
)

type PolymerizationDevice struct {
	pairCount map[string]int
	charCount map[string]int
	rules     map[string]string
}

func parse(input *advent.Input) *PolymerizationDevice {
	template := regexp.MustCompile("^[A-Z]+$")
	rule := regexp.MustCompile("^[A-Z]{2} -> [A-Z]$")

	device := new(PolymerizationDevice)
	device.charCount = make(map[string]int)
	device.pairCount = make(map[string]int)
	device.rules = make(map[string]string)

	for _, line := range *input {
		if template.MatchString(line) {
			device.charCount[string(line[0])]++
			for i := 1; i < len(line); i++ {
				device.pairCount[string(line[i-1])+string(line[i])]++
				device.charCount[string(line[i])]++
			}
		}
		if rule.MatchString(line) {
			r := strings.SplitN(line, " -> ", 2)
			device.rules[r[0]] = r[1]
		}
	}
	return device
}

func (pd *PolymerizationDevice) makeInsertions() {
	newCount := make(map[string]int)
	for k, v := range pd.pairCount {
		newCount[k] = v
	}
	for pair, count := range pd.pairCount {
		newPairA := string(pair[0]) + pd.rules[pair]
		newPairB := pd.rules[pair] + string(pair[1])
		pd.charCount[pd.rules[pair]] += count
		newCount[newPairA] += count
		newCount[newPairB] += count
		newCount[pair] -= count
	}
	pd.pairCount = newCount
}

func (pd *PolymerizationDevice) getScore() int {
	max := 0
	min := math.MaxInt64
	for _, c := range pd.charCount {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}
	return max - min
}
