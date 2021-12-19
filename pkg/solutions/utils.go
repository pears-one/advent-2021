package solutions

import (
	"errors"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"math"
	"regexp"
	"sort"
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
	Distance  int
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
	Rows    [5]mapset.Set
	Columns [5]mapset.Set
	All     mapset.Set
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
	Draw      []int
	Boards    []*Board
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
		rawBoard = (*input)[firstRow : firstRow+5]
		board, err := parseBoard(rawBoard)
		if err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}
	return &BingoGame{
		Draw:      draw,
		Boards:    boards,
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

// Day 6

type LanternfishPopulation [9]int

func parsePopulation(input *advent.Input) (*LanternfishPopulation, error) {
	population, err := parseIntList((*input)[0], ",")
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

// Day 7

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// Day 8

// SegmentPattern is a set containing all of the runes that make up the
// signals which constitute a digit on a seven segment display
type SegmentPattern mapset.Set

func parseSegmentPattern(s string) SegmentPattern {
	r := mapset.NewSet()
	for _, runeValue := range s {
		r.Add(runeValue)
	}
	return r
}

// SegmentKey contains all of the possible segment patterns 0-9. They are stored
// in a map, where the key is the pattern length.
type SegmentKey map[int][]SegmentPattern

func parseSegmentKey(pattern string) SegmentKey {
	p := make(SegmentKey)
	for _, s := range strings.SplitN(pattern, " ", 10) {
		segmentPattern := parseSegmentPattern(s)
		p[segmentPattern.Cardinality()] = append(p[segmentPattern.Cardinality()], segmentPattern)
	}
	return p
}

func (p *SegmentKey) GetCipher() SegmentCipher {
	c := make(SegmentCipher)
	c[(*p)[2][0]] = 1
	c[(*p)[3][0]] = 7
	c[(*p)[4][0]] = 4
	c[(*p)[7][0]] = 8
	sixes := (*p)[6]
	for i, pattern := range sixes {
		four := c.GetPattern(4)
		if four.IsSubset(pattern) {
			c[pattern] = 9
			sixes = append(sixes[:i], sixes[i+1:]...)
		}
	}
	for i, pattern := range sixes {
		if c.GetPattern(1).IsSubset(pattern) {
			c[pattern] = 0
			sixes = append(sixes[:i], sixes[i+1:]...)
		}
	}
	c[sixes[0]] = 6
	fives := (*p)[5]
	for i, pattern := range fives {
		if c.GetPattern(1).IsSubset(pattern) {
			c[pattern] = 3
			fives = append(fives[:i], fives[i+1:]...)
		}
	}
	for i, pattern := range fives {
		if pattern.IsSubset(c.GetPattern(6)) {
			c[pattern] = 5
			fives = append(fives[:i], fives[i+1:]...)
		}
	}
	c[fives[0]] = 2
	return c
}

// Message contains the encrypted four segment patterns in the message.
type Message [4]SegmentPattern

func parseMessage(msg string) Message {
	var m Message
	for i, s := range strings.SplitN(msg, " ", 4) {
		m[i] = parseSegmentPattern(s)
	}
	return m
}

type EncryptedMessage struct {
	Key     SegmentKey
	Message Message
}

func (s *EncryptedMessage) Decrypt() DecryptedMessage {
	c := s.Key.GetCipher()
	var message DecryptedMessage
	for i, pattern := range s.Message {
		message[i] = c.Decode(pattern)
	}
	return message
}

func parseEncryptedMessage(input string) EncryptedMessage {
	splits := strings.SplitN(input, " | ", 2)
	return EncryptedMessage{
		Key:     parseSegmentKey(splits[0]),
		Message: parseMessage(splits[1]),
	}
}

type DecryptedMessage [4]int

func (m *DecryptedMessage) ToInt() int {
	s := 0
	for i := 0; i < 4; i++ {
		s += (*m)[i] * int(math.Pow(10, float64(3-i)))
	}
	return s
}

// SegmentCipher is a map from encrypted segment patterns to their decrypted
// integer values.
type SegmentCipher map[SegmentPattern]int

func (c *SegmentCipher) Decode(p SegmentPattern) int {
	for k, v := range *c {
		if k.Equal(p) {
			return v
		}
	}
	return 0
}

func (c *SegmentCipher) GetPattern(n int) SegmentPattern {
	for k, v := range *c {
		if v == n {
			return k
		}
	}
	return nil
}

// Day 9

type HeightMap [][]int

func (m *HeightMap) IsLowPoint(pt Point) bool {
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

func (m *HeightMap) HeightAt(pt Point) int {
	return (*m)[pt.y][pt.x]
}

func (m *HeightMap) RiskLevel(pt Point) int {
	return m.HeightAt(pt) + 1
}

func (m *HeightMap) GetAdjacentPoints(pt Point) []Point {
	ap := []Point{ // adjacent points
		{pt.x, pt.y - 1}, {pt.x, pt.y + 1}, {pt.x - 1, pt.y}, {pt.x + 1, pt.y},
	}
	validPts := make([]Point, 0, 4) // filter points off the map
	for i := range ap {
		if m.IsOnMap(ap[i]) {
			validPts = append(validPts, ap[i])
		}
	}
	return validPts
}

func (m *HeightMap) IsOnMap(pt Point) bool {
	return pt.x >= 0 && pt.x < m.Width() && pt.y >= 0 && pt.y < m.Height()
}

func (m *HeightMap) Height() int {
	return len(*m)
}

func (m *HeightMap) Width() int {
	return len((*m)[0])
}

func getRuneSlice(s string) []rune {
	r := make([]rune, len(s))
	for i, runeValue := range s {
		r[i] = runeValue
	}
	return r
}

func runesToInts(runes []rune) []int {
	ints := make([]int, len(runes))
	for i, runeValue := range runes {
		ints[i] = int(runeValue - '0')
	}
	return ints
}

func parseHeightMap(input *advent.Input) *HeightMap {
	m := make(HeightMap, len(*input))
	for i, line := range *input {
		m[i] = runesToInts(getRuneSlice(line))
	}
	return &m
}

type Basin mapset.Set

type BasinFinder struct {
	hm      *HeightMap
	visited [][]bool
	basins  []Basin
}

func (bf *BasinFinder) InBasin(pt Point) bool {
	if bf.hm.HeightAt(pt) < 9 {
		return true
	}
	return false
}

func (bf *BasinFinder) HasVisited(pt Point) bool {
	return bf.visited[pt.y][pt.x]
}

func (bf *BasinFinder) Visit(pt Point) {
	bf.visited[pt.y][pt.x] = true
}

// FindBasinFrom recursively visits all points in a basin from some starting point
func (bf *BasinFinder) FindBasinFrom(pt Point) {
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
			pt := Point{col, row}
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

// Day 10

type BracketLinter struct {
	rightsByLefts     map[rune]rune
	rightBracketStack *Stack
}

type LintError struct {
	syntax     bool
	incomplete bool
	pos        int
	message    string
}

func (e *LintError) Error() string {
	return e.message
}

func NewLinter(rightsByLefts map[rune]rune) *BracketLinter {
	return &BracketLinter{rightsByLefts: rightsByLefts}
}

func (l *BracketLinter) IsLeft(bracket rune) bool {
	_, ok := l.rightsByLefts[bracket]
	return ok
}

func (l *BracketLinter) Lint(s string) error {
	l.rightBracketStack = NewStack()
	for i, bracket := range s {
		if l.IsLeft(bracket) {
			l.rightBracketStack.Push(l.rightsByLefts[bracket])
			continue
		}
		if l.rightBracketStack.Len() > 0 && l.rightBracketStack.Peek() == bracket {
			l.rightBracketStack.Pop()
			continue
		}
		return &LintError{
			syntax:     true,
			incomplete: false,
			pos:        i,
			message:    fmt.Sprintf("syntax error at pos: %d", i),
		}
	}
	if l.rightBracketStack.Len() > 0 {
		return &LintError{
			syntax:     false,
			incomplete: true,
			pos:        -1,
			message:    fmt.Sprintf("line incomplete, %d chunks unclosed", l.rightBracketStack.Len()),
		}
	}
	return nil
}

func (l *BracketLinter) Autocomplete(s string) (string, error) {
	err := l.Lint(s)
	if err.(*LintError).incomplete {
		return autocomplete(s, l.rightBracketStack), nil
	}
	if err == nil {
		return s, nil
	}
	return s, errors.New("cannot autocomplete this string")
}

func autocomplete(s string, bracketStack *Stack) string {
	for bracketStack.Len() > 0 {
		s = s + string(bracketStack.Pop().(rune))
	}
	return s
}

func median(arr []int) (int, error) {
	l := len(arr)
	sort.Ints(arr)
	if l == 0 {
		return 0, errors.New("empty array, cannot find median")
	}
	if l%2 == 0 {
		return (arr[l/2] + arr[l/2-1]) / 2, nil
	}
	return arr[l/2], nil
}

// Day 11

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
		l := runesToInts(getRuneSlice(line))
		for j, energyLevel := range l {
			g.energy[i][j] = energyLevel
		}
	}
	return g
}

// Day 12

type Vertex string

func (s *Vertex) isLarge() bool {
	for _, letter := range *s {
		if int(letter) >= 97 && int(letter) <= 122 {
			return false
		}
	}
	return true
}

type Edge struct {
	X Vertex
	Y Vertex
}

type CaveGraph struct {
	V mapset.Set
	E mapset.Set
}

type Path []Vertex

func (p *Path) hasDoubleSmallVisit() bool {
	counts := make(map[Vertex]int)
	for _, v := range *p {
		if v.isLarge() {
			continue
		}
		counts[v]++
		if counts[v] == 2 {
			return true
		}
	}
	return false
}

func (p *Path) Contains(v Vertex) bool {
	for _, p := range *p {
		if p == v {
			return true
		}
	}
	return false
}

func (p *Path) IsComplete() bool {
	return (*p)[len(*p)-1] == "end"
}

func parseCaveGraph(input *advent.Input) *CaveGraph {
	cg := new(CaveGraph)
	cg.V = mapset.NewSet()
	cg.E = mapset.NewSet()
	for _, line := range *input {
		vertices := strings.SplitN(line, "-", 2)
		x, y := Vertex(vertices[0]), Vertex(vertices[1])
		cg.V.Add(x)
		cg.V.Add(y)
		cg.E.Add(Edge{X: x, Y: y})
	}
	return cg
}

func (cg *CaveGraph) GetAdjacent(v Vertex) []Vertex {
	var adj []Vertex
	for _, edge := range cg.E.ToSlice() {
		e := edge.(Edge)
		if e.X == v {
			adj = append(adj, e.Y)
		}
		if e.Y == v {
			adj = append(adj, e.X)
		}
	}
	return adj
}

func PartAIsValid(path Path, vertex Vertex) bool {
	return vertex.isLarge() || !path.Contains(vertex)
}

func PartBIsValid(path Path, vertex Vertex) bool {
	if vertex.isLarge() {
		return true
	}
	if vertex == "start" {
		return false
	}
	if path.Contains(vertex) {
		if path.hasDoubleSmallVisit() {
			return false
		}
	}
	return true
}

func (cg *CaveGraph) Walk(paths []Path, isValid func(Path, Vertex) bool) []Path {
	allComplete := true
	var validPaths []Path
	for _, path := range paths {
		if path.IsComplete() {
			validPaths = append(validPaths, path)
			continue
		}
		allComplete = false
		for _, vertex := range cg.GetAdjacent(path[len(path)-1]) {
			if !isValid(path, vertex) {
				continue
			}
			validPath := make(Path, len(path)+1)
			copy(validPath, append(path, vertex))
			validPaths = append(validPaths, validPath)
		}
	}
	if allComplete {
		return validPaths
	}
	return cg.Walk(validPaths, isValid)
}

// Day 13

type Dot [2]int

type Fold struct {
	axis int
	at   int
}

type Origami struct {
	dots  mapset.Set
	folds []Fold
}

func parseOrigami(input *advent.Input) *Origami {
	dot := regexp.MustCompile("^[0-9]+,[0-9]+$")
	fold := regexp.MustCompile("^fold along [x|y]=[0-9]+$")
	var o Origami
	o.dots = mapset.NewSet()
	for _, line := range *input {
		if dot.MatchString(line) {
			coords, _ := stringsToInts(strings.SplitN(line, ",", 2))
			o.dots.Add(Dot{coords[0], coords[1]})
			continue
		}
		if fold.MatchString(line) {
			chunks := strings.SplitN(line, " ", 3)
			spec := strings.SplitN(chunks[2], "=", 2)
			at, _ := strconv.Atoi(spec[1])
			axis := map[string]int{"x": 0, "y": 1}[spec[0]]
			o.folds = append(o.folds, Fold{axis, at})
		}
	}
	return &o
}

func (o *Origami) Fold(n int) {
	if n < 0 || n >= len(o.folds) {
		return
	}
	fold := o.folds[n]
	for _, d := range o.dots.ToSlice() {
		dot := d.(Dot)
		if dot[fold.axis] > fold.at {
			newDot := dot
			newDot[fold.axis] = 2*fold.at - dot[fold.axis]
			o.dots.Add(newDot)
			o.dots.Remove(dot)
		}
	}
}

func (o *Origami) Print() {
	width := math.MaxInt64
	height := math.MaxInt64
	for _, fold := range o.folds {
		if fold.axis == 0 && fold.at < width {
			width = fold.at
		}
		if fold.axis == 1 && fold.at < height {
			height = fold.at
		}
	}
	pretty := make([][]string, height)
	for row := range pretty {
		for col := 0; col < width; col++ {
			pretty[row] = append(pretty[row], " ")
			if o.dots.Contains(Dot{col, row}) {
				pretty[row][col] = "#"
			}
		}
	}
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			fmt.Print(pretty[row][col])
		}
		fmt.Print("\n")
	}
}
