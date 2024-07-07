package main

import (
	"fmt"
	"strings"

	"github.com/zenoix/advent-of-code-go/utils"
)

type inputType = []string

func main() {
	input := utils.GetInput({{.Year}}, {{.Day}})

	parsedInput := parseInput(input)

	ans1 := part1(parsedInput)
	ans2 := part2(parsedInput)

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}

func part1(input inputType) (ans int) {
	return
}

func part2(input inputType) (ans int) {
	return
}

func parseInput(input string) inputType {
	parsed := make([]string, 0, 2048)

	for _, line := range strings.Split(input, "\n") {
		parsed = append(parsed, line)
	}

	return parsed
}
