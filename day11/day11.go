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
func NewGrid(g [][]rune) Grid {
	return Grid{g}
}

func (g Grid) print() {
	for i := 0; i < len(g.G); i++ {
		for j := 0; j < len(g.G[0]); j++ {
			fmt.Print(string(g.G[i][j]))
		}
		fmt.Println()
	}
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
	// grid = expand(grid)
	galaxies := locateGalaxies(grid)
	distArr := getShortestDistances(galaxies, 1, grid)

	sum := 0
	for _, dist := range distArr {
		sum += dist
	}

	fmt.Println("Part 1:", sum)

	part2(grid, galaxies)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func part2(g Grid, galaxies []Galaxy) {

	distArr := getShortestDistances(galaxies, 1e6-1, g)

	sum := 0

	for _, dist := range distArr {
		sum += dist
	}

	fmt.Println("Part 2", sum)

}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func getShortestDistances(galaxies []Galaxy, expansionMultiplier int, g Grid) []int {
	emptyRows := findEmptyRows(g)
	emptyCols := findEmptyCols(g)

	distances := []int{}
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			newDist := abs(galaxies[j].x-galaxies[i].x) + getNumOfEmptyColsBetween(galaxies[i], galaxies[j], emptyCols)*expansionMultiplier + abs(galaxies[j].y-galaxies[i].y) + getNumOfEmptyRowsBetween(galaxies[i], galaxies[j], emptyRows)*expansionMultiplier

			distances = append(distances, newDist)
		}
	}

	return distances
}

func getNumOfEmptyRowsBetween(g1, g2 Galaxy, emptyRows []int) int {
	count := 0

	if g1.y < g2.y {
		for i := g1.y; i < g2.y; i++ {
			if arrayContains(emptyRows, i) {
				count++
			}
		}
	} else {
		for i := g2.y; i < g1.y; i++ {
			if arrayContains(emptyRows, i) {
				count++
			}
		}
	}

	return count
}

func getNumOfEmptyColsBetween(g1, g2 Galaxy, emptyCols []int) int {
	count := 0

	if g1.x < g2.x {
		for i := g1.x; i < g2.x; i++ {
			if arrayContains(emptyCols, i) {
				count++
			}
		}
	} else {
		for i := g2.x; i < g1.x; i++ {
			if arrayContains(emptyCols, i) {
				count++
			}
		}
	}

	return count
}

func arrayContains(arr []int, el int) bool {
	for _, arrEl := range arr {
		if arrEl == el {
			return true
		}
	}
	return false
}

func locateGalaxies(g Grid) []Galaxy {
	galaxies := []Galaxy{}
	for i, row := range g.G {
		for j, ch := range row {
			if ch == '#' {
				galaxies = append(galaxies, Galaxy{j, i, len(galaxies) + 1})
			}
		}
	}

	return galaxies

}

type Galaxy struct {
	x, y int
	ord  int
}

func expand(g Grid) Grid {
	emptyRows := findEmptyRows(g)
	emptyCols := findEmptyCols(g)

	g = g.expandRows(emptyRows)
	g = g.expandCols(emptyCols)

	return g
}

func (g Grid) expandCols(emptyCols []int) Grid {
	for i, index := range emptyCols {
		g = g.insertColAtIndex(index + i)
	}

	return g
}

func (g Grid) insertColAtIndex(index int) Grid {
	for idx, row := range g.G {
		row = append(row[:index+1], row[index:]...)
		row[index] = '.'

		g.G[idx] = row
	}

	return g
}

func (g Grid) expandRows(emptyRows []int) Grid {
	for i, index := range emptyRows {
		g = g.insertRowAtIndex(index + i)
	}
	return g
}

func (g Grid) insertRowAtIndex(index int) Grid {
	newEmptyRow := []rune{}

	for i := 0; i < len(g.G[0]); i++ {
		newEmptyRow = append(newEmptyRow, '.')
	}
	g.G = append(g.G[:index+1], g.G[index:]...)
	g.G[index] = newEmptyRow

	return g
}

func findEmptyRows(g Grid) []int {
	emptyRowIndices := []int{}
	for idx, row := range g.G {
		isEmpty := true

		for _, ch := range row {
			if ch != '.' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyRowIndices = append(emptyRowIndices, idx)
		}
	}

	return emptyRowIndices
}

func findEmptyCols(g Grid) []int {
	emptyColIndices := []int{}

	for i := 0; i < len(g.G[0]); i++ {
		isEmpty := true
		for _, row := range g.G {
			if row[i] != '.' {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			emptyColIndices = append(emptyColIndices, i)
		}
	}

	return emptyColIndices
}
