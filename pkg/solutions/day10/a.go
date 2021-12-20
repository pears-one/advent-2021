package day10

import "github.com/evanfpearson/advent-2021/pkg/advent"

func A(input *advent.Input) (int, error) {
	linter := NewLinter(closeByOpens)
	s := 0
	for _, line := range *input {
		err := linter.Lint(line)
		lintErr := err.(*LintError)
		if lintErr.syntax {
			pos := err.(*LintError).pos
			s += scoreByBracketA[rune(line[pos])]
		}
	}
	return s, nil
}
