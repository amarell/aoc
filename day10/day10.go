package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type Grid struct {
	G [][]rune
}

func (g Grid) getHeight() int {
	return len(g.G)
}

func (g Grid) getWidth() int {
	return len(g.G[0])
}

func (g Grid) getAtXY(x, y int) (string, error) {
	if x < 0 || y < 0 || x > len(g.G[0])-1 || y > len(g.G)-1 {
		return "", fmt.Errorf("Invalid coordinates")
	}

	return string(g.G[y][x]), nil
}

func (g Grid) getAtStep(step Step) (string, error) {
	if step.x < 0 || step.y < 0 || step.x > len(g.G[0])-1 || step.y > len(g.G)-1 {
		return "", fmt.Errorf("Invalid coordinates")
	}

	return string(g.G[step.y][step.x]), nil
}

func NewGrid(g [][]rune) Grid {
	return Grid{g}
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	G := [][]rune{}

	for scanner.Scan() {
		rowArr := []rune{}
		for _, b := range scanner.Text() {
			rowArr = append(rowArr, b)
		}
		G = append(G, rowArr)
	}

	grid := NewGrid(G)

	findFarthestDistance(grid)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func stepInDirection(direction int, g Grid, currentStep Step) (string, bool) {
	switch direction {
	case RIGHT:
		{
			nextPos, err := g.getAtXY(currentStep.x+1, currentStep.y)

			if err != nil {
				return nextPos, false
			}

			step := Step{currentStep.x + 1, currentStep.y}
			return nextPos, isValidNextStep(LEFT, step, g)
		}
	case LEFT:
		{
			nextPos, err := g.getAtXY(currentStep.x-1, currentStep.y)

			if err != nil {
				return nextPos, false
			}

			step := Step{currentStep.x - 1, currentStep.y}
			return nextPos, isValidNextStep(RIGHT, step, g)

		}

	case UP:
		{
			nextPos, err := g.getAtXY(currentStep.x, currentStep.y-1)

			if err != nil {
				return nextPos, false
			}

			step := Step{currentStep.x, currentStep.y - 1}

			return nextPos, isValidNextStep(DOWN, step, g)
		}
	case DOWN:
		{
			nextPos, err := g.getAtXY(currentStep.x, currentStep.y+1)

			if err != nil {
				return nextPos, false
			}

			step := Step{currentStep.x, currentStep.y + 1}

			return nextPos, isValidNextStep(UP, step, g)

		}

	default:
		return ".", false
	}
}

func isValidNextStep(cameFromDirection int, newStep Step, g Grid) bool {

	newStepValue, _ := g.getAtStep(newStep)

	switch cameFromDirection {
	case LEFT:
		prevStep := Step{newStep.x - 1, newStep.y}
		prevStepValue, _ := g.getAtStep(prevStep)

		if prevStepValue != "L" && prevStepValue != "-" && prevStepValue != "F" && prevStepValue != "S" {
			return false
		}

		return newStepValue == "-" || newStepValue == "J" || newStepValue == "7" || newStepValue == "S"
	case RIGHT:
		prevStep := Step{newStep.x + 1, newStep.y}
		prevStepValue, _ := g.getAtStep(prevStep)

		if prevStepValue != "-" && prevStepValue != "J" && prevStepValue != "7" && prevStepValue != "S" {
			return false
		}

		return newStepValue == "-" || newStepValue == "L" || newStepValue == "F" || newStepValue == "S"
	case UP:
		prevStep := Step{newStep.x, newStep.y - 1}
		prevStepValue, _ := g.getAtStep(prevStep)

		if prevStepValue != "|" && prevStepValue != "7" && prevStepValue != "F" && prevStepValue != "S" {
			return false
		}
		return newStepValue == "|" || newStepValue == "L" || newStepValue == "J" || newStepValue == "S"
	case DOWN:
		prevStep := Step{newStep.x, newStep.y + 1}
		prevStepValue, _ := g.getAtStep(prevStep)

		if prevStepValue != "|" && prevStepValue != "L" && prevStepValue != "J" && prevStepValue != "S" {
			return false
		}

		return newStepValue == "|" || newStepValue == "7" || newStepValue == "F" || newStepValue == "S"
	default:
		return false
	}
}

func findNextSteps(g Grid, currentPos Step) []Step {
	dirs := []int{UP, DOWN, RIGHT, LEFT}
	validNextSteps := []Step{}

	for _, dir := range dirs {
		_, valid := stepInDirection(dir, g, currentPos)
		if valid {
			switch dir {
			case UP:
				validNextSteps = append(validNextSteps, Step{currentPos.x, currentPos.y - 1})
			case DOWN:
				validNextSteps = append(validNextSteps, Step{currentPos.x, currentPos.y + 1})
			case LEFT:
				validNextSteps = append(validNextSteps, Step{currentPos.x - 1, currentPos.y})
			case RIGHT:
				validNextSteps = append(validNextSteps, Step{currentPos.x + 1, currentPos.y})
			}
		}
	}

	return validNextSteps

}

type Step struct {
	x, y int
}

func (s1 Step) isEqual(s2 Step) bool {
	return s1.x == s2.x && s1.y == s2.y
}

func findFarthestDistance(g Grid) int {
	x, y := findS(g)
	origin := Step{x, y}
	loops := findLoops(g, origin)

	for _, loop := range loops {
		fmt.Println(loop)
	}

	maxLength := len(loops[0])
	for _, loop := range loops {
		if len(loop) > maxLength {
			maxLength = len(loop)
		}
	}

	fmt.Println("Part 1:", maxLength/2)

	return 0
}

func findLoops(g Grid, origin Step) [][]Step {
	loops := [][]Step{}

	// we get 2-4 valid options!
	// nextSteps := findNextSteps(g, origin)

	nextSteps := findNextSteps(g, origin)

	for _, step := range nextSteps {
		steps := []Step{origin, step}
		for !steps[len(steps)-1].isEqual(origin) {
			steps = append(steps, propagate(g, steps[len(steps)-1], steps[len(steps)-2]))
		}

		loops = append(loops, steps)
	}

	return loops
}

func propagate(g Grid, current Step, prev Step) Step {
	nextSteps := findNextSteps(g, current)
	if len(nextSteps) == 1 {
		return nextSteps[0]
	}
	if nextSteps[0].isEqual(prev) {
		return nextSteps[1]
	}
	return nextSteps[0]
}

func findS(g Grid) (int, int) {
	for y := 0; y < len(g.G); y++ {
		for x := 0; x < len(g.G[y]); x++ {
			if g.G[y][x] == 'S' {
				return x, y
			}
		}
	}
	return 0, 0
}
