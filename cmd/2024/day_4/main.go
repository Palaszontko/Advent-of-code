package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	Part2()
}

type word int

const (
	XMAS word = iota
	SMAX
	NOT_FOUND
)

func Part1() {
	fmt.Println("Part 1")

	//   0 1 2 3 4 5 6 7 8 9
	//0  M M M S X X M A S M
	//1  M S A M X M S M S A
	//2  A M X S X M A A M M
	//3  M S A M A S M S M X
	//4  X M A S A M X A M M
	//5  X X A M M X X A M A
	//6  S M S M S A S X S S
	//7  S A X A M A S A A A
	//8  M A M M M X M M M M
	//9  M X M X A X M A S X

	input := utils.ReadFile("cmd/2024/day_4/input.txt")
	inputSplit := strings.Split(input, "\n")
	amount := 0

	// Check for left diagonals XMAS ( \ )
	for i := 0; i < len(inputSplit)-3; i += 1 {

		for j := 0; j < len(inputSplit[0])-3; j += 1 {
			letter1 := inputSplit[i][j]
			letter2 := inputSplit[i+1][j+1]
			letter3 := inputSplit[i+2][j+2]
			letter4 := inputSplit[i+3][j+3]

			if isXMAS(letter1, letter2, letter3, letter4) {
				amount += 1
			}
		}
	}
	// Check for right diagonals XMAS ( / )
	for i := 0; i < len(inputSplit)-3; i += 1 {

		for j := 3; j < len(inputSplit[0]); j += 1 {
			letter1 := inputSplit[i][j]
			letter2 := inputSplit[i+1][j-1]
			letter3 := inputSplit[i+2][j-2]
			letter4 := inputSplit[i+3][j-3]

			if isXMAS(letter1, letter2, letter3, letter4) {
				amount += 1
			}

		}
	}
	// Check for vertical XMAS
	for i := 0; i < len(inputSplit)-3; i++ {
		for j := 0; j < len(inputSplit[0]); j += 1 {
			letter1 := inputSplit[i][j]
			letter2 := inputSplit[i+1][j]
			letter3 := inputSplit[i+2][j]
			letter4 := inputSplit[i+3][j]
			if isXMAS(letter1, letter2, letter3, letter4) {
				amount += 1
			}

		}

	}
	//Check for horizontal XMAS
	for _, line := range inputSplit {
		amount += strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
	}

	fmt.Println(amount)

}

func isXMAS(letter1 byte, letter2 byte, letter3 byte, letter4 byte) bool {
	if (letter1 == 'X' && letter2 == 'M' && letter3 == 'A' && letter4 == 'S') || (letter1 == 'S' && letter2 == 'A' && letter3 == 'M' && letter4 == 'X') {
		return true
	}

	return false
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2024/day_4/input.txt")

	inputSplit := strings.Split(input, "\n")

	amount := 0

	for i := 0; i < len(inputSplit)-2; i++ {
		for j := 0; j < len(inputSplit[0])-2; j += 1 {
			square := string(inputSplit[i][j:j+3]) +
				string(inputSplit[i+1][j:j+3]) +
				string(inputSplit[i+2][j:j+3])
			if isCrossMAS(square) {
				amount++
			}
		}
	}

	fmt.Println(amount)

}

func isCrossMAS(square string) bool {
	pattern := regexp.MustCompile(`^(M.M.A.S.S|S.S.A.M.M|M.S.A.M.S|S.M.A.S.M)$`)
	return pattern.Match([]byte(square))
}
