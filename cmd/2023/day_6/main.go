package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
}

func intersection(a, b, c, y float64) []float64 {

	delta := b*b - 4*a*(c-y)

	if delta > 0 {
		delta = math.Sqrt(delta)

		x_1 := (-b + delta) / (2 * a)
		x_2 := (-b - delta) / (2 * a)

		return []float64{x_1, x_2}

	}
	return []float64{}
}

func f(x, a int) int {
	return x * (a - x)
}

func Part1() {
	fmt.Println("Part 1")

	input := strings.TrimSpace(utils.ReadFile("cmd/2023/day_6/input.txt"))
	pattern := regexp.MustCompile(`(\d+)`)

	allNumbers := pattern.FindAllString(input, -1)

	times := utils.StringToIntSlice(allNumbers[:len(allNumbers)/2])
	distances := utils.StringToIntSlice(allNumbers[len(allNumbers)/2:])

	boatMap := map[float64]float64{}

	for i := 0; i < len(times); i += 1 {
		boatMap[float64(times[i])] = float64(distances[i])
	}

	//function(x) = x * (a - x)
	//function(x) = -x^2 + x*a

	result := 1
	for time, record := range boatMap {
		xs := intersection(-1, time, 0, record)

		x_1 := int(xs[0]) + 1
		x_2 := int(xs[1])

		if xs[1] == math.Trunc(xs[1]) {
			x_2 -= 1
		}

		result *= (x_2 - x_1 + 1)
	}

	fmt.Println(result)

}

func Part2() {
	fmt.Println("Part 2")

	input := strings.TrimSpace(utils.ReadFile("cmd/2023/day_6/input.txt"))

	pattern := regexp.MustCompile(`(\d+)`)

	allNumbers := pattern.FindAllString(input, -1)

	times := utils.MustAtoi(strings.Join(allNumbers[:len(allNumbers)/2], ""))
	distances := utils.MustAtoi(strings.Join(allNumbers[len(allNumbers)/2:], ""))

	boatMap := map[float64]float64{
		float64(times): float64(distances),
	}

	//function(x) = x * (a - x)
	//function(x) = -x^2 + x*a

	result := 1
	for time, record := range boatMap {
		xs := intersection(-1, time, 0, record)

		x_1 := int(xs[0]) + 1
		x_2 := int(xs[1])

		if xs[1] == math.Trunc(xs[1]) {
			x_2 -= 1
		}

		result *= (x_2 - x_1 + 1)
	}

	fmt.Println(result)
}
