package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
}

type Point struct {
	i int64
	j int64
}

func createTerrain(input string) int64 {
	graph := make(map[Point]Point)

	startingPoint := Point{0, 0}

	currentPoint := startingPoint

	for _, line := range strings.Split(input, "\n") {
		direction := strings.Split(line, " ")[0]
		meters := int64(utils.MustAtoi(strings.Split(line, " ")[1]))

		next_i, next_j := currentPoint.i, currentPoint.j
		switch direction {
		case "R":
			next_j += meters
		case "L":
			next_j -= meters
		case "U":
			next_i -= meters
		case "D":
			next_i += meters
		}

		graph[Point{currentPoint.i, currentPoint.j}] = Point{next_i, next_j}
		currentPoint = Point{next_i, next_j}
	}

	min_i, min_j := int64(0), int64(0)

	for key, _ := range graph {
		min_i = min(min_i, key.i)
		min_j = min(min_j, key.j)
	}

	//move each point by min_i/j to normalize grid
	normalizedGraph := make(map[Point]Point)
	max_i, max_j := int64(0), int64(0)

	for key, val := range graph {
		normalizedKey := Point{key.i + int64(math.Abs(float64(min_i))), key.j + int64(math.Abs(float64(min_j)))}
		normalizedVal := Point{val.i + int64(math.Abs(float64(min_i))), val.j + int64(math.Abs(float64(min_j)))}

		normalizedGraph[normalizedKey] = normalizedVal

		max_i = max(max_i, normalizedKey.i)
		max_i = max(max_i, normalizedVal.i)
		max_j = max(max_j, normalizedKey.j)
		max_j = max(max_j, normalizedVal.j)
	}

	grid := [][]string{}

	for i := int64(0); i <= max_i; i += 1 {
		grid = append(grid, []string{})
		for j := int64(0); j <= max_j; j += 1 {
			grid[i] = append(grid[i], ".")
		}
	}

	for key, val := range normalizedGraph {
		fillLine(grid, key, val)
	}

	//flood fill around to find amount of dots outside of loop
	//top and bottom
	for j := int64(0); j < int64(len(grid[0])); j += 1 {
		floodFill(grid, 0, j)
		floodFill(grid, int64(len(grid))-1, j)
	}

	//left and right
	for i := int64(0); i < int64(len(grid)); i += 1 {
		floodFill(grid, i, 0)
		floodFill(grid, i, int64(len(grid[0]))-1)
	}

	return int64(utils.CountIn2DSlice(grid, func(x string) bool {
		return x == "#" || x == "."
	}))

}

func withinBounds(i, j, width, height int64) bool {
	if 0 <= i && i < height && 0 <= j && j < width {
		return true
	}
	return false
}

func floodFill(grid [][]string, i, j int64) {
	if !withinBounds(i, j, int64(len(grid[0])), int64(len(grid))) {
		return
	}
	if grid[i][j] != "." {
		return
	}

	grid[i][j] = "X"

	floodFill(grid, i+1, j)
	floodFill(grid, i-1, j)
	floodFill(grid, i, j+1)
	floodFill(grid, i, j-1)
}

func fillLine(grid [][]string, point1, point2 Point) {
	for i := min(point1.i, point2.i); i <= max(point1.i, point2.i); i += 1 {
		for j := min(point1.j, point2.j); j <= max(point1.j, point2.j); j += 1 {
			grid[i][j] = "#"
		}
	}

}
func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_18/input.txt")

	result := createTerrain(input)

	fmt.Println(result)
}

func createTerrainPart2(input string) int64 {
	directionMap := map[int]string{
		0: "R",
		1: "D",
		2: "L",
		3: "U",
	}

	_ = directionMap

	startingPoint := Point{0, 0}

	currentPoint := startingPoint
	order := []Point{startingPoint}
	perimeter := int64(0)

	for _, line := range strings.Split(input, "\n") {
		color := strings.Split(line, " ")[2]
		color = color[2 : len(color)-1]

		meters, _ := strconv.ParseInt(color[:5], 16, 64)
		direction := directionMap[utils.MustAtoi(color[5:])]

		next_i, next_j := currentPoint.i, currentPoint.j
		switch direction {
		case "R":
			next_j += meters
		case "L":
			next_j -= meters
		case "U":
			next_i -= meters
		case "D":
			next_i += meters
		}

		perimeter += meters
		currentPoint = Point{next_i, next_j}
		order = append(order, currentPoint)
	}

	//Calculate area using shoelaces formula + Pick's theorem
	area := shoelacesFormula(order)

	// Pick's theorem: https://en.wikipedia.org/wiki/Pick%27s_theorem
	// from a = i + b/2 - 1 to => b + i = a + b/2 + 1
	totalPoints := int64(area) + perimeter/2 + 1
	return totalPoints
}

func shoelacesFormula(points []Point) float64 {
	points = append(points, points[0])
	sum := int64(0)
	for index := 0; index < len(points)-1; index += 1 {
		sum += points[index].i*points[index+1].j - points[index+1].i*points[index].j
	}

	if sum < 0 {
		sum = -sum
	}
	return float64(sum) / 2
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_18/input.txt")
	result := createTerrainPart2(input)
	fmt.Println(result)
}
