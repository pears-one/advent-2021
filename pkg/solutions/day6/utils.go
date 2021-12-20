package day6

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
)

type LanternfishPopulation [9]int

func parsePopulation(input *advent.Input) (*LanternfishPopulation, error) {
	population, err := utils.ParseIntList((*input)[0], ",")
	if err != nil {
		return nil, err
	}
	var model LanternfishPopulation
	for _, numDays := range population {
		model[numDays]++
	}
	return &model, nil
}

func (p *LanternfishPopulation) NextDay() {
	var nextDay LanternfishPopulation
	for i := 0; i < 9; i++ {
		nextDay[i] = (*p)[(i+1)%9]
	}
	nextDay[6] += (*p)[0]
	*p = nextDay
}

func (p *LanternfishPopulation) Size() int {
	s := 0
	for i := 0; i < 9; i++ {
		s += (*p)[i]
	}
	return s
}

func (p *LanternfishPopulation) After(numDays int) {
	for i := 0; i < numDays; i++ {
		p.NextDay()
	}
}