package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	Value string
	Edges []*Edge
}

type Edge struct {
	To     *Node
	Weight int
}

type NodeHeap struct {
	nodes    []*Node
	index    map[*Node]int
	distance map[*Node]int
}

func NewNodeHeap() *NodeHeap {
	return &NodeHeap{
		nodes:    []*Node{},
		index:    make(map[*Node]int),
		distance: make(map[*Node]int),
	}
}

func (h *NodeHeap) Len() int {
	return len(h.nodes)
}

func (h *NodeHeap) Less(i, j int) bool {
	return h.distance[h.nodes[i]] < h.distance[h.nodes[j]]
}

func (h *NodeHeap) Swap(i, j int) {
	h.index[h.nodes[i]], h.index[h.nodes[j]] = j, i
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *NodeHeap) Push(x interface{}) {
	node := x.(*Node)
	h.nodes = append(h.nodes, node)
}

func (h *NodeHeap) Pop() interface{} {
	old := h.nodes
	n := len(old)
	node := old[n-1]
	h.nodes = old[0 : n-1]
	return node
}

func (h *NodeHeap) Update(node *Node, distance int) {
	h.distance[node] = distance
	heap.Fix(h, h.index[node])
}

func Dijkstra(from *Node) map[*Node]int {
	distanceMap := make(map[*Node]int)
	distanceMap[from] = 0
	selectedNodes := make(map[*Node]struct{})

	nodeHeap := NewNodeHeap()
	nodeHeap.Push(from)
	nodeHeap.distance[from] = 0
	heap.Init(nodeHeap)

	for nodeHeap.Len() > 0 {
		minNode := heap.Pop(nodeHeap).(*Node)
		distance := distanceMap[minNode]

		for _, edge := range minNode.Edges {
			toNode := edge.To
			newDist := distance + edge.Weight
			if _, selected := selectedNodes[toNode]; !selected {
				if existingDist, exists := distanceMap[toNode]; !exists || newDist < existingDist {
					distanceMap[toNode] = newDist
					nodeHeap.Update(toNode, newDist)
				}
			}
		}
		selectedNodes[minNode] = struct{}{}
	}

	return distanceMap
}

func main() {
	// Example usage
	n1 := &Node{Value: "A"}
	n2 := &Node{Value: "B"}
	n3 := &Node{Value: "C"}

	n1.Edges = append(n1.Edges, &Edge{To: n2, Weight: 1})
	n1.Edges = append(n1.Edges, &Edge{To: n3, Weight: 4})
	n2.Edges = append(n2.Edges, &Edge{To: n3, Weight: 2})

	result := Dijkstra(n1)
	for node, dist := range result {
		fmt.Printf("Distance to %s: %d\n", node.Value, dist)
	}
}
