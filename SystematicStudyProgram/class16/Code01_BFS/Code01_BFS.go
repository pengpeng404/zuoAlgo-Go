package main

import (
	"fmt"
	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	matrix := [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 1, 7},
		{0, 2, 5},
		{0, 5, 6},
		{0, 6, 3},
		{0, 4, 3},
		{0, 3, 4},
	}
	graph := generateGraph(matrix)
	root := graph.nodes[1]
	bfs(root)
}

// bfs 宽度优先遍历
func bfs(root *Node) {
	if root == nil {
		return
	}
	set := hashset.New()
	queue := linkedlistqueue.New()
	queue.Enqueue(root)
	set.Add(root)
	for !queue.Empty() {
		cur, _ := queue.Dequeue()
		curNode := cur.(*Node)
		fmt.Print(curNode.val)
		nexts := curNode.nexts
		for i := range nexts {
			if !set.Contains(nexts[i]) {
				set.Add(nexts[i])
				queue.Enqueue(nexts[i])
			}
		}
	}
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
