package day15

import (
	"strconv"
	"strings"
)

type mirror struct {
	Label       string
	FocalLength int
}

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

func removeFromBox(label string, boxes [][]mirror) {
	hash := calculateHash(label)

	index := -1
	for i := 0; i < len(boxes[hash]); i++ {
		if boxes[hash][i].Label == label {
			index = i
			break
		}
	}

	if index > -1 {
		temp := boxes[hash]
		boxes[hash] = make([]mirror, 0)
		boxes[hash] = append(boxes[hash], temp[:index]...)
		boxes[hash] = append(boxes[hash], temp[index+1:]...)
	}
}

func insertIntoBox(label string, focalLength int, boxes [][]mirror) {
	hash := calculateHash(label)

	for i := 0; i < len(boxes[hash]); i++ {
		if boxes[hash][i].Label == label {
			boxes[hash][i].FocalLength = focalLength
			return
		}
	}

	boxes[hash] = append(boxes[hash], mirror{Label: label, FocalLength: focalLength})
}

func SolvePart2(input []string) int {
	boxes := make([][]mirror, 256)
	for i := 0; i < len(boxes); i++ {
		boxes[i] = make([]mirror, 0)
	}

	for _, cmd := range strings.Split(input[0], ",") {
		if strings.Contains(cmd, "-") {
			label := strings.Split(cmd, "-")[0]
			removeFromBox(label, boxes)
		} else {
			label := strings.Split(cmd, "=")[0]
			focalLength, err := strconv.Atoi(strings.Split(cmd, "=")[1])
			if err != nil {
				panic(err)
			}
			insertIntoBox(label, focalLength, boxes)
		}
	}

	focusingPower := 0

	for i := 0; i < 256; i++ {
		if len(boxes[i]) > 0 {
			boxNumber := i + 1

			for j := 0; j < len(boxes[i]); j++ {
				focusingPower += boxNumber * (j + 1) * boxes[i][j].FocalLength
			}
		}

	}

	return focusingPower
}
