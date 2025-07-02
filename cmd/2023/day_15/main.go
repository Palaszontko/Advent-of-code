package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func hashing(text string) int {
	result := 0
	for _, letter := range text {
		result += int(letter)
		result *= 17
		result %= 256
	}

	return result
}

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_15/input.txt")

	result := 0
	for _, word := range strings.Split(input, ",") {

		result += hashing(word)
	}

	fmt.Println(result)
}

type Lens struct {
	label       string
	focalLength int
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_15/input.txt")

	var box [][]Lens

	for i := 0; i < 256; i += 1 {
		box = append(box, make([]Lens, 0))
	}

	for _, word := range strings.Split(input, ",") {
		var boxLabel string
		if strings.Contains(word, "-") {
			boxLabel = strings.Split(word, "-")[0]
		} else {
			boxLabel = strings.Split(word, "=")[0]
		}

		currentBoxId := hashing(boxLabel)

		if strings.Count(word, "-") > 0 {
			for i, sLen := range box[currentBoxId] {
				if sLen.label == boxLabel {
					box[currentBoxId] = slices.Delete(box[currentBoxId], i, i+1)
					break
				}
			}
		} else {
			focalLength := utils.MustAtoi(strings.Split(word, "=")[1])

			//find same lens label in this box

			changed := false

			for i, len := range box[currentBoxId] {
				if len.label == boxLabel {
					box[currentBoxId][i] = Lens{label: boxLabel, focalLength: focalLength}
					changed = true
					break
				}
			}

			if !changed {
				box[currentBoxId] = append(box[currentBoxId], Lens{label: boxLabel, focalLength: focalLength})
			}
		}
	}

	result := 0
	for i, specBox := range box {
		for j, lenStruct := range specBox {
			result += (i + 1) * (j + 1) * lenStruct.focalLength
		}
	}

	fmt.Println(result)
}
