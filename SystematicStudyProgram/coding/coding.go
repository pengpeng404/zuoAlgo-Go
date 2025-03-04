package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line1))
	data := make([][]int, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line) // 按空格拆分
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		data[i] = make([]int, 2)
		data[i][0] = start
		data[i][1] = end
	}
	// 首先排序
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	he := newHeap(n)
	ans := 0
	for i := 0; i < n; i++ {
		for !he.isE() && he.arr[0] <= data[i][0] {
			he.pop()
		}
		he.add(data[i][1])
		if he.size > ans {
			ans = he.size
		}
	}
	fmt.Println(ans)
}

func (h *heap) add(n int) {
	h.arr[h.size] = n
	h.size++
	h.heapInsert(h.size - 1)
}

func (h *heap) pop() int {
	ans := h.arr[0]
	swap(h.arr, 0, h.size-1)
	h.size--
	h.heapify(0)
	return ans
}

func (h *heap) isE() bool {
	if h.size == 0 {
		return true
	}
	return false
}

func (h *heap) heapify(index int) {
	small := 2*index + 1
	for small < h.size {
		if small+1 < h.size && h.arr[small+1] < h.arr[small] {
			small++
		}
		if h.arr[index] > h.arr[small] {
			swap(h.arr, index, small)
			index = small
			small = 2*index + 1
		} else {
			break
		}
	}
}

func (h *heap) heapInsert(index int) {
	for index != 0 && h.arr[index] < h.arr[(index-1)/2] {
		swap(h.arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

type heap struct {
	arr   []int
	size  int
	limit int
}

func newHeap(limit int) *heap {
	return &heap{
		arr:   make([]int, limit),
		size:  0,
		limit: limit,
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
