package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	// Part2()
}

type Point struct {
	i int
	j int
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_16/input.txt")

	grid := [][]byte{}

	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, []byte(row))
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

	keys := []Point{}

	for key, _ := range graph {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if keys[i].i == keys[j].i {
			return keys[i].j < keys[j].j
		}
		return keys[i].i < keys[j].i
	})

	i, j := locateStart(grid)
	start := Point{i, j}

	paths := FindAllPaths(graph, start, grid)

	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	result := math.MaxInt64

	for _, path := range paths {
		moves, rot := countMoves(path)
		if moves*1+rot*1000 < result {
			result = moves*1 + rot*1000
		}
	}

	fmt.Println(result)
}

type Direction int

const (
	up Direction = iota
	down
	left
	right
)

// 1st return is steps forward 2nd is amount of turns
func countMoves(path []Point) (int, int) {

	moves := []Direction{}
	rotations := 0

	for i := 0; i < len(path)-1; i += 1 {
		x1, y1 := path[i].i, path[i].j
		x2, y2 := path[i+1].i, path[i+1].j

		if x1 == x2 {
			if y1+1 == y2 {
				moves = append(moves, down)
			} else if y1-1 == y2 {
				moves = append(moves, up)
			}
		} else if y1 == y2 {
			if x1+1 == x2 {
				moves = append(moves, right)
			} else if x1-1 == x2 {
				moves = append(moves, left)
			}
		}
	}

	for i := 0; i < len(moves)-1; i += 1 {
		if moves[i] != moves[i+1] {
			rotations += 1
		}
	}

	if path[0].j+1 != path[1].j {
		rotations += 1
	}

	return len(moves), rotations
}

func printPathOnGrid(grid [][]byte, path []Point) {
	copyGrid := make([][]byte, len(grid))

	for i := range grid {
		copyGrid[i] = make([]byte, len(grid[i]))
		copy(copyGrid[i], grid[i])
	}

	for _, point := range path {
		copyGrid[point.i][point.j] = 'x'
	}

	printMaze(copyGrid)
}

func FindAllPaths(graph map[Point][]Point, start Point, grid [][]byte) [][]Point {
	allPaths := [][]Point{}
	visited := map[Point]bool{}

	currentPath := []Point{start}

	var dfs func(current Point)
	dfs = func(current Point) {
		currentPath = append(currentPath, current)
		visited[current] = true

		if grid[current.i][current.j] == 'E' {
			pathCopy := make([]Point, len(currentPath))
			copy(pathCopy, currentPath)
			allPaths = append(allPaths, pathCopy)
		} else {
			for _, next := range graph[current] {
				if !visited[next] {
					dfs(next)
				}
			}
		}

		currentPath = currentPath[:len(currentPath)-1]
		visited[current] = false

	}

	dfs(start)
	return allPaths
}

func isSafe(grid [][]byte, i int, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func printMaze(grid [][]byte) {
	for _, v := range grid {
		fmt.Printf("%s\n", v)
	}
}

func locateStart(grid [][]byte) (int, int) {
	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			if grid[i][j] == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func Part2() {
	fmt.Println("Part 2")
}
