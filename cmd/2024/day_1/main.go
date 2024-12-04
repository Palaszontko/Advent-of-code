package main

import (
	"fmt"
	"math"
	"sort"
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

	input := utils.ReadFile("cmd/2024/day_1/input.txt")

	lines := strings.Split(string(input), "\n")

	list1 := []int{}
	list2 := []int{}

	for _, line := range lines {
		values := strings.Fields(line)
		list1 = append(list1, utils.MustAtoi(values[0]))
		list2 = append(list2, utils.MustAtoi(values[1]))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var result float64

	for i := 0; i < len(list1); i++ {
		result += math.Abs(float64(list1[i] - list2[i]))
	}

	fmt.Println(int(result))

}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_1/input.txt")

	lines := strings.Split(string(input), "\n")

	list1 := []int{}
	map2 := map[int]int{}

	for _, line := range lines {
		values := strings.Fields(line)
		list1 = append(list1, utils.MustAtoi(values[0]))
		map2[utils.MustAtoi(values[1])]++
	}

	result := 0

	for _, value := range list1 {
		result += value * map2[value]
	}

	fmt.Println(result)
}
