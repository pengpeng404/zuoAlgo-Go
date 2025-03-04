package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 5}
	fmt.Println(netherlandsFlag(arr, 0, len(arr)-1))
	fmt.Println(arr)
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
