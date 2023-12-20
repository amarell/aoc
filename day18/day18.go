package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ######......----

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

func (g Grid) countHash() int {
	count := 0
	for _, row := range g.G {
		for _, ch := range row {
			if ch == '#' {
				count++
			}
		}
	}

	return count
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

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	getArea(lines)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

const (
	RIGHT = iota
	LEFT
	UP
	DOWN
)

type Instruction struct {
	dir, val int
}

func parseDir(str string) int {
	switch str {
	case "D":
		{
			return DOWN
		}
	case "U":
		{

			return UP
		}
	case "L":
		{
			return LEFT
		}

	case "R":
		{
			return RIGHT
		}
	}

	// should never reach this
	return -1
}

func ternary(cond bool, val1 int, val2 int) int {

	if cond {

		return val1
	}

	return val2
}

func getArea(lines []string) int {
	instructions := []Instruction{}

	for _, line := range lines {
		s := strings.Split(line, " ")

		dir := parseDir(s[0])
		val := ternary(dir == LEFT || dir == UP, atoi(s[1])*-1, atoi(s[1]))
		instructions = append(instructions, Instruction{dir, val})
	}

	width := 0
	currentWidth := 0
	minWidth := 0

	for i := 0; i < len(instructions); i += 2 {
		instr := instructions[i]

		currentWidth += instr.val

		if currentWidth > width {
			width = currentWidth
		}

		if currentWidth < minWidth {
			minWidth = currentWidth
		}
	}

	height := 0
	currentHeight := 0
	minHeight := 0
	for i := 1; i < len(instructions); i += 2 {
		instr := instructions[i]

		currentHeight += instr.val

		if currentHeight > height {
			height = currentHeight
		}

		if currentHeight < minHeight {
			minHeight = currentHeight
		}

	}
	height++
	width++

	// min widht has to be either negative or 0
	// count := 0

	G := [][]rune{}

	bounds := createGrid(width+(-1)*minWidth, height+(-1)*minHeight, instructions, minWidth, minHeight)

	startPointForFlood := Point{strings.Index(bounds[0], "#") + 1, 1}

	for _, line := range bounds {
		runes := []rune{}
		for _, ch := range line {
			runes = append(runes, ch)
		}

		G = append(G, runes)
	}
	G = fillGrid(G, startPointForFlood)
	grid := Grid{G}
	// grid.print()
	fmt.Println(grid.countHash())
	return 0
}

func fillGrid(G [][]rune, startPoint Point) [][]rune {
	// flood filldfjklsjdflk
	G = floodFill(G, startPoint.x, startPoint.y)

	return G
}

func floodFill(G [][]rune, x, y int) [][]rune {
	if G[y][x] == '#' {
		return G
	}

	G[y][x] = '#'

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range dirs {
		G = floodFill(G, x+dir[0], y+dir[1])
	}

	return G
}

type Point struct {
	x, y int
}

func createGrid(width, height int, instructions []Instruction, minWidth, minHeight int) []string {
	G := []string{}
	for i := 0; i < height; i++ {
		G = append(G, strings.Repeat(".", width))
	}

	points := []Point{{0 + abs(minWidth), 0 + abs(minHeight)}}

	currentHeight := 0
	currentWidth := 0

	for _, inst := range instructions {
		if inst.dir == LEFT || inst.dir == RIGHT {
			currentWidth += inst.val
		} else {
			currentHeight += inst.val
		}
		points = append(points, Point{currentWidth + abs(minWidth), currentHeight + abs(minHeight)})
	}

	// fmt.Println(points)
	// connect the dots, but literally :dsfjdsklfjsdafl

	for idx, p := range points {
		if idx == len(points)-1 {
			break
		}

		// making a horizontal line
		nextPoint := points[idx+1]
		if p.y == nextPoint.y {
			minX := ternary(p.x < nextPoint.x, p.x, nextPoint.x)
			maxX := ternary(p.x > nextPoint.x, p.x, nextPoint.x)
			G[p.y] = G[p.y][:minX] + strings.Repeat("#", abs(p.x-nextPoint.x)+1) + G[p.y][maxX+1:]
		} else {
			// making a vertical line
			minY := ternary(p.y < nextPoint.y, p.y, nextPoint.y)
			// maxY := ternary(p.y > nextPoint.y, p.y, nextPoint.y)
			for i := 0; i < abs(nextPoint.y-p.y); i++ {
				G[i+minY+1] = G[i+minY+1][:p.x] + "#" + G[i+minY+1][p.x+1:]
			}
		}
		//
		// for _, line := range G {
		//
		// 	fmt.Println(line)
		// }
	}

	return G
}

func abs(num int) int {
	if num < 0 {
		return num * (-1)
	}
	return num
}

func atoi(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return num
}
