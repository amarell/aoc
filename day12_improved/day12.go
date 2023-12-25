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

	parsedLines := parseLines(lines)

	arrs := []int{}

	cache := make(map[string]int)

	for _, line := range parsedLines {
		arrs = append(arrs, solve(line.line, line.arrangement, cache))
	}

	p2 := []int{}

	for _, line := range parsedLines {
		arr := line.arrangement

		for i := 0; i < 4; i++ {
			arr = append(arr, line.arrangement...)
		}
		// arr = append(arr, arr..., arr..., arr..., arr...)
		// nl := line.line + "?" + line.line
		// it's crazy but this works for test input ffs
		// p2 = append(p2, arrs[idx]*pow(solve(nl, arr)/arrs[idx], 4))

		nl := strings.Repeat(line.line+"?", 5)
		nl = nl[:len(nl)-1]
		p2 = append(p2, solve(nl, arr, cache))
	}

	fmt.Println("Part 1:", sum(arrs))
	fmt.Println("Part 2:", sum(p2))

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func pow(base, times int) int {
	res := 1

	for i := 0; i < times; i++ {
		res *= base
	}

	return res
}

type ParsedLine struct {
	line        string
	arrangement []int
}

func parseLines(lines []string) []ParsedLine {

	parsedLines := []ParsedLine{}

	for _, line := range lines {

		S := strings.Split(line, " ")

		numStrs := strings.Split(S[1], ",")

		numInts := []int{}

		for _, num := range numStrs {

			numInts = append(numInts, atoi(num))
		}

		parsedLines = append(parsedLines, ParsedLine{S[0], numInts})
	}

	return parsedLines
}

func atoi(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func arrToString(arr []int) string {
	res := ""
	for _, el := range arr {
		res += fmt.Sprint(el)
	}
	return res
}

func solve(line string, arrangements []int, cache map[string]int) int {
	val, ok := cache[line+arrToString(arrangements)]

	if ok {
		return val
	}

	if len(line) == 0 {
		if len(arrangements) == 0 {
			return 1
		}
		return 0
	}

	if len(arrangements) == 0 {
		if strings.Contains(line, "#") {
			return 0
		}
		return 1
	}

	if len(line) < sum(arrangements)+len(arrangements)-1 {
		return 0
	}

	if line[0] == '.' {
		res := solve(line[1:], arrangements, cache)
		cache[line+arrToString(arrangements)] = res
		return res
	}

	if line[0] == '#' {
		for i := 0; i < arrangements[0]; i++ {
			if line[i] == '.' {
				return 0
			}
		}

		if len(line) > arrangements[0] {
			if line[arrangements[0]] == '#' {
				return 0
			}
		}

		newLine := ""

		if len(line) > arrangements[0]+1 {
			newLine = line[arrangements[0]+1:]
		}

		res := solve(newLine, arrangements[1:], cache)
		cache[line+arrToString(arrangements)] = res
		return res
	}

	res1 := solve("#"+line[1:], arrangements, cache)
	res2 := solve("."+line[1:], arrangements, cache)

	return res1 + res2
}

func sum(arr []int) int {

	sum := 0
	for _, el := range arr {
		sum += el
	}

	return sum
}
