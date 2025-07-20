package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	// Part2()
}

type instructionStruct struct {
	argument string
	sign     string
	value    int
	result   string
}

type workflowStruct struct {
	name         string
	instructions []instructionStruct
	final        string
}

func executeTransactionHelper(instruction instructionStruct, argument int) (string, bool) {
	if instruction.sign == "<" {
		if argument < instruction.value {
			return instruction.result, true
		}
	} else {
		if argument > instruction.value {
			return instruction.result, true
		}
	}
	return "", false
}

func (ws *workflowStruct) executeInstrucion(rating ratingStruct) string {
	for _, instrucion := range ws.instructions {
		switch instrucion.argument {
		case "x":
			if result, ok := executeTransactionHelper(instrucion, rating.x); ok {
				return result
			}
		case "m":
			if result, ok := executeTransactionHelper(instrucion, rating.m); ok {
				return result
			}
		case "a":
			if result, ok := executeTransactionHelper(instrucion, rating.a); ok {
				return result
			}
		case "s":
			if result, ok := executeTransactionHelper(instrucion, rating.s); ok {
				return result
			}
		}
	}

	return ws.final
}

type ratingStruct struct {
	x int
	m int
	a int
	s int
}

func parseWorkflow(workflow_string string) workflowStruct {
	var final string
	name := strings.Split(workflow_string, "{")[0]
	insides := strings.Split(workflow_string, "{")[1]
	insides_splitted := strings.Split(insides[:len(insides)-1], ",")

	instructions := []instructionStruct{}

	for _, inside := range insides_splitted {
		if strings.Contains(inside, ":") {
			argument := string(inside[0])
			sign := string(inside[1])
			value := utils.MustAtoi(strings.Split(strings.Split(inside, sign)[1], ":")[0])
			result := strings.Split(inside, ":")[1]
			instructions = append(instructions, instructionStruct{
				argument: argument,
				sign:     sign,
				value:    value,
				result:   result,
			})
		} else {
			final = inside
		}
	}

	parsed_workflow := workflowStruct{
		name:         name,
		instructions: instructions,
		final:        final,
	}

	return parsed_workflow
}

func parseInput(input string) (map[string]workflowStruct, []ratingStruct) {
	workflows_splitted := strings.Split(strings.Split(input, "\n\n")[0], "\n")
	ratings_splitted := strings.Split(strings.Split(input, "\n\n")[1], "\n")

	workflowsMap := map[string]workflowStruct{}
	ratingsSlice := []ratingStruct{}

	for _, workflow := range workflows_splitted {
		parsed_workflow := parseWorkflow(workflow)
		workflowsMap[parsed_workflow.name] = parsed_workflow
	}

	pattern := regexp.MustCompile(`.=\d+`)
	for _, rating := range ratings_splitted {
		var newRating ratingStruct
		for _, assigned := range pattern.FindAllString(rating, -1) {
			switch strings.Split(assigned, "=")[0] {
			case "x":
				newRating.x = utils.MustAtoi(strings.Split(assigned, "=")[1])
			case "m":
				newRating.m = utils.MustAtoi(strings.Split(assigned, "=")[1])
			case "a":
				newRating.a = utils.MustAtoi(strings.Split(assigned, "=")[1])
			case "s":
				newRating.s = utils.MustAtoi(strings.Split(assigned, "=")[1])
			}
		}

		ratingsSlice = append(ratingsSlice, newRating)
	}

	return workflowsMap, ratingsSlice
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_19/input.txt")

	workflowsMap, ratingsSlice := parseInput(input)

	result := 0

	for _, rating := range ratingsSlice {
		start := workflowsMap["in"]
		resultOfInstruction := start.executeInstrucion(rating)
		for resultOfInstruction != "A" && resultOfInstruction != "R" {
			nextInstruction := workflowsMap[resultOfInstruction]
			resultOfInstruction = nextInstruction.executeInstrucion(rating)
		}

		if resultOfInstruction == "A" {
			result += rating.x + rating.m + rating.a + rating.s
		}
	}

	fmt.Println(result)

}

func Part2() {
	fmt.Println("Part 2")
}
