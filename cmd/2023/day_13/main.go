package main

import (
	"fmt"
	"reflect"
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

	input := utils.ReadFile("cmd/2023/day_13/input.txt")

	var data [][][]string

	for _, pattern := range strings.Split(input, "\n\n") {
		newPattern := [][]string{}
		for _, row := range strings.Split(pattern, "\n") {
			newPattern = append(newPattern, strings.Split(row, ""))
		}
		data = append(data, newPattern)
	}

	result := 0
	for _, pattern := range data {

		amountOfRowsAbove := findReflection(pattern)
		amountOfColumnsLeft := findReflection(utils.Rotate2DMatrixBy90Deg(pattern))

		if amountOfColumnsLeft != -1 {
			result += amountOfColumnsLeft
		} else {
			result += amountOfRowsAbove * 100
		}

	}

	fmt.Println(result)

}

// return amount of rows(above)/columns(left)
func findReflection(pattern [][]string) int {
	bestReflection := -1
	for i := 0; i < len(pattern)-1; i += 1 {

		for j := 0; j < i+1; j += 1 {
			above_start := i - j
			above_end := i + 1
			under_start := i + 1
			under_end := min(i+j+2, len(pattern))
			if above_end-above_start <= under_end-under_start {
				above := pattern[above_start:above_end]
				under := pattern[under_start:under_end]

				// fmt.Println(above_start, above_end, "<->", under_start, under_end)
				// fmt.Println(above)
				// fmt.Println(under)

				if isEqualReverse(above, under) && (under_end == len(pattern) || above_start == 0) {
					bestReflection = max(bestReflection, i+1)
				}
			}

		}

	}

	return bestReflection
}

func isEqualReverse(pattern1 [][]string, pattern2 [][]string) bool {
	for i := 0; i < len(pattern1); i += 1 {
		if !reflect.DeepEqual(pattern1[i], pattern2[len(pattern2)-1-i]) {
			return false
		}
	}
	return true
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_13/input.txt")

	var data [][][]string

	for _, pattern := range strings.Split(input, "\n\n") {
		newPattern := [][]string{}
		for _, row := range strings.Split(pattern, "\n") {
			newPattern = append(newPattern, strings.Split(row, ""))
		}
		data = append(data, newPattern)
	}

	result := 0
	for _, pattern := range data {
		amountOfRowsAbove := findReflectionPart2(pattern)
		amountOfColumnsLeft := findReflectionPart2(utils.Rotate2DMatrixBy90Deg(pattern))

		if amountOfColumnsLeft != -1 {
			result += amountOfColumnsLeft
		} else {
			result += amountOfRowsAbove * 100
		}

	}

	fmt.Println(result)
}

func findReflectionPart2(pattern [][]string) int {
	for i := 0; i < len(pattern)-1; i += 1 {

		for j := 0; j < i+1; j += 1 {
			above_start := i - j
			above_end := i + 1
			under_start := i + 1
			under_end := min(i+j+2, len(pattern))
			if above_end-above_start <= under_end-under_start {
				above := pattern[above_start:above_end]
				under := pattern[under_start:under_end]

				if dif_index := findAmountOfDifferences(above, under); dif_index[0] != -1 && (under_end == len(pattern) || above_start == 0) {
					fmt.Println(above, under)
					return i + 1

				}
			}
		}
	}

	return -1
}

func findAmountOfDifferences(pattern1, pattern2 [][]string) []int {
	indexes := [][]int{}
	for i := 0; i < len(pattern1); i += 1 {
		for j := 0; j < len(pattern1[i]); j += 1 {
			if pattern1[i][j] != pattern2[len(pattern2)-1-i][j] {
				indexes = append(indexes, []int{i, j})
			}
		}
	}

	if len(indexes) == 1 {
		return indexes[0]
	} else {
		return []int{-1, -1}
	}
}
