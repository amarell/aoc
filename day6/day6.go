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

	part1(lines)
	part2(lines)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func parseInput(lines []string) ([]int, []int) {

	times := []int{}
	distances := []int{}

	timesA := strings.Split(lines[0], " ")
	distancesA := strings.Split(lines[1], " ")

	for idx, time := range timesA {
		if idx > 1 && time != "" && time != " " {
			timeInt, err := strconv.Atoi(time)

			if err != nil {
				log.Fatal(err)
			}

			times = append(times, timeInt)
		}
	}

	for idx, distance := range distancesA {
		if idx > 1 && distance != "" && distance != " " {
			distanceInt, err := strconv.Atoi(distance)

			if err != nil {
				log.Fatal(err)
			}

			distances = append(distances, distanceInt)
		}
	}

	return times, distances
}

func part1(lines []string) int {
	times, distances := parseInput(lines)
	waysToWinArr := []int{}

	for idx, time := range times {
		distanceTraveled := 0
		waysToWin := 0

		for i := 1; i < time; i++ {
			distanceTraveled = i * (time - i)
			if distanceTraveled > distances[idx] {
				waysToWin++
			}
		}

		waysToWinArr = append(waysToWinArr, waysToWin)
	}

	factor := 1

	for _, f := range waysToWinArr {
		factor *= f
	}

	fmt.Println("Part 1", factor)
	return factor

}

func part2(lines []string) int {
	timesNew := []int{}
	timesNew = append(timesNew, 51926890)

	distancesNew := []int{}
	distancesNew = append(distancesNew, 222203111261225)

	waysToWinArr := []int{}
	for idx, time := range timesNew {
		distanceTraveled := 0
		waysToWin := 0

		for i := 1; i < time; i++ {
			distanceTraveled = i * (time - i)
			if distanceTraveled > distancesNew[idx] {
				waysToWin++
			}
		}

		waysToWinArr = append(waysToWinArr, waysToWin)
	}

	factor := 1

	for _, f := range waysToWinArr {
		factor *= f
	}

	fmt.Println("Part 2", factor)
	return factor

}
