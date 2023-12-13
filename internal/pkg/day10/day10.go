package day10

import (
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func findStart(input [][]rune) (int, int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'S' {
				return j, i
			}
		}
	}
	return -1, -1
}

func getNextStep(pipes [][]rune, curPipe rune, curPos []int) []int {
	switch curPipe {
	case '|':
		if pipes[curPos[1]-1][curPos[0]] == 'X' {
			return []int{curPos[0], curPos[1] + 1}
		}
		return []int{curPos[0], curPos[1] - 1}
	case '-':
		if pipes[curPos[1]][curPos[0]-1] == 'X' {
			return []int{curPos[0] + 1, curPos[1]}
		}
		return []int{curPos[0] - 1, curPos[1]}
	case 'L':
		if pipes[curPos[1]-1][curPos[0]] == 'X' {
			return []int{curPos[0] + 1, curPos[1]}
		}
		return []int{curPos[0], curPos[1] - 1}
	case 'J':
		if pipes[curPos[1]-1][curPos[0]] == 'X' {
			return []int{curPos[0] - 1, curPos[1]}
		}
		return []int{curPos[0], curPos[1] - 1}
	case '7':
		if pipes[curPos[1]+1][curPos[0]] == 'X' {
			return []int{curPos[0] - 1, curPos[1]}
		}
		return []int{curPos[0], curPos[1] + 1}
	case 'F':
		if pipes[curPos[1]+1][curPos[0]] == 'X' {
			return []int{curPos[0] + 1, curPos[1]}
		}
		return []int{curPos[0], curPos[1] + 1}
	}
	return []int{-1, -1}
}

func SolvePart1(input []string) int {
	pipes := util.ToRune(input)
	startX, startY := findStart(pipes)

	a := []int{-1, -1}
	b := []int{-1, -1}

	if startY > 0 {
		if pipes[startY-1][startX] == '|' ||
			pipes[startY-1][startX] == '7' ||
			pipes[startY-1][startX] == 'F' {
			if a[0] == -1 {
				a[0] = startX
				a[1] = startY - 1
			} else {
				b[0] = startX
				b[1] = startY - 1
			}
		}
	}
	if startY < len(pipes)-1 {
		if pipes[startY+1][startX] == '|' ||
			pipes[startY+1][startX] == 'J' ||
			pipes[startY+1][startX] == 'L' {
			if a[0] == -1 {
				a[0] = startX
				a[1] = startY + 1
			} else {
				b[0] = startX
				b[1] = startY + 1
			}
		}
	}
	if startX > 0 {
		if pipes[startY][startX-1] == '-' ||
			pipes[startY][startX-1] == 'L' ||
			pipes[startY][startX-1] == 'F' {
			if a[0] == -1 {
				a[0] = startX - 1
				a[1] = startY
			} else {
				b[0] = startX - 1
				b[1] = startY
			}
		}
	}
	if startX < len(pipes[startY])-1 {
		if pipes[startY][startX+1] == '-' ||
			pipes[startY][startX+1] == '7' ||
			pipes[startY][startX+1] == 'J' {
			if a[0] == -1 {
				a[0] = startX + 1
				a[1] = startY
			} else {
				b[0] = startX + 1
				b[1] = startY
			}
		}
	}

	steps := 1
	pipes[startY][startX] = 'X'

	for {
		curA := pipes[a[1]][a[0]]
		curB := pipes[b[1]][b[0]]
		pipes[a[1]][a[0]] = 'X'
		pipes[b[1]][b[0]] = 'X'
		a = getNextStep(pipes, curA, a)
		b = getNextStep(pipes, curB, b)
		steps++

		if a[0] == b[0] && a[1] == b[1] {
			return steps
		}
	}
}

func floodFill(pipes [][]rune, x, y int) {
	if x < 0 || y < 0 || y == len(pipes) || x == len(pipes[y]) {
		return
	}
	if pipes[y][x] == '#' || pipes[y][x] == '0' {
		return
	}

	pipes[y][x] = '0'

	floodFill(pipes, x-1, y)
	floodFill(pipes, x+1, y)
	floodFill(pipes, x, y-1)
	floodFill(pipes, x, y+1)
}

