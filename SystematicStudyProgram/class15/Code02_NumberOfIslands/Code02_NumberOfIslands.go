package main

import "fmt"

func main() {
	// islands 3
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

// https://leetcode.com/problems/number-of-islands/
// numIslands 改写并查集 遍历二维数组 对每一个有 1 的位置进行初始化
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	uf := newUnionFind(m * n)
	// 加入 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				uf.add(i*n + j)
			}
		}
	}
	for i := 0; i < n-1; i++ {
		if grid[0][i] == '1' && grid[0][i+1] == '1' {
			uf.union(i, i+1)
		}
		if grid[m-1][i] == '1' && grid[m-1][i+1] == '1' {
			uf.union(i+(m-1)*n, i+1+(m-1)*n)
		}
	}
	for i := 0; i < m-1; i++ {
		if grid[i][0] == '1' && grid[i+1][0] == '1' {
			uf.union(i*n, (i+1)*n)
		}
		if grid[i][n-1] == '1' && grid[i+1][n-1] == '1' {
			uf.union(i*n+n-1, (i+1)*n+n-1)
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == '1' && grid[i][j-1] == '1' {
				uf.union(i*n+j, i*n+j-1)
			}
			if grid[i][j] == '1' && grid[i][j+1] == '1' {
				uf.union(i*n+j, i*n+j+1)
			}
			if grid[i][j] == '1' && grid[i+1][j] == '1' {
				uf.union(i*n+j, (i+1)*n+j)
			}
			if grid[i][j] == '1' && grid[i-1][j] == '1' {
				uf.union(i*n+j, (i-1)*n+j)
			}
		}
	}
	return uf.sets
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
