package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/zenoix/advent-of-code-go/utils"
)

type inputType = []string

func main() {
	input := utils.GetInput(2023, 1)

	parsedInput := parseInput(input)

	ans1 := part1(parsedInput)
	ans2 := part2(parsedInput)

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}

func part1(input inputType) (ans int) {
	for _, line := range input {
		firstFound := -1
		lastFound := -1
		for _, char := range line {
			if unicode.IsDigit(char) {
				literal, err := strconv.ParseUint(string(char), 10, 0)
				if err != nil {
					fmt.Println(err)
				}
				if firstFound == -1 {
					firstFound = int(literal)
				}
				lastFound = int(literal)
			}
		}
		ans += firstFound*10 + lastFound
	}
	return
}

func part2(input inputType) (ans int) {
	return
}

func parseInput(input string) inputType {
	return strings.Split(input, "\n")
}
