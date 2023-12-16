package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
			fmt.Print(string(g.G[i][j]))
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

func (g Grid) getNumOfHash() int {
	count := 0
	for i := 0; i < g.getHeight(); i++ {
		for j := 0; j < g.getWidth(); j++ {
			if g.G[i][j] == '#' {
				count++
			}
		}
	}
	return count
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

	emptyGrid := EmptyGrid(grid.getWidth(), grid.getHeight())
	steps := []Step{}

	emptyGrid, _ = getEnergizedTiles(grid, emptyGrid, 0, 0, RIGHT, steps)
	fmt.Println("Part 1:", emptyGrid.getNumOfHash())

	A := []int{}

	// each row -> right
	for i := 0; i < grid.getHeight(); i++ {
		steps := []Step{}
		emptyGrid := EmptyGrid(grid.getWidth(), grid.getHeight())

		emptyGrid, _ = getEnergizedTiles(grid, emptyGrid, 0, i, RIGHT, steps)
		A = append(A, emptyGrid.getNumOfHash())
	}

	// each row <- left
	for i := 0; i < grid.getHeight(); i++ {
		steps := []Step{}
		emptyGrid := EmptyGrid(grid.getWidth(), grid.getHeight())

		emptyGrid, _ = getEnergizedTiles(grid, emptyGrid, grid.getWidth()-1, i, LEFT, steps)
		A = append(A, emptyGrid.getNumOfHash())
	}

	// each column -> down
	for i := 0; i < grid.getWidth(); i++ {
		steps := []Step{}
		emptyGrid := EmptyGrid(grid.getWidth(), grid.getHeight())

		emptyGrid, _ = getEnergizedTiles(grid, emptyGrid, i, 0, DOWN, steps)
		A = append(A, emptyGrid.getNumOfHash())

	}

	// each column -> down
	for i := 0; i < grid.getWidth(); i++ {
		steps := []Step{}
		emptyGrid := EmptyGrid(grid.getWidth(), grid.getHeight())

		emptyGrid, _ = getEnergizedTiles(grid, emptyGrid, i, grid.getHeight()-1, UP, steps)
		A = append(A, emptyGrid.getNumOfHash())
	}

	maxNum := A[0]

	for _, num := range A {
		if num > maxNum {
			maxNum = num
		}
	}

	fmt.Println("Part 2:", maxNum)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

const (
	LEFT = iota
	RIGHT
	DOWN
	UP
)

type Step struct {
	x, y int
	dir  int
}

func arrayContains(arr []Step, step Step) bool {
	for _, el := range arr {
		if el.x == step.x && el.y == step.y && el.dir == step.dir {
			return true
		}
	}
	return false
}

func getEnergizedTiles(g, emptyGrid Grid, x, y, direction int, steps []Step) (Grid, []Step) {
	for validCoords(x, y, g) && !arrayContains(steps, Step{x, y, direction}) {
		steps = append(steps, Step{x, y, direction})
		emptyGrid = energizeTile(x, y, emptyGrid)

		if validCoords(x, y, g) {
			val := g.getAtXY(x, y)

			// fmt.Println("(", x, y, ")", "=", string(val))

			switch val {
			case '-':
				{
					switch direction {
					case UP, DOWN:
						{

							emptyGrid, steps = getEnergizedTiles(g, emptyGrid, x, y, RIGHT, steps)
							emptyGrid, steps = getEnergizedTiles(g, emptyGrid, x, y, LEFT, steps)
							return emptyGrid, steps
						}
					}
				}
			case '|':
				{
					switch direction {
					case LEFT, RIGHT:
						{
							emptyGrid, steps = getEnergizedTiles(g, emptyGrid, x, y, UP, steps)
							emptyGrid, steps = getEnergizedTiles(g, emptyGrid, x, y, DOWN, steps)
							return emptyGrid, steps
						}

					}

				}
			case '/':
				{
					switch direction {
					case LEFT:
						{
							direction = DOWN
						}
					case RIGHT:
						{

							direction = UP
						}

					case UP:
						{
							direction = RIGHT
						}
					case DOWN:
						{
							direction = LEFT
						}
					}
				}
			case '\\':
				{
					switch direction {
					case LEFT:
						{
							direction = UP
						}

					case RIGHT:
						{
							direction = DOWN
						}

					case DOWN:
						{
							direction = RIGHT
						}
					case UP:
						{
							direction = LEFT
						}
					}
				}
			}
		}
		x, y = stepInDirection(direction, g, x, y)
	}

	return emptyGrid, steps
}

func stepInDirection(direction int, g Grid, currentX, currentY int) (int, int) {
	switch direction {
	case RIGHT:
		{
			return currentX + 1, currentY
		}
	case LEFT:
		{
			return currentX - 1, currentY
		}
	case UP:
		{
			return currentX, currentY - 1

		}
	case DOWN:
		{
			return currentX, currentY + 1
		}
	}

	return -1, -1
}

func validCoords(x, y int, g Grid) bool {
	return x >= 0 && y >= 0 && x < g.getWidth() && y < g.getHeight()
}

func energizeTile(x, y int, g Grid) Grid {
	g.setAtXY(y, x, '#')

	return g
}
