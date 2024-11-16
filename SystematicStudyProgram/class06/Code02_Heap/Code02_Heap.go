package main

import "fmt"

func main() {
	heap := newMaxHeap(4)
	heap.add(1)
	fmt.Println(heap.pop())
	heap.add(4)
	heap.add(88)
	fmt.Println(heap.pop())
	heap.add(5)
	fmt.Println(heap.pop())
	fmt.Println(heap.pop())
}

// 堆实现 大根堆

type maxHeap struct {
	heap     []int
	heapSize int
	limit    int
}

func newMaxHeap(limit int) *maxHeap {
	return &maxHeap{
		heap:     make([]int, limit),
		heapSize: 0,
		limit:    limit,
	}
}

func (h *maxHeap) add(num int) {
	if h.heapSize == h.limit {
		fmt.Println("heap is full")
		return
	}
	h.heap[h.heapSize] = num
	h.heapSize++
	h.heapInsert(h.heapSize - 1)
}

func (h *maxHeap) pop() int {
	if h.heapSize == 0 {
		fmt.Println("heap is empty")
		return 0
	}
	ans := h.heap[0]
	h.heapSize--
	swap(h.heap, 0, h.heapSize)
	h.heapify(0)
	return ans
}

func (h *maxHeap) heapify(index int) {
	left := index*2 + 1
	for left < h.heapSize {
		largest := left
		if index*2+2 < h.heapSize && h.heap[index*2+2] > h.heap[index*2+1] {
			largest = index*2 + 2
		}
		if h.heap[index] >= h.heap[largest] {
			largest = index
		}
		if largest == index {
			break
		}
		swap(h.heap, index, largest)
		index = largest
		left = index*2 + 1
	}
}

func (h *maxHeap) heapInsert(index int) {
	for h.heap[index] > h.heap[(index-1)/2] {
		swap(h.heap, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
