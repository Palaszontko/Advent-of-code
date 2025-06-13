package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
}

type Cubes struct {
	Red   []int
	Green []int
	Blue  []int
}

func parseGame(game string) Cubes {
	patternRed := regexp.MustCompile(`\d+.red`)
	patternGreen := regexp.MustCompile(`\d+.green`)
	patternBlue := regexp.MustCompile(`(\d+).blue`)

	var cubes Cubes

	parseColor := func(pattern *regexp.Regexp, colorValue *[]int, val string) {
		colorString := strings.Split(pattern.FindString(val), " ")
		if len(colorString) > 1 {
			*colorValue = append(*colorValue, utils.MustAtoi(colorString[0]))
		}
	}

	for _, val := range strings.Split(game, ";") {
		parseColor(patternBlue, &cubes.Blue, val)
		parseColor(patternRed, &cubes.Red, val)
		parseColor(patternGreen, &cubes.Green, val)
	}

	return cubes
}

func Part1() {
	fmt.Println("Part 1")

	MAX_RED_CUBES := 12
	MAX_GREEN_CUBES := 13
	MAX_BLUE_CUBES := 14

	input := utils.ReadFile("cmd/2023/day_2/input.txt")

	result := 0

	for id, row := range strings.Split(input, "\n") {
		parsedGame := parseGame(row)
		if slices.Max(parsedGame.Red) <= MAX_RED_CUBES && slices.Max(parsedGame.Blue) <= MAX_BLUE_CUBES && slices.Max(parsedGame.Green) <= MAX_GREEN_CUBES {
			result += id + 1
		}

	}

	fmt.Println(result)

}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_2/input.txt")

	result := 0

	for _, row := range strings.Split(input, "\n") {
		parsedGame := parseGame(row)
		result += slices.Max(parsedGame.Red) * slices.Max(parsedGame.Blue) * slices.Max(parsedGame.Green)

	}

	fmt.Println(result)
}
