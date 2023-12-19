package day12

import (
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"strconv"
	"strings"
)

type memoItem struct {
	s      string
	counts string
}

func getArrangements(s string, counts []int, memo map[memoItem]int) int {
	s = strings.TrimLeft(s, ".")

	m := memoItem{
		s:      s,
		counts: util.JoinIntList(counts, ","),
	}
	if c, ok := memo[m]; ok {
		return c
	}

	if len(s) == 0 {
		if len(counts) == 0 {
			memo[m] = 1
			return 1
		}
		memo[m] = 0
		return 0
	}

	if len(counts) == 0 {
		if strings.Count(s, "#") > 0 {
			memo[m] = 0
			return 0
		}
		memo[m] = 1
		return 1
	}

	if s[0] == '#' {
		if len(s) < counts[0] {
			memo[m] = 0
			return 0
		}
		if strings.Count(s[:counts[0]], ".") > 0 {
			memo[m] = 0
			return 0
		}
		if len(s) == counts[0] {
			if len(counts) > 1 {
				memo[m] = 0
				return 0
			}
			memo[m] = 1
			return 1
		}
		if len(s) > counts[0] {
			if s[counts[0]] == '#' {
				memo[m] = 0
				return 0
			}
			if len(s) > counts[0] {
				c := getArrangements(s[counts[0]+1:], counts[1:], memo)
				memo[m] = c
				return c
			}
			c := getArrangements("", counts[1:], memo)
			memo[m] = c
			return c
		}
	}

	c := getArrangements("."+s[1:], counts, memo) + getArrangements("#"+s[1:], counts, memo)
	memo[m] = c
	return c
}

func SolvePart1(input []string) int {
	count := 0

	for _, l := range input {
		s := strings.Split(l, " ")[0]
		c := strings.Split(l, " ")[1]
		counts := make([]int, 0)
		for _, n := range strings.Split(c, ",") {
			n, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			counts = append(counts, n)
		}

		memo := make(map[memoItem]int)
		count += getArrangements(s, counts, memo)
	}

	return count
}

func SolvePart2(input []string) int {
	count := 0

	for _, l := range input {
		s := strings.Split(l, " ")[0]
		c := strings.Split(l, " ")[1]
		tempList := make([]string, 5)
		for i := 0; i < 5; i++ {
			tempList[i] = s
		}
		s = strings.Join(tempList, "?")
		tempCounts := make([]int, 0)
		for _, n := range strings.Split(c, ",") {
			n, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			tempCounts = append(tempCounts, n)
		}
		counts := make([]int, 0)
		for i := 0; i < 5; i++ {
			counts = append(counts, tempCounts...)
		}

		memo := make(map[memoItem]int)
		count += getArrangements(s, counts, memo)
	}

	return count
}
