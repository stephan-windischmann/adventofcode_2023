package day5

import (
	"slices"
	"strconv"
	"strings"
)

func parseInput(input []string) ([]int, [][]int) {
	seeds := make([]int, 0)
	for _, seed := range strings.Fields(strings.Split(input[0], ":")[1]) {
		s, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, s)
	}

	maps := make([][]int, 0)
	for i := 2; i < len(input); i++ {
		i++
		m := make([]int, 0)
		for i < len(input) && len(input[i]) > 0 {
			values := strings.Fields(input[i])
			dest, err := strconv.Atoi(values[0])
			if err != nil {
				panic(err)
			}
			src, err := strconv.Atoi(values[1])
			if err != nil {
				panic(err)
			}
			rangeLen, err := strconv.Atoi(values[2])
			if err != nil {
				panic(err)
			}
			m = append(m, dest, src, rangeLen)
			i++
		}
		maps = append(maps, m)
	}

	return seeds, maps
}

func mapLookup(seed int, m []int) int {
	i := 0
	for i < len(m)-2 {
		if seed >= m[i+1] && seed <= m[i+1]+m[i+2]-1 {
			return m[i] + (seed - m[i+1])
		}
		i += 3
	}
	return seed
}

func lookup(seed int, maps [][]int) int {
	for _, m := range maps {
		seed = mapLookup(seed, m)
	}

	return seed
}

func SolvePart1(input []string) int {
	seeds, maps := parseInput(input)

	locations := make([]int, 0)
	for _, seed := range seeds {
		seed = lookup(seed, maps)

		locations = append(locations, seed)
	}

	return slices.Min(locations)
}

func SolvePart2(input []string) int {
	seeds, maps := parseInput(input)

	locations := make([]int, 0)

	for i := 0; i < len(seeds)-1; i += 2 {
		l := make([]int, 0)
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			loc := lookup(seed, maps)
			l = append(l, loc)
		}

		locations = append(locations, slices.Min(l))
	}
	return slices.Min(locations)
}
