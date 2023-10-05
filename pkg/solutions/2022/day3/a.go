package day3

import (
	"fmt"

	"github.com/evanfpearson/advent-2021/pkg/advent"
)

func A(input *advent.Input) (int, error) {
	rucksacks := make(chan [2][]int)
	go processInputA(input, rucksacks)
	for rucksack := range rucksacks {
		fmt.Println(rucksack)
	}
	return 157, nil
}

func score(rucksacks <-chan [2][]int) int {
	for rucksack := range rucksacks {
		
	}
}

func processInputA(input *advent.Input, rucksacks chan<- [2][]int) {
	for _, line := range *input {
		var rucksack [2][]int
		compartmentSize := len(line)/2
		rucksack[0] = convertToPriorities(line)[:compartmentSize]
		rucksack[1] = convertToPriorities(line)[compartmentSize:]
		rucksacks <- rucksack
	}
	close(rucksacks)
}

func convertToPriorities(s string) (p []int) {
	for _, ch := range s {
		p = append(p, convertToPriority(ch)) 
	}
	return p
} 

func convertToPriority(ch rune) int {
	if (ch > 64 && ch <= 90) {
		return int(ch) - 38 
	}
	if (ch > 96 && ch <= 122) {
		return int(ch) - 96
	}
	return 0
}
