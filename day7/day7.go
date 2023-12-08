package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HIGH_CARD int = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	parts := []bool{true, false}

	for _, p := range parts {
		hands := parseLines(lines, p)
		rankedHands := rankHands(hands, p)

		r := 0
		for idx, handRank := range rankedHands {
			r += (idx + 1) * handRank.hand.bid
		}

		if p {
			fmt.Println("Part 1", r)
		} else {
			fmt.Println("Part 2", r)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func rankHands(hands []Hand, part1 bool) []HandRank {
	handTypes := []int{}
	handRanks := []HandRank{}

	for _, hand := range hands {
		handTypes = append(handTypes, hand.getHandType(part1))
	}

	for idx, hand := range hands {
		handRanks = append(handRanks, HandRank{hand, handTypes[idx]})
	}

	sort.Slice(handRanks[:], func(i, j int) bool {
		if handRanks[i].t == handRanks[j].t {
			for idx, card := range handRanks[i].hand.cards {
				if card.value < handRanks[j].hand.cards[idx].value {
					return true
				} else if card.value > handRanks[j].hand.cards[idx].value {
					return false
				}
			}
		}
		return handRanks[i].t < handRanks[j].t
	})

	return handRanks
}

type HandRank struct {
	hand Hand
	t    int
}

func (hand Hand) getHandType(part1 bool) int {
	M := make(map[Card]int)
	for _, card := range hand.cards {
		M[card] = M[card] + 1
	}
	// 2
	// full house
	// four of a kind

	// 3
	// three of a kind
	// two pair

	// 4
	// one pair

	T := 0

	if len(M) == 1 {
		T = FIVE_OF_A_KIND
	} else if len(M) == 5 {
		T = HIGH_CARD
	} else if len(M) == 4 {
		T = ONE_PAIR
	} else if len(M) == 3 {
		isThree := false
		for _, v := range M {
			if v == 3 {
				isThree = true
			}
		}
		if isThree {
			T = THREE_OF_A_KIND
		} else {
			T = TWO_PAIR
		}

	} else {
		isFour := false

		for _, v := range M {
			if v == 4 {
				isFour = true
			}
		}

		if isFour {
			T = FOUR_OF_A_KIND
		} else {
			T = FULL_HOUSE
		}
	}

	if M[Card{1}] > 0 && !part1 {
		if T == FOUR_OF_A_KIND {
			return FIVE_OF_A_KIND
		} else if T == FULL_HOUSE {
			return FIVE_OF_A_KIND
		} else if T == THREE_OF_A_KIND {
			return FOUR_OF_A_KIND
		} else if T == TWO_PAIR {
			if M[Card{1}] == 1 {
				return FULL_HOUSE
			}
			return FOUR_OF_A_KIND
		} else if T == ONE_PAIR {
			return THREE_OF_A_KIND
		} else if T == HIGH_CARD {
			return ONE_PAIR
		}
	}

	return T
}

type Hand struct {
	cards []Card
	bid   int
}

type Card struct {
	value int
}

func sortCardsInHands(hands []Hand) []Hand {
	newHands := []Hand{}

	for _, hand := range hands {
		sort.Slice(hand.cards[:], func(i, j int) bool {
			return hand.cards[i].value < hand.cards[j].value
		})

		newHands = append(newHands, Hand{hand.cards, hand.bid})
	}

	return newHands
}

func parseLines(lines []string, part1 bool) []Hand {
	cardValueMap := getCardMap(part1)

	hands := []Hand{}
	for _, line := range lines {
		cards := []Card{}

		split := strings.Split(line, " ")

		handStr, bidStr := split[0], split[1]

		bidInt, err := strconv.Atoi(bidStr)

		if err != nil {
			log.Fatal(err)
		}

		for _, ch := range handStr {
			cards = append(cards, Card{cardValueMap[ch]})
		}

		hands = append(hands, Hand{cards, bidInt})

	}

	return hands
}

func getCardMap(part1 bool) map[rune]int {
	cardValueMap := make(map[rune]int)

	cardValueMap['2'] = 2
	cardValueMap['3'] = 3
	cardValueMap['4'] = 4
	cardValueMap['5'] = 5
	cardValueMap['6'] = 6
	cardValueMap['7'] = 7
	cardValueMap['8'] = 8
	cardValueMap['9'] = 9
	cardValueMap['T'] = 10
	cardValueMap['J'] = ternary(part1, 11, 1)
	cardValueMap['Q'] = 12
	cardValueMap['K'] = 13
	cardValueMap['A'] = 14

	return cardValueMap
}

func ternary(cond bool, val1, val2 int) int {
	if cond {
		return val1
	}

	return val2
}
