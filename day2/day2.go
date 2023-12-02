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
	sum := 0

	powerSum := 0
	for scanner.Scan() {
		gameId, isPossible, power := part1(scanner.Text())

		if isPossible {
			sum += gameId
		}

		powerSum += power
	}

	fmt.Println(sum)
	fmt.Println(powerSum)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func part2(moves []Move) int {

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, move := range moves {
		if move.red > maxRed {
			maxRed = move.red
		}

		if move.green > maxGreen {
			maxGreen = move.green
		}

		if move.blue > maxBlue {
			maxBlue = move.blue
		}
	}

	return maxBlue * maxGreen * maxRed
}

func part1(input string) (int, bool, int) {
	gameId, moves := parseInput(input)

	fmt.Println("moves: ", moves, " isPossible: ", isPossible(moves))

	return gameId, isPossible(moves), part2(moves)
}

type Move struct {
	red, green, blue int
}

func parseInput(input string) (int, []Move) {
	words := strings.Split(input, " ")

	gameIdStr := words[1]

	gameId, err := strconv.Atoi(gameIdStr[:len(gameIdStr)-1])
	info := strings.Split(input, ": ")

	moves := strings.Split(info[1], ";")

	movesArr := []Move{}
	for _, move := range moves {
		movesArr = append(movesArr, parseMove(move))
	}
	if err != nil {
		log.Fatal(err)
	}

	return gameId, movesArr
}

func isPossible(moves []Move) bool {
	// configuration:
	RED_LIMIT := 12
	GREEN_LIMIT := 13
	BLUE_LIMIT := 14

	for _, move := range moves {
		if move.red > RED_LIMIT || move.green > GREEN_LIMIT || move.blue > BLUE_LIMIT {
			return false
		}
	}
	return true
}

func parseMove(move string) Move {
	red, green, blue := 0, 0, 0

	cubes := strings.Split(move, ",")

	//	fmt.Println(cubes)
	//	fmt.Println("This is the move ", move)
	for _, cube := range cubes {
		cubArr := strings.Split(strings.Trim(cube, " "), " ")

		//	fmt.Println(cubArr)
		val, err := strconv.Atoi(cubArr[0])

		if err != nil {
			log.Fatal(err)
		}
		switch cubArr[1] {
		case "red":
			red = val
		case "blue":
			blue = val
		case "green":
			green = val
		}
	}

	// fmt.Println("This is the move parsed: (red, green, blue)", red, green, blue)
	return Move{red, green, blue}
}
