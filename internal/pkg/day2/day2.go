package day2

import (
	"strconv"
	"strings"
)

func isPossible(game string, bag map[string]int) bool {
	for _, round := range strings.Split(strings.Split(game, ":")[1], ";") {
		for _, color := range strings.Split(round, ",") {
			colorName := strings.Split(strings.TrimSpace(color), " ")[1]
			colorCount, err := strconv.Atoi(strings.Split(strings.TrimSpace(color), " ")[0])
			if err != nil {
				panic(err)
			}
			if colorCount > bag[colorName] {
				return false
			}
		}
	}
	return true
}

func SolvePart1(input []string) int {
	sumGameId := 0

	bag := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, game := range input {
		gameId, err := strconv.Atoi(strings.Split(strings.Split(game, ":")[0], " ")[1])
		if err != nil {
			panic(err)
		}

		if isPossible(game, bag) {
			sumGameId += gameId
		}
	}

	return sumGameId
}

func SolvePart2(input []string) int {
	sumPower := 0

	for _, game := range input {
		minCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range strings.Split(strings.Split(game, ":")[1], ";") {
			for _, color := range strings.Split(round, ",") {
				colorName := strings.Split(strings.TrimSpace(color), " ")[1]
				colorCount, err := strconv.Atoi(strings.Split(strings.TrimSpace(color), " ")[0])
				if err != nil {
					panic(err)
				}

				minCubes[colorName] = max(minCubes[colorName], colorCount)
			}
		}

		sumPower += minCubes["red"] * minCubes["green"] * minCubes["blue"]
	}

	return sumPower
}
