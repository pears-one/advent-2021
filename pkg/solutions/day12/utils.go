package day12

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"strings"
)

const (
	start = "start"
	end = "end"
)

type Vertex string

func (s *Vertex) isLarge() bool {
	for _, letter := range *s {
		if int(letter) >= 97 && int(letter) <= 122 {
			return false
		}
	}
	return true
}

type Edge struct {
	X Vertex
	Y Vertex
}

type CaveGraph struct {
	V mapset.Set
	E mapset.Set
}

type Path []Vertex

func (p *Path) hasDoubleSmallVisit() bool {
	counts := make(map[Vertex]int)
	for _, v := range *p {
		if v.isLarge() {
			continue
		}
		counts[v]++
		if counts[v] == 2 {
			return true
		}
	}
	return false
}

func (p *Path) Contains(v Vertex) bool {
	for _, p := range *p {
		if p == v {
			return true
		}
	}
	return false
}

func (p *Path) IsComplete() bool {
	return (*p)[len(*p)-1] == end
}

func parseCaveGraph(input *advent.Input) *CaveGraph {
	cg := new(CaveGraph)
	cg.V = mapset.NewSet()
	cg.E = mapset.NewSet()
	for _, line := range *input {
		vertices := strings.SplitN(line, "-", 2)
		x, y := Vertex(vertices[0]), Vertex(vertices[1])
		cg.V.Add(x)
		cg.V.Add(y)
		cg.E.Add(Edge{X: x, Y: y})
	}
	return cg
}

func (cg *CaveGraph) GetAdjacent(v Vertex) []Vertex {
	var adj []Vertex
	for _, edge := range cg.E.ToSlice() {
		e := edge.(Edge)
		if e.X == v {
			adj = append(adj, e.Y)
		}
		if e.Y == v {
			adj = append(adj, e.X)
		}
	}
	return adj
}

func PartAIsValid(path Path, vertex Vertex) bool {
	return vertex.isLarge() || !path.Contains(vertex)
}

func PartBIsValid(path Path, vertex Vertex) bool {
	if vertex.isLarge() {
		return true
	}
	if vertex == "start" {
		return false
	}
	if path.Contains(vertex) {
		if path.hasDoubleSmallVisit() {
			return false
		}
	}
	return true
}

func (cg *CaveGraph) Walk(paths []Path, isValid func(Path, Vertex) bool) []Path {
	allComplete := true
	var validPaths []Path
	for _, path := range paths {
		if path.IsComplete() {
			validPaths = append(validPaths, path)
			continue
		}
		allComplete = false
		for _, vertex := range cg.GetAdjacent(path[len(path)-1]) {
			if !isValid(path, vertex) {
				continue
			}
			validPath := make(Path, len(path)+1)
			copy(validPath, append(path, vertex))
			validPaths = append(validPaths, validPath)
		}
	}
	if allComplete {
		return validPaths
	}
	return cg.Walk(validPaths, isValid)
}