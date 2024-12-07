package main

import (
	"fmt"
	"math"
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

	input := utils.ReadFile("cmd/2024/day_7/input.txt")

	result := 0

	for _, row := range strings.Split(input, "\n") {
		outcome := utils.MustAtoi(strings.Split(row, ":")[0])
		values := utils.StringToIntSlice(strings.Split(row, " ")[1:])

		if solveEquation(values, generateBinaryRepresentation(len(values)-1), outcome) {
			result += outcome
		}

	}

	fmt.Println(result)

}

func solveEquation(values []int, signsVariants [][]string, outcome int) bool {
	for _, signs := range signsVariants {
		result := values[0]

		for i, indexSign := 1, 0; i < len(values); i, indexSign = i+1, indexSign+1 {
			if signs[indexSign] == "+" {
				result += values[i]
			} else {
				result *= values[i]
			}

		}

		if result == outcome {
			return true
		}
	}
	return false
}

func generateBinaryRepresentation(size int) [][]string {
	amountOfSeq := 1 << size

	sequences := [][]string{}

	for i := 0; i < amountOfSeq; i += 1 {
		sequence := make([]string, size)

		for j := 0; j < size; j += 1 {
			if i&(1<<j) != 0 {
				sequence[size-j-1] = "+"
			} else {
				sequence[size-j-1] = "*"
			}
		}

		sequences = append(sequences, sequence)

	}

	return sequences
}

func generateBase3Representation(size int) [][]string {

	sequences := [][]string{}

	signs := []string{"||", "+", "*"}

	var number int

	amountOfSeq := int(math.Pow(float64(3), float64(size)))

	for i := 0; i < amountOfSeq; i += 1 {
		number = i
		base3rep := make([]string, size)

		for j := 0; j < size; j += 1 {
			base3rep[size-j-1] = signs[number%3]
			number /= 3
		}

		sequences = append(sequences, base3rep)

	}

	return sequences
}

func solveEquationBase3(values []int, signsVariants [][]string, outcome int) bool {
	for _, signs := range signsVariants {
		result := values[0]

		for i, indexSign := 1, 0; i < len(values); i, indexSign = i+1, indexSign+1 {
			if signs[indexSign] == "+" {
				result += values[i]
			} else if signs[indexSign] == "||" {
				result = utils.MustAtoi(fmt.Sprintf("%v%v", result, values[i]))
			} else {
				result *= values[i]
			}

		}

		if result == outcome {
			return true
		}
	}
	return false
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_7/input.txt")

	result := 0

	for _, row := range strings.Split(input, "\n") {
		outcome := utils.MustAtoi(strings.Split(row, ":")[0])
		values := utils.StringToIntSlice(strings.Split(row, " ")[1:])

		if solveEquationBase3(values, generateBase3Representation(len(values)-1), outcome) {
			result += outcome
		}

	}

	fmt.Println(result)
}
