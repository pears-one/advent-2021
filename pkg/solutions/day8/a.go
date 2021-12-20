package day8

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
)

func A(input *advent.Input) (int, error) {
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
