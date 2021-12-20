package day2

import (
	"strconv"
	"strings"
)

const (
	up      = "up"
	down    = "down"
	forward = "forward"
)

type Instruction struct {
	Direction string
	Distance  int
}

func parseInstruction(line string) (Instruction, error) {
	s := strings.SplitN(line, " ", 2)
	dist, err := strconv.Atoi(s[1])
	if err != nil {
		return Instruction{}, err
	}
	return Instruction{
		Direction: s[0],
		Distance:  dist,
	}, nil
}
