package main

import "fmt"

func main() {
	operators := make([]*Point, 4)
	operators[0] = &Point{1, 1}
	operators[1] = &Point{0, 1}
	operators[2] = &Point{3, 3}
	operators[3] = &Point{3, 4}
	n := 4
	m := 5
	// return [1 1 2 2]
	fmt.Println(NumIslands2(n, m, operators))
}

// https://www.lintcode.com/problem/434

type Point struct {
	X, Y int
}

/**
 * @param n: An integer
 * @param m: An integer
 * @param operators: an array of point
 * @return: an integer array
 */

func NumIslands2(n int, m int, operators []*Point) []int {
	// write your code here
	ans := make([]int, len(operators))
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
	}
	uf := newUnionFind(n * m)
	for i, op := range operators {
		if grid[op.X][op.Y] == 0 {
			uf.add(op.X*m + op.Y)
			grid[op.X][op.Y] = 1
		}
		if op.Y-1 >= 0 && grid[op.X][op.Y-1] == 1 {
			uf.union(op.X*m+op.Y-1, op.X*m+op.Y)
		}
		if op.Y+1 < m && grid[op.X][op.Y+1] == 1 {
			uf.union(op.X*m+op.Y+1, op.X*m+op.Y)
		}
		if op.X-1 >= 0 && grid[op.X-1][op.Y] == 1 {
			uf.union((op.X-1)*m+op.Y, op.X*m+op.Y)
		}
		if op.X+1 < n && grid[op.X+1][op.Y] == 1 {
			uf.union((op.X+1)*m+op.Y, op.X*m+op.Y)
		}
		ans[i] = uf.sets
	}
	return ans
}

type unionFind struct {
	father []int
	size   []int
	help   []int
	sets   int
}

func newUnionFind(n int) *unionFind {
	return &unionFind{
		father: make([]int, n),
		size:   make([]int, n),
		help:   make([]int, n),
		sets:   0,
	}
}

// add 加入 index 到并查集中
func (uf *unionFind) add(index int) {
	uf.father[index] = index
	uf.size[index] = 1
	uf.sets++
}

func (uf *unionFind) find(index int) int {
	helpSize := 0
	for index != uf.father[index] {
		uf.help[helpSize] = index
		helpSize++
		index = uf.father[index]
	}
	for i := 0; i < helpSize; i++ {
		uf.father[uf.help[i]] = index
	}
	return index
}

func (uf *unionFind) union(index1, index2 int) {
	root1 := uf.find(index1)
	root2 := uf.find(index2)
	if root1 != root2 {
		if uf.size[root1] > uf.size[root2] {
			uf.father[root2] = root1
			uf.size[root1] += uf.size[root2]
		} else {
			uf.father[root1] = root2
			uf.size[root2] += uf.size[root1]
		}
		uf.sets--
	}
}
