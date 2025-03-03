package main

import "fmt"

func main() {
	arr := []int{4, 5, 2, 7, 6, 9, 3, 8, 7, 10, 4, 7}
	fmt.Println(insertionSort(arr))
}

func insertionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
