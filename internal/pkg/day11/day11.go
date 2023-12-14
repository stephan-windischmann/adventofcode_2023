package day11

import (
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func getDistance(
	galaxies [][]int, i, j int, rowNone *util.Set,
	colNone *util.Set, expansion int) int {
	minX := min(galaxies[i][0], galaxies[j][0])
	maxX := max(galaxies[i][0], galaxies[j][0])

	minY := min(galaxies[i][1], galaxies[j][1])
	maxY := max(galaxies[i][1], galaxies[j][1])

	distance := 0

	for i := minX; i < maxX; i++ {
		distance++
		if colNone.Contains(i) {
			distance += expansion
		}
	}

	for i := minY; i < maxY; i++ {
		distance++
		if rowNone.Contains(i) {
			distance += expansion
		}
	}

	return distance
}

func SolvePart1(input []string, expansion int) int {
	galaxies := make([][]int, 0)

	rowNoGalaxies := util.NewSet()
	colNoGalaxies := util.NewSet()

	for i, row := range input {
		noGalInRow := true

		for j := 0; j < len(row); j++ {
			if row[j] == '#' {
				noGalInRow = false
				galaxies = append(galaxies, []int{j, i})
			}
		}

		if noGalInRow {
			rowNoGalaxies.Add(i)
		}
	}

	for i := 0; i < len(input[0]); i++ {
		noGalInCol := true
		for j := 0; j < len(input); j++ {
			if input[j][i] == '#' {
				noGalInCol = false
				break
			}
		}
		if noGalInCol {
			colNoGalaxies.Add(i)
		}
	}

	sumDistances := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sumDistances += getDistance(
				galaxies, i, j, rowNoGalaxies, colNoGalaxies, expansion)
		}
	}

	return sumDistances
}
