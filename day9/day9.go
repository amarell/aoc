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

	startNums := parseInput(lines)

	nexts := []int{}

	for _, startNum := range startNums {
		nexts = append(nexts, getNext(startNum))
	}

	sum := 0
	for _, n := range nexts {
		sum += n
	}

	fmt.Println("Part 1", sum)

	prevs := []int{}

	for _, startNum := range startNums {
		prevs = append(prevs, getPrev(startNum))
	}

	sum = 0
	for _, n := range prevs {
		sum += n
	}

	fmt.Println("Part 2", sum)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func parseInput(lines []string) [][]int {
	N := [][]int{}
	for _, l := range lines {
		s := strings.Split(l, " ")
		nums := []int{}
		for _, str := range s {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}

		N = append(N, nums)
	}

	return N
}

func getNext(startNums []int) int {
	D := [][]int{}
	D = append(D, startNums)

	for !arrSame(D[len(D)-1], 0) {
		D = append(D, getDiffsForArray(D[len(D)-1]))
	}

	return getNextFromDiffs(D)
}

func getPrev(startNums []int) int {
	D := [][]int{}
	D = append(D, startNums)

	for !arrSame(D[len(D)-1], 0) {
		D = append(D, getDiffsForArray(D[len(D)-1]))
	}

	return getPrevFromDiffs(D)
}

func getPrevFromDiffs(D [][]int) int {

	current := 0

	for i := len(D) - 1; i >= 1; i-- {
		current = D[i-1][0] - current
	}

	return current
}

func getNextFromDiffs(D [][]int) int {
	r := 0
	for _, diff := range D {
		r += diff[len(diff)-1]
	}

	return r
}

func getDiffsForArray(arr []int) []int {
	diffs := []int{}

	for i := 1; i < len(arr); i++ {
		diffs = append(diffs, arr[i]-arr[i-1])
	}

	return diffs

}

func arrSame(arr []int, el int) bool {
	for _, e := range arr {
		if e != el {
			return false
		}
	}
	return true
}
