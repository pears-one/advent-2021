package day11

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
)

type Position struct {
	row int
	col int
}

type OctopusGrid struct {
	energy          *[10][10]int
	stage           int
	totalFlashCount int
}

func (g *OctopusGrid) onGrid(p *Position) bool {
	if p.row >= 0 && p.row < 10 && p.col >= 0 && p.col < 10 {
		return true
	}
	return false
}

func (g *OctopusGrid) GetAdjacent(p *Position) []*Position {
	var adjacentPositions []*Position
	for rowDiff := -1; rowDiff <= 1; rowDiff++ {
		for colDiff := -1; colDiff <= 1; colDiff++ {
			ap := &Position{p.row + rowDiff, p.col + colDiff}
			if g.onGrid(ap) && *ap != *p {
				adjacentPositions = append(adjacentPositions, ap)
			}
		}
	}
	return adjacentPositions
}

func (g *OctopusGrid) IncreaseEnergy(p *Position) {
	// We only want energy to increase if the octopus has not flashed
	if g.stage == 0 || g.energy[p.row][p.col] > 0 {
		g.energy[p.row][p.col]++
	}
}

func (g *OctopusGrid) Flash(p *Position) {
	g.totalFlashCount++
	g.energy[p.row][p.col] = 0
	for _, ap := range g.GetAdjacent(p) {
		g.IncreaseEnergy(ap)
	}
}

func (g *OctopusGrid) Step() {
	g.stage = 0
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			p := &Position{row, col}
			g.IncreaseEnergy(p)
		}
	}
	g.stage = 1
	iterationFlashCount := 1
	for iterationFlashCount > 0 {
		iterationFlashCount = 0
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				p := &Position{row, col}
				if g.energy[row][col] >= 10 {
					iterationFlashCount++
					g.Flash(p)
				}
			}
		}
	}
}

func (g *OctopusGrid) isSynchronised() bool {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if g.energy[row][col] != 0 {
				return false
			}
		}
	}
	return true
}

func parseOctopusGrid(input *advent.Input) *OctopusGrid {
	g := new(OctopusGrid)
	g.energy = new([10][10]int)
	for i, line := range *input {
		l := utils.RunesToInts(utils.GetRuneSlice(line))
		for j, energyLevel := range l {
			g.energy[i][j] = energyLevel
		}
	}
	return g
}
