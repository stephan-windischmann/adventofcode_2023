package day3

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func isPartNumber(table [][]rune, x, endX, y int) bool {
	if x > 0 {
		if table[y][x-1] != '.' && !unicode.IsDigit(table[y][x-1]) {
			return true
		}
		if y > 0 {
			if table[y-1][x-1] != '.' && !unicode.IsDigit(table[y-1][x-1]) {
				return true
			}
		}
		if y < len(table)-1 {
			if table[y+1][x-1] != '.' && !unicode.IsDigit(table[y+1][x-1]) {
				return true
			}
		}
	}
	for i := x; i < endX; i++ {
		if y > 0 {
			if table[y-1][i] != '.' && !unicode.IsDigit(table[y-1][i]) {
				return true
			}
		}
		if y < len(table)-1 {
			if table[y+1][i] != '.' && !unicode.IsDigit(table[y+1][i]) {
				return true
			}
		}
	}
	if endX < len(table[y]) {
		if table[y][endX] != '.' && !unicode.IsDigit(table[y][endX]) {
			return true
		}
		if y > 0 {
			if table[y-1][endX] != '.' && !unicode.IsDigit(table[y-1][endX]) {
				return true
			}
		}
		if y < len(table)-1 {
			if table[y+1][endX] != '.' && !unicode.IsDigit(table[y+1][endX]) {
				return true
			}
		}
	}
	return false
}

func parsePartNumber(table [][]rune, x, y int) (int, int) {
	num := 0

	for x < len(table[y]) && unicode.IsDigit(table[y][x]) {
		num = (num * 10) + util.RuneToInt(table[y][x])
		x++
	}

	return x, num
}

func SolvePart1(input []string) int {
	table := util.CreateRuneTable(input)
	sum := 0

	for y := 0; y < len(table); y++ {
		for x := 0; x < len(table[y]); x++ {
			endX, num := parsePartNumber(table, x, y)
			if endX > x && isPartNumber(table, x, endX, y) {
				sum += num
			}
			x = endX
		}
	}

	return sum
}

func parseNum(table [][]rune, x, y int) int {
	var sb strings.Builder
	// First go left to find the beginning of the number
	for x > 0 && unicode.IsDigit(table[y][x-1]) {
		x--
	}
	for x < len(table[y]) && unicode.IsDigit(table[y][x]) {
		sb.WriteRune(table[y][x])
		table[y][x] = 'X'
		x++
	}

	n, _ := strconv.Atoi(sb.String())
	return n
}

func processGear(table [][]rune, x, y int) []int {
	partNums := make([]int, 0)

	if y > 0 && x > 0 {
		if unicode.IsDigit(table[y-1][x-1]) {
			partNums = append(partNums, parseNum(table, x-1, y-1))
		}
	}
	if y > 0 {
		if unicode.IsDigit(table[y-1][x]) {
			partNums = append(partNums, parseNum(table, x, y-1))
		}
	}
	if y > 0 && x < len(table[y])-1 {
		if unicode.IsDigit(table[y-1][x+1]) {
			partNums = append(partNums, parseNum(table, x+1, y-1))
		}
	}
	if x < len(table[y])-1 {
		if unicode.IsDigit(table[y][x+1]) {
			partNums = append(partNums, parseNum(table, x+1, y))
		}
	}
	if y < len(table)-1 && x < len(table[y])-1 {
		if unicode.IsDigit(table[y+1][x+1]) {
			partNums = append(partNums, parseNum(table, x+1, y+1))
		}
	}
	if y < len(table)-1 {
		if unicode.IsDigit(table[y+1][x]) {
			partNums = append(partNums, parseNum(table, x, y+1))
		}
	}
	if y < len(table)-1 && x > 0 {
		if unicode.IsDigit(table[y+1][x-1]) {
			partNums = append(partNums, parseNum(table, x-1, y+1))
		}
	}
	if x > 0 {
		if unicode.IsDigit(table[y][x-1]) {
			partNums = append(partNums, parseNum(table, x-1, y))
		}
	}

	return partNums
}

func SolvePart2(input []string) int {
	table := util.CreateRuneTable(input)
	sum := 0

	for y := 0; y < len(table); y++ {
		for x := 0; x < len(table[y]); x++ {
			if table[y][x] == '*' {
				partNums := processGear(table, x, y)
				if len(partNums) == 2 {
					sum += partNums[0] * partNums[1]
				}
			}
		}
	}

	return sum
}
