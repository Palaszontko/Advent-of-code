package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Palaszontko/advent-of-code/cmd/utils"
)

func main() {
	fmt.Println("Advent of Code 2023!")
	Part1()
	// Part2()
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

func handType(cards string) HandType {
	handMap := map[string]int{}

	for _, card := range cards {
		handMap[string(card)] += 1
	}

	amounts := []int{}

	for _, val := range handMap {
		amounts = append(amounts, val)
	}

	sort.Slice(amounts, func(i, j int) bool {
		return amounts[i] > amounts[j]
	})

	switch {
	case amounts[0] == 5:
		return FiveOfKind
	case amounts[0] == 4:
		return FourOfKind
	case amounts[0] == 3 && amounts[1] == 2:
		return FullHouse
	case amounts[0] == 3:
		return ThreeOfKind
	case amounts[0] == 2 && amounts[1] == 2:
		return TwoPair
	case amounts[0] == 2:
		return OnePair
	default:
		return HighCard
	}

}

func compareHands(hand1 string, hand2 string) bool {
	valueCardMap := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	// fmt.Printf("hand1: %v\n", hand1)
	// fmt.Printf("hand2: %v\n", hand2)

	for i := 0; i < len(hand1); i += 1 {
		card1 := valueCardMap[string(hand1[i])]
		card2 := valueCardMap[string(hand2[i])]
		if card1 > card2 {
			return true
		} else if card1 < card2 {
			return false
		}
	}
	return false
}

type Hand struct {
	Cards string
	Bid   int
}

func Part1() {
	fmt.Println("Part 1")

	input := utils.ReadFile("cmd/2023/day_7/input.txt")

	handsMap := map[HandType][]Hand{}

	for _, line := range strings.Split(input, "\n") {
		cards := strings.Split(line, " ")[0]
		bid := utils.MustAtoi(strings.Split(line, " ")[1])
		handType := handType(cards)
		hand := Hand{Cards: cards, Bid: bid}
		handsMap[handType] = append(handsMap[handType], hand)
	}

	finalHands := []Hand{}

	for i := HighCard; i <= FiveOfKind; i += 1 {
		if len(handsMap[i]) == 1 {
			finalHands = append(finalHands, handsMap[i]...)
		} else if len(handsMap[i]) > 1 {
			hands := handsMap[i]
			sort.Slice(hands, func(j, k int) bool {
				return compareHands(hands[k].Cards, hands[j].Cards)
			})
			finalHands = append(finalHands, hands...)
		}
	}

	result := 0

	for index_rank, hand := range finalHands {
		result += (index_rank + 1) * hand.Bid
	}

	fmt.Println(result)
}

func Part2() {
	fmt.Println("Part 2")
}
