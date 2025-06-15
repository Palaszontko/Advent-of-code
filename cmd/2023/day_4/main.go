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

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_4/input.txt")

	result := 0
	for _, line := range strings.Split(input, "\n") {
		value := 0
		winningNumbersString := strings.Split(strings.Split(line, ":")[1], "|")[0]
		yourNumbersString := strings.Split(strings.Split(line, ":")[1], "|")[1]

		pattern := regexp.MustCompile(`(\d)+`)

		winningNumbers := utils.StringToIntSlice(pattern.FindAllString(winningNumbersString, -1))
		yourNumbers := utils.StringToIntSlice(pattern.FindAllString(yourNumbersString, -1))

		for _, yourNumber := range yourNumbers {
			if slices.Contains(winningNumbers, yourNumber) {
				if value == 0 {
					value += 1
				} else {
					value *= 2
				}
			}
		}
		result += value
	}

	fmt.Println(result)

}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_4/input.txt")

	cardsWinMap := make(map[int]int)

	for i, line := range strings.Split(input, "\n") {
		amountOfWinningNumbers := 0

		winningNumbersString := strings.Split(strings.Split(line, ":")[1], "|")[0]
		yourNumbersString := strings.Split(strings.Split(line, ":")[1], "|")[1]

		pattern := regexp.MustCompile(`(\d)+`)

		winningNumbers := utils.StringToIntSlice(pattern.FindAllString(winningNumbersString, -1))
		yourNumbers := utils.StringToIntSlice(pattern.FindAllString(yourNumbersString, -1))

		for _, yourNumber := range yourNumbers {
			if slices.Contains(winningNumbers, yourNumber) {
				amountOfWinningNumbers += 1
			}
		}
		cardsWinMap[i+1] = amountOfWinningNumbers
	}

	copyOfMap := make(map[int]int)
	for k := range len(cardsWinMap) {
		copyOfMap[k+1] = 1
	}

	for i := 1; i <= len(cardsWinMap); i += 1 {
		for j := i + 1; j <= i+cardsWinMap[i]; j += 1 {
			copyOfMap[j] += copyOfMap[i]
		}
	}

	result := 0

	for _, val := range copyOfMap {
		result += val
	}

	fmt.Println(result)
}
