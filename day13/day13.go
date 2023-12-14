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
		isVertical, vertIdx := isVerticalReflection(pattern)
		isHorizontal, horiIdx := isHorizontalReflection(pattern)

		if isVertical {
			sum += vertIdx
		}
		if isHorizontal {
			sum += 100 * horiIdx
		}
	}

	fmt.Println("Part 1:", sum)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func getPatterns(lines []string) []Pattern {
	patterns := []Pattern{}
	patternLines := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			patterns = append(patterns, Pattern{patternLines, false, false})
			patternLines = []string{}
		} else {
			patternLines = append(patternLines, line)
		}
	}
	patterns = append(patterns, Pattern{patternLines, false, false})
	return patterns
}

func processPatterns(patterns []Pattern) []Pattern {
	newPatterns := []Pattern{}

	return newPatterns
}

func isHorizontalReflection(pattern Pattern) (bool, int) {
	lines := getColumnsAsLines(pattern)

	for i := 1; i < len(lines[0]); i++ {
		allLinesReflective := true
		for _, line := range lines {
			if !isLineReflective(line, i) {
				allLinesReflective = false
				break
			}
		}

		if allLinesReflective {
			return true, i
		}
	}
	return false, -1
}

func isVerticalReflection(pattern Pattern) (bool, int) {
	lineLength := len(pattern.lines[0])
	for i := 1; i < lineLength; i++ {
		allLinesReflective := true
		for _, line := range pattern.lines {
			if !isLineReflective(line, i) {
				allLinesReflective = false
				break
			}
		}

		if allLinesReflective {
			return true, i
		}
	}
	return false, -1
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
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

func isLineReflective(line string, index int) bool {
	reversed := Reverse(line)
	normal := line

	if index <= len(line)/2 {
		offset := len(line) - (index * 2)

		for i := 0; i < offset; i++ {
			normal = " " + normal
		}

		for i := offset; i < index*2+offset; i++ {
			if reversed[i] != normal[i] {
				return false
			}
		}
	} else {
		offset := len(line) - (len(line)-index)*2

		for i := 0; i < offset; i++ {
			reversed = " " + reversed
		}

		for i := offset; i < (len(line)-index)*2+offset; i++ {
			if reversed[i] != normal[i] {
				return false
			}
		}

	}

	return true
}

type Pattern struct {
	lines                  []string
	reflectingVertical     bool
	reflectingHorizontally bool
}
