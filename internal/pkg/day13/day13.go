package day13

import (
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

type pattern struct {
	pattern [][]rune
}

func (p *pattern) isVerticalReflection(col, width, diffWanted int) bool {
	numDiff := 0
	for i := width; i > 0; i-- {
		for j := 0; j < len(p.pattern); j++ {
			if p.pattern[j][col-i] != p.pattern[j][col+i-1] {
				numDiff++
			}
		}
	}
	return numDiff == diffWanted
}

func (p *pattern) isHorizontalReflection(row, width, diffWanted int) bool {
	numDiff := 0
	for i := width; i > 0; i-- {
		for j := 0; j < len(p.pattern[0]); j++ {
			if p.pattern[row-i][j] != p.pattern[row+i-1][j] {
				numDiff++
			}
		}
	}
	return numDiff == diffWanted
}

func (p *pattern) VerticalReflection() int {
	for i := 1; i < len(p.pattern[0]); i++ {
		width := min(i, len(p.pattern[0])-i)
		if p.isVerticalReflection(i, width, 0) {
			return i
		}
	}
	return 0
}

func (p *pattern) HorizontalReflection() int {
	for i := 1; i < len(p.pattern); i++ {
		width := min(i, len(p.pattern)-i)
		if p.isHorizontalReflection(i, width, 0) {
			return i
		}
	}
	return 0
}

func (p *pattern) VerticalReflectionSmudge() int {
	for i := 1; i < len(p.pattern[0]); i++ {
		width := min(i, len(p.pattern[0])-i)
		if p.isVerticalReflection(i, width, 1) {
			return i
		}
	}
	return 0
}

func (p *pattern) HorizontalReflectionSmudge() int {
	for i := 1; i < len(p.pattern); i++ {
		width := min(i, len(p.pattern)-i)
		if p.isHorizontalReflection(i, width, 1) {
			return i
		}
	}
	return 0
}

func parseInput(input []string) []pattern {
	patterns := make([]pattern, 0)
	lines := make([]string, 0)

	for _, line := range input {
		if len(line) > 0 {
			lines = append(lines, line)
		} else {
			p := pattern{
				pattern: util.CreateRuneTable(lines),
			}
			patterns = append(patterns, p)
			lines = make([]string, 0)
		}
	}
	p := pattern{
		pattern: util.CreateRuneTable(lines),
	}
	patterns = append(patterns, p)

	return patterns
}

func SolvePart1(input []string) int {
	patterns := parseInput(input)

	total := 0
	for _, p := range patterns {
		i := p.VerticalReflection()
		if i > 0 {
			total += i
		} else {
			i := p.HorizontalReflection()
			total += 100 * i
		}
	}

	return total
}

func SolvePart2(input []string) int {
	patterns := parseInput(input)

	total := 0
	for _, p := range patterns {
		i := p.VerticalReflectionSmudge()
		if i > 0 {
			total += i
		} else {
			i := p.HorizontalReflectionSmudge()
			total += 100 * i
		}
	}

	return total
}
