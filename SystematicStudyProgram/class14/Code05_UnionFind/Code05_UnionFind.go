package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 假设我们测试10000个元素
	n := 10000
	uf := newUnionFind(n)

	// 定义测试的union和find操作次数
	operations := 10000000
	sameSetCount := 0

	// 我们加入一个 map 来记录手动合并后的集合，用来做正确性的验证
	setMap := make(map[int]int) // 每个元素映射到它的根

	// 初始化setMap
	for i := 0; i < n; i++ {
		setMap[i] = i
	}

	// 测试操作
	for i := 0; i < operations; i++ {
		// 随机生成两个节点
		a := rand.Intn(n)
		b := rand.Intn(n)

		// 随机选择是执行union还是isSameSet操作
		if rand.Float64() < 0.5 {
			uf.union(a, b)

			// 更新setMap, 合并两个集合
			rootA := findSetRoot(setMap, a)
			rootB := findSetRoot(setMap, b)
			if rootA != rootB {
				// 合并两个集合
				for k, v := range setMap {
					if v == rootA {
						setMap[k] = rootB
					}
				}
			}
		} else {
			// 验证isSameSet的结果
			ufResult := uf.isSameSet(a, b)
			mapResult := findSetRoot(setMap, a) == findSetRoot(setMap, b)
			if ufResult != mapResult {
				fmt.Printf("错误: isSameSet(%d, %d) 应该返回 %v，但返回了 %v\n", a, b, mapResult, ufResult)
			} else {
				if ufResult {
					sameSetCount++
				}
			}
		}
	}

	fmt.Printf("Completed %d operations. Found %d same sets.\n", operations, sameSetCount)
}

// 辅助函数：找到元素的根节点（用来模拟正确的合并操作）
func findSetRoot(setMap map[int]int, x int) int {
	for setMap[x] != x {
		x = setMap[x]
	}
	return x
}

// 手写并查集

type unionFind struct {
	father []int
	size   []int
	help   []int
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
	}
}

func (uf *unionFind) isSameSet(ch1, ch2 int) bool {
	return uf.find(ch1) == uf.find(ch2)
}
