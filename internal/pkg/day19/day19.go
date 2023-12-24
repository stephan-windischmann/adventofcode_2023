package day19

import (
	"sort"
	"strconv"
	"strings"
)

type part struct {
	X int
	M int
	A int
	S int
}

func parseInput(input []string) (map[string]string, []part) {
	workflows := make(map[string]string)
	parts := make([]part, 0)

	i := 0
	for input[i] != "" {
		workflowName := strings.Split(input[i], "{")[0]
		workflowRules := strings.Split(input[i], "{")[1]
		workflows[workflowName] = workflowRules[:len(workflowRules)-1]

		i++
	}
	i++

	for i < len(input) {
		ratings := strings.Split(input[i][1:len(input[i])-1], ",")
		p := part{}
		for _, rating := range ratings {
			r := strings.Split(rating, "=")[0]
			v, err := strconv.Atoi(strings.Split(rating, "=")[1])
			if err != nil {
				panic(err)
			}
			switch r {
			case "x":
				p.X = v
			case "m":
				p.M = v
			case "a":
				p.A = v
			case "s":
				p.S = v
			}
		}
		parts = append(parts, p)

		i++
	}

	return workflows, parts
}

func runTest(test string, p part) bool {
	var op string
	if strings.Contains(test, ">") {
		op = ">"
	} else {
		op = "<"
	}

	rating := strings.Split(test, op)[0]
	limit, err := strconv.Atoi(strings.Split(test, op)[1])
	if err != nil {
		panic(err)
	}
	switch rating {
	case "x":
		if op == ">" {
			return p.X > limit
		} else {
			return p.X < limit
		}
	case "m":
		if op == ">" {
			return p.M > limit
		} else {
			return p.M < limit
		}
	case "a":
		if op == ">" {
			return p.A > limit
		} else {
			return p.A < limit
		}
	case "s":
		if op == ">" {
			return p.S > limit
		} else {
			return p.S < limit
		}
	}

	return false
}

func runWorkflow(workflow string, p part) string {
	rules := strings.Split(workflow, ",")

	for i := 0; i < len(rules)-1; i++ {
		test := strings.Split(rules[i], ":")[0]
		nextState := strings.Split(rules[i], ":")[1]

		if runTest(test, p) {
			return nextState
		}
	}

	return rules[len(rules)-1]
}

func sumAllParts(parts []part) int {
	sum := 0

	for _, p := range parts {
		sum += p.X + p.M + p.A + p.S
	}

	return sum
}

func SolvePart1(input []string) int {
	workflows, parts := parseInput(input)

	accepted := make([]part, 0)
	for _, p := range parts {
		nextState := runWorkflow(workflows["in"], p)
		for nextState != "A" && nextState != "R" {
			nextState = runWorkflow(workflows[nextState], p)
		}
		if nextState == "A" {
			accepted = append(accepted, p)
		}
	}

	return sumAllParts(accepted)
}

type TreeNode struct {
	state      string
	ranges     map[string][][]int
	nextStates []*TreeNode
}

func getNewRange(test string, ranges map[string][][]int) map[string][][]int {
	var op string
	if strings.Contains(test, "<") {
		op = "<"
	} else {
		op = ">"
	}

	rating := strings.Split(test, op)[0]
	limit, err := strconv.Atoi(strings.Split(test, op)[1])
	if err != nil {
		panic(err)
	}

	newRanges := make(map[string][][]int)

	for k := range ranges {
		sort.Slice(ranges[k], func(i, j int) bool {
			return ranges[k][i][0] < ranges[k][j][0]
		})
		newRanges[k] = make([][]int, 0)
		if k == rating {
			for _, r := range ranges[k] {
				if op == "<" {
					if r[1] < limit {
						newRanges[k] = append(newRanges[k], []int{r[0], r[1]})
					} else if r[0] < limit {
						newRanges[k] = append(newRanges[k], []int{r[0], limit - 1})
					}
				}
				if op == ">" {
					if r[0] > limit {
						newRanges[k] = append(newRanges[k], []int{r[0], r[1]})
					} else if r[1] > limit {
						newRanges[k] = append(newRanges[k], []int{limit + 1, r[1]})
					}
				}
			}
		} else {
			newRanges[k] = make([][]int, len(ranges[k]))
			for i, r := range ranges[k] {
				newRanges[k][i] = []int{r[0], r[1]}
			}
		}
	}

	return newRanges
}

