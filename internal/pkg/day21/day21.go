package day21

import (
	"strconv"
	"strings"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day9"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func SolvePart1(input []string, maxSteps int) int {
	m := util.CreateRuneTable(input)

	for i := 0; i < maxSteps; i++ {
		newMap := make([][]rune, len(m))
		for i := 0; i < len(m); i++ {
			newMap[i] = make([]rune, len(m[0]))
			for j := 0; j < len(m[0]); j++ {
				newMap[i][j] = m[i][j]
			}
		}

		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[0]); j++ {
				if m[i][j] == 'O' || m[i][j] == 'S' {
					if i > 0 && m[i-1][j] != '#' && m[i-1][j] != 'O' {
						newMap[i-1][j] = 'O'
					}

					if i < len(m)-1 && m[i+1][j] != '#' && m[i+1][j] != 'O' {
						newMap[i+1][j] = 'O'
					}

					if j > 0 && m[i][j-1] != '#' && m[i][j-1] != 'O' {
						newMap[i][j-1] = 'O'
					}

					if j < len(m[0])-1 && m[i][j+1] != '#' && m[i][j+1] != 'O' {
						newMap[i][j+1] = 'O'
					}
					newMap[i][j] = '.'
				}
			}
		}

		m = newMap
	}

	count := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 'O' {
				count++
			}
		}
	}

	return count
}

func SolvePart2(input []string, maxSteps int) int {
	inputNoStart := make([]string, len(input)*5)
	for i := 0; i < len(input); i++ {
		if strings.Contains(input[i], "S") {
			inputNoStart[i] =
				input[i][:strings.Index(input[i], "S")] +
					"." + input[i][strings.Index(input[i], "S")+1:]
		} else {
			inputNoStart[i] = input[i]
		}
	}

	// Expand input
	expandedIput := make([]string, len(input)*5)
	for i := 0; i < 5; i++ {
		for j := 0; j < len(input); j++ {
			if i == 2 {
				expandedIput[(i*len(input))+j] = strings.Repeat(inputNoStart[j], 2) + input[j] + strings.Repeat(inputNoStart[j], 2)
			} else {
				expandedIput[(i*len(input))+j] = strings.Repeat(inputNoStart[j], 5)
			}
		}
	}

	// Get visited plots after 65, 65 + 131, 65 + (2 * 131)
	visited := make([]string, 0)
	i := 0
	for i < 3 {
		visited = append(visited, strconv.Itoa(SolvePart1(expandedIput, 65+(i*131))))
		i++
	}

	// Now use code from day 9 to get result. Thanks to
	// https://www.reddit.com/r/adventofcode/comments/18orn0s/2023_day_21_part_2_links_between_days/
	// for the idea.
	// WARNING: Slow, runtime ~25 minutes.
	var totalPlots int
	for i < 202301 {
		totalPlots = day9.SolvePart1([]string{strings.Join(visited, " ")})
		visited = append(visited, strconv.Itoa(totalPlots))
		i++
	}

	return totalPlots
}
