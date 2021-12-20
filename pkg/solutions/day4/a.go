package day4

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
)

func A(input *advent.Input) (int, error) {
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
