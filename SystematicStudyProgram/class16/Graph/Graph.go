package main

import "github.com/emirpasic/gods/sets/hashset"

func main() {
	set := hashset.New()   // empty
	set.Add(1)             // 1
	set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Remove(4)          // 5, 3, 2, 1 (random order)
	set.Remove(2, 3)       // 1, 5 (random order)
	set.Contains(1)        // true
	set.Contains(1, 5)     // true
	set.Contains(1, 6)     // false
	_ = set.Values()       // []int{5,1} (random order)
	set.Clear()            // empty
	set.Empty()            // true
	set.Size()             // 0
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
