package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
}

func generateLevels(data []int) [][]int {
	//check 0
	set := map[int]bool{}

	for _, val := range data {
		set[val] = true
	}

	if len(set) != 1 {
		newList := []int{}
		for i := 0; i < len(data)-1; i++ {
			newList = append(newList, data[i+1]-data[i])
		}
		return append([][]int{data}, generateLevels(newList)...)
	}
	return [][]int{data}
}

func predictNextValue(data []int) int {
	levels := generateLevels(data)
	slices.Reverse(levels)

	for i := 1; i < len(levels); i += 1 {
		levels[i] = append(levels[i], levels[i][len(levels[i])-1]+levels[i-1][len(levels[i-1])-1])
	}

	//last element from last list
	return levels[len(levels)-1][len(levels[len(levels)-1])-1]
}

func predictFirstValue(data []int) int {
	levels := generateLevels(data)
	slices.Reverse(levels)

	for i := 1; i < len(levels); i += 1 {
		newValue := levels[i][0] - levels[i-1][0]
		levels[i] = append([]int{newValue}, levels[i]...)
	}

	//first element from last list
	return levels[len(levels)-1][0]
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_9/input.txt")

	result := 0
	for _, line := range strings.Split(input, "\n") {
		values := utils.StringToIntSlice(strings.Split(line, " "))
		result += predictNextValue(values)
	}

	fmt.Println(result)
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_9/input.txt")

	result := 0
	for _, line := range strings.Split(input, "\n") {
		values := utils.StringToIntSlice(strings.Split(line, " "))

		result += predictFirstValue(values)
	}

	fmt.Println(result)
}
