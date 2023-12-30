package main

import (
	"aoc/pkg/file"
	"fmt"
	"strings"
)

func main() {
	input, _ := file.ReadInput("./input.txt")
	parseInput(input)
}

const (
	CONJ = iota
	FLOP
)

type Node struct {
	class int
	name  string
	pulse bool
}

func parseInput(lines []string) {
	flops := make(map[string]bool)
	conjs := make(map[string][]Node)

	// map src to array of destinations
	graph := make(map[string][]string)

	for _, l := range lines {
		s := strings.Split(l, " -> ")
		src := s[0]
		dests := strings.Split(s[1], ", ")

		if src[0] == '%' {
			// initially flops are turned off (false)
			src = src[1:]
			flops[src] = false
		} else if src[0] == '&' {
			src = src[1:]
			conjs[src] = []Node{}
		}

		graph[src] = dests
	}

	for src, dests := range graph {
		for _, dest := range dests {
			_, ok := conjs[dest]

			if ok {
				conjs[dest] = append(conjs[dest], Node{CONJ, src, false})
			}
		}
	}

	tothi, totlo := 0, 0
	nhi, nlo := 0, 0

	init_flops, init_conjs, init_graph := copyMap(flops), copyMap(conjs), copyMap(graph)

	for i := 0; i < 1000; i++ {
		nhi, nlo, flops, conjs, graph, _ = run(flops, conjs, graph, []string{})
		tothi += nhi
		totlo += nlo
	}
	fmt.Println("Part 1: ", tothi*totlo)

	// reset
	periods := findPeriods(init_flops, init_conjs, init_graph)

	res := 1
	for _, per := range periods {
		res *= per
	}

	// should be lcm, but they're prime anyway so idc
	fmt.Println("Part 2: ", res)
}

type anything interface {
}

func copyMap[K comparable, V anything](m map[K]V) map[K]V {

	cp := make(map[K]V)

	for k, v := range m {
		cp[k] = v
	}

	return cp
}

type Signal struct {
	sender, receiver string
	pulse            bool
}

func popLeft(arr []Signal) (Signal, []Signal) {
	el := arr[0]
	arr = arr[1:]

	return el, arr
}

func propagatePulse(flops map[string]bool, conjs map[string][]Node, graph map[string][]string, signal Signal) (map[string]bool, map[string][]Node, map[string][]string, []Signal) {
	_, isFlop := flops[signal.receiver]
	_, isConj := conjs[signal.receiver]
	_, isInGraph := graph[signal.receiver]
	next_pulse := signal.pulse
	skip := false

	if isFlop {
		if signal.pulse {
			// continue
			skip = true
		} else {
			flops[signal.receiver] = !flops[signal.receiver]
			next_pulse = flops[signal.receiver]
		}
	} else if isConj {
		// first update state
		conjs[signal.receiver] = updateNodeInArr(signal.sender, conjs[signal.receiver], signal.pulse)
		next_pulse = !allHigh(conjs[signal.receiver])
	} else if isInGraph {
		// just nothing
	} else {
		skip = true
	}

	// now propagate the next pulse to all modules connected to this one
	addToQ := []Signal{}

	if skip {
		return flops, conjs, graph, addToQ
	}

	for _, dest := range graph[signal.receiver] {
		addToQ = append(addToQ, Signal{signal.receiver, dest, next_pulse})
	}
	return flops, conjs, graph, addToQ
}

func run(flops map[string]bool, conjs map[string][]Node, graph map[string][]string, inputsToRx []string) (int, int, map[string]bool, map[string][]Node, map[string][]string, string) {
	q := []Signal{{"button", "broadcaster", false}}
	nhi := 0
	nlo := 0

	foundThisInput := ""

	for len(q) > 0 {
		signal, nq := popLeft(q)
		q = nq

		if arrContains(inputsToRx, signal.receiver) && !signal.pulse {
			foundThisInput = signal.receiver
		}

		if signal.pulse {
			nhi++
		} else {
			nlo++
		}

		addToQ := []Signal{}
		flops, conjs, graph, addToQ = propagatePulse(flops, conjs, graph, signal)
		q = append(q, addToQ...)

	}

	return nhi, nlo, flops, conjs, graph, foundThisInput
}

func findPeriods(flops map[string]bool, conjs map[string][]Node, graph map[string][]string) []int {
	periodics := make(map[string]int)

	// 1. there is only one rx module
	// 2. there is only one input to rx (call it A)
	// 3. all inputs to this input are conj modules (call it B1, ... Bn)

	// rx needs to receive a low pulse
	// that means that B1, ... Bn need to all send high pulses
	// that happens whenever B1, ... Bn  receive a low pulse

	// find A
	rx_src := ""
	for src, dests := range graph {
		if arrContains(dests, "rx") {
			// make sure it's a conj
			_, isConj := conjs[src]

			if !isConj {
				panic("A is not a conj module")
			}

			if len(dests) != 1 {
				panic("A has more than one destination")
			}

			rx_src = src
		}
	}

	for src, dest := range graph {
		if arrContains(dest, rx_src) {
			_, isConj := conjs[src]

			if !isConj {
				panic("Input to A is not a conj")
			}

			periodics[src] = -1
		}
	}

	inputsToRx := []string{}

	for k := range periodics {
		inputsToRx = append(inputsToRx, k)
	}

	periods := []int{}

	i := 0
	for len(inputsToRx) != 0 {
		foundThisInput := ""
		_, _, flops, conjs, graph, foundThisInput = run(flops, conjs, graph, inputsToRx)

		if len(foundThisInput) > 0 {
			inputsToRx = removeElFromArray(inputsToRx, foundThisInput)
			periods = append(periods, i+1)
		}
		i++
	}

	return periods
}

func removeElFromArray(arr []string, el string) []string {
	newArr := []string{}

	for _, ele := range arr {
		if ele != el {
			newArr = append(newArr, ele)
		}
	}

	return newArr
}

func arrContains(arr []string, el string) bool {
	for _, element := range arr {
		if el == element {
			return true
		}
	}
	return false
}

func updateNodeInArr(name string, arr []Node, val bool) []Node {
	newNodes := []Node{}

	for _, node := range arr {
		if name == node.name {
			node.pulse = val
		}
		newNodes = append(newNodes, node)
	}

	return newNodes
}

func allHigh(arr []Node) bool {
	for _, el := range arr {
		if !el.pulse {
			return false
		}
	}

	return true
}
