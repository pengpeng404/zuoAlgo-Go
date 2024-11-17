package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var maxLen = 20
	var maxValue = 20
	var testTimes = 5000000
	for i := 0; i < testTimes; i++ {
		arr := getRandomArr(maxLen, maxValue)
		arrCopy := CopyArr(arr)
		heapSort(arr)
		sort.Ints(arrCopy)
		if !EqualArr(arr, arrCopy) {
			PrintArr(arr)
			PrintArr(arrCopy)
			panic("shit bro sort is error")
		}
	}
	fmt.Println("test good")

}

/*
如何使用大根堆排序
1、把整个数组搞成大根堆 [0 - n-1]   heapSize == N
2、swap(0, n-1) heapSize--
3、heapify(0) swap(0, n-1) heapSize--
4、loop 直到 heapSize == 0
*/

/*
从上向下建堆 N*logN
从下向上建堆 N
*/

func heapSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	heapSize := len(arr)
	// 1、建立大根堆
	//for i := 0; i < len(arr); i++ {
	//	heapInsert(arr, i)
	//}
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, heapSize)
	}
	// 2、交换
	heapSize--
	swap(arr, 0, heapSize)
	// 3、重建大根堆并交换
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		heapSize--
		swap(arr, 0, heapSize)
	}
}

func heapify(arr []int, index, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		largest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}
		if arr[index] >= arr[largest] {
			largest = index
		}
		if largest == index {
			break
		}
		swap(arr, index, largest)
		index = largest
		left = index*2 + 1
	}
}

func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
func getRandomArr(maxLen int, maxValue int) []int {
	// 随机生成数组长度，范围是 [0, maxLen]
	length := rand.Intn(maxLen + 1)
	// 创建切片，初始容量为计算得到的长度
	array := make([]int, length)
	// 填充切片
	for i := 0; i < length; i++ {
		array[i] = rand.Intn(maxValue + 1) // 生成 [0, maxValue] 范围内的随机数
	}
	return array
}
func CopyArr(arr []int) []int {
	// 创建目标切片，并分配足够的空间
	dst := make([]int, len(arr))
	// 使用copy函数复制切片
	copy(dst, arr)
	return dst
}
func PrintArr(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
func EqualArr(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
