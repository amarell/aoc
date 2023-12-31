package main

import (
	"aoc/pkg/file"
	"aoc/pkg/grid"
	"fmt"
)

func main() {
	lines, _ := file.ReadInput("./input.txt")

	G := [][]rune{}

	for _, line := range lines {
		runes := []rune{}
		for _, ch := range line {
			runes = append(runes, ch)
		}
		G = append(G, runes)
	}

	g := grid.NewGrid(G)

	startStep := getElfLoc(g)

	fmt.Println("Part 1: ", getSteps(g, startStep))

}

func getSteps(g grid.Grid, startStep Step) int {
	steps := []Step{startStep}
	ans := make(map[Step]bool)
	seen := make(map[Point]bool)
	for len(steps) > 0 {
		p := Step{}
		steps, p = popLeft(steps)
		_, isSeen := seen[Point{p.x, p.y}]
		if !(p.x >= 0 && p.x < g.GetWidth() && p.y >= 0 && p.y < g.GetHeight() && g.G[p.y][p.x] != '#') || isSeen {
			continue
		}

		seen[Point{p.x, p.y}] = true

		if p.step%2 == 0 && p.step <= 64 {
			ans[p] = true
		}
		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for _, dir := range dirs {
			nr := dir[0] + p.x
			nc := dir[1] + p.y

			steps = append(steps, Step{nr, nc, p.step + 1})
		}
	}

	return len(ans)
}

func popLeft(arr []Step) ([]Step, Step) {
	p := arr[0]
	newArr := arr[1:]

	return newArr, p
}

type Point struct {
	x, y int
}

type Step struct {
	x, y, step int
}

func getElfLoc(g grid.Grid) Step {
	for i := 0; i < g.GetHeight(); i++ {
		for j := 0; j < g.GetWidth(); j++ {
			if g.G[i][j] == 'S' {
				return Step{j, i, 0}
			}
		}
	}
	// should never reach this
	return Step{-1, -1, -1}
}
