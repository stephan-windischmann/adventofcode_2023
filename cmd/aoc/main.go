package main

import (
	"fmt"
	"os"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day1"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day2"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day3"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day4"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day5"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func solveDay1() {
	input := util.LoadInput("../../input/day1/input")

	r1 := day1.SolvePart1(input)
	fmt.Println("Day 1 Part 1 solution: ", r1)

	r2 := day1.SolvePart2(input)
	fmt.Println("Day 1 Part 2 solution: ", r2)
}

func solveDay2() {
	input := util.LoadInput("../../input/day2/input")

	r1 := day2.SolvePart1(input)
	fmt.Println("Day 2 Part 1 solution: ", r1)

	r2 := day2.SolvePart2(input)
	fmt.Println("Day 2 Part 2 solution: ", r2)
}

func solveDay3() {
	input := util.LoadInput("../../input/day3/input")

	r1 := day3.SolvePart1(input)
	fmt.Println("Day 3 Part 1 solution: ", r1)

	r2 := day3.SolvePart2(input)
	fmt.Println("Day 3 Part 2 solution: ", r2)
}

func solveDay4() {
	input := util.LoadInput("../../input/day4/input")

	r1 := day4.SolvePart1(input)
	fmt.Println("Day 4 Part 1 solution: ", r1)

	r2 := day4.SolvePart2(input)
	fmt.Println("Day 4 Part 2 solution: ", r2)
}

func solveDay5() {
	input := util.LoadInput("../../input/day5/input")

	r1 := day5.SolvePart1(input)
	fmt.Println("Day 5 Part 1 solution: ", r1)

	r2 := day5.SolvePart2(input)
	fmt.Println("Day 5 Part 2 solution: ", r2)
}

func main() {
	switch os.Args[1] {
	case "1":
		solveDay1()
	case "2":
		solveDay2()
	case "3":
		solveDay3()
	case "4":
		solveDay4()
	case "5":
		solveDay5()
	}
}
