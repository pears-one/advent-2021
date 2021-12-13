package solutions

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"math"
	"sort"
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

// Day 7

type CrabPositions []int

func (c *CrabPositions) Sort() {
	sort.Ints(*c)
}

func (c *CrabPositions) FindOptimalPosition() int {
	c.Sort()
	return (*c)[len(*c)/2]
}

func LinearUsage(diff int) int {
	return abs(diff)
}

func TriangularUsage(diff int) int {
	return abs(diff)*(abs(diff)+1)/2
}

func (c *CrabPositions) FuelUsageToPosition(p int, usage func(int) int) int {
	s := 0
	for _, cp := range *c {
		s += usage(cp-p)
	}
	return s
}

func FuelMinUsageA(input *advent.Input) (int, error) {
	var cp CrabPositions
	cp, err := parseIntList((*input)[0], ",")
	if err != nil {
		return 0, err
	}
	return cp.FuelUsageToPosition(cp.FindOptimalPosition(), LinearUsage), nil
}

func FuelMinUsageB(input *advent.Input) (int, error) {
	var cp CrabPositions
	cp, err := parseIntList((*input)[0], ",")
	if err != nil {
		return 0, err
	}
	cp.Sort()
	s := cp[0]
	pu := cp.FuelUsageToPosition(s, TriangularUsage)
	for {
		nu := cp.FuelUsageToPosition(s, TriangularUsage)
		if nu > pu {
			return pu, nil
		}
		pu = nu
		s++
	}
}

// Day 8

func EasySevenSegments(input *advent.Input) (int, error) {
	uniques := mapset.NewSet(1, 4, 7, 8)
	s := 0
	for _, line := range *input {
		encryptedMessage := parseEncryptedMessage(line)
		digits := encryptedMessage.Decrypt()
		for i := 0; i < 4; i++ {
			if uniques.Contains(digits[i]) {
				s++
			}
		}
	}
	return s, nil
}

func HardSevenSegments(input *advent.Input) (int, error) {
	s := 0
	for _, line := range *input {
		encryptedMessage := parseEncryptedMessage(line)
		digits := encryptedMessage.Decrypt()
		s += digits.ToInt()
	}
	return s, nil
}

// Day 9

func LowPointRiskLevels(input *advent.Input) (int, error) {
	s := 0
	hm := parseHeightMap(input)
	for row := 0; row < hm.Height(); row++ {
		for col := 0; col < hm.Width(); col++ {
			pt := Point{x: col, y: row}
			if hm.IsLowPoint(pt) {
				s += hm.RiskLevel(pt)
			}
		}
	}
	return s, nil
}

func FindBasins(input *advent.Input) (int, error) {
	bf := NewBasinFinder(input)
	bf.FindAll()
	basinSizes := make([]int, len(bf.basins))
	for i, basin := range bf.basins {
		basinSizes[i] = basin.Cardinality()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2], nil
}