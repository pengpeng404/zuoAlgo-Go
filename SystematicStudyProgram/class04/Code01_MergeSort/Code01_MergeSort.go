package main

import "fmt"

func main() {
	arr := []int{3, 4, 5, 1, 2, 3}
	f(arr)
	fmt.Println(arr)
}

func f(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

func process(arr []int, l, r int) {
	// base case
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	process(arr, l, mid)
	process(arr, mid+1, r)
	MergeSort(arr, l, mid, r)
}

func MergeSort(arr []int, l, mid, r int) {
	help := make([]int, r-l+1)
	i := 0
	p1 := l
	p2 := mid + 1
	for p1 <= mid && p2 <= r {
		if arr[p1] < arr[p2] {
			help[i] = arr[p1]
			p1++
			i++
		} else {
			help[i] = arr[p2]
			p2++
			i++
		}
	}
	for p1 <= mid {
		help[i] = arr[p1]
		p1++
		i++
	}
	for p2 <= r {
		help[i] = arr[p2]
		p2++
		i++
	}

	for i2 := range help {
		arr[l+i2] = help[i2]
	}
}
