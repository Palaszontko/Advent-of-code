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

	input := utils.ReadFile("cmd/2024/day_9/input.txt")

	inputSplit := strings.Split(input, "")

	disk := []string{}

	space := false

	for i, j := 0, 0; i < len(inputSplit); i += 1 {
		if space {
			value := "."
			for k := 0; k < utils.MustAtoi(inputSplit[i]); k += 1 {
				disk = append(disk, value)
			}
		} else {
			for k := 0; k < utils.MustAtoi(inputSplit[i]); k += 1 {
				disk = append(disk, fmt.Sprint(j))
			}
			j += 1
		}
		space = !space
	}

	i, j := len(disk)-1, 0
	for i > j {
		if disk[i] == "." {
			i -= 1
		} else if disk[j] != "." {
			j += 1
		}

		if disk[i] != "." && disk[j] == "." {
			disk[i], disk[j] = disk[j], disk[i]

		}
	}

	result := checkSum(disk)
	fmt.Println(result)

}

func checkSum(data []string) int64 {
	var amount int64

	for i := 0; i < len(data); i += 1 {
		if data[i] != "." {
			amount += int64(i * utils.MustAtoi(string(data[i])))
		}
	}

	return amount
}

func Part2() {
	fmt.Println("Part 2")
}