func buildEpxandedMap(m [][]rune) [][]rune {
	expandedMap := make([][]rune, len(m)*3)
	for i := 0; i < len(expandedMap); i++ {
		expandedMap[i] = make([]rune, len(m[0])*3)
	}

	/*
	 *             ...
	 * '-' becomes ###
	 *             ...
	 *
	 *             .#.
	 * 'L' becomes .##
	 *             ...
	 *
	 * Etc...
	 */

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			switch m[i][j] {
			case '-':
				expandedMap[i*3][j*3] = '.'
				expandedMap[i*3][(j*3)+1] = '.'
				expandedMap[i*3][(j*3)+2] = '.'
				expandedMap[(i*3)+1][j*3] = '#'
				expandedMap[(i*3)+1][(j*3)+1] = '#'
				expandedMap[(i*3)+1][(j*3)+2] = '#'
				expandedMap[(i*3)+2][j*3] = '.'
				expandedMap[(i*3)+2][(j*3)+1] = '.'
				expandedMap[(i*3)+2][(j*3)+2] = '.'
			case '|':
				expandedMap[i*3][j*3] = '.'
				expandedMap[i*3][(j*3)+1] = '#'
				expandedMap[i*3][(j*3)+2] = '.'
				expandedMap[(i*3)+1][j*3] = '.'
				expandedMap[(i*3)+1][(j*3)+1] = '#'
				expandedMap[(i*3)+1][(j*3)+2] = '.'
				expandedMap[(i*3)+2][j*3] = '.'
				expandedMap[(i*3)+2][(j*3)+1] = '#'
				expandedMap[(i*3)+2][(j*3)+2] = '.'
			case 'L':
				expandedMap[i*3][j*3] = '.'
				expandedMap[i*3][(j*3)+1] = '#'
				expandedMap[i*3][(j*3)+2] = '.'
				expandedMap[(i*3)+1][j*3] = '.'
				expandedMap[(i*3)+1][(j*3)+1] = '#'
				expandedMap[(i*3)+1][(j*3)+2] = '#'
				expandedMap[(i*3)+2][j*3] = '.'
				expandedMap[(i*3)+2][(j*3)+1] = '.'
				expandedMap[(i*3)+2][(j*3)+2] = '.'
			case 'J':
				expandedMap[i*3][j*3] = '.'
				expandedMap[i*3][(j*3)+1] = '#'
				expandedMap[i*3][(j*3)+2] = '.'
				expandedMap[(i*3)+1][j*3] = '#'
				expandedMap[(i*3)+1][(j*3)+1] = '#'
				expandedMap[(i*3)+1][(j*3)+2] = '.'
				expandedMap[(i*3)+2][j*3] = '.'
				expandedMap[(i*3)+2][(j*3)+1] = '.'
				expandedMap[(i*3)+2][(j*3)+2] = '.'
			case '7':
				expandedMap[i*3][j*3] = '.'
				expandedMap[i*3][(j*3)+1] = '.'
				expandedMap[i*3][(j*3)+2] = '.'
				expandedMap[(i*3)+1][j*3] = '#'
				expandedMap[(i*3)+1][(j*3)+1] = '#'
				expandedMap[(i*3)+1][(j*3)+2] = '.'
				expandedMap[(i*3)+2][j*3] = '.'
				expandedMap[(i*3)+2][(j*3)+1] = '#'
				expandedMap[(i*3)+2][(j*3)+2] = '.'
			case 'F':
				expandedMap[i*3][j*3] = '.'
				expandedMap[i*3][(j*3)+1] = '.'
				expandedMap[i*3][(j*3)+2] = '.'
				expandedMap[(i*3)+1][j*3] = '.'
				expandedMap[(i*3)+1][(j*3)+1] = '#'
				expandedMap[(i*3)+1][(j*3)+2] = '#'
				expandedMap[(i*3)+2][j*3] = '.'
				expandedMap[(i*3)+2][(j*3)+1] = '#'
				expandedMap[(i*3)+2][(j*3)+2] = '.'
			default:
				expandedMap[i*3][j*3] = '.'
				expandedMap[i*3][(j*3)+1] = '.'
				expandedMap[i*3][(j*3)+2] = '.'
				expandedMap[(i*3)+1][j*3] = '.'
				expandedMap[(i*3)+1][(j*3)+1] = '.'
				expandedMap[(i*3)+1][(j*3)+2] = '.'
				expandedMap[(i*3)+2][j*3] = '.'
				expandedMap[(i*3)+2][(j*3)+1] = '.'
				expandedMap[(i*3)+2][(j*3)+2] = '.'
			}
		}
	}

	return expandedMap
}

