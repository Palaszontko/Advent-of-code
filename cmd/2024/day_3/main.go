package main

import (
	"fmt"
	"github.com/Palaszontko/advent-of-code/cmd/utils"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_3/input.txt")

	input = strings.Replace(input, "\n", "", -1)

	pattern := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)

	result := 0

	instructions := pattern.FindAllString(input, -1)

	for _, instruction := range instructions {
		result += instructionCompiler(instruction)
	}

	fmt.Println(result)

}

func instructionCompiler(instrucion string) int {
	pattern := regexp.MustCompile(`(\d)+`)

	result := 1

	for _, value := range pattern.FindAllString(instrucion, -1) {
		result *= utils.MustAtoi(value)
	}

	return result
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_3/input.txt")

	input = strings.Replace(input, "\n", "", -1)

	result := 0

	patternMul := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)

	patternReplace := regexp.MustCompile(`don't\(\).*?do\(\)`)

	replaced := patternReplace.ReplaceAllString(input, "")

	// if text still contain one don't at the end
	patternReplace = regexp.MustCompile(`don't\(\).*`)
	replaced = patternReplace.ReplaceAllString(replaced, "")

	for _, part := range patternMul.FindAllString(replaced, -1) {
		result += instructionCompiler(part)
	}

	fmt.Println(result)
}
