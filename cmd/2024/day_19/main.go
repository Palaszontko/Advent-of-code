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
	input := utils.ReadFile("cmd/2024/day_19/input.txt")

	patterns := map[string]bool{}
	for _, pat := range strings.Split(strings.Split(input, "\n")[0], ", ") {
		patterns[pat] = true
	}

	designs := []string{}
	for _, row := range strings.Split(input, "\n")[1:] {
		if row != "" {
			designs = append(designs, row)
		}
	}

	result := 0
	for _, design := range designs {
		if canMakePattern(design, patterns, 0, make(map[int]bool)) {
			result += 1
		}
	}

	fmt.Println(result)
}

func canMakePattern(design string, patterns map[string]bool, pos int, cache map[int]bool) bool {
	if pos == len(design) {
		return true
	}

	if val, exists := cache[pos]; exists {
		return val
	}

	for pattern := range patterns {
		if pos+len(pattern) <= len(design) && design[pos:pos+len(pattern)] == pattern {
			if canMakePattern(design, patterns, pos+len(pattern), cache) {
				cache[pos] = true
				return true
			}
		}
	}

	cache[pos] = false
	return false
}

func findAllWays(design string, patterns map[string]bool, pos int, cache map[int]int) int {
	if pos == len(design) {
		return 1
	}

	if val, exists := cache[pos]; exists {
		return val
	}

	ways := 0
	for pattern := range patterns {
		if pos+len(pattern) <= len(design) && design[pos:pos+len(pattern)] == pattern {
			ways += findAllWays(design, patterns, pos+len(pattern), cache)
		}
	}

	cache[pos] = ways
	return ways
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_19/input.txt")

	patterns := map[string]bool{}
	for _, pat := range strings.Split(strings.Split(input, "\n")[0], ", ") {
		patterns[pat] = true
	}

	designs := []string{}
	for _, row := range strings.Split(input, "\n")[1:] {
		if row != "" {
			designs = append(designs, row)
		}
	}

	result := 0
	for _, design := range designs {
		ways := findAllWays(design, patterns, 0, make(map[int]int))
		result += ways
	}

	fmt.Println(result)
}
