package main

import (
	"fmt"
	"github.com/Palaszontko/advent-of-code/cmd/utils"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	// Part2()
}

type Direction int

const (
	up Direction = iota
	right
	down
	left
)

func Part1() {
	fmt.Println("Part 1")
	input := utils.ReadFile("cmd/2024/day_6/input.txt")

	grid := [][]byte{}
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []byte(line))
	}

	visitedSquares := make([][]int, len(grid))
	for i := range visitedSquares {
		visitedSquares[i] = make([]int, len(grid[0]))
	}

	i, j, direction := findGuard(grid)
	for !moveGuard(&grid, i, j, direction, &visitedSquares) {
		// printGridAnimation(grid)
		i, j, direction = findGuard(grid)
	}

	amount := utils.SliceSum2D(visitedSquares) + 1

	fmt.Println(amount)

}

// func printGridAnimation(grid [][]byte) {
// 	fmt.Print("\033[2J")
// 	fmt.Print("\033[H")

// 	for _, line := range grid {
// 		fmt.Printf("%s\n", line)
// 	}

// 	time.Sleep(20 * time.Millisecond)
// }

func findGuard(grid [][]byte) (int, int, Direction) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case '^':
				return i, j, up
			case 'v':
				return i, j, down
			case '>':
				return i, j, right
			case '<':
				return i, j, left
			default:
				continue
			}
		}
	}
	return -1, -1, -1
}

// return true if guard escaped
func moveGuard(grid *[][]byte, i int, j int, direction Direction, visitedSquares *[][]int) bool {
	mapDirection := map[Direction]byte{up: '^', down: 'v', left: '<', right: '>'}

	switch direction {
	case up:
		if !isSafe(*grid, i-1, j) {
			return true
		}
		if (*grid)[i-1][j] == '#' {
			(*grid)[i][j] = mapDirection[(direction+1)%4]
		} else if (*grid)[i-1][j] == '.' {
			(*grid)[i][j], (*grid)[i-1][j] = (*grid)[i-1][j], (*grid)[i][j]
			(*visitedSquares)[i][j] = 1
		}
	case right:
		if !isSafe(*grid, i, j+1) {
			return true
		}
		if (*grid)[i][j+1] == '#' {
			(*grid)[i][j] = mapDirection[(direction+1)%4]
		} else if (*grid)[i][j+1] == '.' {
			(*grid)[i][j], (*grid)[i][j+1] = (*grid)[i][j+1], (*grid)[i][j]
			(*visitedSquares)[i][j] = 1
		}
	case down:
		if !isSafe(*grid, i+1, j) {
			return true
		}
		if (*grid)[i+1][j] == '#' {
			(*grid)[i][j] = mapDirection[(direction+1)%4]
		} else if (*grid)[i+1][j] == '.' {
			(*grid)[i][j], (*grid)[i+1][j] = (*grid)[i+1][j], (*grid)[i][j]
			(*visitedSquares)[i][j] = 1
		}
	case left:
		if !isSafe(*grid, i, j-1) {
			return true
		}
		if (*grid)[i][j-1] == '#' {
			(*grid)[i][j] = mapDirection[(direction+1)%4]
		} else if (*grid)[i][j-1] == '.' {
			(*grid)[i][j], (*grid)[i][j-1] = (*grid)[i][j-1], (*grid)[i][j]
			(*visitedSquares)[i][j] = 1
		}
	}

	return false
}

func isSafe(grid [][]byte, i, j int) bool {
	if i >= 0 && i < len(grid) {
		if j >= 0 && j < len(grid[0]) {
			return true
		}
	}
	return false
}
