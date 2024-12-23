package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	// Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	patterns := strings.Split(strings.Split(input, "\n")[0], ", ")

	designs := []string{}

	for _, row := range strings.Split(input, "\n")[1:] {
		if row != "" {
			designs = append(designs, row)
			fmt.Println("SEP")
		}
	}

	fmt.Println(patterns)
	fmt.Println(designs)
}

func Part2() {
	fmt.Println("Part 2")
}
