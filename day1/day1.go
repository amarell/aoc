package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")

	sum := 0
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num := getDigitsPart2(scanner.Text())
		// num := getDigits(scanner.Text())
		sum += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The sum is: ", sum)

}

func runeToInt(r rune) int {
	return int(r) - int('0')
}

func getDigitsPart2(str string) int {
	var ultFirstDigit int
	var ultLastDigit int

	firstIdx1, lastIdx1, digit1, digit2 := getFirstAndLastIndexFromSubstrings(str)
	firstIdx2, lastIdx2, digit3, digit4 := getFirstAndLastIndices(str)

	if firstIdx1 < firstIdx2 {
		ultFirstDigit = digit1
	} else {
		ultFirstDigit = digit3
	}

	if lastIdx1 > lastIdx2 {
		ultLastDigit = digit2
	} else {
		ultLastDigit = digit4
	}

	return ultFirstDigit*10 + ultLastDigit
}

func getFirstAndLastIndexFromSubstrings(str string) (int, int, int, int) {
	numNames := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	minIndex := 999
	minNumber := -1
	maxIndex := -1
	maxNumber := -1

	for idx, num := range numNames {
		firstInstance := strings.Index(str, num)
		lastInstance := strings.LastIndex(str, num)

		if firstInstance < minIndex && firstInstance >= 0 {
			minIndex = firstInstance
			minNumber = idx
		}

		if lastInstance > maxIndex && lastInstance >= 0 {
			maxIndex = lastInstance
			maxNumber = idx
		}
	}

	return minIndex, maxIndex, minNumber, maxNumber
}

func getFirstAndLastIndices(str string) (int, int, int, int) {
	var firstDigitIndex int
	var lastDigitIndex int
	var firstDigit int
	var lastDigit int

	for idx, char := range str {
		if unicode.IsDigit(char) {
			firstDigitIndex = idx
			firstDigit = runeToInt(char)
			break
		}
	}

	for idx, char := range str {
		if unicode.IsDigit(char) {
			lastDigitIndex = idx
			lastDigit = runeToInt(char)
		}
	}

	return firstDigitIndex, lastDigitIndex, firstDigit, lastDigit
}

func getDigits(str string) int {
	var digits []int

	for _, character := range str {
		if unicode.IsDigit(character) {
			digits = append(digits, runeToInt(character))
		}
	}

	return digits[0]*10 + digits[len(digits)-1]
}
