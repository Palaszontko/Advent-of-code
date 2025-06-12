package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_11/input.txt")

	var splittedData [][]string

	for _, row := range strings.Split(input, "\n") {
		splittedData = append(splittedData, strings.Split(row, ""))
	}

	galaxy := expandGalaxy(splittedData)

	fmt.Println(sumOfDistancesGalaxies(galaxy))
	//works with second approach with 0 as argument too ;)
	// fmt.Println(calculateDistanceWithExpandIndexes(galaxy, 0))

}

func expandGalaxy(dataOfGalaxy [][]string) [][]string {
	var rowsToExpandIndexes []int

	data := utils.DeepCopy2D(dataOfGalaxy)

	for i, row := range data {
		if strings.Count(strings.Join(row, ""), ".") == len(row) {
			rowsToExpandIndexes = append(rowsToExpandIndexes, i)
		}
	}

	var columnToExpandIndexes []int

	for i := 0; i < len(data[0]); i += 1 {
		var column string
		for j := 0; j < len(data); j += 1 {
			column += data[j][i]
		}
		if strings.Count(column, ".") == len(column) {
			columnToExpandIndexes = append(columnToExpandIndexes, i)
		}
	}

	rowSlice := strings.Split(strings.Repeat(".", len(data[0])), "")

	for step, rowIndex := range rowsToExpandIndexes {
		data = utils.InsertAtIndexIn2DSlice(data, rowSlice, rowIndex+step)
	}

	for step, colIndex := range columnToExpandIndexes {
		for i := 0; i < len(data); i += 1 {
			data[i] = slices.Insert(data[i], colIndex+step, ".")
		}
	}

	return data

}

func findAllGalaxies(data [][]string) map[int][]int {
	galaxyMap := make(map[int][]int)
	id := 0
	for i := 0; i < len(data); i += 1 {
		for j := 0; j < len(data[i]); j += 1 {
			if data[i][j] == "#" {
				galaxyMap[id] = []int{i, j}
				id += 1
			}
		}
	}

	return galaxyMap
}

func manhattannDistance(galaxyA_i int, galaxyA_j int, galaxyB_i int, galaxyB_j int) int {
	return int(math.Abs(float64(galaxyA_i)-float64(galaxyA_j)) + math.Abs(float64(galaxyB_i)-float64(galaxyB_j)))
}

func sumOfDistancesGalaxies(data [][]string) int {
	cords := findAllGalaxies(data)
	sum := 0

	for i := 0; i < len(cords); i += 1 {
		for j := i + 1; j < len(cords); j += 1 {
			sum += manhattannDistance(cords[i][0], cords[j][0], cords[i][1], cords[j][1])
		}
	}

	return sum
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_11/input.txt")

	var galaxy [][]string
	for _, row := range strings.Split(input, "\n") {
		galaxy = append(galaxy, strings.Split(row, ""))
	}

	expandSize := 1000000
	expandSize -= 1

	fmt.Println(calculateDistanceWithExpandIndexes(galaxy, expandSize))

}

type FreeRowsColumns struct {
	Columns []int
	Rows    []int
}

func findFreeRowsAndColumnIndexes(data [][]string) FreeRowsColumns {
	var result FreeRowsColumns

	for index, row := range data {
		if strings.Count(strings.Join(row, ""), ".") == len(row) {
			result.Rows = append(result.Rows, index)
		}
	}

	for i := 0; i < len(data[0]); i += 1 {
		columnMap := make(map[string]bool)
		for j := 0; j < len(data); j += 1 {
			columnMap[data[j][i]] = true
		}

		if len(columnMap) == 1 {
			result.Columns = append(result.Columns, i)
		}
	}

	return result
}

func calculateDistanceWithExpandIndexes(data [][]string, expandSize int) int {
	cords := findAllGalaxies(data)
	freeRowsColumns := findFreeRowsAndColumnIndexes(data)

	sum := 0

	for i := 0; i < len(cords); i += 1 {
		for j := i + 1; j < len(cords); j += 1 {
			A_i := cords[i][0]
			A_j := cords[i][1]

			B_i := cords[j][0]
			B_j := cords[j][1]
			manDist := manhattannDistance(A_i, B_i, A_j, B_j)

			//check rows expands
			row_amount := 0
			for row := 0; row < len(freeRowsColumns.Rows); row += 1 {
				if min(A_i, B_i) < freeRowsColumns.Rows[row] && freeRowsColumns.Rows[row] < max(A_i, B_i) {
					row_amount += 1
				}
			}

			//check column exapnds
			col_amount := 0
			for col := 0; col < len(freeRowsColumns.Columns); col += 1 {
				if min(A_j, B_j) < freeRowsColumns.Columns[col] && freeRowsColumns.Columns[col] < max(A_j, B_j) {
					col_amount += 1
				}
			}

			finalDistance := manDist + row_amount*expandSize + col_amount*expandSize

			sum += finalDistance
		}
	}

	return sum
}
