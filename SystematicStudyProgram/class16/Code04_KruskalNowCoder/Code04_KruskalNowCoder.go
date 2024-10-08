package main

import (
	"bufio"
	"fmt"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"os"
	"strconv"
	"strings"
)

// https://www.nowcoder.com/questionTerminal/c23eab7bb39748b6b224a8a3afbe396b
/*
牛客不支持gods包 复制最后注释掉的部分
*/

func main() {
	// 高效读取输入
	reader := bufio.NewReader(os.Stdin)

	// 读取第一行，包含节点数量和边数量
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0]) // 节点数
	m, _ := strconv.Atoi(parts[1]) // 边数

	uf := newUnionFind(n + 1)
	priqueue := pq.NewWith(byPriority)
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		w, _ := strconv.Atoi(parts[2])
		priqueue.Enqueue(&edge{
			from:   u,
			to:     v,
			weight: w,
		})
	}
	ans := 0
	for !priqueue.Empty() {
		cur, _ := priqueue.Dequeue()
		curEdge := cur.(*edge)
		if !uf.isSameSet(curEdge.from, curEdge.to) {
			ans += curEdge.weight
			uf.union(curEdge.from, curEdge.to)
		}
	}
	// 输出结果
	writer := bufio.NewWriter(os.Stdout)
	_, _ = fmt.Fprintln(writer, ans)
	_ = writer.Flush() // 确保输出刷新到控制台
}

func byPriority(a, b interface{}) int {
	priorityA := a.(*edge).weight
	priorityB := b.(*edge).weight
	return utils.IntComparator(priorityA, priorityB)
}

type edge struct {
	from   int
	to     int
	weight int
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

/*
复制下面
*/
//package main
//
//import (
//	"bufio"
//	"container/heap"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	// 高效读取输入
//	reader := bufio.NewReader(os.Stdin)
//
//	// 读取第一行，包含节点数量和边数量
//	line, _ := reader.ReadString('\n')
//	line = strings.TrimSpace(line)
//	parts := strings.Split(line, " ")
//	n, _ := strconv.Atoi(parts[0]) // 节点数
//	m, _ := strconv.Atoi(parts[1]) // 边数
//
//	uf := newUnionFind(n + 1)
//	priqueue := &PriorityQueue{}
//	heap.Init(priqueue)
//
//	for i := 0; i < m; i++ {
//		line, _ := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		parts := strings.Split(line, " ")
//		u, _ := strconv.Atoi(parts[0])
//		v, _ := strconv.Atoi(parts[1])
//		w, _ := strconv.Atoi(parts[2])
//		heap.Push(priqueue, &edge{
//			from:   u,
//			to:     v,
//			weight: w,
//		})
//	}
//	ans := 0
//	for priqueue.Len() > 0 {
//		cur := heap.Pop(priqueue).(*edge)
//		if !uf.isSameSet(cur.from, cur.to) {
//			ans += cur.weight
//			uf.union(cur.from, cur.to)
//		}
//	}
//
//	// 输出结果
//	writer := bufio.NewWriter(os.Stdout)
//	_, _ = fmt.Fprintln(writer, ans)
//	_ = writer.Flush() // 确保输出刷新到控制台
//}
//
//// 优先队列的实现
//type edge struct {
//	from   int
//	to     int
//	weight int
//}
//
//// PriorityQueue implements a priority queue using container/heap
//type PriorityQueue []*edge
//
//func (pq PriorityQueue) Len() int { return len(pq) }
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	return pq[i].weight < pq[j].weight // 最小堆，按边的权重排序
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	item := x.(*edge)
//	*pq = append(*pq, item)
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	item := old[n-1]
//	*pq = old[0 : n-1]
//	return item
//}
//
//// 手写并查集
//
//type unionFind struct {
//	father []int
//	size   []int
//	help   []int
//}
//
//func newUnionFind(n int) *unionFind {
//	f := make([]int, n)
//	s := make([]int, n)
//	h := make([]int, n)
//	for i := 0; i < n; i++ {
//		s[i] = 1
//		f[i] = i
//	}
//	return &unionFind{
//		father: f,
//		size:   s,
//		help:   h,
//	}
//}
//
//func (uf *unionFind) find(ch int) int {
//	// 使用size记录沿途有多少孩子
//	size := 0
//	for ch != uf.father[ch] {
//		uf.help[size] = ch
//		size++
//		ch = uf.father[ch]
//	}
//	for i := 0; i < size; i++ {
//		uf.father[uf.help[i]] = ch
//	}
//	return ch
//}
//
//func (uf *unionFind) union(ch1, ch2 int) {
//	root1 := uf.find(ch1)
//	root2 := uf.find(ch2)
//	if root1 != root2 {
//		sizeRoot1 := uf.size[root1]
//		sizeRoot2 := uf.size[root2]
//		if sizeRoot1 > sizeRoot2 {
//			uf.father[root2] = root1
//			uf.size[root1] += sizeRoot2
//		} else {
//			uf.father[root1] = root2
//			uf.size[root2] += sizeRoot1
//		}
//	}
//}
//
//func (uf *unionFind) isSameSet(ch1, ch2 int) bool {
//	return uf.find(ch1) == uf.find(ch2)
//}
