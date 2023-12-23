package day17

import (
	"sort"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

type coordinate struct {
	X         int
	Y         int
	Count     int
	Direction rune
}

type block struct {
	Coordinate coordinate
	HeatLoss   int
}

func getNextMovesPart1(cur block, input []string) []block {
	possibleMoves := make([]block, 0)
	if cur.Coordinate.Direction == 'D' {
		if cur.Coordinate.Count < 2 && cur.Coordinate.Y < len(input)-1 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y + 1,
					Count:     cur.Coordinate.Count + 1,
					Direction: 'D',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y+1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.X > 0 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X - 1,
					Y:         cur.Coordinate.Y,
					Count:     0,
					Direction: 'L',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X-1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.X < len(input[0])-1 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X + 1,
					Y:         cur.Coordinate.Y,
					Count:     0,
					Direction: 'R',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X+1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
	}
	if cur.Coordinate.Direction == 'U' {
		if cur.Coordinate.Count < 2 && cur.Coordinate.Y > 0 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y - 1,
					Count:     cur.Coordinate.Count + 1,
					Direction: 'U',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y-1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.X > 0 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X - 1,
					Y:         cur.Coordinate.Y,
					Count:     0,
					Direction: 'L',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X-1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.X < len(input[0])-1 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X + 1,
					Y:         cur.Coordinate.Y,
					Count:     0,
					Direction: 'R',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X+1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
	}
	if cur.Coordinate.Direction == 'L' {
		if cur.Coordinate.Count < 2 && cur.Coordinate.X > 0 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X - 1,
					Y:         cur.Coordinate.Y,
					Count:     cur.Coordinate.Count + 1,
					Direction: 'L',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X-1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.Y > 0 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y - 1,
					Count:     0,
					Direction: 'U',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y-1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.Y < len(input)-1 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y + 1,
					Count:     0,
					Direction: 'D',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y+1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
	}
	if cur.Coordinate.Direction == 'R' {
		if cur.Coordinate.Count < 2 && cur.Coordinate.X < len(input[0])-1 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X + 1,
					Y:         cur.Coordinate.Y,
					Count:     cur.Coordinate.Count + 1,
					Direction: 'R',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X+1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.Y > 0 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y - 1,
					Count:     0,
					Direction: 'U',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y-1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.Y < len(input)-1 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y + 1,
					Count:     0,
					Direction: 'D',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y+1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
	}
	return possibleMoves
}

func getNextMovesPart2(cur block, input []string) []block {
	possibleMoves := make([]block, 0)
	if cur.Coordinate.Direction == 'D' {
		if cur.Coordinate.Count < 9 && cur.Coordinate.Y < len(input)-1 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y + 1,
					Count:     cur.Coordinate.Count + 1,
					Direction: 'D',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y+1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.X > 0 && cur.Coordinate.Count > 2 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X - 1,
					Y:         cur.Coordinate.Y,
					Count:     0,
					Direction: 'L',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X-1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.X < len(input[0])-1 && cur.Coordinate.Count > 2 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X + 1,
					Y:         cur.Coordinate.Y,
					Count:     0,
					Direction: 'R',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X+1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
	}
	if cur.Coordinate.Direction == 'U' {
		if cur.Coordinate.Count < 9 && cur.Coordinate.Y > 0 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y - 1,
					Count:     cur.Coordinate.Count + 1,
					Direction: 'U',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y-1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.X > 0 && cur.Coordinate.Count > 2 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X - 1,
					Y:         cur.Coordinate.Y,
					Count:     0,
					Direction: 'L',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X-1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.X < len(input[0])-1 && cur.Coordinate.Count > 2 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X + 1,
					Y:         cur.Coordinate.Y,
					Count:     0,
					Direction: 'R',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X+1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
	}
	if cur.Coordinate.Direction == 'L' {
		if cur.Coordinate.Count < 9 && cur.Coordinate.X > 0 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X - 1,
					Y:         cur.Coordinate.Y,
					Count:     cur.Coordinate.Count + 1,
					Direction: 'L',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X-1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.Y > 0 && cur.Coordinate.Count > 2 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y - 1,
					Count:     0,
					Direction: 'U',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y-1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.Y < len(input)-1 && cur.Coordinate.Count > 2 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y + 1,
					Count:     0,
					Direction: 'D',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y+1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
	}
	if cur.Coordinate.Direction == 'R' {
		if cur.Coordinate.Count < 9 && cur.Coordinate.X < len(input[0])-1 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X + 1,
					Y:         cur.Coordinate.Y,
					Count:     cur.Coordinate.Count + 1,
					Direction: 'R',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y][cur.Coordinate.X+1])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.Y > 0 && cur.Coordinate.Count > 2 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y - 1,
					Count:     0,
					Direction: 'U',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y-1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)
		}
		if cur.Coordinate.Y < len(input)-1 && cur.Coordinate.Count > 2 {
			next := block{
				Coordinate: coordinate{
					X:         cur.Coordinate.X,
					Y:         cur.Coordinate.Y + 1,
					Count:     0,
					Direction: 'D',
				},
				HeatLoss: cur.HeatLoss + int((input[cur.Coordinate.Y+1][cur.Coordinate.X])-'0'),
			}
			possibleMoves = append(possibleMoves, next)

		}
	}
	return possibleMoves
}

func runDijkstra(input []string, nextFunc func(block, []string) []block) int {
	return 0
}

func SolvePart1(input []string) int {
	visited := util.NewSet()
	queue := make([]block, 0)

	start1 := block{
		Coordinate: coordinate{
			X:         1,
			Y:         0,
			Count:     0,
			Direction: 'R',
		},
		HeatLoss: int((input[0][1]) - '0'),
	}
	start2 := block{
		Coordinate: coordinate{
			X:         0,
			Y:         1,
			Count:     0,
			Direction: 'D',
		},
		HeatLoss: int((input[1][0]) - '0'),
	}

	queue = append(queue, start1, start2)
	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].HeatLoss < queue[j].HeatLoss
		})

		cur := queue[0]
		queue = queue[1:]

		if cur.Coordinate.X == len(input[0])-1 && cur.Coordinate.Y == len(input)-1 {
			return cur.HeatLoss
		}

		if visited.Contains(cur.Coordinate) {
			continue
		}

		visited.Add(cur.Coordinate)

		queue = append(queue, getNextMovesPart1(cur, input)...)
	}

	return -1
}

func SolvePart2(input []string) int {
	visited := util.NewSet()
	queue := make([]block, 0)

	start1 := block{
		Coordinate: coordinate{
			X:         1,
			Y:         0,
			Count:     0,
			Direction: 'R',
		},
		HeatLoss: int((input[0][1]) - '0'),
	}
	start2 := block{
		Coordinate: coordinate{
			X:         0,
			Y:         1,
			Count:     0,
			Direction: 'D',
		},
		HeatLoss: int((input[1][0]) - '0'),
	}

	queue = append(queue, start1, start2)
	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].HeatLoss < queue[j].HeatLoss
		})

		cur := queue[0]
		queue = queue[1:]

		if cur.Coordinate.X == len(input[0])-1 && cur.Coordinate.Y == len(input)-1 && cur.Coordinate.Count > 2 {
			return cur.HeatLoss
		}

		if visited.Contains(cur.Coordinate) {
			continue
		}

		visited.Add(cur.Coordinate)

		queue = append(queue, getNextMovesPart2(cur, input)...)
	}

	return -1
}
