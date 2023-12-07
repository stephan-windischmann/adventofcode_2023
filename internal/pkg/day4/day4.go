package day4

import (
	"strings"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func SolvePart1(input []string) int {
	points := 0

	for _, card := range input {
		allNumbers := strings.Split(card, ":")[1]

		winningNums := strings.Fields(strings.Split(allNumbers, "|")[0])
		s := util.NewSet()
		for _, n := range winningNums {
			s.Add(n)
		}

		myNums := strings.Fields(strings.Split(allNumbers, "|")[1])
		p := 0
		for _, n := range myNums {
			if s.Contains(n) {
				if p == 0 {
					p = 1
				} else {
					p *= 2
				}
			}
		}

		points += p
	}

	return points
}

func SolvePart2(input []string) int {
	cardCount := make([]int, len(input))
	for i := 0; i < len(cardCount); i++ {
		cardCount[i] = 1
	}

	for curCard, card := range input {
		allNumbers := strings.Split(card, ":")[1]

		winningNums := strings.Fields(strings.Split(allNumbers, "|")[0])
		s := util.NewSet()
		for _, n := range winningNums {
			s.Add(n)
		}

		myNums := strings.Fields(strings.Split(allNumbers, "|")[1])

		numMatches := 0
		for _, n := range myNums {
			if s.Contains(n) {
				numMatches++
			}
		}
		for i := 1; i <= numMatches; i++ {
			cardCount[curCard+i] += cardCount[curCard]
		}
	}

	return util.SumInt(cardCount)
}
