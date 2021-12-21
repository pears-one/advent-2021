package day14

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"math"
	"regexp"
	"strings"
)

type InsertionRules map[string]string

func parse(input *advent.Input) (string, *InsertionRules) {
	template := regexp.MustCompile("^[A-Z]+$")
	rule := regexp.MustCompile("^[A-Z]{2} -> [A-Z]$")

	var polymerTemplate string
	insertionRules := make(InsertionRules)

	for _, line := range *input {
		if template.MatchString(line) {
			polymerTemplate = line
		}
		if rule.MatchString(line) {
			r := strings.SplitN(line, " -> ", 2)
			insertionRules[r[0]] = r[1]
		}
	}
	return polymerTemplate, &insertionRules
}

func makeInsertions(template string, rules *InsertionRules) string {
	var newTemplate string
	newTemplate += string(template[0])
	for i := 1; i < len(template); i++ {
		injection := (*rules)[string(template[i-1]) + string(template[i])]
		newTemplate += injection
		newTemplate += string(template[i])
	}
	return newTemplate
}

func getScore(template string) int {
	occurrences := make(map[rune]int)
	for _, char := range template {
		occurrences[char] += 1
	}
	max := 0
	min := math.MaxInt64
	for _, c := range occurrences {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}
	return max - min
}