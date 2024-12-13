package main

import (
	"fmt"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	// Part2()
}

type Point struct {
	x  int
	y  int
	id string
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_12/input.txt")

	var grid [][]string

	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(row, ""))
	}

	graph := map[Point][]Point{}

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			graph[Point{i, j, grid[i][j]}] = []Point{}
		}
	}

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			check := func(a int, b int) {
				if isSafe(grid, i+a, j+b) {
					if grid[i][j] == grid[i+a][j+b] {
						graph[Point{i, j, grid[i][j]}] = append(graph[Point{i, j, grid[i][j]}], Point{i + a, j + b, grid[i][j]})
					}
				}
			}
			//up
			check(-1, 0)
			//down
			check(1, 0)
			//left
			check(0, -1)
			//right
			check(0, 1)

		}
	}

	discovered := map[Point]bool{}

	for key, _ := range graph {
		discovered[key] = false
	}

	result := 0

	for key, _ := range graph {
		if fences, plants := DFS1(graph, key, &discovered); fences != 0 && discovered[key] {
			result += (fences * plants)
		}
	}

	fmt.Println(result)
}

// return1 is sumOfFences and return2 is amount of plants
func DFS1(graph map[Point][]Point, current Point, discovered *map[Point]bool) (int, int) {
	S := []Point{current}

	amountOfFences := 0
	amountOfPlants := 0

	for len(S) > 0 {
		v := S[len(S)-1]
		S = S[:len(S)-1]

		if !(*discovered)[v] {
			(*discovered)[v] = true

			amountOfPlants += 1
			amountOfFences += (4 - len(graph[v]))

			for _, w := range graph[v] {
				S = append(S, w)
			}
		}
	}

	return amountOfFences, amountOfPlants
}

func isSafe(grid [][]string, i int, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func Part2() {
	fmt.Println("Part 2")

}