func removeFromRange(test string, ranges map[string][][]int) map[string][][]int {
	var op string
	if strings.Contains(test, "<") {
		op = "<"
	} else {
		op = ">"
	}

	rating := strings.Split(test, op)[0]
	limit, err := strconv.Atoi(strings.Split(test, op)[1])
	if err != nil {
		panic(err)
	}

	newRanges := make(map[string][][]int)

	for k := range ranges {
		sort.Slice(ranges[k], func(i, j int) bool {
			return ranges[k][i][0] < ranges[k][j][0]
		})
		newRanges[k] = make([][]int, 0)
		if k == rating {
			for _, r := range ranges[k] {
				if op == ">" {
					if r[1] <= limit {
						newRanges[k] = append(newRanges[k], []int{r[0], r[1]})
					} else if r[0] <= limit {
						newRanges[k] = append(newRanges[k], []int{r[0], limit})
					}
				}
				if op == "<" {
					if r[0] > limit {
						newRanges[k] = append(newRanges[k], []int{r[0], r[1]})
					} else if r[1] > limit {
						newRanges[k] = append(newRanges[k], []int{limit, r[1]})
					}
				}
			}
		} else {
			newRanges[k] = make([][]int, len(ranges[k]))
			for i, r := range ranges[k] {
				newRanges[k][i] = []int{r[0], r[1]}
			}
		}
	}

	return newRanges
}

func makeTree(state string, ranges map[string][][]int, workflows map[string]string) *TreeNode {
	cur := &TreeNode{
		state:      state,
		ranges:     ranges,
		nextStates: make([]*TreeNode, 0),
	}

	if state == "A" {
		return cur
	}

	defaultRanges := make(map[string][][]int)
	for k := range ranges {
		defaultRanges[k] = make([][]int, len(ranges[k]))
		for i, r := range ranges[k] {
			defaultRanges[k][i] = []int{r[0], r[1]}
		}
	}

	rules := strings.Split(workflows[state], ",")

	for i := 0; i < len(rules)-1; i++ {
		test := strings.Split(rules[i], ":")[0]
		nextState := strings.Split(rules[i], ":")[1]

		newRange := getNewRange(test, defaultRanges)
		if nextState != "R" {
			newNode := makeTree(nextState, newRange, workflows)
			cur.nextStates = append(cur.nextStates, newNode)
		}
		defaultRanges = removeFromRange(test, defaultRanges)
	}

	lastState := rules[len(rules)-1]
	if lastState != "R" {
		newNode := makeTree(lastState, defaultRanges, workflows)
		cur.nextStates = append(cur.nextStates, newNode)
	}

	return cur
}

func getCombinations(treeNode *TreeNode) int {
	c := 0
	if len(treeNode.nextStates) == 0 && treeNode.state == "A" {
		sumX := 0
		sumM := 0
		sumA := 0
		sumS := 0

		for _, r := range treeNode.ranges["x"] {
			sumX += r[1] - (r[0] - 1)
		}
		for _, r := range treeNode.ranges["m"] {
			sumM += r[1] - (r[0] - 1)
		}
		for _, r := range treeNode.ranges["a"] {
			sumA += r[1] - (r[0] - 1)
		}
		for _, r := range treeNode.ranges["s"] {
			sumS += r[1] - (r[0] - 1)
		}
		return sumX * sumM * sumA * sumS
	}
	for _, nextNode := range treeNode.nextStates {
		c += getCombinations(nextNode)
	}
	return c
}

func SolvePart2(input []string) int {
	workflows, _ := parseInput(input)

	ranges := map[string][][]int{
		"x": {{1, 4000}},
		"m": {{1, 4000}},
		"a": {{1, 4000}},
		"s": {{1, 4000}},
	}

	treeNode := makeTree("in", ranges, workflows)

	return getCombinations(treeNode)
}
