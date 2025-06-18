package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	// Part2()
}

type Point struct {
	i int
	j int
}

type Graph struct {
	graph map[Point][]Point
}

func NewGraph() *Graph {
	return &Graph{
		graph: make(map[Point][]Point),
	}
}

func (g *Graph) addEdge(point1, point2 Point) {
	g.graph[point1] = append(g.graph[point1], point2)
}

func (g *Graph) hasEdge(point1, point2 Point) bool {
	return slices.Contains(g.graph[point1], point2)
}

func (g *Graph) getNeigbours(point1 Point) []Point {
	return g.graph[point1]
}

func findStartingPoint(grid [][]string) (i, j int) {
	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			if grid[i][j] == "S" {
				return i, j
			}
		}
	}
	return -1, -1
}

func withinBounds(i, j, width, height int) bool {
	if 0 <= i && i < height && 0 <= j && j < width {
		return true
	}
	return false
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

func checkDirection(i, j int) direction {
	if i == -1 {
		return up
	}

	if i == 1 {
		return down
	}

	if j == -1 {
		return left
	}

	if j == 1 {
		return right
	}

	return up
}

func checkConnection(currentPipe, checkPipe string, i, j int) bool {

	upPipes := map[string]bool{
		"|": true,
		"L": true,
		"J": true,
		"S": true,
	}

	downPipes := map[string]bool{
		"|": true,
		"7": true,
		"F": true,
		"S": true,
	}

	leftPipes := map[string]bool{
		"-": true,
		"J": true,
		"7": true,
		"S": true,
	}

	rightPipes := map[string]bool{
		"-": true,
		"L": true,
		"F": true,
		"S": true,
	}

	direction := checkDirection(i, j)

	switch direction {
	case up:
		return upPipes[currentPipe] && downPipes[checkPipe]
	case down:
		return downPipes[currentPipe] && upPipes[checkPipe]
	case left:
		return leftPipes[currentPipe] && rightPipes[checkPipe]
	case right:
		return rightPipes[currentPipe] && leftPipes[checkPipe]
	}

	return false
}

func Part1() {
	fmt.Println("Part 1")

	input := `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`
	input = utils.ReadFile("cmd/2023/day_10/input.txt")

	var grid [][]string

	for i, row := range strings.Split(input, "\n") {
		grid = append(grid, make([]string, 0))
		for _, val := range strings.Split(row, "") {
			grid[i] = append(grid[i], val)
		}
	}

	neighbours := [][]int{
		{-1, 0},
		{0, -1}, {0, 1},
		{1, 0},
	}

	height := len(grid)
	width := len(grid[0])

	graph := NewGraph()

	//create all possible graphs
	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			currentPipe := grid[i][j]
			for _, neighbour := range neighbours {
				if withinBounds(i+neighbour[0], j+neighbour[1], width, height) {
					pipe := grid[i+neighbour[0]][j+neighbour[1]]
					if checkConnection(currentPipe, pipe, neighbour[0], neighbour[1]) {
						graph.addEdge(Point{i, j}, Point{i + neighbour[0], j + neighbour[1]})
					}
				}
			}
		}
	}

	start_i, start_j := findStartingPoint(grid)
	startingPoint := Point{start_i, start_j}

	que := []Point{startingPoint}
	distances := map[Point]int{}

	for len(que) != 0 {
		current := que[0]
		que = que[1:]

		for _, neighbour := range graph.getNeigbours(current) {
			if _, ok := distances[neighbour]; !ok {
				distances[neighbour] = distances[current] + 1
				que = append(que, neighbour)
			}
		}
	}

	result := -1
	for _, val := range distances {
		if val > result {
			result = val
		}
	}

	fmt.Println(result)
}

func Part2() {
	fmt.Println("Part 2")
}
