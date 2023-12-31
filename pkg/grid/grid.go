package grid

import "fmt"

type Grid struct {
	G [][]rune
}

func (g Grid) GetHeight() int {
	return len(g.G)
}

func (g Grid) GetWidth() int {
	return len(g.G[0])
}

func (g Grid) GetAtXY(x, y int) rune {
	if x < 0 || y < 0 || x > len(g.G[0])-1 || y > len(g.G)-1 {
		panic("invalid coordinates!")
	}
	return g.G[y][x]
}

func (g Grid) SetAtXY(x, y int, val rune) {
	if x < 0 || y < 0 || x > len(g.G[0])-1 || y > len(g.G)-1 {
		panic("invalid coordinates!")
	}

	g.G[y][x] = val
}

func (g Grid) Print() {
	for j := 0; j < g.GetHeight(); j++ {
		for i := 0; i < g.GetWidth(); i++ {
			fmt.Print(string(g.G[j][i]))
		}
		fmt.Println()
	}
}

func (g Grid) countHash() int {
	count := 0
	for _, row := range g.G {
		for _, ch := range row {
			if ch == 'O' {
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