func replaceStart(m [][]rune, x, y int) {
	if x > 0 && x < len(m[y])-1 {
		if (m[y][x-1] == '-' || m[y][x-1] == 'F' || m[y][x-1] == 'L') &&
			(m[y][x+1] == '-' || m[y][x+1] == 'J' || m[y][x+1] == '7') {
			m[y][x] = '-'
			return
		}
	}
	if y > 0 && y < len(m)-1 {
		if (m[y-1][x] == '|' || m[y-1][x] == 'F' || m[y-1][x] == '7') &&
			(m[y+1][x] == '|' || m[y+1][x] == 'L' || m[y+1][x] == 'J') {
			m[y][x] = '|'
			return
		}
	}
	if x > 0 && y > 0 {
		if (m[y][x-1] == '-' || m[y][x-1] == 'F' || m[y][x-1] == 'L') &&
			(m[y-1][x] == '|' || m[y-1][x] == 'F' || m[y-1][x] == '7') {
			m[y][x] = 'J'
			return
		}
	}
	if x < len(m[y])-1 && y > 0 {
		if (m[y][x+1] == '-' || m[y][x+1] == '7' || m[y][x+1] == 'J') &&
			(m[y-1][x] == '|' || m[y-1][x] == '7' || m[y-1][x] == 'F') {
			m[y][x] = 'L'
			return
		}
	}
	if x > 0 && y < len(m)-1 {
		if (m[y][x-1] == '-' || m[y][x-1] == 'F' || m[y][x-1] == 'L') &&
			(m[y+1][x] == '|' || m[y+1][x] == 'J' || m[y+1][x] == 'L') {
			m[y][x] = '7'
			return
		}
	}
	if x < len(m[y])-1 && y < len(m)-1 {
		if (m[y][x+1] == '-' || m[y][x+1] == '7' || m[y][x+1] == 'J') &&
			(m[y+1][x] == '|' || m[y+1][x] == 'L' || m[y+1][x] == 'J') {
			m[y][x] = 'F'
			return
		}
	}
}

func SolvePart2(input []string) int {
	pipes := util.ToRune(input)

	// First find and mark the main loop
	startX, startY := findStart(pipes)

	a := []int{-1, -1}
	b := []int{-1, -1}

	if startY > 0 {
		if pipes[startY-1][startX] == '|' ||
			pipes[startY-1][startX] == '7' ||
			pipes[startY-1][startX] == 'F' {
			if a[0] == -1 {
				a[0] = startX
				a[1] = startY - 1
			} else {
				b[0] = startX
				b[1] = startY - 1
			}
		}
	}
	if startY < len(pipes)-1 {
		if pipes[startY+1][startX] == '|' ||
			pipes[startY+1][startX] == 'J' ||
			pipes[startY+1][startX] == 'L' {
			if a[0] == -1 {
				a[0] = startX
				a[1] = startY + 1
			} else {
				b[0] = startX
				b[1] = startY + 1
			}
		}
	}
	if startX > 0 {
		if pipes[startY][startX-1] == '-' ||
			pipes[startY][startX-1] == 'L' ||
			pipes[startY][startX-1] == 'F' {
			if a[0] == -1 {
				a[0] = startX - 1
				a[1] = startY
			} else {
				b[0] = startX - 1
				b[1] = startY
			}
		}
	}
	if startX < len(pipes[startY])-1 {
		if pipes[startY][startX+1] == '-' ||
			pipes[startY][startX+1] == '7' ||
			pipes[startY][startX+1] == 'J' {
			if a[0] == -1 {
				a[0] = startX + 1
				a[1] = startY
			} else {
				b[0] = startX + 1
				b[1] = startY
			}
		}
	}

	pipes[startY][startX] = 'X'

	for {
		curA := pipes[a[1]][a[0]]
		curB := pipes[b[1]][b[0]]
		pipes[a[1]][a[0]] = 'X'
		pipes[b[1]][b[0]] = 'X'
		if a[0] == b[0] && a[1] == b[1] {
			break
		}
		a = getNextStep(pipes, curA, a)
		b = getNextStep(pipes, curB, b)
	}

	// Create map with only the main loop
	onlyLoop := util.ToRune(input)
	for i := 0; i < len(onlyLoop); i++ {
		for j := 0; j < len(onlyLoop[i]); j++ {
			if pipes[i][j] != 'X' {
				onlyLoop[i][j] = '.'
			}
		}
	}
	// Replace 'S' with required pipe piece.
	replaceStart(onlyLoop, startX, startY)

	// Now expand each tile into 3x3. That way we can easily flood fill.
	expandedMap := buildEpxandedMap(onlyLoop)

	// Now flood fill with '0'
	floodFill(expandedMap, 0, 0)

	area := 0
	for i := 1; i < len(expandedMap); i += 3 {
		for j := 1; j < len(expandedMap[i]); j += 3 {
			if expandedMap[i][j] == '.' {
				area++
			}
		}
	}

	return area
}
