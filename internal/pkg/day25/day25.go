package day25

import (
	"math/rand"
	"sort"
	"strings"
)

func parseInput(input []string) map[string][]string {
	graph := make(map[string][]string)

	for _, l := range input {
		node := strings.Split(l, ":")[0]
		connections := strings.Fields(strings.Split(l, ":")[1])

		if _, ok := graph[node]; ok {
			graph[node] = append(graph[node], connections...)
		} else {
			graph[node] = connections
		}

		for _, n := range connections {
			graph[n] = append(graph[n], node)
		}
	}

	return graph
}

type node struct {
	Node string
	Prev *node
}

func shortestPath(a, b string, graph map[string][]string) []string {
	q := make([]node, 0)

	q = append(q, node{Node: a})

	visited := make(map[node]bool)

	var goal *node

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if _, ok := visited[cur]; ok {
			continue
		}

		if cur.Node == b {
			goal = &cur
			break
		}

		visited[cur] = true

		for _, e := range graph[cur.Node] {
			q = append(q, node{Node: e, Prev: &cur})
		}
	}

	path := make([]string, 0)

	curNode := goal
	path = append(path, curNode.Node)
	for curNode.Node != a {
		curNode = curNode.Prev
		path = append(path, curNode.Node)
	}

	return path
}

func SolvePart1(input []string) int {
	graph := parseInput(input)

	edgeCount := make(map[string]int)

	for i := 0; i < 100; i++ {
		a := rand.Intn(len(graph))
		b := rand.Intn(len(graph))
		for b == a {
			b = rand.Intn(len(graph))
		}

		keys := make([]string, len(graph))
		i := 0
		for k := range graph {
			keys[i] = k
			i++
		}

		keyA := keys[a]
		keyB := keys[b]
		path := shortestPath(keyA, keyB, graph)
		for i := 0; i < len(path)-1; i++ {
			if _, ok := edgeCount[path[i]+"-"+path[i+1]]; ok {
				edgeCount[path[i]+"-"+path[i+1]]++
			} else if _, ok := edgeCount[path[i+1]+"-"+path[i]]; ok {
				edgeCount[path[i+1]+"-"+path[i]]++
			} else {
				edgeCount[path[i]+"-"+path[i+1]] = 1
			}
		}

	}

	type edge struct {
		Edge  string
		Count int
	}

	i := 0
	edges := make([]edge, len(edgeCount))
	for k, v := range edgeCount {
		edges[i] = edge{Edge: k, Count: v}
		i++
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Count > edges[j].Count
	})

	// The 3 most often used edges should be the ones we need to cut.
	for i := 0; i < 3; i++ {
		a := strings.Split(edges[i].Edge, "-")[0]
		b := strings.Split(edges[i].Edge, "-")[1]

		newEdges := make([]string, 0)
		for _, e := range graph[a] {
			if e != b {
				newEdges = append(newEdges, e)
			}
		}
		graph[a] = newEdges

		newEdges = make([]string, 0)
		for _, e := range graph[b] {
			if e != a {
				newEdges = append(newEdges, e)
			}
		}
		graph[b] = newEdges
	}

	// Now perform a BFS over the graph to walk all nodes
	visitedEdges := make([][]string, 0)
	for len(graph) > 1 {
		visitedThisRound := make([]string, 0)

		// Start from a random node
		a := rand.Intn(len(graph))

		keys := make([]string, len(graph))
		i := 0
		for k := range graph {
			keys[i] = k
			i++
		}

		keyA := keys[a]

		q := make([]string, 0)
		q = append(q, keyA)

		for len(q) > 0 {
			curLen := len(q)

			for i := 0; i < curLen; i++ {
				curNode := q[0]
				q = q[1:]

				if _, ok := graph[curNode]; !ok {
					continue
				}

				visitedThisRound = append(visitedThisRound, curNode)
				q = append(q, graph[curNode]...)

				delete(graph, curNode)
			}
		}

		visitedEdges = append(visitedEdges, visitedThisRound)
	}

	if len(visitedEdges) != 2 {
		return -1
	}

	return len(visitedEdges[0]) * len(visitedEdges[1])
}
