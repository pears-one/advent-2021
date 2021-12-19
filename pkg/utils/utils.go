package utils

// Day 1

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
