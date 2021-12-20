package day4

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
	"strconv"
	"strings"
)

type Board struct {
	Rows    [5]mapset.Set
	Columns [5]mapset.Set
	All     mapset.Set
}

func EmptyBoard() *Board {
	board := new(Board)
	for i := 0; i < 5; i++ {
		board.Rows[i] = mapset.NewSet()
		board.Columns[i] = mapset.NewSet()
	}
	board.All = mapset.NewSet()
	return board
}

type BingoGame struct {
	Draw      []int
	Boards    []*Board
	NumBoards int
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
	draw, err := utils.ParseIntList(rawDraw)
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
	return lastDrawn * utils.SumSet(board.All.Difference(draw))
}