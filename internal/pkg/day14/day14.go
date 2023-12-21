package day14

import (
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"strings"
)

func moveRocksNorth(rocks [][]rune) {
	for {
		movements := 0

		for i := 0; i < len(rocks)-1; i++ {
			for j := 0; j < len(rocks[i]); j++ {
				if rocks[i][j] == '.' && rocks[i+1][j] == 'O' {
					rocks[i][j] = 'O'
					rocks[i+1][j] = '.'
					movements++
				}
			}
		}

		if movements == 0 {
			break
		}
	}
}

func moveRocksSouth(rocks [][]rune) {
	for {
		movements := 0

		for i := len(rocks) - 1; i > 0; i-- {
			for j := 0; j < len(rocks[i]); j++ {
				if rocks[i][j] == '.' && rocks[i-1][j] == 'O' {
					rocks[i][j] = 'O'
					rocks[i-1][j] = '.'
					movements++
				}
			}
		}

		if movements == 0 {
			break
		}
	}
}

func moveRocksWest(rocks [][]rune) {
	for {
		movements := 0

		for i := 0; i < len(rocks[0])-1; i++ {
			for j := 0; j < len(rocks); j++ {
				if rocks[j][i] == '.' && rocks[j][i+1] == 'O' {
					rocks[j][i] = 'O'
					rocks[j][i+1] = '.'
					movements++
				}
			}
		}

		if movements == 0 {
			break
		}
	}
}

func moveRocksEast(rocks [][]rune) {
	for {
		movements := 0

		for i := len(rocks[0]) - 1; i > 0; i-- {
			for j := 0; j < len(rocks); j++ {
				if rocks[j][i] == '.' && rocks[j][i-1] == 'O' {
					rocks[j][i] = 'O'
					rocks[j][i-1] = '.'
					movements++
				}
			}
		}

		if movements == 0 {
			break
		}
	}
}

func getLoad(rocks [][]rune) int {
	load := 0

	for i := 0; i < len(rocks); i++ {
		for j := 0; j < len(rocks[i]); j++ {
			if rocks[i][j] == 'O' {
				load += len(rocks) - i
			}
		}
	}

	return load
}

func getCurState(rocks [][]rune) string {
	var state strings.Builder

	for i := 0; i < len(rocks); i++ {
		for j := 0; j < len(rocks[i]); j++ {
			state.WriteRune(rocks[i][j])
		}
		state.WriteString("\n")
	}

	return state.String()
}

func SolvePart1(input []string) int {
	rocks := util.CreateRuneTable(input)

	moveRocksNorth(rocks)
	return getLoad(rocks)
}

func SolvePart2(input []string) int {
	rocks := util.CreateRuneTable(input)

	// First we find the cycle
	memo := make(map[string]int)
	i := 0
	var cycle int
	for i < 1000000000 {
		curState := getCurState(rocks)
		if firstFound, ok := memo[curState]; ok {
			cycle = i - firstFound
			break
		} else {
			memo[curState] = i
		}

		moveRocksNorth(rocks)
		moveRocksWest(rocks)
		moveRocksSouth(rocks)
		moveRocksEast(rocks)
		i++
	}
	// Now we fast forwards
	i = i + ((1000000000-i)/cycle)*cycle

	// And finish with the last few moves
	for i < 1000000000 {
		moveRocksNorth(rocks)
		moveRocksWest(rocks)
		moveRocksSouth(rocks)
		moveRocksEast(rocks)
		i++
	}

	return getLoad(rocks)
}
