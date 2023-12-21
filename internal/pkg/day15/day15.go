package day15

import "strings"

func calculateHash(in string) int {
	hash := 0

	for _, c := range []rune(in) {
		hash += int(c)
		hash *= 17
		hash %= 256
	}

	return hash
}

func SolvePart1(input []string) int {
	sumHash := 0

	for _, cmd := range strings.Split(input[0], ",") {
		sumHash += calculateHash(cmd)
	}

	return sumHash
}
