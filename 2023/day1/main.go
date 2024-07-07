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
				num := runeToInt(char)
				if firstFound == -1 {
					firstFound = num
				}
				lastFound = num
			}
		}
		ans += firstFound*10 + lastFound
	}
	return
}

func part2(input inputType) (ans int) {
	wordToNumber := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range input {

		firstFound := -1
		lastFound := -1

		for len(line) > 0 {
			char := rune(line[0])
			if unicode.IsDigit(char) {
				num := runeToInt(char)
				if firstFound == -1 {
					firstFound = num
				}
				lastFound = num
			} else {
				for word, num := range wordToNumber {
					if len(line) < len(word) || line[:len(word)] != word {
						continue
					}
					if firstFound == -1 {
						firstFound = num
					}
					lastFound = num
				}
			}
			line = line[1:]
		}
		ans += firstFound*10 + lastFound
	}
	return
}

func parseInput(input string) inputType {
	return strings.Split(input, "\n")
}

func runeToInt(r rune) int {
	num, err := strconv.ParseUint(string(r), 10, 0)
	if err != nil {
		fmt.Println(err)
	}

	return int(num)
}
