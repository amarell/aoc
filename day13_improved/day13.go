package main

import (
	"aoc/pkg/file"
	"fmt"
	"path/filepath"
)

func main() {
	abs, _ := filepath.Abs("./input.txt")
	lines, _ := file.ReadInput(abs)

	patterns := getPatterns(lines)

	s1 := 0
	s2 := 0

	for _, pattern := range patterns {
		s1 += calculateScore(pattern, true)
		s2 += calculateScore(pattern, false)
	}

	fmt.Println("Part 1:", s1)
	fmt.Println("Part 2:", s2)
}

func calculateScore(pattern Pattern, part1 bool) int {
	_, hIndex := isHorizontalReflection(pattern, part1)
	isVertical, vIndex := isVerticalReflection(pattern, part1)

	if isVertical {
		return vIndex
	} else {
		return 100 * hIndex
	}
}

func isHorizontalReflection(pattern Pattern, part1 bool) (bool, int) {
	lines := getColumnsAsLines(pattern)

	for i := 1; i < len(lines[0]); i++ {
		wrongsForIndex := []int{}
		allLinesReflective := true
		for _, line := range lines {
			reflective, wrongs := isLineReflective(line, i, part1)

			if !reflective && part1 {
				allLinesReflective = false
				break
			}

			wrongsForIndex = append(wrongsForIndex, wrongs)
		}

		if isSmudge(wrongsForIndex) && !part1 {
			return true, i
		}

		if allLinesReflective && part1 {
			return false, i
		}
	}
	return false, -1
}

func isVerticalReflection(pattern Pattern, part1 bool) (bool, int) {
	lineLength := len(pattern.lines[0])
	for i := 1; i < lineLength; i++ {
		allLinesReflective := true
		wrongsForIndex := []int{}

		for _, line := range pattern.lines {
			reflective, wrongs := isLineReflective(line, i, part1)

			if !reflective && part1 {
				allLinesReflective = false
				break
			}

			wrongsForIndex = append(wrongsForIndex, wrongs)
		}

		if isSmudge(wrongsForIndex) && !part1 {
			return true, i
		}

		if allLinesReflective && part1 {
			return true, i
		}
	}
	return false, -1
}

func isSmudge(wrongsArr []int) bool {
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
