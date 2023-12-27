package day23

type coordinate struct {
	X int
	Y int
}

type node struct {
	Coordinate coordinate
	Visited    bool
	Next       []edge
	End        bool
}

type edge struct {
	Destination *node
	Cost        int
}

func notIntersection(x, y int, input []string) bool {
	count := 0

	if x > 0 && input[y][x-1] != '#' {
		count++
	}
	if y > 0 && input[y-1][x] != '#' {
		count++
	}
	if x < len(input[0])-1 && input[y][x+1] != '#' {
		count++
	}
	if y < len(input)-1 && input[y+1][x] != '#' {
		count++
	}

	return count < 3
}

func getNext(x, y, prevX, prevY int, input []string) []coordinate {
	next := make([]coordinate, 0)

	if x > 0 {
		if input[y][x-1] != '#' && prevX != x-1 {
			next = append(next, coordinate{x - 1, y})
		}
	}
	if y > 0 {
		if input[y-1][x] != '#' && prevY != y-1 {
			next = append(next, coordinate{x, y - 1})
		}
	}
	if x < len(input[0])-1 {
		if input[y][x+1] != '#' && prevX != x+1 {
			next = append(next, coordinate{x + 1, y})
		}
	}
	if y < len(input)-1 {
		if input[y+1][x] != '#' && prevY != y+1 {
			next = append(next, coordinate{x, y + 1})
		}
	}

	return next
}

func edgeNotFound(start, end *node) bool {
	for _, e := range start.Next {
		if e.Destination == end {
			return false
		}
	}
	return true
}

func walkTrails(startNode *node, x, y, prevX, prevY int, input []string, nodes map[coordinate]*node, canSlip bool) {
	steps := 1
	for notIntersection(x, y, input) {
		if x == len(input[0])-2 && y == len(input)-1 { // We have reached the goal
			break
		}

		if canSlip {
			switch input[y][x] {
			case '>':
				if prevX >= x {
					return
				}
			case '<':
				if prevX <= x {
					return
				}
			case 'v':
				if prevY >= y {
					return
				}
			case '^':
				if prevY <= y {
					return
				}
			}
		}

		steps++
		next := getNext(x, y, prevX, prevY, input)
		prevX, prevY = x, y
		x, y = next[0].X, next[0].Y
	}

	var newNode *node
	if n, ok := nodes[coordinate{x, y}]; ok {
		newNode = n
	} else {
		newNode = &node{
			Coordinate: coordinate{X: x, Y: y},
			Next:       make([]edge, 0),
		}
	}

	e := edge{
		Destination: newNode,
		Cost:        steps,
	}
	if edgeNotFound(startNode, newNode) {
		startNode.Next = append(startNode.Next, e)
	}

	if !canSlip {
		e = edge{
			Destination: startNode,
			Cost:        steps,
		}
		if edgeNotFound(newNode, startNode) {
			newNode.Next = append(newNode.Next, e)
		}
	}

	nodes[coordinate{x, y}] = newNode
	if newNode.Visited {
		return
	}

	newNode.Visited = true
	if x == len(input[0])-2 && y == len(input)-1 {
		newNode.End = true
		return
	}
	for _, c := range getNext(x, y, prevX, prevY, input) {
		walkTrails(newNode, c.X, c.Y, x, y, input, nodes, canSlip)
	}
}

func buildGraph(input []string, canSlip bool) map[coordinate]*node {
	nodes := make(map[coordinate]*node)
	startNode := &node{
		Coordinate: coordinate{X: 1, Y: 0},
		Visited:    true,
		Next:       make([]edge, 0),
	}

	nodes[coordinate{1, 0}] = startNode

	walkTrails(startNode, 1, 1, 1, 0, input, nodes, canSlip)

	return nodes
}

func copyMap(in map[coordinate]bool) map[coordinate]bool {
	out := make(map[coordinate]bool)

	for k := range in {
		out[k] = true
	}
	return out
}

func longestPath(start *node, memo map[coordinate]bool) (int, bool) {
	if _, ok := memo[start.Coordinate]; ok {
		return 0, false
	}

	found := false

	if start.End {
		return 0, true
	}

	memo[start.Coordinate] = true
	steps := make([]int, 0)
	for _, e := range start.Next {
		s, endReached := longestPath(e.Destination, copyMap(memo))
		if endReached {
			found = true
			steps = append(steps, s+e.Cost)
		}
	}

	if len(steps) == 0 {
		return 0, false
	}

	m := steps[0]
	for _, s := range steps[1:] {
		m = max(m, s)
	}
	return m, found
}

func SolvePart1(input []string) int {
	nodes := buildGraph(input, true)

	memo := make(map[coordinate]bool)
	steps, _ := longestPath(nodes[coordinate{1, 0}], memo)

	return steps
}

func SolvePart2(input []string) int {
	nodes := buildGraph(input, false)

	memo := make(map[coordinate]bool)
	steps, _ := longestPath(nodes[coordinate{1, 0}], memo)

	return steps
}
