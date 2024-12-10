package main

import (
	"fmt"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1_2()
}

type Point struct {
	x int
	y int
}

func Part1_2() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_10/input.txt")

	// 	input = `89010123
	// 78121874
	// 87430965
	// 96549874
	// 45678903
	// 32019012
	// 01329801
	// 10456732`
	//  0 1 2 3 4 5 6 7
	//0 8 9 0 1 0 1 2 3
	//1 7 8 1 2 1 8 7 4
	//2 8 7 4 3 0 9 6 5
	//3 9 6 5 4 9 8 7 4
	//4 4 5 6 7 8 9 0 3
	//5 3 2 0 1 9 0 1 2
	//6 0 1 3 2 9 8 0 1
	//7 1 0 4 5 6 7 3 2

	var grid [][]int

	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, utils.StringToIntSlice(strings.Split(row, "")))
	}

	graph := map[Point][]Point{}

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			if isSafe(grid, i-1, j) {
				if grid[i][j]+1 == grid[i-1][j] {
					graph[Point{i, j}] = append(graph[Point{i, j}], Point{i - 1, j})
				}
			}
			if isSafe(grid, i+1, j) {
				if grid[i][j]+1 == grid[i+1][j] {
					graph[Point{i, j}] = append(graph[Point{i, j}], Point{i + 1, j})
				}
			}
			if isSafe(grid, i, j-1) {
				if grid[i][j]+1 == grid[i][j-1] {
					graph[Point{i, j}] = append(graph[Point{i, j}], Point{i, j - 1})
				}
			}
			if isSafe(grid, i, j+1) {
				if grid[i][j]+1 == grid[i][j+1] {
					graph[Point{i, j}] = append(graph[Point{i, j}], Point{i, j + 1})
				}
			}
		}
	}

	startingPoints := []Point{}

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			if grid[i][j] == 0 {
				startingPoints = append(startingPoints, Point{i, j})
			}
		}
	}

	resultPart1 := 0
	resultPart2 := 0

	for _, val := range startingPoints {
		resultPart1 += findPaths(val, graph, grid)
		resultPart2 += findPathsDFS(graph, grid, val, 0)
	}

	fmt.Println(resultPart1)
	fmt.Println("Part 2")
	fmt.Println(resultPart2)

}

func findPaths(start Point, graph map[Point][]Point, grid [][]int) int {
	que := []Point{start}

	paths := 0

	visited := map[Point]bool{}

	currentValues := map[Point]int{}
	currentValues[start] = 0

	for len(que) > 0 {
		current := que[0]
		que = que[1:]
		currentValue := currentValues[current]

		if currentValue == 9 {
			paths += 1
			continue
		}

		for _, next := range graph[current] {
			nextVal := grid[next.x][next.y]

			if !visited[next] && nextVal == currentValue+1 {
				visited[next] = true
				que = append(que, next)
				currentValues[next] = nextVal
			}
		}

	}

	return paths

}

func findPathsDFS(graph map[Point][]Point, grid [][]int, currentPoint Point, currentValue int) int {
	if currentValue == 9 {
		return 1
	}

	paths := 0

	for _, next := range graph[currentPoint] {
		nextVal := grid[next.x][next.y]

		if nextVal == currentValue+1 {
			paths += findPathsDFS(graph, grid, next, nextVal)
		}
	}

	return paths
}

func isSafe(grid [][]int, i int, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}
