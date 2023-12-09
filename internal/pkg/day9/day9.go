package day9

import (
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func getNextInSequence(seq []int) int {
	t := make([][]int, 0)
	t = append(t, seq)

	i := 0
	for len(t[i]) > 1 {
		newSeq := make([]int, len(t[i])-1)
		for j := 0; j < len(newSeq); j++ {
			newSeq[j] = t[i][j+1] - t[i][j]
		}
		t = append(t, newSeq)

		i++
		if util.AllEquals(t[i], 0) {
			break
		}
	}

	for i := len(t) - 1; i >= 1; i-- {
		t[i-1] = append(t[i-1], t[i][len(t[i])-1]+t[i-1][len(t[i-1])-1])
	}

	return t[0][len(t[0])-1]
}

func getPrevInSequence(seq []int) int {
	t := make([][]int, 0)
	t = append(t, seq)

	i := 0
	for len(t[i]) > 1 {
		newSeq := make([]int, len(t[i])-1)
		for j := 0; j < len(newSeq); j++ {
			newSeq[j] = t[i][j+1] - t[i][j]
		}
		t = append(t, newSeq)

		i++
		if util.AllEquals(t[i], 0) {
			break
		}
	}

	prev := 0

	for i := len(t) - 1; i >= 0; i-- {
		prev = t[i][0] - prev
	}

	return prev
}

func SolvePart1(input []string) int {
	sumValues := 0

	for _, line := range input {
		seq := util.StringToIntSlice(line)
		next := getNextInSequence(seq)
		sumValues += next
	}

	return sumValues
}

func SolvePart2(input []string) int {
	sumValues := 0

	for _, line := range input {
		seq := util.StringToIntSlice(line)
		prev := getPrevInSequence(seq)
		sumValues += prev
	}

	return sumValues
}
