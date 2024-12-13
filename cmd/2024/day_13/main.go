package main

import (
	"fmt"
	"math"
	"regexp"
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

	input := utils.ReadFile("cmd/2024/day_13/input.txt")

	inputSplit := strings.Split(input, "\n")

	result := 0

	for i := 0; i < len(inputSplit)-2; i += 4 {
		if inputSplit[i] == "" {
			continue
		}

		pattern := regexp.MustCompile(`(X|Y)\+\d+`)
		button1 := pattern.FindAllString(inputSplit[i], -1)
		button2 := pattern.FindAllString(inputSplit[i+1], -1)

		pattern = regexp.MustCompile(`(X|Y)\=\d+`)
		prize := pattern.FindAllString(inputSplit[i+2], -1)

		button1_x := utils.MustAtoi(button1[0][2:])
		button1_y := utils.MustAtoi(button1[1][2:])

		button2_x := utils.MustAtoi(button2[0][2:])
		button2_y := utils.MustAtoi(button2[1][2:])

		prize_x := utils.MustAtoi(prize[0][2:])
		prize_y := utils.MustAtoi(prize[1][2:])

		A := []int{button1_x, button1_y}
		B := []int{button2_x, button2_y}
		final := []int{prize_x, prize_y}

		matrix := [][]int{{A[0], B[0], final[0]}, {A[1], B[1], final[1]}}

		a, b, determinated := cramer(matrix)

		if determinated && isPositiveInteger(a) && isPositiveInteger(b) && a <= 100 && b <= 100 {
			result += 3*int(a) + int(b)
		}
	}

	fmt.Println(result)

}

func isPositiveInteger(x float64) bool {
	return float64(int(x)) == x && x > 0
}

func cramer(matrix [][]int) (float64, float64, bool) {
	D := matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]

	if D == 0 {
		return -1, -1, false
	}

	Da := matrix[0][2]*matrix[1][1] - matrix[0][1]*matrix[1][2]
	Db := matrix[0][0]*matrix[1][2] - matrix[0][2]*matrix[1][0]

	return float64(Da) / float64(D), float64(Db) / float64(D), true
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_13/input.txt")

	inputSplit := strings.Split(input, "\n")

	result := 0

	for i := 0; i < len(inputSplit)-2; i += 4 {
		if inputSplit[i] == "" {
			continue
		}

		pattern := regexp.MustCompile(`(X|Y)\+\d+`)
		button1 := pattern.FindAllString(inputSplit[i], -1)
		button2 := pattern.FindAllString(inputSplit[i+1], -1)

		pattern = regexp.MustCompile(`(X|Y)\=\d+`)
		prize := pattern.FindAllString(inputSplit[i+2], -1)

		button1_x := utils.MustAtoi(button1[0][2:])
		button1_y := utils.MustAtoi(button1[1][2:])

		button2_x := utils.MustAtoi(button2[0][2:])
		button2_y := utils.MustAtoi(button2[1][2:])

		prize_x := utils.MustAtoi(prize[0][2:]) + int(math.Pow(10, 13))
		prize_y := utils.MustAtoi(prize[1][2:]) + int(math.Pow(10, 13))

		A := []int{button1_x, button1_y}
		B := []int{button2_x, button2_y}
		final := []int{prize_x, prize_y}

		matrix := [][]int{{A[0], B[0], final[0]}, {A[1], B[1], final[1]}}

		a, b, determinated := cramer(matrix)

		// to solve part 2 just add 10^13 to the result and remove the condition a <= 100 && b <= 100
		if determinated && isPositiveInteger(a) && isPositiveInteger(b) {
			result += 3*int(a) + int(b)
		}
	}

	fmt.Println(result)

}
