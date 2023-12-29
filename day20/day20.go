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
	for i := 0; i < 1000; i++ {
		nhi, nlo, f, c, g := run(flops, conjs, graph)
		flops = f
		conjs = c
		graph = g
		tothi += nhi
		totlo += nlo
	}

	fmt.Println("Part 1:", tothi*totlo)
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

func run(flops map[string]bool, conjs map[string][]Node, graph map[string][]string) (int, int, map[string]bool, map[string][]Node, map[string][]string) {
	q := []Signal{{"button", "broadcaster", false}}
	nhi := 0
	nlo := 0

	for len(q) > 0 {
		signal, nq := popLeft(q)
		next_pulse := signal.pulse
		q = nq

		if signal.pulse {
			nhi++
		} else {
			nlo++
		}

		_, isFlop := flops[signal.receiver]
		_, isConj := conjs[signal.receiver]
		_, isInGraph := graph[signal.receiver]

		if isFlop {
			if signal.pulse {
				continue
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
			continue
		}

		// now propagate the next pulse to all modules connected to this one
		for _, dest := range graph[signal.receiver] {
			q = append(q, Signal{signal.receiver, dest, next_pulse})
		}
	}

	return nhi, nlo, flops, conjs, graph
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
