package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{3, 5, 7, 3, 6, 9, 6, 5, 7, 1, 2, 9, 7, 6, 7, 9, 8, 7, 6, 5, 6, 7, 8, 9, 3, 4, 23, 2, 4, 435, 6, 7, 7}
	process(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func process(arr []int, l, r int) {
	if l >= r {
		return
	}
	index := rand.Intn(r - l + 1)
	swap(arr, l+index, r)
	area := netherlandsFlag(arr, l, r)
	process(arr, l, area[0]-1)
	process(arr, area[1]+1, r)
}

// 荷兰国旗问题

func netherlandsFlag(arr []int, l, r int) []int {
	if l == r {
		return []int{l, r}
	}
	less := l - 1
	more := r
	index := l
	for index < more {
		if arr[index] < arr[r] {
			swap(arr, index, less+1)
			less++
			index++
		} else if arr[index] == arr[r] {
			index++
		} else {
			swap(arr, index, more-1)
			more--
		}
	}
	swap(arr, r, more)
	return []int{less + 1, more}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
