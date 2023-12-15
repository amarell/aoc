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

	string := ""
	for scanner.Scan() {
		string = scanner.Text()
	}

	parseLine(string)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func parseLine(line string) {
	S := strings.Split(line, ",")

	M := make(map[string]int)
	A := []string{}

	sum := 0
	for _, s := range S {
		if strings.Contains(s, "-") {
			label := s[:len(s)-1]
			M[label] = 0
			A = removeFromArray(label, A)

		} else if strings.Contains(s, "=") {
			label := s[:len(s)-2]

			if M[label] == 0 {
				A = append(A, label)
			}
			M[label] = atoi(string(s[len(s)-1]))

		}
		sum += getHash(s)
	}

	part2Sum := 0
	for k, v := range M {
		boxNum := getHash(k)
		part2Sum += (boxNum + 1) * findBoxSlot(k, A) * v
	}

	fmt.Println("Part 2: ", part2Sum)

	fmt.Println("Part 1: ", sum)
}

func removeFromArray(el string, arr []string) []string {
	newArr := arr
	for idx, element := range arr {
		if element == el {
			newArr = append(arr[:idx], arr[idx+1:]...)
		}
	}

	return newArr
}

func findBoxSlot(key string, arr []string) int {
	hash := getHash(key)

	count := 0

	for _, el := range arr {
		if getHash(el) == hash {
			count++
		}
		if el == key {
			return count
		}
	}

	return count
}

func atoi(str string) int {
	strInt, _ := strconv.Atoi(str)
	return strInt
}

type Lens struct {
	id          string
	focalLength int
	boxNum      int
}

func getHash(str string) int {
	res := 0

	for _, ch := range str {
		res += int(ch)
		res *= 17
		res %= 256
	}

	return res
}
