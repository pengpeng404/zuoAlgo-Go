package main

import (
	"fmt"
	"github.com/emirpasic/gods/queues/linkedlistqueue"
)

func main() {
	matrix := [][]int{
		{0, 2, 1},
		{0, 3, 1},
		{0, 4, 2},
		{0, 7, 2},
		{0, 4, 3},
		{0, 5, 3},
		{0, 8, 4},
		{0, 6, 5},
	}
	graph := generateGraph(matrix)
	fmt.Println(topologySort(graph))
}

func topologySort(graph *Graph) []int {
	if graph == nil {
		return nil
	}
	ans := make([]int, 0)
	inMap := make(map[*Node]int)
	zeroQueue := linkedlistqueue.New()
	for _, node := range graph.nodes {
		if node.in == 0 {
			zeroQueue.Enqueue(node)
		} else {
			inMap[node] = node.in
		}
	}
	for !zeroQueue.Empty() {
		cur, _ := zeroQueue.Dequeue()
		curNode := cur.(*Node)
		ans = append(ans, curNode.val)
		for i := range curNode.nexts {
			inMap[curNode.nexts[i]]--
			if inMap[curNode.nexts[i]] == 0 {
				zeroQueue.Enqueue(curNode.nexts[i])
			}
		}
	}
	return ans
}

func generateGraph(matrix [][]int) *Graph {
	graph := newGraph()
	for i := 0; i < len(matrix); i++ {
		weight := matrix[i][0]
		from := matrix[i][1]
		to := matrix[i][2]
		if _, exist := graph.nodes[from]; !exist {
			graph.nodes[from] = newNode(from)
		}
		if _, exist := graph.nodes[to]; !exist {
			graph.nodes[to] = newNode(to)
		}
		fromNode := graph.nodes[from]
		toNode := graph.nodes[to]
		fromNode.out++
		toNode.in++
		fromNode.nexts = append(fromNode.nexts, toNode)
		newedge := newEdge(weight, fromNode, toNode)
		fromNode.edges = append(fromNode.edges, newedge)
		graph.edges[newedge] = weight
	}
	return graph
}

type Node struct {
	val   int
	in    int
	out   int
	nexts []*Node
	edges []*Edge
}

func newNode(val int) *Node {
	return &Node{
		val: val,
	}
}

type Edge struct {
	weight int
	from   *Node
	to     *Node
}

func newEdge(weight int, from, to *Node) *Edge {
	return &Edge{
		weight: weight,
		from:   from,
		to:     to,
	}
}

type Graph struct {
	nodes map[int]*Node
	edges map[*Edge]int
}

func newGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
		edges: make(map[*Edge]int),
	}
}
