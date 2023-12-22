package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Grid struct {
	G [][]rune
}

func (g Grid) getHeight() int {
	return len(g.G)
}

func (g Grid) getWidth() int {
	return len(g.G[0])
}

func (g Grid) getAtXY(x, y int) rune {
	if x < 0 || y < 0 || x > len(g.G[0])-1 || y > len(g.G)-1 {
		panic("invalid coordinates!")
	}
	return g.G[y][x]
}

func (g Grid) setAtXY(x, y int, val rune) {
	if x < 0 || y < 0 || x > len(g.G[0])-1 || y > len(g.G)-1 {
		panic("invalid coordinates!")
	}

	g.G[y][x] = val
}

func (g Grid) print() {
	for j := 0; j < g.getHeight(); j++ {
		for i := 0; i < g.getWidth(); i++ {
			fmt.Print(string(g.G[j][i]))
		}
		fmt.Println()
	}
}

func NewGrid(g [][]rune) Grid {
	return Grid{g}
}

func EmptyGrid(width, height int) Grid {
	g := [][]rune{}

	for i := 0; i < height; i++ {
		row := []rune{}
		for j := 0; j < width; j++ {
			row = append(row, '.')
		}

		g = append(g, row)
	}

	return NewGrid(g)
}

type State struct {
	r, c   int
	dr, dc int
	indir  int
}

// An Item is something we manage in a priority queue.
type Item struct {
	heatLoss int // The value of the item; arbitrary.

	state State
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// this is a function which returns the element with the highest priority
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, heatLoss int, priority int) {
	item.heatLoss = heatLoss
	heap.Fix(pq, item.index)
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	G := [][]rune{}

	for scanner.Scan() {
		rowArr := []rune{}
		for _, b := range scanner.Text() {
			rowArr = append(rowArr, b)
		}
		G = append(G, rowArr)
	}

	grid := NewGrid(G)

	fmt.Println("Part 1: ", dijkstra(grid, true))
	// fmt.Println("Part 2: ", dijkstra(grid, false))
}

type NodeA struct {
	x, y   int
	dr, dc int
	indir  int
}

func dijkstra(g Grid, part1 bool) int {
	seen := make(map[NodeA]bool) // New empty set

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{0, State{0, 0, 0, 0, 0}, 0}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.state.r == g.getHeight()-1 && item.state.c == g.getWidth()-1 {
			return item.heatLoss
		}

		nodeA := NodeA{item.state.c, item.state.r, item.state.dr, item.state.dc, item.state.indir}

		if seen[nodeA] {
			continue
		}

		seen[nodeA] = true
		loss, state := item.heatLoss, item.state

		// {8 11 1 0 1}
		if state.indir < 3 && !(state.dr == 0 && state.dc == 0) {
			nr := state.r + state.dr
			nc := state.c + state.dc

			if nr >= 0 && nr < g.getWidth() && nc >= 0 && nc < g.getHeight() {
				heap.Push(&pq, &Item{heatLoss: loss + atoi(g.G[nc][nr]), state: State{nr, nc, state.dr, state.dc, state.indir + 1}})
			}
		}

		dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for _, dir := range dirs {
			if !(dir[0] == state.dr && dir[1] == state.dc) && !(dir[0] == -state.dr && dir[1] == -state.dc) {
				nr := state.r + dir[0]
				nc := state.c + dir[1]

				if nr >= 0 && nr < g.getWidth() && nc >= 0 && nc < g.getHeight() {
					heap.Push(&pq, &Item{heatLoss: loss + atoi(g.G[nc][nr]), state: State{nr, nc, dir[0], dir[1], 1}})
				}
			}
		}
	}
	return -1
}

func atoi(c rune) int {
	cInt, _ := strconv.Atoi(string(c))

	return cInt
}
