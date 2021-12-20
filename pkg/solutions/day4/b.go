package day4

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
)

func B(input *advent.Input) (int, error) {
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
