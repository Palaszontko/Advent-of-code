package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_5/input.txt")

	inputSplit := strings.Split(input, "\n")

	rules := [][]int{}

	orders := [][]int{}

	for _, line := range inputSplit {
		if strings.Contains(line, "|") {
			rules = append(rules, utils.StringToIntSlice(strings.Split(line, "|")))
		} else if strings.Contains(line, ",") {
			orders = append(orders, utils.StringToIntSlice(strings.Split(line, ",")))
		}
	}

	graphSize := slices.Max(slices.Concat(rules...)) + 1

	graph := utils.CreateGraph(graphSize)

	for _, pair := range rules {
		graph.AddEdge(pair[0], pair[1])

	}

	result := 0

	for _, order := range orders {
		check := func(order []int) bool {
			for i := 0; i < len(order)-1; i++ {
				if !graph.HasEdge(order[i], order[i+1]) {
					return false
				}
			}
			return true
		}
		if check(order) {
			result += order[len(order)/2]
		}
	}

	fmt.Println(result)

}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_5/input.txt")

	inputSplit := strings.Split(input, "\n")

	rules := [][]int{}

	orders := [][]int{}

	for _, line := range inputSplit {
		if strings.Contains(line, "|") {
			rules = append(rules, utils.StringToIntSlice(strings.Split(line, "|")))
		} else if strings.Contains(line, ",") {
			orders = append(orders, utils.StringToIntSlice(strings.Split(line, ",")))
		}
	}

	graphSize := slices.Max(slices.Concat(rules...)) + 1

	graph := utils.CreateGraph(graphSize)

	for _, pair := range rules {
		graph.AddEdge(pair[0], pair[1])
	}

	result := 0

	for _, order := range orders {
		amountOfIncomingEdges := map[int]int{}

		for _, val := range order {
			amountOfIncomingEdges[val] = 0
		}

		for _, val1 := range order {
			for _, val2 := range order {
				if graph.HasEdge(val2, val1) {
					amountOfIncomingEdges[val1] += 1
				}
			}
		}

		newOrder := make([]int, len(order))

		copy(newOrder, order)

		slices.SortFunc(newOrder, func(i, j int) int {
			return amountOfIncomingEdges[i] - amountOfIncomingEdges[j]
		})

		if slices.Compare(order, newOrder) != 0 {
			result += newOrder[len(newOrder)/2]
		}

	}

	fmt.Println(result)

}
