package day6

import (
	"strconv"
	"strings"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func processRace(race []int) int {
	beats := 0

	for i := 0; i <= race[0]; i++ {
		if i*(race[0]-i) > race[1] {
			beats++
		}
	}

	return beats
}

func SolvePart1(input []string) int {
	times := strings.Fields(strings.Split(input[0], ":")[1])
	distances := strings.Fields(strings.Split(input[1], ":")[1])

	races := make([][]int, len(times))

	for i := 0; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			panic(err)
		}
		dist, err := strconv.Atoi(distances[i])
		if err != nil {
			panic(err)
		}
		races[i] = []int{time, dist}
	}

	beats := make([]int, 0, len(races))
	for _, race := range races {
		beat := processRace(race)

		beats = append(beats, beat)
	}

	return util.ProdInt(beats)
}

func SolvePart2(input []string) int {
	time, err := strconv.Atoi(strings.Join(strings.Fields(strings.Split(input[0], ":")[1]), ""))
	if err != nil {
		panic(err)
	}
	distance, err := strconv.Atoi(strings.Join(strings.Fields(strings.Split(input[1], ":")[1]), ""))
	if err != nil {
		panic(err)
	}

	return processRace([]int{time, distance})
}
