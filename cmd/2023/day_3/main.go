package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	Part2()
}

func withinBounds(i, j, width, height int) bool {
	if 0 <= i && i < width && 0 <= j && j < height {
		return true
	}
	return false
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_3/input.txt")

	pattern := regexp.MustCompile(`\d+`)

	grid := strings.Split(input, "\n")

	height := len(grid)
	width := len(grid[0])

	check := func(i int, j int) bool {
		if withinBounds(i, j, width, height) {
			if grid[i][j] != '.' {
				return true
			}
		}
		return false
	}

	result := 0
	for i, row := range grid {
		numbersIndexes := pattern.FindAllStringIndex(row, -1)
		for _, index := range numbersIndexes {
			valid := false
			for j := index[0]; j < index[1]; j += 1 {

				if j == index[0] {
					valid = valid || check(i-1, j-1)
					valid = valid || check(i, j-1)
					valid = valid || check(i+1, j-1)
				}

				if j == index[1]-1 {
					valid = valid || check(i-1, j+1)
					valid = valid || check(i, j+1)
					valid = valid || check(i+1, j+1)
				}

				valid = valid || check(i-1, j)
				valid = valid || check(i+1, j)
			}
			if valid {
				result += utils.MustAtoi(row[index[0]:index[1]])
			}
		}
	}
	fmt.Println(result)
}

type NumberCords struct {
	number int
	row    int
	space  []int
}

func (nb NumberCords) String() string {
	return fmt.Sprintf("%d_%d-%d", nb.row, nb.space[0], nb.space[len(nb.space)-1])
}

func generateRange(start int, end int) []int {
	var tmp []int
	for i := start; i < end; i += 1 {
		tmp = append(tmp, i)
	}
	return tmp
}

func Part2() {
	fmt.Println("Part 2")

	input := utils.ReadFile("cmd/2023/day_3/input.txt")

	pattern := regexp.MustCompile(`\d+`)

	grid := strings.Split(input, "\n")

	allNumbersIndexesMap := make(map[string]NumberCords)

	for i, row := range grid {
		numbersIndexes := pattern.FindAllStringIndex(row, -1)
		for _, index := range numbersIndexes {
			nb := NumberCords{
				number: utils.MustAtoi(row[index[0]:index[1]]),
				row:    i,
				space:  generateRange(index[0], index[1]+1),
			}
			allNumbersIndexesMap[nb.String()] = nb
		}
	}

	checkSurroundings := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	height := len(grid)
	width := len(grid[0])

	result := 0

	for i := 0; i < height; i += 1 {
		for j := 0; j < width; j += 1 {
			foundGears := make(map[string]bool)

			if grid[i][j] == '*' {
				for _, cordsToCheck := range checkSurroundings {
					if withinBounds(i+cordsToCheck[0], j+cordsToCheck[1], width, height) {
						if unicode.IsDigit(rune(grid[i+cordsToCheck[0]][j+cordsToCheck[1]])) {
							// fmt.Println("digit", string(grid[i+cordsToCheck[0]][j+cordsToCheck[1]]), "cords", i+cordsToCheck[0], j+cordsToCheck[1])
							for hash, _ := range allNumbersIndexesMap {
								if utils.MustAtoi(strings.Split(hash, "_")[0]) == i+cordsToCheck[0] {

									start := utils.MustAtoi(strings.Split(strings.Split(hash, "_")[1], "-")[0])
									end := utils.MustAtoi(strings.Split(strings.Split(hash, "_")[1], "-")[1])

									// fmt.Println(hash)
									// fmt.Printf("start: %v ", start)
									// fmt.Printf("end: %v\n", end)

									if start <= j+cordsToCheck[1] && j+cordsToCheck[1] <= end {
										foundGears[hash] = true
									}
								}
							}
						}
					}
				}
			}

			if len(foundGears) == 2 {
				tmp := 1
				for gearHash, _ := range foundGears {
					tmp *= allNumbersIndexesMap[gearHash].number
				}
				result += tmp
			}
		}
	}

	fmt.Println(result)

}
