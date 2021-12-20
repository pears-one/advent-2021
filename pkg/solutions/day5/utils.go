package day5

import (
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
	"strings"
)

type VentMap map[utils.Point]int

func parsePoint(pointStr string) (*utils.Point, error) {
	coords, err := utils.ParseIntList(pointStr)
	if err != nil {
		return nil, err
	}
	return &utils.Point{X: coords[0], Y: coords[1]}, nil
}

func parsePoints(inputLine string) (*utils.Point, *utils.Point, error) {
	points := strings.SplitN(inputLine, " -> ", 2)
	a, err := parsePoint(points[0])
	if err != nil {
		return nil, nil, err
	}
	b, err := parsePoint(points[1])
	return a, b, err
}

func verticalLineFromPoints(a, b *utils.Point) []utils.Point {
	numPoints := utils.Diff(a.Y, b.Y) + 1
	points := make([]utils.Point, numPoints)
	s := utils.Min(a.Y, b.Y)
	for i := range points {
		points[i] = utils.Point{
			X: a.X,
			Y: s + i,
		}
	}
	return points
}

func horizontalLineFromPoints(a, b *utils.Point) []utils.Point {
	numPoints := utils.Diff(a.X, b.X) + 1
	points := make([]utils.Point, numPoints)
	s := utils.Min(a.X, b.X)
	for i := range points {
		points[i] = utils.Point{
			X: s + i,
			Y: a.Y,
		}
	}
	return points
}

func diagonalLineFromPoints(a, b *utils.Point) []utils.Point {
	xDirection := 1
	yDirection := 1
	if a.X > b.X {
		xDirection = -1
	}
	if a.Y > b.Y {
		yDirection = -1
	}
	numPoints := utils.Diff(a.X, b.X) + 1
	points := make([]utils.Point, numPoints)
	for i := range points {
		points[i] = utils.Point{
			X: a.X + (i * xDirection),
			Y: a.Y + (i * yDirection),
		}
	}
	return points
}

func lineFromPoints(a, b *utils.Point, diag bool) []utils.Point {
	if a.X == b.X {
		return verticalLineFromPoints(a, b)
	}
	if a.Y == b.Y {
		return horizontalLineFromPoints(a, b)
	}
	if diag {
		return diagonalLineFromPoints(a, b)
	}
	return nil
}

func ParseVentMap(input *advent.Input, diag bool) (VentMap, error) {
	m := make(VentMap)
	for _, inputLine := range *input {
		a, b, err := parsePoints(inputLine)
		if err != nil {
			return nil, err
		}
		line := lineFromPoints(a, b, diag)
		for _, point := range line {
			m[point]++
		}
	}
	return m, nil
}
