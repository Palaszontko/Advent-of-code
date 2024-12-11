package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

var per int = 1

func main() {
	fmt.Println("Advent of Code 2024!")
	// Part1()
	Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_11/input.txt")

	stones := utils.StringToIntSlice(strings.Split(input, " "))

	blinks := 25

	for blink := 0; blink < blinks; blink += 1 {
		for i := 0; i < len(stones); i += 1 {
			if stones[i] == 0 {
				stones[i] = 1
			} else if lengthOfInt(stones[i])%2 == 0 {
				half := int(math.Pow(10, float64((lengthOfInt(stones[i]) / 2))))
				left := stones[i] / half
				right := stones[i] % half
				if i == len(stones)-1 {
					stones = stones[:len(stones)-1]
					stones = append(stones, left)
					stones = append(stones, right)
					i += 1
				} else {
					stones = slices.Delete(stones, i, i+1)
					stones = slices.Insert(stones, i, left)
					stones = slices.Insert(stones, i+1, right)
					i += 1
				}
			} else {
				stones[i] *= 2024
			}
		}
		fmt.Println(blink)
	}

	result := len(stones)

	fmt.Println(result)

}

func lengthOfInt(val int) int {
	tmp := val
	length := 0

	if val == 0 {
		return 1
	}

	for tmp >= 1 {
		tmp /= 10
		length += 1
	}

	return length
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_11/input.txt")

	stones := utils.StringToIntSlice(strings.Split(input, " "))

	stonesMap := map[string]int{}

	blinks := 75
	result := 0

	for _, stone := range stones {
		result += countRecursive(stone, blinks, stonesMap)
	}

	fmt.Println(result)

}

func countRecursive(stone int, blinksLeft int, cache map[string]int) int {

	key := fmt.Sprintf("%d_%d", stone, blinksLeft)

	if val, exist := cache[key]; exist {
		return val
	}

	var result int
	if blinksLeft == 0 {
		per += 1
		fmt.Println(per)
		result = 1
	} else if stone == 0 {
		result = countRecursive(1, blinksLeft-1, cache)
	} else if lengthOfInt(stone)%2 == 0 {
		half := int(math.Pow(10, float64((lengthOfInt(stone) / 2))))
		left := stone / half
		right := stone % half

		result = countRecursive(left, blinksLeft-1, cache) + countRecursive(right, blinksLeft-1, cache)
	} else {
		result = countRecursive(stone*2024, blinksLeft-1, cache)
	}

	cache[key] = result
	return result

}
