package day16

import (
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

type position struct {
	X int
	Y int
}

type visited struct {
	Position  position
	Direction rune
}

func shootBeam(x, y int, direction rune, tiles [][]rune, v map[visited]bool) {
	if x < 0 || y < 0 || x == len(tiles[0]) || y == len(tiles) {
		return
	}

	curVisited := visited{
		Position:  position{X: x, Y: y},
		Direction: direction,
	}
	if _, ok := v[curVisited]; ok {
		return
	}
	v[curVisited] = true

	if tiles[y][x] == '|' && (direction == 'R' || direction == 'L') {
		shootBeam(x, y-1, 'U', tiles, v)
		shootBeam(x, y+1, 'D', tiles, v)
		return
	}
	if tiles[y][x] == '-' && (direction == 'U' || direction == 'D') {
		shootBeam(x-1, y, 'L', tiles, v)
		shootBeam(x+1, y, 'R', tiles, v)
		return
	}
	if tiles[y][x] == '\\' {
		switch direction {
		case 'R':
			direction = 'D'
		case 'L':
			direction = 'U'
		case 'D':
			direction = 'R'
		case 'U':
			direction = 'L'
		}
	}
	if tiles[y][x] == '/' {
		switch direction {
		case 'R':
			direction = 'U'
		case 'L':
			direction = 'D'
		case 'D':
			direction = 'L'
		case 'U':
			direction = 'R'
		}
	}

	switch direction {
	case 'R':
		shootBeam(x+1, y, direction, tiles, v)
	case 'L':
		shootBeam(x-1, y, direction, tiles, v)
	case 'U':
		shootBeam(x, y-1, direction, tiles, v)
	case 'D':
		shootBeam(x, y+1, direction, tiles, v)
	}
}

func SolvePart1(input []string) int {
	tiles := util.CreateRuneTable(input)

	v := make(map[visited]bool)

	shootBeam(0, 0, 'R', tiles, v)

	s := util.NewSet()
	for k, _ := range v {
		s.Add(k.Position)
	}

	return s.Size()
}

func SolvePart2(input []string) int {
	maxEnergized := 0

	tiles := util.CreateRuneTable(input)
	// Left edge
	for i := 0; i < len(tiles); i++ {
		v := make(map[visited]bool)

		shootBeam(0, i, 'R', tiles, v)

		s := util.NewSet()
		for k, _ := range v {
			s.Add(k.Position)
		}

		maxEnergized = max(maxEnergized, s.Size())
	}

	// Right edge
	for i := 0; i < len(tiles); i++ {
		v := make(map[visited]bool)

		shootBeam(len(tiles[0])-1, i, 'L', tiles, v)

		s := util.NewSet()
		for k, _ := range v {
			s.Add(k.Position)
		}

		maxEnergized = max(maxEnergized, s.Size())
	}

	// Top edge
	for i := 0; i < len(tiles[0]); i++ {
		v := make(map[visited]bool)

		shootBeam(i, 0, 'D', tiles, v)

		s := util.NewSet()
		for k, _ := range v {
			s.Add(k.Position)
		}

		maxEnergized = max(maxEnergized, s.Size())
	}

	// Bottom edge
	for i := 0; i < len(tiles[0]); i++ {
		v := make(map[visited]bool)

		shootBeam(i, len(tiles)-1, 'U', tiles, v)

		s := util.NewSet()
		for k, _ := range v {
			s.Add(k.Position)
		}

		maxEnergized = max(maxEnergized, s.Size())
	}

	return maxEnergized
}
