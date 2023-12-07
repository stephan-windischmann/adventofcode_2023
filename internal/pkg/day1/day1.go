package day1

import (
	"strings"
	"unicode"
)

func SolvePart1(input []string) int {
	total := 0

	for _, line := range input {
		var l, r int
		for _, c := range line {
			if unicode.IsDigit(c) {
				if l == 0 {
					l = int(c - '0')
				}
				r = int(c - '0')
			}
		}
		total += (l * 10) + r
	}

	return total
}

func SolvePart2(input []string) int {
	total := 0

	for _, line := range input {
		var curLine []int

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				curLine = append(curLine, int(line[i]-'0'))
			}

			switch {
			case strings.HasPrefix(line[i:], "one"):
				curLine = append(curLine, 1)
			case strings.HasPrefix(line[i:], "two"):
				curLine = append(curLine, 2)
			case strings.HasPrefix(line[i:], "three"):
				curLine = append(curLine, 3)
			case strings.HasPrefix(line[i:], "four"):
				curLine = append(curLine, 4)
			case strings.HasPrefix(line[i:], "five"):
				curLine = append(curLine, 5)
			case strings.HasPrefix(line[i:], "six"):
				curLine = append(curLine, 6)
			case strings.HasPrefix(line[i:], "seven"):
				curLine = append(curLine, 7)
			case strings.HasPrefix(line[i:], "eight"):
				curLine = append(curLine, 8)
			case strings.HasPrefix(line[i:], "nine"):
				curLine = append(curLine, 9)
			}
		}

		total += (curLine[0] * 10) + curLine[len(curLine)-1]
	}

	return total
}
