package utils

import (
	mapset "github.com/deckarep/golang-set"
	"strconv"
	"strings"
)

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func SumSet(set mapset.Set) int {
	result := 0
	for set.Cardinality() > 0 {
		result += set.Pop().(int)
	}
	return result
}

func StringsToInts(s []string) ([]int, error) {
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

func ParseIntList(lst string, sep ...string) ([]int, error) {
	s := ","
	if len(sep) > 0 {
		s = sep[0]
	}
	strs := strings.Split(lst, s)
	return StringsToInts(strs)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

type Point struct {
	X int
	Y int
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func GetRuneSlice(s string) []rune {
	r := make([]rune, len(s))
	for i, runeValue := range s {
		r[i] = runeValue
	}
	return r
}

func RuneToInt(r rune) int {
	return int(r - '0')
}

func RunesToInts(runes []rune) []int {
	ints := make([]int, len(runes))
	for i, runeValue := range runes {
		ints[i] = RuneToInt(runeValue)
	}
	return ints
}
