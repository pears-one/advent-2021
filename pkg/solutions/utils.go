package solutions

import (
	"errors"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"math"
	"strconv"
	"strings"
)

// Day 1

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func sumSet(set mapset.Set) int {
	result := 0
	for set.Cardinality() > 0 {
		result += set.Pop().(int)
	}
	return result
}

// Day 2

type Instruction struct {
	Direction string
	Distance int
}

func parse(instruction string) (Instruction, error) {
	s := strings.SplitN(instruction, " ", 2)
	dist, err := strconv.Atoi(s[1])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{
		Direction: s[0],
		Distance:  dist,
	}, nil
}

// Day 3

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

// Day 4

type Board struct {
	Rows [5]mapset.Set
	Columns [5]mapset.Set
	All mapset.Set
}

func EmptyBoard() *Board {
	rows := [5]mapset.Set{
		mapset.NewSet(),
		mapset.NewSet(),
		mapset.NewSet(),
		mapset.NewSet(),
		mapset.NewSet(),
	}
	cols := [5]mapset.Set{
		mapset.NewSet(),
		mapset.NewSet(),
		mapset.NewSet(),
		mapset.NewSet(),
		mapset.NewSet(),
	}
	return &Board{
		Rows:    rows,
		Columns: cols,
		All:     mapset.NewSet(),
	}
}

type BingoGame struct {
	Draw []int
	Boards []*Board
	NumBoards int
}

func stringsToInts(s []string) ([]int, error) {
	ints := make([]int, len(s))
	for i := range s {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			return ints, err
		}
		ints[i] = n
	}
	return ints, nil
}

func parseIntList(lst string, sep ...string) ([]int, error) {
	s := ","
	if len(sep) > 0 {
		s = sep[0]
	}
	strs := strings.Split(lst, s)
	return stringsToInts(strs)
}

func parseBoard(rawBoard []string) (*Board, error) {
	board := EmptyBoard()
	for row, rawRow := range rawBoard {
		rowStrings := strings.Fields(rawRow)
		for col := 0; col < 5; col++ {
			n, err := strconv.Atoi(rowStrings[col])
			if err != nil {
				return nil, err
			}
			board.Rows[row].Add(n)
			board.Columns[col].Add(n)
			board.All.Add(n)
		}
	}
	return board, nil
}

func ParseBingoGame(input *advent.Input) (*BingoGame, error) {
	rawDraw := (*input)[0]
	draw, err := parseIntList(rawDraw)
	if err != nil {
		return nil, err
	}
	var boards []*Board
	var rawBoard []string
	for firstRow := 2; firstRow < len(*input); firstRow += 6 {
		rawBoard = (*input)[firstRow:firstRow+5]
		board, err := parseBoard(rawBoard)
		if err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}
	return &BingoGame{
		Draw:   draw,
		Boards: boards,
		NumBoards: len(boards),
	}, nil
}

func (b *Board) IsWon(draw mapset.Set) bool {
	for i := 0; i < 5; i++ {
		if b.Rows[i].IsSubset(draw) || b.Columns[i].IsSubset(draw) {
			return true
		}
	}
	return false
}

func CalculateScore(board *Board, draw mapset.Set, lastDrawn int) int {
	return lastDrawn * sumSet(board.All.Difference(draw))
}

// day 5

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

type Point struct {
	x int
	y int
}

type VentMap map[Point]int

func parsePoint(pointStr string) (*Point, error) {
	coords, err := parseIntList(pointStr)
	if err != nil {
		return nil, err
	}
	return &Point{x: coords[0], y: coords[1]}, nil
}

func parsePoints(inputLine string) (*Point, *Point, error) {
	points := strings.SplitN(inputLine, " -> ", 2)
	a, err := parsePoint(points[0])
	if err != nil {
		return nil, nil, err
	}
	b, err := parsePoint(points[1])
	return a, b, err
}

func verticalLineFromPoints(a, b *Point) []Point {
	numPoints := diff(a.y, b.y) + 1
	points := make([]Point, numPoints)
	s := min(a.y, b.y)
	for i := range points {
		points[i] = Point{
			x: a.x,
			y: s + i,
		}
	}
	return points
}

func horizontalLineFromPoints(a, b *Point) []Point {
	numPoints := diff(a.x, b.x) + 1
	points := make([]Point, numPoints)
	s := min(a.x, b.x)
	for i := range points {
		points[i] = Point{
			x: s + i,
			y: a.y,
		}
	}
	return points
}

func diagonalLineFromPoints(a, b *Point) []Point {
	xDirection := 1
	yDirection := 1
	if a.x > b.x {
		xDirection = -1
	}
	if a.y > b.y {
		yDirection = -1
	}
	numPoints := diff(a.x, b.x) + 1
	points := make([]Point, numPoints)
	for i := range points {
		points[i] = Point{
			x: a.x + (i * xDirection),
			y: a.y + (i * yDirection),
		}
	}
	return points
}

func lineFromPoints(a, b *Point, diag bool) []Point {
	if a.x == b.x {
		return verticalLineFromPoints(a, b)
	}
	if a.y == b.y {
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