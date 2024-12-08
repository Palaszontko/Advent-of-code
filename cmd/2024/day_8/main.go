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

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_8/input.txt")

	grid := make([][]byte, len(strings.Split(input, "\n")))

	for i, row := range strings.Split(input, "\n") {
		grid[i] = make([]byte, len(row))
		copy(grid[i], row)
	}

	placeSignasMap := make([][]int, len(grid))

	for i := range placeSignasMap {
		placeSignasMap[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[i]); j += 1 {
			if grid[i][j] != '.' && grid[i][j] != '#' {
				antennas := findAntennas(grid, i, j, grid[i][j])

				for _, antenna := range antennas {
					vector_i, vector_j := antenna[0]-i, antenna[1]-j
					if placeSignal(&grid, i-vector_i, j-vector_j) {
						placeSignasMap[i-vector_i][j-vector_j] += 1
					}
				}
			}
		}
	}

	result := utils.CountIn2DSlice[int](placeSignasMap, func(i int) bool { return i > 0 })

	fmt.Println(result)

}

func findAntennas(grid [][]byte, index_i int, index_j int, antennaId byte) [][]int {
	otherAntennas := [][]int{}

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[i]); j += 1 {
			if grid[i][j] != '.' && grid[i][j] != '#' && (i != index_i || j != index_j) && grid[i][j] == antennaId {
				otherAntennas = append(otherAntennas, []int{i, j})
			}
		}
	}

	return otherAntennas
}

func placeSignal(grid *[][]byte, i int, j int) bool {
	if isSafe(*grid, i, j) {
		if (*grid)[i][j] == '.' {
			(*grid)[i][j] = '#'
		}
		return true
	}
	return false
}

func isSafe(grid [][]byte, i int, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len((grid)[0])
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_8/input.txt")

	input = strings.Replace(input, "#", ".", -1)

	grid := make([][]byte, len(strings.Split(input, "\n")))

	for i, row := range strings.Split(input, "\n") {
		grid[i] = make([]byte, len(row))
		copy(grid[i], row)
	}

	placeSignasMap := make([][]int, len(grid))

	for i := range placeSignasMap {
		placeSignasMap[i] = make([]int, len(grid[0]))
	}

	result := 0

	for i := 0; i < len(grid); i += 1 {
		for j := 0; j < len(grid[i]); j += 1 {
			if grid[i][j] != '.' && grid[i][j] != '#' {
				antennas := findAntennas(grid, i, j, grid[i][j])

				for _, antenna := range antennas {
					placeSignasMap[i][j] += 1
					vector_i, vector_j := antenna[0]-i, antenna[1]-j

					for multp := 1; placeSignal(&grid, i-(multp*vector_i), j-(multp*vector_j)); multp += 1 {
						if isSafe(grid, i-(multp*vector_i), j-(multp*vector_j)) {
							placeSignasMap[i-(multp*vector_i)][j-(multp*vector_j)] += 1
						}
					}
				}
			}
		}
	}

	result += utils.CountIn2DSlice[int](placeSignasMap, func(i int) bool { return i > 0 })

	fmt.Println(result)

}
