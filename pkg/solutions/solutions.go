package solutions

import (
	mapset "github.com/deckarep/golang-set"
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

// Day 4

func WinningBingoCardScore(input *advent.Input) (int, error) {
	game, err := ParseBingoGame(input)
	if err != nil {
		return 0, err
	}
	drawn := mapset.NewSet()
	for _, n := range game.Draw {
		drawn.Add(n)
		for _, board := range game.Boards {
			if board.IsWon(drawn) {
				return CalculateScore(board, drawn, n), nil
			}
		}
	}
	return 0, nil
}

func LosingBingoCardScore(input *advent.Input) (int, error) {
	game, err := ParseBingoGame(input)
	if err != nil {
		return 0, err
	}
	drawn := mapset.NewSet()
	var gamesWon int
	boards := game.Boards
	for _, n := range game.Draw {
		drawn.Add(n)
		var remainingBoards []*Board
		for _, board := range boards {
			if board.IsWon(drawn) {
				gamesWon++
			} else {
				remainingBoards = append(remainingBoards, board)
			}
			if gamesWon == game.NumBoards { // this is the last board to be completed
				return CalculateScore(board, drawn, n), nil
			}
		}
		boards = remainingBoards
	}
	return 0, nil
}

// Day 5

func DangerZones(input *advent.Input) (int, error) {
	m, err := ParseVentMap(input, false)
	if err != nil {
		return 0, err
	}
	n := 0
	for _, strength := range m {
		if strength >= 2 {
			n++
		}
	}
	return n, nil
}

func DangerZonesWithDiagonals(input *advent.Input) (int, error) {
	m, err := ParseVentMap(input, true)
	if err != nil {
		return 0, err
	}
	n := 0
	for _, strength := range m {
		if strength >= 2 {
			n++
		}
	}
	return n, nil
}

// Day 6

type LanternfishPopulation []int

func parsePopulation(input *advent.Input) (*LanternfishPopulation, error) {
	population, err := parseIntList((*input)[0], ",")
	if err != nil {
		return nil, err
	}
	model := make(LanternfishPopulation, 9)
	for _, numDays := range population {
		model[numDays]++
	}
	return &model, nil
}

func (p *LanternfishPopulation) NextDay() {
	nextDay := make([]int, 9)
	for i := 0; i < 7; i++ {
		nextDay[i] = (*p)[(i+1)%7]
	}
	nextDay[8] = (*p)[0]
	nextDay[6] += (*p)[7]
	nextDay[7] += (*p)[8]
	*p = nextDay
}

func (p *LanternfishPopulation) Size() int {
	return sum(*p)
}

func (p *LanternfishPopulation) After(numDays int) {
	for i := 0; i < numDays; i++ {
		p.NextDay()
	}
}

func LanternfishPopulation80Days(input *advent.Input) (int, error) {
	population, err := parsePopulation(input)
	if err != nil {
		return 0, err
	}
	population.After(80)
	return population.Size(), nil
}


func LanternfishPopulation256Days(input *advent.Input) (int, error) {
	population, err := parsePopulation(input)
	if err != nil {
		return 0, err
	}
	population.After(256)
	return population.Size(), nil
}