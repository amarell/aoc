package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	// part1(lines)
	part2(lines)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func part2(lines []string) {
	matches := []int{}
	for _, line := range lines {
		nums := strings.Split(line, "|")

		winning, my := strings.Trim(nums[0], " "), strings.Trim(nums[1], " ")

		winningNums := strings.Split(winning, " ")
		myNums := strings.Split(my, " ")
		match := 0

		for _, num := range myNums {
			if slices.Index(winningNums, num) >= 0 && num != "" && num != " " {
				match++
			}
		}
		matches = append(matches, match)
	}

	// I have at least this amount of cards
	// res := len(lines)

	// find number of copies of each card

	numOfCopies := []int{}

	for i := 0; i < len(lines); i++ {
		numOfCopies = append(numOfCopies, 1)
	}

	//matches [4 2 3 1 0 0]
	//init state [1,1,1, 1,1,1]
	//
	//prva iteracija
	//[1, 2, 2, 2, 2, 1]
	//
	//druga iteracija
	//[1, 2, 4, 4, 2, 1]
	//
	//treca iteracija
	//[1, 2, 4, 8, 6, 5]
	//
	//cetvrta iteracija
	//[1, 2, 4, 8, 14, 5]

	for idx, m := range matches {
		for i := 0; i < m; i++ {
			numOfCopies[idx+1+i] = numOfCopies[idx+1+i] + numOfCopies[idx]
		}
	}

	res2 := 0
	for _, num := range numOfCopies {
		res2 += num
	}

	fmt.Println("res 2", res2)
	// fmt.Println("res", res)
}

func part1(lines []string) {
	finalPoints := 0

	for _, line := range lines {
		nums := strings.Split(line, "|")

		winning, my := strings.Trim(nums[0], " "), strings.Trim(nums[1], " ")

		winningNums := strings.Split(winning, " ")

		fmt.Println("winning", winningNums)

		myNums := strings.Split(my, " ")
		fmt.Println("my", myNums)

		points := 0

		for _, num := range myNums {
			if slices.Index(winningNums, num) >= 0 {
				if points == 0 && num != "" && num != " " {
					fmt.Println("winning num is ", num)
					points = 1
				} else if num != "" && num != " " {
					points *= 2
				}
			}
		}

		fmt.Println(points)

		finalPoints += points
	}

	fmt.Println(finalPoints)
}
