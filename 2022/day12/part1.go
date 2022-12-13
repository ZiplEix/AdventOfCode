package main

import (
	hp "container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findStart(rows []string) (int, int) {
	for i, row := range rows {
		for j, char := range row {
			if char == 'S' {
				return i, j
			}
		}
	}
	return 0, 0
}

func findEnd(rows []string) (int, int) {
	for i, row := range rows {
		for j, char := range row {
			if char == 'E' {
				return i, j
			}
		}
	}
	return 0, 0
}

type path struct {
	value int
	nodes []string
}

type minPath []path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type heap struct {
	values *minPath
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}

func (h *heap) push(p path) {
	hp.Push(h.values, p)
}

func (h *heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}

type edge struct {
	node   string
	weight int
}

type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
	// g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, weight: weight})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (g *graph) getPath(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func toStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

func cleanRows(rows []string) []string {
	for i := 0; i < len(rows); i++ {
		rows[i] = strings.ReplaceAll(rows[i], "S", "a")
		rows[i] = strings.ReplaceAll(rows[i], "E", "z")
	}
	return rows
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}
	inputString := string(input)
	rows := strings.Split(inputString, "\n")

	xS, yS := findStart(rows)
	xE, yE := findEnd(rows)

	rows = cleanRows(rows)

	graph := newGraph()
	for i, row := range rows {
		for j, _ := range row {
			current := toStr(i) + "," + toStr(j)
			if i-1 >= 0 && rows[i-1][j] <= rows[i][j]+1 {
				next := toStr(i-1) + "," + toStr(j)
				graph.addEdge(current, next, 1)
				// fmt.Println("added", current, "to", next)
			}
			if i+1 < len(rows) && rows[i+1][j] <= rows[i][j]+1 {
				next := toStr(i+1) + "," + toStr(j)
				graph.addEdge(current, next, 1)
				// fmt.Println("added", current, "to", next)
			}
			if j-1 >= 0 && rows[i][j-1] <= rows[i][j]+1 {
				next := toStr(i) + "," + toStr(j-1)
				graph.addEdge(current, next, 1)
				// fmt.Println("added", current, "to", next)
			}
			if j+1 < len(rows[i]) && rows[i][j+1] <= rows[i][j]+1 {
				next := toStr(i) + "," + toStr(j+1)
				graph.addEdge(current, next, 1)
				// fmt.Println("added", current, "to", next)
			}
		}
	}
	fmt.Println(graph.getPath(toStr(xS)+","+toStr(yS), toStr(xE)+","+toStr(yE)))
}
