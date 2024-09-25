package main

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/linkedliststack"
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
	dfs(root)
}

// dfs 深度优先遍历
func dfs(root *Node) {
	if root == nil {
		return
	}
	// 栈中保存的是当前的路径
	stack := linkedliststack.New()
	set := hashset.New()
	stack.Push(root)
	fmt.Print(root.val)
	set.Add(root)
	for !stack.Empty() {
		cur, _ := stack.Pop()
		curNode := cur.(*Node)
		nexts := curNode.nexts
		for i := range nexts {
			if !set.Contains(nexts[i]) {
				stack.Push(curNode)
				stack.Push(nexts[i])
				set.Add(nexts[i])
				fmt.Print(nexts[i].val)
				break
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
