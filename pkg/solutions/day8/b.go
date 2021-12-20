package day8

import "github.com/evanfpearson/advent-2021/pkg/advent"

func B(input *advent.Input) (int, error) {
	s := 0
	for _, line := range *input {
		encryptedMessage := parseEncryptedMessage(line)
		digits := encryptedMessage.Decrypt()
		s += digits.ToInt()
	}
	return s, nil
}
