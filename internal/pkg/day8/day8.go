package day8

import (
	"strings"

	"github.com/jdavid5815/extmath"
)

func SolvePart1(input []string) int {
	instructions := input[0]

	// Parse network into map
	m := make(map[string][]string, len(input)-2)
	for i := 2; i < len(input); i++ {
		cur := strings.TrimSpace(strings.Split(input[i], "=")[0])
		rhs := strings.TrimSpace(strings.Split(input[i], "=")[1])
		elements := strings.Split(rhs[1:len(rhs)-1], ", ")
		m[cur] = []string{elements[0], elements[1]}
	}

	steps := 0
	i := 0
	cur := "AAA"
	for {
		if cur == "ZZZ" {
			return steps
		}
		if instructions[i] == 'L' {
			cur = m[cur][0]
		} else {
			cur = m[cur][1]
		}
		steps++
		if i < len(instructions)-1 {
			i++
		} else {
			i = 0
		}
	}
}

func isAtEnd(positions map[string]string) bool {
	for _, v := range positions {
		if !strings.HasSuffix(v, "Z") {
			return false
		}
	}
	return true
}

func SolvePart2(input []string) int {
	instructions := input[0]

	// Parse network into map
	m := make(map[string][]string, len(input)-2)
	for i := 2; i < len(input); i++ {
		cur := strings.TrimSpace(strings.Split(input[i], "=")[0])
		rhs := strings.TrimSpace(strings.Split(input[i], "=")[1])
		elements := strings.Split(rhs[1:len(rhs)-1], ", ")
		m[cur] = []string{elements[0], elements[1]}
	}

	// Determine all starting positions
	positions := make([]string, 0)
	for k, _ := range m {
		if strings.HasSuffix(k, "A") {
			positions = append(positions, k)
		}
	}

	// Determine steps needed from each start to end
	allSteps := make([]int, 0)

	for _, start := range positions {
		steps := 0
		i := 0
		cur := start
		for {
			if strings.HasSuffix(cur, "Z") {
				allSteps = append(allSteps, steps)
				break
			}
			if instructions[i] == 'L' {
				cur = m[cur][0]
			} else {
				cur = m[cur][1]
			}
			steps++
			if i < len(instructions)-1 {
				i++
			} else {
				i = 0
			}
		}
	}

	// Now we need to find the LCM.
	lcm := uint(allSteps[0])
	for i := 1; i < len(allSteps); i++ {
		lcm = extmath.Lcm(lcm, uint(allSteps[i]))
	}

	return int(lcm)
}
