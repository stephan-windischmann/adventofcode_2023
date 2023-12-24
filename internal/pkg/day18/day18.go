package day18

import (
	"math"
	"strconv"
	"strings"
)

func calculateArea(coordinates [][]int) int {
	// Shoelace Algorithm
	sumA := 0
	sumB := 0
	for i := 0; i < len(coordinates)-1; i++ {
		sumA += coordinates[i][0] * coordinates[i+1][1]
		sumB += coordinates[i][1] * coordinates[i+1][0]
	}

	// Get length of perimeter
	perimeter := 0
	for i := 0; i < len(coordinates)-1; i++ {
		if coordinates[i][0] == coordinates[i+1][0] {
			perimeter += int(math.Abs(float64(coordinates[i+1][1] - coordinates[i][1])))
		} else {
			perimeter += int(math.Abs(float64(coordinates[i+1][0] - coordinates[i][0])))
		}
	}

	// Now add Pick's formula
	return int(0.5*math.Abs(float64(sumA-sumB))) + int(0.5*float64(perimeter)) + 1
}

func SolvePart1(input []string) int {
	coordinates := make([][]int, 1)
	coordinates[0] = []int{0, 0}

	for i, line := range input {
		direction := strings.Split(line, " ")[0]
		steps, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "R":
			coordinates = append(coordinates, []int{coordinates[i][0] + steps, coordinates[i][1]})
		case "L":
			coordinates = append(coordinates, []int{coordinates[i][0] - steps, coordinates[i][1]})
		case "U":
			coordinates = append(coordinates, []int{coordinates[i][0], coordinates[i][1] - steps})
		case "D":
			coordinates = append(coordinates, []int{coordinates[i][0], coordinates[i][1] + steps})
		}
	}

	return calculateArea(coordinates)
}

func SolvePart2(input []string) int {
	coordinates := make([][]int, 1)
	coordinates[0] = []int{0, 0}

	for i, line := range input {
		hex := strings.Split(line, " ")[2][1:8]

		steps64, err := strconv.ParseInt(hex[1:6], 16, 64)
		if err != nil {
			panic(err)
		}
		steps := int(steps64)

		var direction string
		switch hex[6] {
		case '0':
			direction = "R"
		case '1':
			direction = "D"
		case '2':
			direction = "L"
		case '3':
			direction = "U"
		}

		switch direction {
		case "R":
			coordinates = append(coordinates, []int{coordinates[i][0] + steps, coordinates[i][1]})
		case "L":
			coordinates = append(coordinates, []int{coordinates[i][0] - steps, coordinates[i][1]})
		case "U":
			coordinates = append(coordinates, []int{coordinates[i][0], coordinates[i][1] - steps})
		case "D":
			coordinates = append(coordinates, []int{coordinates[i][0], coordinates[i][1] + steps})
		}
	}

	return calculateArea(coordinates)
}
