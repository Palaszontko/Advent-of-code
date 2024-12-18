package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

const (
	adv = 0
	bxl = 1
	bst = 2
	jnz = 3
	bxc = 4
	out = 5
	bdv = 6
	cdv = 7
)

func main() {
	fmt.Println("Advent of Code 2024!")
	Part1()
	// Part2()
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2024/day_17/input.txt")

	inputSplit := strings.Split(input, "\n")

	registerA := utils.MustAtoi(inputSplit[0][12:])
	registerB := utils.MustAtoi(inputSplit[1][12:])
	registerC := utils.MustAtoi(inputSplit[2][12:])

	program := utils.StringToIntSlice(strings.Split(inputSplit[4][9:], ","))

	i := 0
	output := []int{}

	for i < len(program) {
		opcode := program[i]
		operand := program[i+1]
		instructions(opcode, operand, &registerA, &registerB, &registerC, &i, &output)
	}

	result := strings.Join(strings.Fields(fmt.Sprint(output)[1:len(fmt.Sprint(output))-1]), ",")
	fmt.Println(result)

}

func Part2() {
	fmt.Println("Part 2")
}

func instructions(opcode int, operand int, registerA, registerB, registerC, index *int, output *[]int) {
	getValue := func(operand int) int {
		if operand <= 3 {
			return operand
		} else if operand == 4 {
			return *registerA
		} else if operand == 5 {
			return *registerB
		} else if operand == 6 {
			return *registerC
		} else {
			return 7
		}
	}

	switch opcode {
	case adv:
		*registerA = int(*registerA / int(math.Pow(float64(2), float64(getValue(operand)))))
	case bxl:
		*registerB = *registerB ^ operand
	case bst:
		*registerB = getValue(operand) % 8
	case jnz:
		if *registerA != 0 {
			*index = operand
			return
		}
	case bxc:
		*registerB = *registerB ^ *registerC
	case out:
		(*output) = append((*output), getValue(operand)%8)
	case bdv:
		*registerB = int(*registerA / int(math.Pow(float64(2), float64(getValue(operand)))))
	case cdv:
		*registerC = int(*registerA / int(math.Pow(float64(2), float64(getValue(operand)))))
	}
	*index += 2
}
