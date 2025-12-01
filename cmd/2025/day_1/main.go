package main

import (
	"fmt"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2025!")
	Part1()
	// Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2025/day_1/input.txt")

	start := 50
	amount := 0

	for _, code := range strings.Split(input, "\n") {
		letter, value := code[0], utils.MustAtoi(code[1:])

		value %= 100

		if letter == 'L' {
			if start-value < 0 {
				start = 100 - (value - start)
			} else {
				start = (start - value) % 100
			}
		} else {
			start = (start + value) % 100
		}

		if start == 0 {
			amount += 1
		}

	}

	fmt.Println(amount)
}

func Part2() {
	fmt.Println("Part 2")
}
