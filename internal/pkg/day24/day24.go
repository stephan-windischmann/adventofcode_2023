package day24

import (
	"strconv"
	"strings"
)

type coordinate struct {
	X int
	Y int
	Z int
}

type hail struct {
	Point    coordinate
	Velocity coordinate
}

func parseInput(input []string) []hail {
	hailstones := make([]hail, len(input))

	for i := 0; i < len(input); i++ {
		point := strings.Replace(
			strings.TrimSpace(strings.Split(input[i], "@")[0]),
			",", " ", -1,
		)
		velocity := strings.Replace(
			strings.TrimSpace(strings.Split(input[i], "@")[1]),
			",", " ", -1,
		)

		h := hail{}

		var err error
		h.Point.X, err = strconv.Atoi(strings.Fields(point)[0])
		if err != nil {
			panic(err)
		}
		h.Point.Y, err = strconv.Atoi(strings.Fields(point)[1])
		if err != nil {
			panic(err)
		}
		h.Point.Z, err = strconv.Atoi(strings.Fields(point)[2])
		if err != nil {
			panic(err)
		}

		h.Velocity.X, err = strconv.Atoi(strings.Fields(velocity)[0])
		if err != nil {
			panic(err)
		}
		h.Velocity.Y, err = strconv.Atoi(strings.Fields(velocity)[1])
		if err != nil {
			panic(err)
		}
		h.Velocity.Z, err = strconv.Atoi(strings.Fields(velocity)[2])
		if err != nil {
			panic(err)
		}

		hailstones[i] = h
	}

	return hailstones
}

func intersectionInArea(h1, h2 hail, min, max int) bool {
	// High school math

	// First put lines into the form y=a*x+b
	slope_l1 := float64(h1.Velocity.Y) / float64(h1.Velocity.X)
	intercept_l1 := float64(h1.Point.Y) - slope_l1*float64(h1.Point.X)
	slope_l2 := float64(h2.Velocity.Y) / float64(h2.Velocity.X)
	intercept_l2 := float64(h2.Point.Y) - slope_l2*float64(h2.Point.X)

	if slope_l1 == slope_l2 {
		return false
	}

	// Now solve the equation for  y = a*x + b and y = c*x + d
	// a*x + b = c*x + d
	// a*x - c*x = d - b
	// x*(a - c) = d - b
	// x = (d - b) / (a - c)

	x := float64(intercept_l2-intercept_l1) / (slope_l1 - slope_l2)
	y := slope_l1*x + intercept_l1

	// Check if intersection was in past
	if (x-float64(h1.Point.X))/float64(h1.Velocity.X) < 0 ||
		(x-float64(h2.Point.X))/float64(h2.Velocity.X) < 0 {
		return false
	}

	if (x >= float64(min) && x <= float64(max)) &&
		(y >= float64(min) && y <= float64(max)) {
		return true
	}

	return false
}

func SolvePart1(input []string, min, max int) int {
	hailstones := parseInput(input)

	intersections := 0

	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if intersectionInArea(hailstones[i], hailstones[j], min, max) {
				intersections++
			}
		}
	}

	return intersections
}
