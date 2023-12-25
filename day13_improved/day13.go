package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	patterns := getPatterns(lines)

	sum := 0
	for _, pattern := range patterns {
		_, hIndex := isHorizontalReflection(pattern)
		isV, vIndex := isVerticalReflection(pattern)

		if isV {
			sum += vIndex
		} else {
			sum += 100 * hIndex
		}

	}

	fmt.Println("Part 1", sum)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func isHorizontalReflection(pattern Pattern) (bool, int) {
	lines := getColumnsAsLines(pattern)

	for i := 1; i < len(lines[0]); i++ {
		wrongsForIndex := []int{}
		// allLinesReflective := true
		for _, line := range lines {
			reflective, wrongs := isLineReflective(line, i, false)
			if !reflective {
				// allLinesReflective = false
				// break
			}

			wrongsForIndex = append(wrongsForIndex, wrongs)
		}

		// fmt.Println(isSmidge(wrongsForIndex))
		if isSmidge(wrongsForIndex) {
			return true, i
		}
		//
		// if allLinesReflective {
		// 	return false, i
		// }
	}
	return false, -1
}

func isVerticalReflection(pattern Pattern) (bool, int) {
	lineLength := len(pattern.lines[0])
	for i := 1; i < lineLength; i++ {
		// allLinesReflective := true
		wrongsForIndex := []int{}
		for _, line := range pattern.lines {
			reflective, wrongs := isLineReflective(line, i, false)
			if !reflective {
				// allLinesReflective = false
				// break
			}

			wrongsForIndex = append(wrongsForIndex, wrongs)
		}

		// fmt.Println(wrongsForIndex)
		// fmt.Println(isSmidge(wrongsForIndex))

		if isSmidge(wrongsForIndex) {
			return true, i
		}

		// if allLinesReflective {
		// 	return true, i
		// }
	}
	return false, -1
}

func isSmidge(wrongsArr []int) bool {
	countOnes := 0

	for _, i := range wrongsArr {
		if i == 1 {
			countOnes++

			if countOnes > 1 {
				return false
			}
		} else if i != 0 {
			return false
		}
	}

	return countOnes == 1
}

func getColumnsAsLines(pattern Pattern) []string {
	lines := []string{}

	for i := 0; i < len(pattern.lines[0]); i++ {
		line := ""
		for j := 0; j < len(pattern.lines); j++ {
			line += string(pattern.lines[j][i])
		}
		lines = append(lines, line)
	}

	return lines
}

func isLineReflective(line string, index int, part1 bool) (bool, int) {
	countWrongs := 0
	hi := index
	lo := index - 1

	for lo >= 0 && hi < len(line) {
		if line[hi] != line[lo] {
			countWrongs++
			if part1 {
				return false, -1
			}
		}
		lo--
		hi++
	}

	if countWrongs > 0 {
		return false, countWrongs
	}

	return true, countWrongs
}

func getPatterns(lines []string) []Pattern {
	patterns := []Pattern{}
	patternLines := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			patterns = append(patterns, Pattern{patternLines})
			patternLines = []string{}
		} else {
			patternLines = append(patternLines, line)
		}
	}
	patterns = append(patterns, Pattern{patternLines})

	return patterns
}

type Pattern struct {
	lines []string
}
