package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	Part2()
}

type Point struct {
	i int
	j int
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_18/input.txt")

	grid := [][]byte{}

	for i := 0; i < 71; i += 1 {
		grid = append(grid, []byte(strings.Repeat(".", 71)))
	}

	for count, row := range strings.Split(input, "\n") {
		i := utils.MustAtoi(strings.Split(row, ",")[0])
		j := utils.MustAtoi(strings.Split(row, ",")[1])

		grid[j][i] = '#'

		count += 1
		if count == 1024 {
			break
		}

	}

	graph := map[Point][]Point{}

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			connect := func(a, b int) {
				if isSafe(grid, i+a, j+b) {
					if (grid[i][j] == '.' || grid[i][j] == 'S' || grid[i][j] == 'E') && (grid[i+a][j+b] == '.' || grid[i+a][j+b] == 'S' || grid[i+a][j+b] == 'E') {
						graph[Point{i, j}] = append(graph[Point{i, j}], Point{i + a, j + b})
					}
				}
			}
			connect(-1, 0)
			connect(1, 0)
			connect(0, -1)
			connect(0, 1)
		}
	}

	fmt.Println(dijakstra(graph, Point{0, 0}, Point{70, 70}))

}

func dijakstra(graph map[Point][]Point, start, end Point) int {
	distances := map[Point]int{}

	for key := range graph {
		distances[key] = math.MaxInt
	}

	distances[start] = 0

	toVisit := []Point{start}

	for len(toVisit) > 0 {
		current := toVisit[0]
		shortestPath := distances[current]
		lowestIndex := 0

		for i, point := range toVisit {
			if distances[point] < shortestPath {
				current = point
				shortestPath = distances[point]
				lowestIndex = i
			}
		}

		toVisit = append(toVisit[:lowestIndex], toVisit[lowestIndex+1:]...)

		for _, neig := range graph[current] {
			newDistance := distances[current] + 1

			if newDistance < distances[neig] {
				distances[neig] = newDistance
				toVisit = append(toVisit, neig)
			}
		}

	}

	return distances[end]

}

func isSafe(grid [][]byte, i int, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_18/input.txt")

	grid := [][]byte{}

	for i := 0; i < 71; i += 1 {
		grid = append(grid, []byte(strings.Repeat(".", 71)))
	}

	for count, row := range strings.Split(input, "\n") {
		i := utils.MustAtoi(strings.Split(row, ",")[0])
		j := utils.MustAtoi(strings.Split(row, ",")[1])

		grid[j][i] = '#'

		graph := map[Point][]Point{}

		for i := 0; i < len(grid); i += 1 {
			for j := 0; j < len(grid[0]); j += 1 {
				connect := func(a, b int) {
					if isSafe(grid, i+a, j+b) {
						if (grid[i][j] == '.' || grid[i][j] == 'S' || grid[i][j] == 'E') && (grid[i+a][j+b] == '.' || grid[i+a][j+b] == 'S' || grid[i+a][j+b] == 'E') {
							graph[Point{i, j}] = append(graph[Point{i, j}], Point{i + a, j + b})
						}
					}
				}
				connect(-1, 0)
				connect(1, 0)
				connect(0, -1)
				connect(0, 1)
			}
		}

		if dijakstra(graph, Point{0, 0}, Point{70, 70}) == math.MaxInt {
			fmt.Printf("%v,%v\n", i, j)
			break
		}

		count += 1
	}
}
