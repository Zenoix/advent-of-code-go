package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zenoix/advent-of-code-go/utils"
)

type inputType = []string

func main() {
	input := utils.GetInput(2023, 2)

	parsedInput := parseInput(input)

	ans1 := part1(parsedInput)
	ans2 := part2(parsedInput)

	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans2)
}

func part1(input inputType) (ans int) {
	for _, line := range input {
		b, game, _ := strings.Cut(line, ": ")
		gameID, _ := strconv.ParseInt(b, 10, 0)

		games := strings.ReplaceAll(game, ",", "")
		splitGames := strings.Split(games, "; ")

		isValidGameFlag := true
		for _, g := range splitGames {
			numColAlt := strings.Split(g, " ")
			isValidGameFlag = isValidGame(numColAlt)
			if !isValidGameFlag {
				break
			}
		}
		if isValidGameFlag {
			ans += int(gameID)
		}
	}
	return
}

func part2(input inputType) (ans int) {
	for _, line := range input {
		_, game, _ := strings.Cut(line, ": ")

		games := strings.ReplaceAll(game, ",", "")
		splitGames := strings.Split(games, "; ")

		var maxR, maxG, maxB int = 0, 0, 0
		for _, g := range splitGames {
			numCubes := 0
			for i, x := range strings.Split(g, " ") {
				if i%2 == 0 {
					v, _ := strconv.ParseInt(x, 10, 0)
					numCubes = int(v)
				} else {
					switch x {
					case "red":
						if numCubes > maxR {
							maxR = numCubes
						}
					case "green":
						if numCubes > maxG {
							maxG = numCubes
						}
					case "blue":
						if numCubes > maxB {
							maxB = numCubes
						}
					}
				}
			}
		}
		ans += maxR * maxG * maxB
	}
	return
}

func isValidGame(game []string) bool {
	numCubes := 0
	for i, x := range game {
		if i%2 == 0 {
			v, _ := strconv.ParseInt(x, 10, 0)
			numCubes = int(v)
		} else {
			if !isValidRound(x, numCubes, 12, 13, 14) {
				return false
			}
		}
	}
	return true
}

func isValidRound(
	currColor string,
	numCubes,
	redLimit,
	greenLimit,
	blueLimit int) bool {
	return numCubes <= redLimit && currColor == "red" ||
		numCubes <= greenLimit && currColor == "green" ||
		numCubes <= blueLimit && currColor == "blue"
}

func parseInput(input string) inputType {
	parsed := make([]string, 0, 2048)

	for _, line := range strings.Split(input, "\n") {
		parsed = append(parsed, line[5:])
	}

	return parsed
}
