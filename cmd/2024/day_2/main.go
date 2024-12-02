package main

import (
	"fmt"
	"github.com/Palaszontko/advent-of-code/cmd/utils"
	"math"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_2/input.txt")

	inputSplit := strings.Split(input, "\n")

	result := 0

	for _, line := range inputSplit {
		if line == "" {
			continue
		}

		levels := func(ss []string) []int {
			intValues := make([]int, len(ss))
			for i, s := range ss {
				intValues[i] = utils.MustAtoi(s)
			}
			return intValues
		}(strings.Split(line, " "))

		if isMonotonicStrict(levels) && isSafeDistance(levels) {
			result += 1
		}
	}

	fmt.Println(result)

}

func isMonotonicStrict(values []int) bool {
	//increasing
	funcA := func(values []int) bool {
		for i := 0; i < len(values)-1; i++ {
			if values[i] >= values[i+1] {
				return false
			}
		}
		return true
	}

	//decreasing
	funcB := func(values []int) bool {
		for i := 0; i < len(values)-1; i++ {
			if values[i] <= values[i+1] {
				return false
			}
		}
		return true
	}

	return (funcA(values) || funcB(values))
}

func isSafeDistance(values []int) bool {
	for i := 0; i < len(values)-1; i++ {
		if math.Abs(float64(values[i]-values[i+1])) > 3 {
			return false
		}
	}
	return true
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_2/input.txt")

	inputSplit := strings.Split(input, "\n")

	result := 0

	for _, line := range inputSplit {
		if line == "" {
			continue
		}

		levels := func(ss []string) []int {
			intValues := make([]int, len(ss))
			for i, s := range ss {
				intValues[i] = utils.MustAtoi(s)
			}
			return intValues
		}(strings.Split(line, " "))

		if isMonotonicStrict(levels) && isSafeDistance(levels) {
			result += 1
		} else {
			for _, levelsVariant := range generateVariants(levels) {
				if isMonotonicStrict(levelsVariant) && isSafeDistance(levelsVariant) {
					result += 1
					break
				}
			}
		}
	}

	fmt.Println(result)

}

func generateVariants(levels []int) [][]int {
	var levelsVariants [][]int

	for i := 0; i < len(levels); i++ {
		levelsVariants = append(levelsVariants, slices.Concat(levels[:i], levels[i+1:]))
	}

	return levelsVariants
}
