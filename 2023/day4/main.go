package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/zenoix/advent-of-code-go/utils"
)

type inputType = [][2]string

func main() {
	input := utils.GetInput(2023, 4)

	parsedInput := parseInput(input)

	ans1 := part1(parsedInput)
	ans2 := part2(parsedInput)

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}

func part1(input inputType) (ans int) {
	for _, cards := range input {
		winning := map[string]struct{}{}
		for _, c := range strings.Split(cards[0], " ") {
			winning[c] = struct{}{}
		}

		numSame := getNumberOfMatches(winning, strings.Split(cards[1], " "))

		if numSame != 0 {
			ans += int(math.Pow(2, float64(numSame-1)))
		}
	}
	return
}

func part2(input inputType) (ans int) {
	numCards := len(input)
	cardCounts := make([]int, numCards)
	for i := range input {
		cardCounts[i] = 1
	}

	for i, cards := range input {
		winning := map[string]struct{}{}
		for _, c := range strings.Split(cards[0], " ") {
			winning[c] = struct{}{}
		}

		numSame := getNumberOfMatches(winning, strings.Split(cards[1], " "))

		for j := range numSame {
			cardCounts[i+j+1] = cardCounts[i] + cardCounts[i+j+1]
		}

	}

	for _, s := range cardCounts {
		ans += s
	}

	return
}

func getNumberOfMatches(winning map[string]struct{}, ours []string) (numMatches int) {
	for _, num := range ours {
		if _, ok := winning[num]; ok {
			numMatches += 1
		}
	}
	return
}

func parseInput(input string) inputType {
	parsed := make([][2]string, 0, 2048)

	for _, line := range strings.Split(strings.ReplaceAll(input, "  ", " "), "\n") {
		_, nums, _ := strings.Cut(line, ": ")
		winning, ours, _ := strings.Cut(nums, " | ")

		parsed = append(parsed, [2]string{winning, ours})
	}

	return parsed
}
