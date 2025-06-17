package main

import (
	"fmt"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	// Part2()
}

type Graph struct {
	graph map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		graph: make(map[string][]string),
	}
}

func (g *Graph) addEdge(from string, to []string) {
	if from == to[0] && from == to[1] {
		g.graph[from] = []string{}
	} else if from == to[0] {
		g.graph[from] = []string{to[1]}
	} else if from == to[1] {
		g.graph[from] = []string{to[0]}
	} else {
		g.graph[from] = to
	}
}

func (g *Graph) get(from, inst string) string {
	if inst == "L" {
		return g.graph[from][0]
	} else {
		return g.graph[from][1]
	}
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_8/input.txt")

	instructions := strings.Split(input, "\n")[0]

	graph := NewGraph()

	for _, line := range strings.Split(input, "\n")[2:] {
		if line != "" {
			name := strings.TrimSpace(strings.Split(line, "=")[0])
			values := strings.Split(strings.Split(line, "=")[1][2:len(strings.Split(line, "=")[1])-1], ", ")
			graph.addEdge(name, values)
		}
	}

	currentNetwork := "AAA"

	result := 0

	for currentNetwork != "ZZZ" {
		result += 1
		inst := string(instructions[0])
		instructions = instructions[1:]

		currentNetwork = graph.get(currentNetwork, inst)

		instructions += inst
	}

	fmt.Println(result)

}

func Part2() {
	fmt.Println("Part 2")
}
