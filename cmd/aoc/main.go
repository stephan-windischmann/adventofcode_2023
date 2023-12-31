package main

import (
	"fmt"
	"os"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day18"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day19"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day20"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day21"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day22"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day23"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day24"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day25"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day13"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day14"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day15"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day16"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day17"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day1"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day10"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day11"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day12"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day2"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day3"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day4"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day5"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day6"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day7"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day8"
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/day9"
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

func solveDay6() {
	input := util.LoadInput("../../input/day6/input")

	r1 := day6.SolvePart1(input)
	fmt.Println("Day 6 Part 1 solution: ", r1)

	r2 := day6.SolvePart2(input)
	fmt.Println("Day 6 Part 2 solution: ", r2)
}

func solveDay7() {
	input := util.LoadInput("../../input/day7/input")

	r1 := day7.SolvePart1(input)
	fmt.Println("Day 7 Part 1 solution: ", r1)

	r2 := day7.SolvePart2(input)
	fmt.Println("Day 7 Part 2 solution: ", r2)
}

func solveDay8() {
	input := util.LoadInput("../../input/day8/input")

	r1 := day8.SolvePart1(input)
	fmt.Println("Day 8 Part 1 solution: ", r1)

	r2 := day8.SolvePart2(input)
	fmt.Println("Day 8 Part 2 solution: ", r2)
}

func solveDay9() {
	input := util.LoadInput("../../input/day9/input")

	r1 := day9.SolvePart1(input)
	fmt.Println("Day 9 Part 1 solution: ", r1)

	r2 := day9.SolvePart2(input)
	fmt.Println("Day 9 Part 2 solution: ", r2)
}

func solveDay10() {
	input := util.LoadInput("../../input/day10/input")
	r1 := day10.SolvePart1(input)
	fmt.Println("Day 10 Part 1 solution: ", r1)

	input = util.LoadInput("../../input/day10/input")
	r2 := day10.SolvePart2(input)
	fmt.Println("Day 10 Part 2 solution: ", r2)
}

func solveDay11() {
	input := util.LoadInput("../../input/day11/input")

	r1 := day11.SolvePart1(input, 1)
	fmt.Println("Day 11 Part 1 solution: ", r1)

	r2 := day11.SolvePart1(input, 999999)
	fmt.Println("Day 11 Part 2 solution: ", r2)
}

func solveDay12() {
	input := util.LoadInput("../../input/day12/input")

	r1 := day12.SolvePart1(input)
	fmt.Println("Day 12 Part 1 solution: ", r1)

	r2 := day12.SolvePart2(input)
	fmt.Println("Day 12 Part 2 solution: ", r2)
}

func solveDay13() {
	input := util.LoadInput("../../input/day13/input")

	r1 := day13.SolvePart1(input)
	fmt.Println("Day 13 Part 1 solution: ", r1)

	r2 := day13.SolvePart2(input)
	fmt.Println("Day 13 Part 2 solution: ", r2)
}

func solveDay14() {
	input := util.LoadInput("../../input/day14/input")

	r1 := day14.SolvePart1(input)
	fmt.Println("Day 14 Part 1 solution: ", r1)

	r2 := day14.SolvePart2(input)
	fmt.Println("Day 14 Part 2 solution: ", r2)
}

func solveDay15() {
	input := util.LoadInput("../../input/day15/input")

	r1 := day15.SolvePart1(input)
	fmt.Println("Day 15 Part 1 solution: ", r1)

	r2 := day15.SolvePart2(input)
	fmt.Println("Day 15 Part 2 solution: ", r2)
}

func solveDay16() {
	input := util.LoadInput("../../input/day16/input")

	r1 := day16.SolvePart1(input)
	fmt.Println("Day 16 Part 1 solution: ", r1)

	r2 := day16.SolvePart2(input)
	fmt.Println("Day 16 Part 2 solution: ", r2)
}

func solveDay17() {
	input := util.LoadInput("../../input/day17/input")

	r1 := day17.SolvePart1(input)
	fmt.Println("Day 17 Part 1 solution: ", r1)

	r2 := day17.SolvePart2(input)
	fmt.Println("Day 17 Part 2 solution: ", r2)
}

func solveDay18() {
	input := util.LoadInput("../../input/day18/input")

	r1 := day18.SolvePart1(input)
	fmt.Println("Day 18 Part 1 solution: ", r1)

	r2 := day18.SolvePart2(input)
	fmt.Println("Day 18 Part 2 solution: ", r2)
}

func solveDay19() {
	input := util.LoadInput("../../input/day19/input")

	r1 := day19.SolvePart1(input)
	fmt.Println("Day 19 Part 1 solution: ", r1)

	r2 := day19.SolvePart2(input)
	fmt.Println("Day 19 Part 2 solution: ", r2)
}

func solveDay20() {
	input := util.LoadInput("../../input/day20/input")

	r1 := day20.SolvePart1(input)
	fmt.Println("Day 20 Part 1 solution: ", r1)

	r2 := day20.SolvePart2(input)
	fmt.Println("Day 20 Part 2 solution: ", r2)
}

func solveDay21() {
	input := util.LoadInput("../../input/day21/input")

	r1 := day21.SolvePart1(input, 64)
	fmt.Println("Day 21 Part 1 solution: ", r1)

	r2 := day21.SolvePart2(input, 26501365)
	fmt.Println("Day 21 Part 2 solution: ", r2)
}

func solveDay22() {
	input := util.LoadInput("../../input/day22/input")

	r1 := day22.SolvePart1(input)
	fmt.Println("Day 22 Part 1 solution: ", r1)

	r2 := day22.SolvePart2(input)
	fmt.Println("Day 22 Part 2 solution: ", r2)
}

func solveDay23() {
	input := util.LoadInput("../../input/day23/input")

	r1 := day23.SolvePart1(input)
	fmt.Println("Day 23 Part 1 solution: ", r1)

	r2 := day23.SolvePart2(input)
	fmt.Println("Day 23 Part 2 solution: ", r2)
}

func solveDay24() {
	input := util.LoadInput("../../input/day24/input")

	r1 := day24.SolvePart1(input, 200000000000000, 400000000000000)
	fmt.Println("Day 24 Part 1 solution: ", r1)
}

func solveDay25() {
	input := util.LoadInput("../../input/day25/input")

	r1 := day25.SolvePart1(input)
	fmt.Println("Day 24 Part 1 solution: ", r1)
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
	case "6":
		solveDay6()
	case "7":
		solveDay7()
	case "8":
		solveDay8()
	case "9":
		solveDay9()
	case "10":
		solveDay10()
	case "11":
		solveDay11()
	case "12":
		solveDay12()
	case "13":
		solveDay13()
	case "14":
		solveDay14()
	case "15":
		solveDay15()
	case "16":
		solveDay16()
	case "17":
		solveDay17()
	case "18":
		solveDay18()
	case "19":
		solveDay19()
	case "20":
		solveDay20()
	case "21":
		solveDay21()
	case "22":
		solveDay22()
	case "23":
		solveDay23()
	case "24":
		solveDay24()
	case "25":
		solveDay25()
	}
}
