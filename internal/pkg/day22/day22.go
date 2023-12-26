package day22

import (
	"strconv"
	"strings"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

type block struct {
	Id   int
	Xmin int
	Xmax int
	Ymin int
	Ymax int
	Zmin int
	Zmax int
}

type graphNode struct {
	Id            int
	OnGround      bool
	Supports      *util.Set
	IsSupportedBy *util.Set
}

func newGraphNode(id int) graphNode {
	return graphNode{
		Id:            id,
		Supports:      util.NewSet(),
		IsSupportedBy: util.NewSet(),
	}
}

func parseInputLine(l string) (int, int, int, int, int, int) {
	left := strings.Split(l, "~")[0]
	right := strings.Split(l, "~")[1]

	xL, err := strconv.Atoi(strings.Split(left, ",")[0])
	if err != nil {
		panic(err)
	}
	yL, err := strconv.Atoi(strings.Split(left, ",")[1])
	if err != nil {
		panic(err)
	}
	zL, err := strconv.Atoi(strings.Split(left, ",")[2])
	if err != nil {
		panic(err)
	}
	xR, err := strconv.Atoi(strings.Split(right, ",")[0])
	if err != nil {
		panic(err)
	}
	yR, err := strconv.Atoi(strings.Split(right, ",")[1])
	if err != nil {
		panic(err)
	}
	zR, err := strconv.Atoi(strings.Split(right, ",")[2])
	if err != nil {
		panic(err)
	}

	return xL, yL, zL, xR, yR, zR
}

func parseInput(input []string) map[block]bool {
	blocks := make(map[block]bool)

	blockId := 0
	for _, l := range input {
		xL, yL, zL, xR, yR, zR := parseInputLine(l)

		b := block{
			Id:   blockId,
			Xmin: min(xL, xR),
			Xmax: max(xL, xR),
			Ymin: min(yL, yR),
			Ymax: max(yL, yR),
			Zmin: min(zL, zR),
			Zmax: max(zL, zR),
		}

		blocks[b] = true
		blockId++
	}

	return blocks
}

func canFall(blocks map[block]bool, b block) bool {
	for k := range blocks {
		if k.Zmax >= b.Zmin || k.Zmax < b.Zmin-1 {
			continue
		}
		if ((b.Xmin >= k.Xmin && b.Xmin <= k.Xmax) ||
			(b.Xmax >= k.Xmin && b.Xmax <= k.Xmax) ||
			(k.Xmin >= b.Xmin && k.Xmin <= b.Xmax) ||
			(k.Xmax >= b.Xmin && k.Xmax <= b.Xmax)) &&
			((b.Ymin >= k.Ymin && b.Ymin <= k.Ymax) ||
				(b.Ymax >= k.Ymin && b.Ymax <= k.Ymax) ||
				(k.Ymin >= b.Ymin && k.Ymin <= b.Ymax) ||
				(k.Ymax >= b.Ymin && k.Ymax <= b.Ymax)) {
			return false
		}
	}

	return true
}

func fallBlocks(blocks map[block]bool) {
	falling := true
	for falling {
		falling = false

		allBlocks := make([]block, len(blocks))
		i := 0
		for k := range blocks {
			allBlocks[i] = k
			i++
		}

		for _, b := range allBlocks {
			if b.Zmin == 1 {
				continue
			}
			delete(blocks, b)
			if canFall(blocks, b) {
				b.Zmin--
				b.Zmax--
				falling = true
			}
			blocks[b] = true
		}
	}

}

func addSupports(a int, b int, graph map[int]graphNode) {
	if _, ok := graph[b]; !ok {
		graph[b] = newGraphNode(b)
	}
	graph[b].Supports.Add(a)

	if _, ok := graph[a]; !ok {
		graph[a] = newGraphNode(a)
	}
	graph[a].IsSupportedBy.Add(b)
}

func buildBlockGraph(blocks map[block]bool) map[int]graphNode {
	allBlocks := make([]block, len(blocks))
	graph := make(map[int]graphNode)

	i := 0
	for k := range blocks {
		allBlocks[i] = k
		i++
	}

	for _, b := range allBlocks {
		for k := range blocks {

			if ((b.Xmin >= k.Xmin && b.Xmin <= k.Xmax) ||
				(b.Xmax >= k.Xmin && b.Xmax <= k.Xmax) ||
				(k.Xmin >= b.Xmin && k.Xmin <= b.Xmax) ||
				(k.Xmax >= b.Xmin && k.Xmax <= b.Xmax)) &&
				((b.Ymin >= k.Ymin && b.Ymin <= k.Ymax) ||
					(b.Ymax >= k.Ymin && b.Ymax <= k.Ymax) ||
					(k.Ymin >= b.Ymin && k.Ymin <= b.Ymax) ||
					(k.Ymax >= b.Ymin && k.Ymax <= b.Ymax)) {
				if b.Zmin == k.Zmax+1 {
					addSupports(b.Id, k.Id, graph)
				}
			}
		}
	}

	// Flag all blocks on the ground and add missing blocks
	for _, b := range allBlocks {
		if b.Zmin == 1 {
			if block, ok := graph[b.Id]; ok {
				block.OnGround = true
				graph[b.Id] = block
			} else {
				block := newGraphNode(b.Id)
				block.OnGround = true
				graph[b.Id] = block
			}
		}
	}

	return graph
}

func canBeRemoved(nodeCount int, graph map[int]graphNode) int {
	total := 0

	for nodeId := 0; nodeId < nodeCount; nodeId++ {
		needed := false

		for k, n := range graph {
			if k == nodeId {
				continue
			}
			if n.IsSupportedBy.Contains(nodeId) && n.IsSupportedBy.Size() == 1 {
				needed = true
				break
			}
		}

		if !needed {
			total++
		}
	}

	return total
}

func SolvePart1(input []string) int {
	blocks := parseInput(input)

	fallBlocks(blocks)

	graph := buildBlockGraph(blocks)

	return canBeRemoved(len(blocks), graph)
}

func getCountFall(graph map[int]graphNode, blockId int) int {
	queue := make([]int, 0)

	if _, ok := graph[blockId]; !ok {
		return 0
	}

	for _, n := range graph[blockId].Supports.GetAll() {
		graph[n.(int)].IsSupportedBy.Remove(blockId)
		queue = append(queue, n.(int))
	}

	for len(queue) > 0 {
		batchSize := len(queue)
		for i := 0; i < batchSize; i++ {
			curBlock := queue[0]
			queue = queue[1:]
			if graph[curBlock].IsSupportedBy.Size() == 0 {
				for _, n := range graph[curBlock].Supports.GetAll() {
					graph[n.(int)].IsSupportedBy.Remove(curBlock)
					queue = append(queue, n.(int))
				}
			}
		}
	}

	countFall := 0
	for k := range graph {
		if graph[k].IsSupportedBy.IsEmpty() && !graph[k].OnGround {
			countFall++
		}
	}
	return countFall
}

func SolvePart2(input []string) int {
	blocks := parseInput(input)

	fallBlocks(blocks)

	sum := 0
	for blockId := 0; blockId < len(blocks); blockId++ {
		graph := buildBlockGraph(blocks)
		sum += getCountFall(graph, blockId)
	}

	return sum
}
