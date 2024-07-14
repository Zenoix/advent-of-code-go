package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/zenoix/advent-of-code-go/utils"
)

type inputType = []string

type Coordinate struct {
	row int
	col int
}

type Offset struct {
	rowOffset int
	colOffset int
}

var adjacentOffsets = [...]Offset{
	{rowOffset: -1, colOffset: -1},
	{rowOffset: 0, colOffset: -1},
	{rowOffset: 1, colOffset: -1},
	{rowOffset: -1, colOffset: 0},
	{rowOffset: 1, colOffset: 0},
	{rowOffset: -1, colOffset: 1},
	{rowOffset: 0, colOffset: 1},
	{rowOffset: 1, colOffset: 1},
}

func main() {
	input := utils.GetInput(2023, 3)

	parsedInput := parseInput(input)

	ans1 := part1(parsedInput)
	ans2 := part2(parsedInput)

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}

func part1(input inputType) (ans int) {
	for r, line := range input {
		isPartNumber := false
		currNumber := 0

		for c, char := range line {
			currentCoord := Coordinate{row: r, col: c}
			if unicode.IsDigit(char) {
				currNumber = currNumber*10 + int(char-'0')
				if !isPartNumber {
					isPartNumber = hasAdjacentSymbol(currentCoord, input)
				}
			} else {
				if isPartNumber {
					ans += currNumber
				}
				currNumber = 0
				isPartNumber = false
			}
		}

		if isPartNumber {
			ans += currNumber
		}
	}
	return
}

func part2(input inputType) (ans int) {
	stars := make(map[Coordinate][]int)

	for r, line := range input {
		currNumber := 0
		adjStars := map[Coordinate]struct{}{}

		for c, char := range line {
			currentCoord := Coordinate{row: r, col: c}
			if unicode.IsDigit(char) {
				currNumber = currNumber*10 + int(char-'0')
				charAdjStars := findAdjacentStars(currentCoord, input)
				for _, coord := range charAdjStars {
					adjStars[coord] = struct{}{}
				}

			} else {
				if len(adjStars) != 0 {
					for coord := range adjStars {
						stars[coord] = append(stars[coord], currNumber)
					}
				}
				currNumber = 0
				adjStars = map[Coordinate]struct{}{}
			}
		}

		if len(adjStars) != 0 {
			for coord := range adjStars {
				stars[coord] = append(stars[coord], currNumber)
			}
		}
	}
	for _, nums := range stars {
		if len(nums) == 2 {
			ans += nums[0] * nums[1]
		}
	}
	return
}

func hasAdjacentSymbol(coord Coordinate, s []string) bool {
	numRow := len(s)
	numCol := len(s[0])

	r := coord.row
	c := coord.col

	for _, offset := range adjacentOffsets {
		if (r == 0 && offset.rowOffset == -1) ||
			(c == 0 && offset.colOffset == -1) ||
			(r == numRow-1 && offset.rowOffset == 1) ||
			(c == numCol-1 && offset.colOffset == 1) {
			continue
		}

		x := rune(s[r+offset.rowOffset][c+offset.colOffset])
		if x != '.' && (unicode.IsPunct(x) || unicode.IsSymbol(x)) {
			return true
		}
	}

	return false
}

func findAdjacentStars(coord Coordinate, s []string) []Coordinate {
	adjStarCoords := make([]Coordinate, 0, 8)

	numRow := len(s)
	numCol := len(s[0])

	r := coord.row
	c := coord.col

	for _, offset := range adjacentOffsets {
		if (r == 0 && offset.rowOffset == -1) ||
			(c == 0 && offset.colOffset == -1) ||
			(r == numRow-1 && offset.rowOffset == 1) ||
			(c == numCol-1 && offset.colOffset == 1) {
			continue
		}

		offsetR := r + offset.rowOffset
		offsetC := c + offset.colOffset
		x := rune(s[offsetR][offsetC])
		if x == '*' {
			adjStarCoords = append(adjStarCoords, Coordinate{row: offsetR, col: offsetC})
		}
	}

	return adjStarCoords
}

func parseInput(input string) inputType {
	return strings.Split(input, "\n")
}
