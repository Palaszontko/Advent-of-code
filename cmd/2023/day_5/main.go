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

// destination, source, length slice
type dslSlice struct {
	destinationStart int
	sourceStart      int
	lenght           int
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_5/input.txt")

	splittedInput := strings.Split(input, "\n")

	seeds := utils.StringToIntSlice(strings.Split(splittedInput[0], " ")[1:len(strings.Split(splittedInput[0], " "))])

	pattern := regexp.MustCompile(`(.+).{1}(map:)\n((\d+).(\d+).(\d+)(\n|))+`)

	matches := pattern.FindAllStringSubmatch(input, -1)

	allMaps := [][]dslSlice{}

	for _, match := range matches {
		arrayOfDSL := []dslSlice{}

		valesPattern := regexp.MustCompile(`(\d+ \d+ \d+)`)
		values := valesPattern.FindAllString(match[0], -1)

		for _, value := range values {
			tmp := utils.StringToIntSlice(strings.Split(value, " "))
			arrayOfDSL = append(arrayOfDSL,
				dslSlice{
					destinationStart: tmp[0],
					sourceStart:      tmp[1],
					lenght:           tmp[2],
				})
		}
		allMaps = append(allMaps, arrayOfDSL)
	}

	result := []int64{}

	for _, seed := range seeds {
		result = append(result, mapSeedToLocation(seed, allMaps))
	}

	fmt.Println(slices.Min(result))

}

func mapSeedToLocation(seed int, allMaps [][]dslSlice) int64 {
	actualSeedId := seed
	for _, specyficMaps := range allMaps {
		for _, currentMap := range specyficMaps {
			if currentMap.sourceStart <= actualSeedId && actualSeedId < currentMap.sourceStart+currentMap.lenght {
				actualSeedId = currentMap.destinationStart + (actualSeedId - currentMap.sourceStart)
				break
			}
		}
	}
	return int64(actualSeedId)
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_5/input.txt")

	splittedInput := strings.Split(input, "\n")

	seedsSpliited := utils.StringToIntSlice(strings.Split(splittedInput[0], " ")[1:len(strings.Split(splittedInput[0], " "))])

	seeds := [][]int{}

	for i := 0; i < len(seedsSpliited)-1; i += 2 {
		seeds = append(seeds, []int{seedsSpliited[i], seedsSpliited[i] + seedsSpliited[i+1] - 1})
	}

	pattern := regexp.MustCompile(`(.+).{1}(map:)\n((\d+).(\d+).(\d+)(\n|))+`)

	matches := pattern.FindAllStringSubmatch(input, -1)

	allMaps := [][]dslSlice{}

	for _, match := range matches {
		arrayOfDSL := []dslSlice{}

		valesPattern := regexp.MustCompile(`(\d+ \d+ \d+)`)
		values := valesPattern.FindAllString(match[0], -1)

		for _, value := range values {
			tmp := utils.StringToIntSlice(strings.Split(value, " "))
			arrayOfDSL = append(arrayOfDSL,
				dslSlice{
					destinationStart: tmp[0],
					sourceStart:      tmp[1],
					lenght:           tmp[2],
				})
		}
		allMaps = append(allMaps, arrayOfDSL)
	}

	currentRanges := [][2]int{}
	for _, s := range seeds {
		currentRanges = append(currentRanges, [2]int{s[0], s[1] - s[0] + 1})
	}

	for _, mapping := range allMaps {
		newRanges := [][2]int{}
		for _, r := range currentRanges {
			newRanges = append(newRanges, applyingMap(r[0], r[1], mapping)...)
		}
		currentRanges = newRanges
	}

	min := currentRanges[0][0]
	for _, r := range currentRanges {
		if r[0] < min {
			min = r[0]
		}
	}

	fmt.Println(min)

}

func applyingMap(rStart, rLen int, maps []dslSlice) [][2]int {
	result := make([][2]int, 0)
	stack := [][2]int{{rStart, rLen}}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		applied := false

		for _, m := range maps {
			overLapStart := max(curr[0], m.sourceStart)
			overLapEnd := min(curr[0]+curr[1], m.sourceStart+m.lenght)

			if overLapStart < overLapEnd {
				shift := m.destinationStart - m.sourceStart
				result = append(result, [2]int{overLapStart + shift, overLapEnd - overLapStart})
				applied = true

				if curr[0] < overLapStart {
					stack = append(stack, [2]int{curr[0], overLapStart - curr[0]})
				}

				if overLapEnd < curr[0]+curr[1] {
					stack = append(stack, [2]int{overLapEnd, curr[0] + curr[1] - overLapEnd})
				}
				break
			}
		}
		if !applied {
			result = append(result, curr)
		}

	}

	return result
}
