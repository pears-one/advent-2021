package day13

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Dot [2]int

type Fold struct {
	axis int
	at   int
}

type Origami struct {
	dots  mapset.Set
	folds []Fold
}

func parseOrigami(input *advent.Input) *Origami {
	dot := regexp.MustCompile("^[0-9]+,[0-9]+$")
	fold := regexp.MustCompile("^fold along [x|y]=[0-9]+$")
	var o Origami
	o.dots = mapset.NewSet()
	for _, line := range *input {
		if dot.MatchString(line) {
			coords, _ := utils.StringsToInts(strings.SplitN(line, ",", 2))
			o.dots.Add(Dot{coords[0], coords[1]})
			continue
		}
		if fold.MatchString(line) {
			chunks := strings.SplitN(line, " ", 3)
			spec := strings.SplitN(chunks[2], "=", 2)
			at, _ := strconv.Atoi(spec[1])
			axis := map[string]int{"x": 0, "y": 1}[spec[0]]
			o.folds = append(o.folds, Fold{axis, at})
		}
	}
	return &o
}

func (o *Origami) Fold(n int) {
	if n < 0 || n >= len(o.folds) {
		return
	}
	fold := o.folds[n]
	for _, d := range o.dots.ToSlice() {
		dot := d.(Dot)
		if dot[fold.axis] > fold.at {
			newDot := dot
			newDot[fold.axis] = 2*fold.at - dot[fold.axis]
			o.dots.Add(newDot)
			o.dots.Remove(dot)
		}
	}
}

func (o *Origami) Print() {
	width := math.MaxInt64
	height := math.MaxInt64
	for _, fold := range o.folds {
		if fold.axis == 0 && fold.at < width {
			width = fold.at
		}
		if fold.axis == 1 && fold.at < height {
			height = fold.at
		}
	}
	pretty := make([][]string, height)
	for row := range pretty {
		for col := 0; col < width; col++ {
			pretty[row] = append(pretty[row], " ")
			if o.dots.Contains(Dot{col, row}) {
				pretty[row][col] = "#"
			}
		}
	}
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			fmt.Print(pretty[row][col])
		}
		fmt.Print("\n")
	}
}
