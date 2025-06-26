package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
}

func simulateGravity(transGrid *[][]string) {
	for i := 0; i < len(*transGrid); i += 1 {
		for j := 0; j < len((*transGrid)[i]); j += 1 {
			if (*transGrid)[i][j] == "#" {
				continue
			}
			sortTo := len((*transGrid)[i])

			for k := j + 1; k < len((*transGrid)[i]); k += 1 {
				if (*transGrid)[i][k] == "#" {
					sortTo = k
					break
				}
			}

			sort.Slice((*transGrid)[i][j:sortTo], func(i2, j2 int) bool {
				return (*transGrid)[i][j+i2] == "O" && (*transGrid)[i][j+j2] == "."
			})

			j = sortTo
		}
	}
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_14/input.txt")

	var grid [][]string

	for i, row := range strings.Split(input, "\n") {
		grid = append(grid, make([]string, 0))
		for _, val := range strings.Split(row, "") {
			grid[i] = append(grid[i], val)
		}
	}

	transGrid := utils.Transpose2D(grid)

	simulateGravity(&transGrid)

	start := len(transGrid)
	result := 0
	for i := 0; i < len(transGrid[0]); i += 1 {
		amount := 0
		for j := 0; j < len(transGrid); j += 1 {
			if transGrid[j][i] == "O" {
				amount += 1
			}
		}
		result += amount * (start - i)
	}

	fmt.Println(result)

}

func reverseEachColumn(grid *[][]string) {
	for left, right := 0, len(*grid)-1; left < right; left, right = left+1, right-1 {
		(*grid)[left], (*grid)[right] = (*grid)[right], (*grid)[left]
	}
}

func reverseEachRow(grid *[][]string) {
	for i := 0; i < len(*grid); i++ {
		for left, right := 0, len((*grid)[i])-1; left < right; left, right = left+1, right-1 {
			(*grid)[i][left], (*grid)[i][right] = (*grid)[i][right], (*grid)[i][left]
		}
	}
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_14/input.txt")

	var grid [][]string

	for i, row := range strings.Split(input, "\n") {
		grid = append(grid, make([]string, 0))
		for _, val := range strings.Split(row, "") {
			grid[i] = append(grid[i], val)
		}
	}

	transGrid := utils.DeepCopy2D(grid)

	cycles := map[string]int{}

	for i := 0; i < 1_000_000_000; i++ {
		//north
		transGrid = utils.Transpose2D(transGrid)
		simulateGravity(&transGrid)
		transGrid = utils.Transpose2D(transGrid)
		// printGrid(transGrid)

		//west
		simulateGravity(&transGrid)
		// printGrid(transGrid)

		//south
		reverseEachColumn(&transGrid)
		transGrid = utils.Transpose2D(transGrid)
		simulateGravity(&transGrid)
		transGrid = utils.Transpose2D(transGrid)
		reverseEachColumn(&transGrid)
		// printGrid(transGrid)

		//east
		reverseEachRow(&transGrid)
		simulateGravity(&transGrid)
		reverseEachRow(&transGrid)
		// printGrid(transGrid)

		state := buildStringGrid(transGrid)

		if saw, ok := cycles[state]; ok {
			cycleLength := i - saw
			remaining := (1_000_000_000 - saw - 1) % cycleLength

			for j := 0; j < remaining; j++ {
				//north
				transGrid = utils.Transpose2D(transGrid)
				simulateGravity(&transGrid)
				transGrid = utils.Transpose2D(transGrid)
				// printGrid(transGrid)

				//west
				simulateGravity(&transGrid)
				// printGrid(transGrid)

				//south
				reverseEachColumn(&transGrid)
				transGrid = utils.Transpose2D(transGrid)
				simulateGravity(&transGrid)
				transGrid = utils.Transpose2D(transGrid)
				reverseEachColumn(&transGrid)
				// printGrid(transGrid)

				//east
				reverseEachRow(&transGrid)
				simulateGravity(&transGrid)
				reverseEachRow(&transGrid)
				// printGrid(transGrid)
			}

			result := 0

			for i, row := range transGrid {
				result += (len(transGrid[0]) - i) * utils.CountInSlice(row, "O")
			}

			fmt.Println(result)
			return
		} else {
			cycles[state] = i
		}

	}

}

func buildStringGrid(grid [][]string) string {
	result := ""
	for _, row := range grid {
		result += strings.Join(row, "")
	}

	return result
}
