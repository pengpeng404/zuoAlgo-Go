package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 5, 2, 7, 6, 9, 3, 8, 7, 10, 4, 7}
	fmt.Println(selectionSort(arr))
}

func selectionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := len(arr) - 1; i > 0; i-- {
		index := i
		for j := 0; j <= i; j++ {
			if arr[j] > arr[index] {
				index = j
			}
		}
		swap(arr, i, index)
	}
	return arr
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
