package day9

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
)

type HeightMap [][]int

func (m *HeightMap) IsLowPoint(pt utils.Point) bool {
	aps := m.GetAdjacentPoints(pt)
	for _, ap := range aps {
		ph := m.HeightAt(pt)
		ah := m.HeightAt(ap)
		if ah <= ph {
			return false
		}
	}
	return true
}

func (m *HeightMap) HeightAt(pt utils.Point) int {
	return (*m)[pt.Y][pt.X]
}

func (m *HeightMap) RiskLevel(pt utils.Point) int {
	return m.HeightAt(pt) + 1
}

func (m *HeightMap) GetAdjacentPoints(pt utils.Point) []utils.Point {
	ap := []utils.Point{ // adjacent points
		{pt.X, pt.Y - 1}, {pt.X, pt.Y + 1}, {pt.X - 1, pt.Y}, {pt.X + 1, pt.Y},
	}
	validPts := make([]utils.Point, 0, 4) // filter points off the map
	for i := range ap {
		if m.IsOnMap(ap[i]) {
			validPts = append(validPts, ap[i])
		}
	}
	return validPts
}

func (m *HeightMap) IsOnMap(pt utils.Point) bool {
	return pt.X >= 0 && pt.X < m.Width() && pt.Y >= 0 && pt.Y < m.Height()
}

func (m *HeightMap) Height() int {
	return len(*m)
}

func (m *HeightMap) Width() int {
	return len((*m)[0])
}

func parseHeightMap(input *advent.Input) *HeightMap {
	m := make(HeightMap, len(*input))
	for i, line := range *input {
		m[i] = utils.RunesToInts(utils.GetRuneSlice(line))
	}
	return &m
}

type Basin mapset.Set

type BasinFinder struct {
	hm      *HeightMap
	visited [][]bool
	basins  []Basin
}

func (bf *BasinFinder) InBasin(pt utils.Point) bool {
	if bf.hm.HeightAt(pt) < 9 {
		return true
	}
	return false
}

func (bf *BasinFinder) HasVisited(pt utils.Point) bool {
	return bf.visited[pt.Y][pt.X]
}

func (bf *BasinFinder) Visit(pt utils.Point) {
	bf.visited[pt.Y][pt.X] = true
}

// FindBasinFrom recursively visits all points in a basin from some starting point
func (bf *BasinFinder) FindBasinFrom(pt utils.Point) {
	if bf.InBasin(pt) && !bf.HasVisited(pt) {
		bf.basins[len(bf.basins)-1].Add(pt)
		bf.Visit(pt)
		for _, point := range bf.hm.GetAdjacentPoints(pt) {
			bf.FindBasinFrom(point)
		}
	}
}

func (bf *BasinFinder) FindAll() {
	for row := 0; row < bf.hm.Height(); row++ {
		for col := 0; col < bf.hm.Width(); col++ {
			pt := utils.Point{X: col, Y: row}
			if bf.InBasin(pt) && !bf.HasVisited(pt) {
				bf.basins = append(bf.basins, mapset.NewSet())
				bf.FindBasinFrom(pt)
			}
		}
	}
}

func NewBasinFinder(input *advent.Input) BasinFinder {
	hm := parseHeightMap(input)
	visited := make([][]bool, hm.Height())
	for i := range *hm {
		visited[i] = make([]bool, hm.Width())
	}
	return BasinFinder{
		hm:      hm,
		visited: visited,
		basins:  []Basin{},
	}
}
