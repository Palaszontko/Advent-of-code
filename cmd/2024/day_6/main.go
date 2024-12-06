package main

import (
	"fmt"
	"github.com/Palaszontko/advent-of-code/cmd/utils"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
}

type Direction int

const (
	up Direction = iota
	right
	down
	left
)

func Part1() {
	input := utils.ReadFile("cmd/2024/day_6/input.txt")

	grid := [][]byte{}

	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, []byte(row))
	}

	_, _, _, visitedSquares := moveGuard(grid)

	result := utils.CountIn2DSlice(visitedSquares, func(val int) bool { return val > 0 })

	fmt.Println(result)

}

func moveGuard(grid [][]byte) (bool, int, int, [][]int) {

	visitedSquares := make([][]int, len(grid))

	for i := range visitedSquares {
		visitedSquares[i] = make([]int, len(grid[0]))
	}

	i, j := findGuard(grid)
	visitedSquares[i][j] += 1

	direction := up

	for isSafe(grid, i, j) {

		switch direction {
		case up:
			if isSafe(grid, i-1, j) {
				if grid[i-1][j] == '#' {
					direction = (direction + 1) % 4
				} else {
					i -= 1
				}
			} else {
				return true, i, j, visitedSquares
			}
		case down:
			if isSafe(grid, i+1, j) {
				if grid[i+1][j] == '#' {
					direction = (direction + 1) % 4
				} else {
					i += 1
				}
			} else {
				return true, i, j, visitedSquares
			}
		case left:
			if isSafe(grid, i, j-1) {
				if grid[i][j-1] == '#' {
					direction = (direction + 1) % 4
				} else {
					j -= 1
				}
			} else {
				return true, i, j, visitedSquares
			}
		case right:
			if isSafe(grid, i, j+1) {
				if grid[i][j+1] == '#' {
					direction = (direction + 1) % 4
				} else {
					j += 1
				}
			} else {
				return true, i, j, visitedSquares
			}
		}

		visitedSquares[i][j] += 1

	}

	return true, i, j, visitedSquares

}

func isSafe(grid [][]byte, i int, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0])
}

func findGuard(grid [][]byte) (int, int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case '^':
				return i, j
			case 'v':
				return i, j
			case '>':
				return i, j
			case '<':
				return i, j
			default:
				continue
			}
		}
	}
	return -1, -1
}
