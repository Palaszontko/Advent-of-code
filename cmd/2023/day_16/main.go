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

type Point struct {
	i int
	j int
}
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

var directionMap = map[Direction][2]int{
	Up:    {-1, 0},
	Down:  {1, 0},
	Left:  {0, -1},
	Right: {0, 1},
}

func withinBounds(i, j, width, height int) bool {
	if 0 <= i && i < height && 0 <= j && j < width {
		return true
	}
	return false
}

func energizedPath(grid [][]string, i, j int, direction Direction, visitedPoints map[Point]map[Direction]bool) {
	visitedPoints[Point{i, j}][direction] = true

	nextMove := func(dir Direction) {
		next_i := i + directionMap[dir][0]
		next_j := j + directionMap[dir][1]

		if withinBounds(next_i, next_j, len(grid), len(grid[0])) {
			if !visitedPoints[Point{next_i, next_j}][dir] {
				visitedPoints[Point{next_i, next_j}][dir] = true
				energizedPath(grid, next_i, next_j, dir, visitedPoints)
			}
		}

	}

	currentTile := grid[i][j]

	if currentTile == "." {
		nextMove(direction)
	} else {
		switch currentTile {
		case "|":
			switch direction {
			case Right, Left:
				nextMove(Down)
				nextMove(Up)
			case Up:
				nextMove(Up)
			case Down:
				nextMove(Down)
			}

		case "-":
			switch direction {
			case Up, Down:
				nextMove(Right)
				nextMove(Left)
			case Left:
				nextMove(Left)
			case Right:
				nextMove(Right)
			}
		case "/":
			switch direction {
			case Right:
				nextMove(Up)
			case Left:
				nextMove(Down)
			case Down:
				nextMove(Left)
			case Up:
				nextMove(Right)
			}

		case `\`:
			switch direction {
			case Right:
				nextMove(Down)
			case Left:
				nextMove(Up)
			case Down:
				nextMove(Right)
			case Up:
				nextMove(Left)
			}
		}

	}

}

func Part1() {
	fmt.Println("Part 1")

	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	input = utils.ReadFile("cmd/2023/day_16/input.txt")

	var grid [][]string

	for i, row := range strings.Split(input, "\n") {
		grid = append(grid, make([]string, len(row)))
		for j, ch := range row {
			grid[i][j] = string(ch)
		}
	}

	visitedPoints := make(map[Point]map[Direction]bool)
	for i := range grid {
		for j := range grid[i] {
			visitedPoints[Point{i, j}] = make(map[Direction]bool)
		}
	}
	energizedPath(grid, 0, 0, Right, visitedPoints)

	result := 0

	for _, directions := range visitedPoints {
		if len(directions) > 0 {
			result += 1
		}
	}
	fmt.Println(result)

}

func Part2() {
	fmt.Println("Part 2")
}
