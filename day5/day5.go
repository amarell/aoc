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

	// fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func getSeeds(lines []string) []int {
	seeds := []int{}
	seedStrs := strings.Split(lines[0], " ")
	for idx, seedStr := range seedStrs {
		if idx != 0 {
			num, err := strconv.Atoi(seedStr)

			if err != nil {
				log.Fatal(err)
			}

			seeds = append(seeds, num)
		}
	}

	return seeds
}

func getSeedsPart2(lines []string) []int {
	seeds := getSeeds(lines)

	newSeeds := []int{}
	for idx, seed := range seeds {
		if idx%2 == 0 && idx > 1 && idx < 4 {
			for i := 0; i < seeds[idx+1]; i++ {
				newSeeds = append(newSeeds, seed+i)
				fmt.Println("appending", seed+i)
				fmt.Println(i, "/", seeds[idx+1])
			}
		}
	}

	fmt.Println("new seeds", newSeeds)
	return newSeeds
}

func part2(lines []string) int {
	seeds := getSeeds(lines)

	return getMinLocationFromSeeds(seeds, lines)
}

func getSmallestRange(rngs []Rng) int {

	return 0
}

func getMinLocationFromSeeds(seeds []int, lines []string) int {
	emptyLinesIndices := []int{}

	for idx, line := range lines {
		if line == "" {
			emptyLinesIndices = append(emptyLinesIndices, idx)
		}
	}

	seedToSoilLines := lines[emptyLinesIndices[0]+2 : emptyLinesIndices[1]]
	soilToFertilizerLines := lines[emptyLinesIndices[1]+2 : emptyLinesIndices[2]]
	fertilizerToWaterLines := lines[emptyLinesIndices[2]+2 : emptyLinesIndices[3]]
	waterToLightLines := lines[emptyLinesIndices[3]+2 : emptyLinesIndices[4]]
	lightToTemperatureLines := lines[emptyLinesIndices[4]+2 : emptyLinesIndices[5]]
	temperatureToHumidityLines := lines[emptyLinesIndices[5]+2 : emptyLinesIndices[6]]
	humidityToLocationLines := lines[emptyLinesIndices[6]+2:]

	seedToSoilRngs := createRngs(seedToSoilLines)
	soilToFertilizerRngs := createRngs(soilToFertilizerLines)
	fertilizerToWaterRngs := createRngs(fertilizerToWaterLines)
	waterToLightRngs := createRngs(waterToLightLines)
	lightToTemperatureRngs := createRngs(lightToTemperatureLines)
	temperatureToHumidityRngs := createRngs(temperatureToHumidityLines)
	humidityToLocationRngs := createRngs(humidityToLocationLines)

	fmt.Println(humidityToLocationRngs)

	// goal => min range humidity to Location => get ranges for temperature to humidity

	minRangeHumidity := Rng{}

	minBeginDest := humidityToLocationRngs[0].beginDest

	for _, rng := range humidityToLocationRngs {
		if rng.beginDest < minBeginDest {
			minBeginDest = rng.beginDest
			minRangeHumidity = rng
		}
	}

	if minBeginDest > 0 {
		minRangeHumidity = Rng{0, 0, minBeginDest}
	}

	fmt.Println(seeds)
	for i := 0; i < minRangeHumidity.rng; i++ {
		location := minRangeHumidity.beginDest + i
		humidity := getPrevInput(humidityToLocationRngs, location)
		temperature := getPrevInput(temperatureToHumidityRngs, humidity)
		light := getPrevInput(lightToTemperatureRngs, temperature)
		water := getPrevInput(waterToLightRngs, light)
		fertilizer := getPrevInput(fertilizerToWaterRngs, water)
		soil := getPrevInput(soilToFertilizerRngs, fertilizer)
		seed := getPrevInput(seedToSoilRngs, soil)

		// fmt.Println("Seed nm: ", seed)
		for j := 0; j < len(seeds); j += 2 {
			if seed >= seeds[j] && seed < seeds[j]+seeds[j+1] {
				fmt.Println("Part 2 amar :", seed)
				fmt.Println("Part 2 location:", location)
				return location
			}
		}
	}
	// xD
	//fmt.Println("humidity", humidity)
	//fmt.Println("temperature", temperature)
	//fmt.Println("light", light)
	//fmt.Println("water", water)

	locations := []int{}

	for _, seed := range seeds {
		soil := getNextInput(seedToSoilRngs, seed)
		fertilizer := getNextInput(soilToFertilizerRngs, soil)
		water := getNextInput(fertilizerToWaterRngs, fertilizer)
		light := getNextInput(waterToLightRngs, water)
		temperature := getNextInput(lightToTemperatureRngs, light)
		humidity := getNextInput(temperatureToHumidityRngs, temperature)
		location := getNextInput(humidityToLocationRngs, humidity)

		locations = append(locations, location)

	}

	min := locations[0]

	for _, location := range locations {
		if location < min {
			min = location
		}
	}

	return min
}

func part1(lines []string) int {
	seeds := getSeeds(lines)

	return getMinLocationFromSeeds(seeds, lines)
}

func createRngs(lines []string) []Rng {
	rngs := []Rng{}
	for _, line := range lines {
		split := strings.Split(line, " ")

		beginValue, beginKey, rng := split[0], split[1], split[2]

		beginValueInt, err := strconv.Atoi(beginValue)

		if err != nil {
			log.Fatal(err)
		}

		beginKeyInt, err := strconv.Atoi(beginKey)

		if err != nil {
			log.Fatal(err)
		}

		rngInt, err := strconv.Atoi(rng)

		if err != nil {
			log.Fatal(err)
		}
		rngs = append(rngs, Rng{beginValueInt, beginKeyInt, rngInt})
	}

	return rngs
}

func getPrevInput(rngs []Rng, output int) int {
	input := output
	for _, rng := range rngs {
		offset := output - rng.beginDest
		if output >= rng.beginDest && output < rng.beginDest+rng.rng {
			input = rng.beginSrc + offset
			break
		}
	}

	return input
}

func getNextInput(rngs []Rng, input int) int {
	output := input
	for _, rng := range rngs {
		offset := input - rng.beginSrc
		if input >= rng.beginSrc && input < rng.beginSrc+rng.rng {
			output = rng.beginDest + offset
			break
		}
	}

	return output
}

type Rng struct {
	beginDest, beginSrc, rng int
}

func createMap(lines []string) map[int]int {

	newMap := make(map[int]int)

	for _, line := range lines {
		split := strings.Split(line, " ")

		beginValue, beginKey, rng := split[0], split[1], split[2]

		beginKeyInt, err := strconv.Atoi(beginKey)

		if err != nil {
			log.Fatal(err)
		}

		beginValueInt, err := strconv.Atoi(beginValue)

		if err != nil {
			log.Fatal(err)
		}

		rangeInt, err := strconv.Atoi(rng)

		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < rangeInt; i++ {
			newMap[beginKeyInt+i] = beginValueInt + i
		}

	}

	return newMap
}
