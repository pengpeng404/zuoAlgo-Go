package main

import "fmt"

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
	graph.PrintGraph()
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

// 打印Graph结构的信息
func (g *Graph) PrintGraph() {
	fmt.Println("Graph Information:")

	// 打印每个节点的信息
	fmt.Println("Nodes:")
	for _, node := range g.nodes {
		fmt.Printf("Node %d: In-degree: %d, Out-degree: %d\n", node.val, node.in, node.out)
		fmt.Print("  Next nodes: ")
		for _, next := range node.nexts {
			fmt.Printf("%d ", next.val)
		}
		fmt.Println()

		// 打印当前节点的边
		fmt.Println("  Edges:")
		for _, edge := range node.edges {
			fmt.Printf("    Edge from %d to %d, weight: %d\n", edge.from.val, edge.to.val, edge.weight)
		}
	}

	// 打印所有边的信息
	fmt.Println("Edges:")
	for edge, weight := range g.edges {
		fmt.Printf("Edge from %d to %d, weight: %d\n", edge.from.val, edge.to.val, weight)
	}

}
