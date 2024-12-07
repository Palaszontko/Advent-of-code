package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	Part2()
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

func Part2() {
	fmt.Println("Part 2")

	start := time.Now()
	input := utils.ReadFile("cmd/2024/day_6/input.txt")

	grid := [][]byte{}

	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, []byte(row))
	}

	_, _, _, visitedSquares := moveGuard(grid)

	guard_index_i, guard_index_j := findGuard(grid)

	result := 0

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[i]); j += 1 {
			if visitedSquares[i][j] > 0 {
				grid[i][j] = '#'

				if findCycle(grid) {
					// for _, row := range grid {
					// 	fmt.Printf("%s\n", strings.Split(string(row), ""))
					// }
					result += 1
				}

				grid[i][j] = '.'
				grid[guard_index_i][guard_index_j] = '^'
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("%v - Execution completed in: %02dh:%02dm:%02ds:%03dms (%v)\n",
		result,
		int(elapsed.Hours()),
		int(elapsed.Minutes())%60,
		int(elapsed.Seconds())%60,
		elapsed.Milliseconds()%1000,
		elapsed,
	)
}

func findCycle(grid [][]byte) bool {
	i, j := findGuard(grid)

	visitedCorners := make([][]int, len(grid))

	for i := range visitedCorners {
		visitedCorners[i] = make([]int, len(grid[0]))
	}

	direction := up

	for isSafe(grid, i, j) {
		switch direction {
		case up:
			if isSafe(grid, i-1, j) {
				if grid[i-1][j] == '#' {
					direction = (direction + 1) % 4
					visitedCorners[i][j] += 1
					if utils.Contains2DSlice(visitedCorners, func(val int) bool { return val == 3 }) {
						return true
					}
				} else {
					i -= 1
				}
			} else {
				return false
			}
		case down:
			if isSafe(grid, i+1, j) {
				if grid[i+1][j] == '#' {
					direction = (direction + 1) % 4
					visitedCorners[i][j] += 1
					if utils.Contains2DSlice(visitedCorners, func(val int) bool { return val == 3 }) {
						return true
					}
				} else {
					i += 1
				}
			} else {
				return false
			}
		case left:
			if isSafe(grid, i, j-1) {
				if grid[i][j-1] == '#' {
					direction = (direction + 1) % 4
					visitedCorners[i][j] += 1
					if utils.Contains2DSlice(visitedCorners, func(val int) bool { return val == 3 }) {
						return true
					}
				} else {
					j -= 1
				}
			} else {
				return false
			}
		case right:
			if isSafe(grid, i, j+1) {
				if grid[i][j+1] == '#' {
					direction = (direction + 1) % 4
					visitedCorners[i][j] += 1
					if utils.Contains2DSlice(visitedCorners, func(val int) bool { return val == 3 }) {
						return true
					}
				} else {
					j += 1
				}
			} else {
				return false
			}
		}
	}
	return false

}
