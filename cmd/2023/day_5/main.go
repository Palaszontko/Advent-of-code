package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

var input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	// Part2()
}

// destination, source, length slice
type dslSlice struct {
	destinationStart int
	sourceStart      int
	lenght           int
}

func Part1() {
	fmt.Println("Part 1")

	input = utils.ReadFile("cmd/2023/day_5/input.txt")

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
		actualSeedId := seed
		// fmt.Println("start: ", actualSeedId)
		for _, specyficMaps := range allMaps {
			for _, currentMap := range specyficMaps {
				if currentMap.sourceStart <= actualSeedId && actualSeedId < currentMap.sourceStart+currentMap.lenght {
					actualSeedId = currentMap.destinationStart + (actualSeedId - currentMap.sourceStart)
					break
				}
			}
		}
		result = append(result, int64(actualSeedId))
	}

	fmt.Println(slices.Min(result))

}
