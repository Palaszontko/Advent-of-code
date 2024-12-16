package main

import (
	"fmt"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

type Point struct {
	x int
	y int
}

func newPoint(values []int) Point {
	return Point{values[0], values[1]}
}

type Vector struct {
	x int
	y int
}

func newVector(values []int) Vector {
	return Vector{values[0], values[1]}
}

type Robot struct {
	position Point
	move     Vector
}

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_14/input.txt")

	inputSplit := strings.Split(input, "\n")

	robots := []Robot{}

	for _, row := range inputSplit {
		point := newPoint(utils.StringToIntSlice(strings.Split(strings.Split(row, " ")[0][2:], ",")))
		vector := newVector(utils.StringToIntSlice(strings.Split(strings.Split(row, " ")[1][2:], ",")))

		robots = append(robots, Robot{point, vector})
		// fmt.Println(point, vector)
	}

	grid := make([][]int, 103)

	for i := 0; i < len(grid); i += 1 {
		grid[i] = make([]int, 101)
	}

	for _, robot := range robots {
		grid[robot.position.y][robot.position.x] += 1
	}

	for i := 0; i < 100; i += 1 {

		for j := 0; j < len(robots); j += 1 {
			moveRobot(&robots[j], &grid)

		}
	}

	leftUp := 0
	leftDown := 0

	rightUp := 0
	rightDown := 0

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[0]); j += 1 {
			if i == len(grid)/2 || j == len(grid[0])/2 {
				// fmt.Print(" ")
			} else if grid[i][j] == 0 {
				// fmt.Print(".")
			} else {
				// fmt.Print(grid[i][j])

				if i < len(grid)/2 && j < len(grid[0])/2 {
					leftUp += grid[i][j]
				} else if i < len(grid)/2 && j > len(grid[0])/2 {
					rightUp += grid[i][j]
				} else if i > len(grid)/2 && j < len(grid[0])/2 {
					leftDown += grid[i][j]
				} else {
					rightDown += grid[i][j]
				}
			}
		}
		// fmt.Println()
	}

	result := leftDown * leftUp * rightDown * rightUp

	fmt.Println(result)

}

func moveRobot(robot *Robot, grid *[][]int) {
	x := robot.position.x
	y := robot.position.y

	vector_x := robot.move.x
	vector_y := robot.move.y

	move_x := (x + vector_x) % len((*grid)[0])
	move_y := (y + vector_y) % len(*grid)

	if move_x < 0 {
		move_x += len((*grid)[0])
	} else if move_x >= len((*grid)[0]) {
		move_x -= len((*grid)[0])
	}

	if move_y < 0 {
		move_y += len(*grid)
	} else if move_y >= len(*grid) {
		move_y -= len(*grid)
	}

	if (*grid)[y][x] <= 1 {
		(*grid)[y][x] = 0
	} else {
		(*grid)[y][x] -= 1
	}

	(*grid)[move_y][move_x] += 1

	robot.position.x = move_x
	robot.position.y = move_y

}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_14/input.txt")

	inputSplit := strings.Split(input, "\n")

	robots := []Robot{}

	for _, row := range inputSplit {
		point := newPoint(utils.StringToIntSlice(strings.Split(strings.Split(row, " ")[0][2:], ",")))
		vector := newVector(utils.StringToIntSlice(strings.Split(strings.Split(row, " ")[1][2:], ",")))
		robots = append(robots, Robot{point, vector})
	}

	grid := make([][]int, 103)

	for i := 0; i < len(grid); i += 1 {
		grid[i] = make([]int, 101)
	}

	for _, robot := range robots {
		grid[robot.position.y][robot.position.x] += 1
	}

	for i := 0; i < 100_000; i += 1 {
		for j := 0; j < len(robots); j += 1 {
			moveRobot(&robots[j], &grid)

		}
		if findLine(grid) {
			fmt.Println(i)
			printGrid(grid)
		}

	}

}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, v := range row {
			if v == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func findLine(grid [][]int) bool {
	for _, row := range grid {
		str := ""

		for _, v := range row {
			if v == 0 {
				str += "."
			} else {
				str += "#"
			}
		}

		if strings.Contains(str, "########") {
			return true
		}
	}
	return false
}
