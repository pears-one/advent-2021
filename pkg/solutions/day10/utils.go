package day10

import (
	"errors"
	"fmt"
	"github.com/evanfpearson/advent-2021/pkg/utils"
	"sort"
)

var (
	closeByOpens = map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
		'<': '>',
	}
	scoreByBracketA = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	scoreByBracketB = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

type BracketLinter struct {
	rightsByLefts     map[rune]rune
	rightBracketStack *utils.Stack
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
	l.rightBracketStack = utils.NewStack()
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

func autocomplete(s string, bracketStack *utils.Stack) string {
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
