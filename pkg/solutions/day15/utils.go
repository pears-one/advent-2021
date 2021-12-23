package day15

import (
	"container/heap"
	"fmt"
	"github.com/evanfpearson/advent-2021/pkg/advent"
	"github.com/evanfpearson/advent-2021/pkg/utils"
	"math"
)

type Edge struct {
	To     string
	Weight int
}

type DirectedGraph map[string][]Edge

func NewCaveGraph(input *advent.Input) *DirectedGraph {
	g := make(DirectedGraph)
	for row := range *input {
		for col := range (*input)[row] {
			from := fmt.Sprintf("%d,%d", row, col)
			if row > 0 {
				e := Edge{
					To:     fmt.Sprintf("%d,%d", row-1, col),
					Weight: utils.RuneToInt(rune((*input)[row-1][col])),
				}
				g[from] = append(g[from], e)
			}
			if col > 0 {
				e := Edge{
					To:     fmt.Sprintf("%d,%d", row, col-1),
					Weight: utils.RuneToInt(rune((*input)[row][col-1])),
				}
				g[from] = append(g[from], e)
			}
			if row < len(*input)-1 {
				e := Edge{
					To:     fmt.Sprintf("%d,%d", row+1, col),
					Weight: utils.RuneToInt(rune((*input)[row+1][col])),
				}
				g[from] = append(g[from], e)
			}
			if col < len((*input)[row])-1 {
				e := Edge{
					To:     fmt.Sprintf("%d,%d", row, col+1),
					Weight: utils.RuneToInt(rune((*input)[row][col+1])),
				}
				g[from] = append(g[from], e)
			}
		}
	}
	return &g
}

func NewBigCaveGraph(input *advent.Input, multiplier int) *DirectedGraph {
	g := make(DirectedGraph)
	size := len(*input)
	for row := 0; row < size*multiplier; row++ {
		for col := 0; col < size*multiplier; col++ {
			from := fmt.Sprintf("%d,%d", row, col)
			if row > 0 {
				w := utils.RuneToInt(rune((*input)[(row-1)%size][col%size]))
				e := Edge{
					To:     fmt.Sprintf("%d,%d", row-1, col),
					Weight: ((w + (row-1)/size + col/size)%10 + (w + (row-1)/size + col/size)/10)%10,
				}
				g[from] = append(g[from], e)
			}
			if col > 0 {
				w := utils.RuneToInt(rune((*input)[row%size][(col-1)%size]))
				e := Edge{
					To:     fmt.Sprintf("%d,%d", row, col-1),
					Weight: ((w + row/size + (col-1)/size)%10 + (w + row/size + (col-1)/size)/10)%10,
				}
				g[from] = append(g[from], e)
			}
			if row < size*multiplier-1 {
				w := utils.RuneToInt(rune((*input)[(row+1)%size][col%size]))
				e := Edge{
					To:     fmt.Sprintf("%d,%d", row+1, col),
					Weight: ((w + (row+1)/size + col/size)%10 + (w + (row+1)/size + col/size)/10)%10,
				}
				g[from] = append(g[from], e)
			}
			if col < size*multiplier-1 {
				w := utils.RuneToInt(rune((*input)[row%size][(col+1)%size]))
				e := Edge{
					To:     fmt.Sprintf("%d,%d", row, col+1),
					Weight: ((w + (col+1)/size + row/size)%10 + (w + (col+1)/size + row/size)/10)%10,
				}
				g[from] = append(g[from], e)
			}
		}
	}
	return &g
}

func (g *DirectedGraph) ShortestPath(src, dest string) int {
	// Init Priority Queue
	unvisitedNodes := make(UnvisitedNodeQueue, 1)
	unvisitedNodes[0] = &UnvisitedNode{src, 0}
	heap.Init(&unvisitedNodes)

	// Init Distances
	distances := make(map[string]int)
	for node := range *g {
		if node != src {
			distances[node] = math.MaxInt64
		}
	}
	for unvisitedNodes.Len() > 0 {
		currentNode := heap.Pop(&unvisitedNodes).(*UnvisitedNode)
		for _, edge := range (*g)[currentNode.name] {
			newDist := currentNode.dist + edge.Weight
			if newDist < distances[edge.To] {
				heap.Push(&unvisitedNodes, &UnvisitedNode{edge.To, newDist})
				distances[edge.To] = newDist
			}
		}
		if currentNode.name == dest {
			break
		}
	}
	return distances[dest]
}

type UnvisitedNode struct {
	name string
	dist int
}

type UnvisitedNodeQueue []*UnvisitedNode

func (q UnvisitedNodeQueue) Len() int {
	return len(q)
}

func (q UnvisitedNodeQueue) Less(i, j int) bool {
	return q[i].dist < q[j].dist
}

func (q UnvisitedNodeQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}


func (q *UnvisitedNodeQueue) Pop() interface{} {
	n := len(*q)
	item := (*q)[n-1]
	*q = (*q)[0 : n-1]
	return item
}

func (q *UnvisitedNodeQueue) Push(x interface{}) {
	item := x.(*UnvisitedNode)
	*q = append(*q, item)
}