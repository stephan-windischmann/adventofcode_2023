package day7

import (
	"slices"
	"strconv"
	"strings"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func countList(m map[rune]int) []int {
	c := make([]int, 0)

	for _, v := range m {
		c = append(c, v)
	}

	slices.Sort(c)
	slices.Reverse(c)
	return c
}

func isFiveKind(hand string) bool {
	counts := countList(util.CountChars(hand))

	return counts[0] == 5
}

func isFourKind(hand string) bool {
	counts := countList(util.CountChars(hand))

	return counts[0] == 4
}

func isFullHouse(hand string) bool {
	counts := countList(util.CountChars(hand))

	return counts[0] == 3 && counts[1] == 2
}

func isThreeKind(hand string) bool {
	counts := countList(util.CountChars(hand))

	return counts[0] == 3
}

func isTwoPairs(hand string) bool {
	counts := countList(util.CountChars(hand))

	return counts[0] == 2 && counts[1] == 2
}

func isPair(hand string) bool {
	counts := countList(util.CountChars(hand))

	return counts[0] == 2
}

func sortFunc(a, b string) int {
	handA := strings.Split(a, " ")[0]
	handB := strings.Split(b, " ")[0]

	cards := []byte{
		'A',
		'K',
		'Q',
		'J',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
	}

	for i := 0; i < len(handA); i++ {
		if handA[i] == handB[i] {
			continue
		}
		idxA := util.IndexByte(handA[i], cards)
		idxB := util.IndexByte(handB[i], cards)

		if idxA < idxB {
			return 1
		}
		return -1
	}
	return 0
}

func sortHands(hands []string) []string {
	fiveKind := make([]string, 0)
	fourKind := make([]string, 0)
	fullHouse := make([]string, 0)
	threeKind := make([]string, 0)
	twoPair := make([]string, 0)
	onePair := make([]string, 0)
	highCard := make([]string, 0)

	for _, h := range hands {
		hand := strings.Fields(h)[0]
		switch {
		case isFiveKind(hand):
			fiveKind = append(fiveKind, h)
		case isFourKind(hand):
			fourKind = append(fourKind, h)
		case isFullHouse(hand):
			fullHouse = append(fullHouse, h)
		case isThreeKind(hand):
			threeKind = append(threeKind, h)
		case isTwoPairs(hand):
			twoPair = append(twoPair, h)
		case isPair(hand):
			onePair = append(onePair, h)
		default:
			highCard = append(highCard, h)
		}
	}

	slices.SortFunc(fiveKind, sortFunc)
	slices.SortFunc(fourKind, sortFunc)
	slices.SortFunc(fullHouse, sortFunc)
	slices.SortFunc(threeKind, sortFunc)
	slices.SortFunc(twoPair, sortFunc)
	slices.SortFunc(onePair, sortFunc)
	slices.SortFunc(highCard, sortFunc)

	sortedHands := highCard
	sortedHands = append(sortedHands, onePair...)
	sortedHands = append(sortedHands, twoPair...)
	sortedHands = append(sortedHands, threeKind...)
	sortedHands = append(sortedHands, fullHouse...)
	sortedHands = append(sortedHands, fourKind...)
	sortedHands = append(sortedHands, fiveKind...)

	return sortedHands
}

func jokerCount(hand string) int {
	return strings.Count(hand, "J")
}

func isFiveKindJoker(hand string) bool {
	counts := countList(util.CountCharsExclude(hand, 'J'))

	if len(counts) == 0 {
		return jokerCount(hand) == 5
	}

	return counts[0] == 5 || counts[0]+jokerCount(hand) == 5
}

func isFourKindJoker(hand string) bool {
	counts := countList(util.CountCharsExclude(hand, 'J'))

	return counts[0] == 4 || counts[0]+jokerCount(hand) == 4
}

func isFullHouseJoker(hand string) bool {
	counts := countList(util.CountCharsExclude(hand, 'J'))
	jokers := jokerCount(hand)
	jokersNeeded := (3 - counts[0]) + (2 - counts[1])

	return (counts[0] == 3 && counts[1] == 2) || jokers == jokersNeeded
}

func isThreeKindJoker(hand string) bool {
	counts := countList(util.CountCharsExclude(hand, 'J'))

	return counts[0] == 3 || counts[0]+jokerCount(hand) == 3
}

func isTwoPairsJoker(hand string) bool {
	counts := countList(util.CountCharsExclude(hand, 'J'))
	jokers := jokerCount(hand)
	jokersNeeded := (2 - counts[0]) + (2 - counts[1])

	return (counts[0] == 2 && counts[1] == 2) || jokers == jokersNeeded
}

func isPairJoker(hand string) bool {
	counts := countList(util.CountCharsExclude(hand, 'J'))

	return counts[0] == 2 || counts[0]+jokerCount(hand) == 2
}

func sortFuncJoker(a, b string) int {
	handA := strings.Split(a, " ")[0]
	handB := strings.Split(b, " ")[0]

	cards := []byte{
		'A',
		'K',
		'Q',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
		'J',
	}

	for i := 0; i < len(handA); i++ {
		if handA[i] == handB[i] {
			continue
		}
		idxA := util.IndexByte(handA[i], cards)
		idxB := util.IndexByte(handB[i], cards)

		if idxA < idxB {
			return 1
		}
		return -1
	}
	return 0
}

func sortHandsJoker(hands []string) []string {
	fiveKind := make([]string, 0)
	fourKind := make([]string, 0)
	fullHouse := make([]string, 0)
	threeKind := make([]string, 0)
	twoPair := make([]string, 0)
	onePair := make([]string, 0)
	highCard := make([]string, 0)

	for _, h := range hands {
		hand := strings.Fields(h)[0]
		switch {
		case isFiveKindJoker(hand):
			fiveKind = append(fiveKind, h)
		case isFourKindJoker(hand):
			fourKind = append(fourKind, h)
		case isFullHouseJoker(hand):
			fullHouse = append(fullHouse, h)
		case isThreeKindJoker(hand):
			threeKind = append(threeKind, h)
		case isTwoPairsJoker(hand):
			twoPair = append(twoPair, h)
		case isPairJoker(hand):
			onePair = append(onePair, h)
		default:
			highCard = append(highCard, h)
		}
	}

	slices.SortFunc(fiveKind, sortFuncJoker)
	slices.SortFunc(fourKind, sortFuncJoker)
	slices.SortFunc(fullHouse, sortFuncJoker)
	slices.SortFunc(threeKind, sortFuncJoker)
	slices.SortFunc(twoPair, sortFuncJoker)
	slices.SortFunc(onePair, sortFuncJoker)
	slices.SortFunc(highCard, sortFuncJoker)

	sortedHands := highCard
	sortedHands = append(sortedHands, onePair...)
	sortedHands = append(sortedHands, twoPair...)
	sortedHands = append(sortedHands, threeKind...)
	sortedHands = append(sortedHands, fullHouse...)
	sortedHands = append(sortedHands, fourKind...)
	sortedHands = append(sortedHands, fiveKind...)

	return sortedHands
}

func SolvePart1(input []string) int {
	sortedHands := sortHands(input)

	winnings := 0

	for i, hand := range sortedHands {
		winning, err := strconv.Atoi(strings.Split(hand, " ")[1])
		if err != nil {
			panic(err)
		}
		winnings += (i + 1) * winning
	}

	return winnings
}

func SolvePart2(input []string) int {
	sortedHands := sortHandsJoker(input)

	winnings := 0

	for i, hand := range sortedHands {
		winning, err := strconv.Atoi(strings.Split(hand, " ")[1])
		if err != nil {
			panic(err)
		}
		winnings += (i + 1) * winning
	}

	return winnings
}
