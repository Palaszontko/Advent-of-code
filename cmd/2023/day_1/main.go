package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
}

func calibrate(values []string) int {
	if len(values) > 0 {
		return utils.MustAtoi(values[0])*10 + utils.MustAtoi(values[len(values)-1])
	} else {
		return 0
	}
}
func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_1/input.txt")

	result := 0

	pattern := regexp.MustCompile(`\d`)

	for _, text := range strings.Split(input, "\n") {
		result += calibrate(pattern.FindAllString(text, -1))
	}

	fmt.Println(result)
}

func Part2() {
	fmt.Println("Part 2")

	stringMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	input := utils.ReadFile("cmd/2023/day_1/input.txt")
	result := 0

	for _, row := range strings.Split(input, "\n") {
		words := make(map[int]int)
		numbers := make(map[int]int)

		for i := 0; i < len(row); i += 1 {
			for j := i; j <= len(row); j += 1 {
				val, ok := stringMap[row[i:j]]
				if ok {
					words[i] = val
					break
				}
			}
			if unicode.IsDigit(rune(row[i])) {
				numbers[i] = int(rune(row[i]) - 48)
			}
		}

		for x := 0; x < len(row); x += 1 {
			val, ok := words[x]
			if ok {
				result += val * 10
				break
			}

			val2, ok2 := numbers[x]
			if ok2 {
				result += val2 * 10
				break
			}
		}

		for y := len(row) - 1; y >= 0; y -= 1 {
			val, ok := words[y]
			if ok {
				result += val
				break
			}

			val2, ok2 := numbers[y]
			if ok2 {
				result += val2
				break
			}
		}

	}

	fmt.Println(result)

}
