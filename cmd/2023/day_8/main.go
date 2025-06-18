package main

import (
	"fmt"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
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

	input := utils.ReadFile("cmd/2023/day_8/input.txt")

	// 	input = `LR

	// 11A = (11B, XXX)
	// 11B = (XXX, 11Z)
	// 11Z = (11B, XXX)
	// 22A = (22B, XXX)
	// 22B = (22C, 22C)
	// 22C = (22Z, 22Z)
	// 22Z = (22B, 22B)
	// XXX = (XXX, XXX)`

	instructions := strings.Split(input, "\n")[0]

	graph := NewGraph()

	startingNodes_A := map[string]string{}

	for _, line := range strings.Split(input, "\n")[2:] {
		if line != "" {
			name := strings.TrimSpace(strings.Split(line, "=")[0])
			values := strings.Split(strings.Split(line, "=")[1][2:len(strings.Split(line, "=")[1])-1], ", ")
			graph.addEdge(name, values)

			if string(name[2]) == "A" {
				startingNodes_A[name] = name
			}
		}
	}

	ind := 0

	foundCycle := map[string]bool{}
	cycleLength := map[string]int64{}

	for key, _ := range startingNodes_A {
		foundCycle[key] = false
	}

	data := []int64{}

	for len(startingNodes_A) != 0 {
		currentInstruction := string(instructions[ind%len(instructions)])

		for key, val := range startingNodes_A {
			startingNodes_A[key] = graph.get(val, currentInstruction)

			if foundCycle[key] {
				if string(startingNodes_A[key][2]) == "Z" {
					data = append(data, cycleLength[key])
					delete(startingNodes_A, key)
				} else {
					cycleLength[key] += 1
				}
			} else {
				if string(startingNodes_A[key][2]) == "Z" {
					foundCycle[key] = true
					cycleLength[key] = 1
				}
			}
		}
		ind += 1
	}

	result := utils.LcmMultipleNumbers(data)

	fmt.Println(result)

}
