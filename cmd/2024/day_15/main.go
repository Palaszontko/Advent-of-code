package main

import (
	"fmt"
	"strings"

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
	down
	left
	right
)

func giveDirection(move byte) Direction {
	directionMap := map[byte]Direction{
		'^': up,
		'v': down,
		'>': right,
		'<': left,
	}

	return directionMap[move]
}
func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_15/input.txt")

	grid := [][]byte{}
	moves := []byte{}

	isGrid := true

	for _, row := range strings.Split(input, "\n") {
		if isGrid {
			if row == "" {
				isGrid = false
				continue
			}
			grid = append(grid, []byte(row))
		} else {
			moves = append(moves, []byte(row)...)
		}
	}

	for i := 0; i < len(moves); i += 1 {
		moveRobot(&grid, moves[i])
		// fmt.Println(string(moves[i]))
		// printGrid(grid)
	}

	result := findAllBoxesGPSSum(grid)

	fmt.Println(result)

}

func moveRobot(grid *[][]byte, directionByte byte) {
	robot_i, robot_j := findRobot(*grid)

	direction := giveDirection(directionByte)

	switch direction {
	case up:
		if isWall(*grid, robot_i-1, robot_j) {
			return
		} else if isDot(*grid, robot_i-1, robot_j) {
			(*grid)[robot_i][robot_j], (*grid)[robot_i-1][robot_j] = (*grid)[robot_i-1][robot_j], (*grid)[robot_i][robot_j]
			return
		} else if isBox(*grid, robot_i-1, robot_j) {
			k := 1
			boxes := 0
			for isBox(*grid, robot_i-k, robot_j) {
				boxes += 1
				k += 1
			}
			lastBox_i, lastBox_j := robot_i-boxes, robot_j

			if isDot(*grid, lastBox_i-1, lastBox_j) {
				for i := 0; i < boxes+1; i += 1 {
					(*grid)[lastBox_i][lastBox_j], (*grid)[lastBox_i-1][lastBox_j] = (*grid)[lastBox_i-1][lastBox_j], (*grid)[lastBox_i][lastBox_j]
					lastBox_i += 1
				}
			}
		}
	case down:

		if isWall(*grid, robot_i+1, robot_j) {
			return
		} else if isDot(*grid, robot_i+1, robot_j) {
			(*grid)[robot_i][robot_j], (*grid)[robot_i+1][robot_j] = (*grid)[robot_i+1][robot_j], (*grid)[robot_i][robot_j]
			return
		} else if isBox(*grid, robot_i+1, robot_j) {

			k := 1
			boxes := 0
			for isBox(*grid, robot_i+k, robot_j) {
				boxes += 1
				k += 1
			}
			lastBox_i, lastBox_j := robot_i+boxes, robot_j

			if isDot(*grid, lastBox_i+1, lastBox_j) {
				for i := 0; i < boxes+1; i += 1 {
					(*grid)[lastBox_i][lastBox_j], (*grid)[lastBox_i+1][lastBox_j] = (*grid)[lastBox_i+1][lastBox_j], (*grid)[lastBox_i][lastBox_j]
					lastBox_i -= 1
				}
			}
		}
	case left:
		if isWall(*grid, robot_i, robot_j-1) {
			return
		} else if isDot(*grid, robot_i, robot_j-1) {
			(*grid)[robot_i][robot_j], (*grid)[robot_i][robot_j-1] = (*grid)[robot_i][robot_j-1], (*grid)[robot_i][robot_j]
			return
		} else if isBox(*grid, robot_i, robot_j-1) {
			k := 1
			boxes := 0
			for isBox(*grid, robot_i, robot_j-k) {
				boxes += 1
				k += 1
			}
			lastBox_i, lastBox_j := robot_i, robot_j-boxes

			if isDot(*grid, lastBox_i, lastBox_j-1) {
				for i := 0; i < boxes+1; i += 1 {
					(*grid)[lastBox_i][lastBox_j], (*grid)[lastBox_i][lastBox_j-1] = (*grid)[lastBox_i][lastBox_j-1], (*grid)[lastBox_i][lastBox_j]
					lastBox_j += 1
				}
			}
		}
	case right:
		if isWall(*grid, robot_i, robot_j+1) {
			return
		} else if isDot(*grid, robot_i, robot_j+1) {
			(*grid)[robot_i][robot_j], (*grid)[robot_i][robot_j+1] = (*grid)[robot_i][robot_j+1], (*grid)[robot_i][robot_j]
			return
		} else if isBox(*grid, robot_i, robot_j+1) {
			k := 1
			boxes := 0
			for isBox(*grid, robot_i, robot_j+k) {
				boxes += 1
				k += 1
			}
			lastBox_i, lastBox_j := robot_i, robot_j+boxes

			if isDot(*grid, lastBox_i, lastBox_j+1) {
				for i := 0; i < boxes+1; i += 1 {
					(*grid)[lastBox_i][lastBox_j], (*grid)[lastBox_i][lastBox_j+1] = (*grid)[lastBox_i][lastBox_j+1], (*grid)[lastBox_i][lastBox_j]
					lastBox_j -= 1
				}
			}
		}
	}

}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Printf("%s\n", row)
	}
}

func findAllBoxesGPSSum(grid [][]byte) int {
	result := 0

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			if grid[i][j] == 'O' {
				result += 100*i + j
			}
		}
	}

	return result
}

func isWall(grid [][]byte, i, j int) bool {
	return grid[i][j] == '#'
}

func isBox(grid [][]byte, i, j int) bool {
	return grid[i][j] == 'O'
}

func isDot(grid [][]byte, i, j int) bool {
	return grid[i][j] == '.'
}

func printMoves(moves []byte) {
	fmt.Printf("%s\n", moves)
}

func findRobot(grid [][]byte) (int, int) {
	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			if grid[i][j] == '@' {
				return i, j
			}
		}
	}

	return -1, -1
}

func Part2() {
	fmt.Println("Part 2")

	input := `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

	grid := [][]byte{}
	moves := []byte{}

	isGrid := true

	for _, row := range strings.Split(input, "\n") {
		if isGrid {
			if row == "" {
				isGrid = false
				continue
			}
			grid = append(grid, []byte(row))
		} else {
			moves = append(moves, []byte(row)...)
		}
	}

	wideGrid := convertToWideGrid(grid)

	printGrid(wideGrid)

}

func convertToWideGrid(grid [][]byte) [][]byte {
	gridWide := []byte{}

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			switch grid[i][j] {
			case '#':
				gridWide = append(gridWide, []byte{'#', '#'}...)
			case 'O':
				gridWide = append(gridWide, []byte{'[', ']'}...)
			case '.':
				gridWide = append(gridWide, []byte{'.', '.'}...)
			case '@':
				gridWide = append(gridWide, []byte{'@', '.'}...)
			}
		}
	}

	gridWide2D := [][]byte{}

	for i := 0; i < len(gridWide); i += len(grid[0]) * 2 {
		gridWide2D = append(gridWide2D, gridWide[i:i+len(grid[0])*2])
	}

	return gridWide2D

}

func isBoxWide(grid [][]byte, i, j int) bool {
	return grid[i][j] == '[' || grid[i][j] == ']'
}

func isBoxLeft(grid [][]byte, i, j int) bool {
	return grid[i][j] == '['
}

func isBoxRight(grid [][]byte, i, j int) bool {
	return grid[i][j] == ']'
}
