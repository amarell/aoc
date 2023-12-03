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
	// this is too low : 80417784
	numArr := [][]Num{}
	for _, line := range lines {
		numArr = append(numArr, parseLine(line))
	}

	fmt.Println(findSumOfEngineParts(numArr, lines))
	// i can like parse the lines => arr of begin, end indices of digits in each line
	findSumOfGearRatios(numArr, lines)
	// end result I need like an array of gear parts that has the number1 and number2

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func findSumOfGearRatios(numArr [][]Num, lines []string) int {
	gears := findGears(lines)

	gears = findNumsAdjacentToGear(numArr, gears)

	filteredGears := []Gear{}

	for _, gear := range gears {
		if len(gear.adjNums) == 2 {
			filteredGears = append(filteredGears, gear)
		}
	}

	sum := 0
	for _, g := range filteredGears {

		fmt.Println("Adding ", g.adjNums[0].val, "+", g.adjNums[1].val)
		sum += g.adjNums[0].val * g.adjNums[1].val
	}
	fmt.Println("Sum of gear ratios is ", sum)
	return sum
}

type Gear struct {
	adjNums []Num
	line    int
	index   int
}

func findNumsAdjacentToGear(numArr [][]Num, gears []Gear) []Gear {

	newGears := []Gear{}

	// line above => finds nums where the gears index is in the num's range
	// current line => finds the num right before the gear and right after the gear
	// line below => finds nums where the geras index is in the num's range

	for _, gear := range gears {
		newGear := Gear{}
		if gear.line == 0 {
			// just check current line and the line below
			for _, num := range numArr[gear.line] {
				if num.begin == gear.index+1 || num.end == gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				}
			}
			for _, num := range numArr[gear.line+1] {
				if num.begin <= gear.index && num.end >= gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				} else if num.end == gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				} else if num.begin-1 == gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				}
			}

		} else if gear.line == len(numArr) {
			// just check current line and the line above
			for _, num := range numArr[gear.line] {
				if num.begin == gear.index+1 || num.end == gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				}
			}
			for _, num := range numArr[gear.line-1] {
				if num.begin <= gear.index && num.end >= gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				} else if num.end == gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				} else if num.begin-1 == gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				}
			}
		} else {
			for _, num := range numArr[gear.line] {
				if num.begin == gear.index+1 || num.end == gear.index {
					fmt.Println(num.val)
					newGear.adjNums = append(newGear.adjNums, num)
				}
			}
			for _, num := range numArr[gear.line+1] {
				if num.begin <= gear.index && num.end >= gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				} else if num.end == gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				} else if num.begin-1 == gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				}
			}
			for _, num := range numArr[gear.line-1] {
				if num.begin <= gear.index && num.end >= gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				} else if num.end == gear.index {
					newGear.adjNums = append(newGear.adjNums, num)
				} else if num.begin-1 == gear.index {

					newGear.adjNums = append(newGear.adjNums, num)
				}
			}

		}
		newGears = append(newGears, newGear)
	}

	fmt.Println(newGears)

	return newGears
}

func findGears(lines []string) []Gear {
	gears := []Gear{}
	for lineIdx, line := range lines {
		for idx, ch := range line {
			if ch == '*' {
				gears = append(gears, Gear{[]Num{}, lineIdx, idx})
			}
		}
	}

	return gears
}

func findSumOfEngineParts(numArr [][]Num, lines []string) int {
	sum := 0
	for i := 0; i < len(numArr); i++ {
		for _, num := range numArr[i] {
			if i == 0 {
				// just check next line
				if checkRangeForSymbols(lines[i+1], num.begin-1, num.end+1) || checkRangeForSymbols(lines[i], num.begin-1, num.end+1) {
					sum += num.val
				}
			} else if i == len(numArr)-1 {
				// only check previous line
				if checkRangeForSymbols(lines[i-1], num.begin-1, num.end+1) || checkRangeForSymbols(lines[i], num.begin-1, num.end+1) {
					sum += num.val
				}
			} else {
				// check both
				if checkRangeForSymbols(lines[i-1], num.begin-1, num.end+1) || checkRangeForSymbols(lines[i+1], num.begin-1, num.end+1) || checkRangeForSymbols(lines[i], num.begin-1, num.end+1) {
					sum += num.val
				}
			}
		}
	}

	return sum
}

func checkRangeForSymbols(str string, begin, end int) bool {
	beginIdx := 0
	endIdx := len(str)

	if begin > 0 {
		beginIdx = begin
	}
	if end < len(str) {
		endIdx = end
	}

	slice := str[beginIdx:endIdx]
	for i := 0; i < len(slice); i++ {
		if isSymbol(string(slice[i])) {
			return true
		}
	}

	return false
}

func isSymbol(str string) bool {
	symbols := "0123456789."
	return strings.Index(symbols, str) == -1
}

func isDigit(str string) bool {
	digits := "0123456789"
	return strings.Index(digits, str) >= 0
}

func parseLine(line string) []Num {
	numArr := []Num{}

	for i := 0; i < len(line); i++ {
		if isDigit(string(line[i])) {
			numStr := ""
			endIndex := len(line)
			for j := i; j < len(line); j++ {
				if isDigit(string(line[j])) {

					numStr += string(line[j])
				} else {
					endIndex = j
					break
				}
			}

			numVal, err := strconv.Atoi(line[i:endIndex])
			if err != nil {
				log.Fatal(err)
			}
			numArr = append(numArr, Num{i, endIndex, false, numVal})

			i = endIndex
		}
	}

	return numArr
}

type Num struct {
	begin, end     int
	isPartOfEngine bool
	val            int
}
