package main

import "fmt"

func main() {
	isConnected := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(isConnected))
}

// https://leetcode.com/problems/friend-circles/

func findCircleNum(isConnected [][]int) int {
	size := len(isConnected)
	uf := newUnionFind(size)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if isConnected[i][j] == 1 {
				uf.union(i, j)
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
	f := make([]int, n)
	s := make([]int, n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = 1
		f[i] = i
	}
	return &unionFind{
		father: f,
		size:   s,
		help:   h,
		sets:   n,
	}
}

func (uf *unionFind) find(ch int) int {
	// 使用size记录沿途有多少孩子
	size := 0
	for ch != uf.father[ch] {
		uf.help[size] = ch
		size++
		ch = uf.father[ch]
	}
	for i := 0; i < size; i++ {
		uf.father[uf.help[i]] = ch
	}
	return ch
}

func (uf *unionFind) union(ch1, ch2 int) {
	root1 := uf.find(ch1)
	root2 := uf.find(ch2)
	if root1 != root2 {
		sizeRoot1 := uf.size[root1]
		sizeRoot2 := uf.size[root2]
		if sizeRoot1 > sizeRoot2 {
			uf.father[root2] = root1
			uf.size[root1] += sizeRoot2
		} else {
			uf.father[root1] = root2
			uf.size[root2] += sizeRoot1
		}
		uf.sets--
	}
}

func (uf *unionFind) isSameSet(ch1, ch2 int) bool {
	return uf.find(ch1) == uf.find(ch2)
}
