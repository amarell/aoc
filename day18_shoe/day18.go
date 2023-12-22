package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	instructions := parseLines(lines, true)
	instructions2 := parseLines(lines, false)
	fmt.Println("Part 1: ", solve(instructions))
	fmt.Println("Part 2: ", solve(instructions2))

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func hexToDec(hex string) int {
	dec, err := strconv.ParseInt(hex, 16, 32)

	// in case of any error
	if err != nil {
		panic(err)
	}

	return int(dec)
}

func parseLines(lines []string, part1 bool) []Instruction {
	instructions := []Instruction{}
	for _, line := range lines {
		S := strings.Split(line, " ")
		dir, val := []int{0, 0}, 0

		if part1 {
			dir, val = parseDir(S[0]), atoi(S[1])
		} else {
			dir, val = parseDir2(S[2][len(S[2])-2:len(S[2])-1]), hexToDec(S[2][2:len(S[2])-2])
		}

		instructions = append(instructions, Instruction{dir, val})
	}

	return instructions
}

func parseDir2(str string) []int {
	switch str {
	case "0":
		{
			return []int{1, 0}
		}
	case "2":
		{

			return []int{-1, 0}

		}
	case "3":
		{
			return []int{0, -1}
		}
	case "1":
		{
			return []int{0, 1}
		}
	}

	panic("invalid direction")

}

func parseDir(str string) []int {
	switch str {
	case "R":
		{
			return []int{1, 0}
		}
	case "L":
		{

			return []int{-1, 0}

		}
	case "U":
		{
			return []int{0, -1}
		}
	case "D":
		{
			return []int{0, 1}
		}
	}

	panic("invalid direction")
}

func abs(num int) int {
	if num < 0 {
		return num * -1

	}
	return num
}

func atoi(str string) int {
	num, _ := strconv.Atoi(str)

	return num
}

type Point struct {
	x, y int
}

type Instruction struct {
	dir []int
	val int
}

func ternary(cond bool, val1, val2 int) int {
	if cond {

		return val1
	}
	return val2
}

func solve(instructions []Instruction) int {
	// shoelace formula + pick's theorem
	points := []Point{{0, 0}}

	for _, instruction := range instructions {
		nx := points[len(points)-1].x + instruction.dir[0]*instruction.val
		ny := points[len(points)-1].y + instruction.dir[1]*instruction.val

		points = append(points, Point{nx, ny})
	}

	A := 0

	for idx, point := range points {
		y1 := points[(idx+1)%len(points)].y
		y2 := 0

		if idx == 0 {
			y2 = points[len(points)-1].y
		} else {
			y2 = points[idx-1].y
		}

		A += point.x * (y1 - y2)
	}

	A /= 2
	P := 0

	for idx, point := range points {
		P += abs(point.x-points[(idx+1)%len(points)].x) + abs(point.y-points[(idx+1)%len(points)].y)
	}

	return P/2 + 1 + A
}
