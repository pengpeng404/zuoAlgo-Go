package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
荷兰国旗问题
*/
// netherlandsFlag 荷兰国旗问题
// <[R] ==[R] >[R]
// 返回等于区域的下标数组
// 每一次做划分的是数组的最后一个数
func netherlandsFlag(arr []int, l, r int) []int {
	if l > r {
		return []int{-1, -1}
	}
	if l == r {
		return []int{l, l}
	}
	less := l - 1
	// 以 r 位置的数作为划分值
	more := r
	index := l
	// index 不进入到大于区域
	for index < more {
		if arr[index] == arr[r] {
			index++
		} else if arr[index] < arr[r] {
			swap(arr, index, less+1)
			less++
			index++
		} else {
			swap(arr, index, more-1)
			more--
			// 此时 index 不动
			// 因为换过来的数还没有看
		}
	}
	// 最后应该让 r 位置与 more 位置的数字交换
	// 也只能这样交换
	swap(arr, r, more)
	// 最后等于区域就是 less+1 和交换之后的 more 位置夹住的区域
	return []int{less + 1, more}
}

func main() {
	var maxLen = 20
	var maxValue = 20
	var testTimes = 5000000
	for i := 0; i < testTimes; i++ {
		arr := getRandomArr(maxLen, maxValue)
		arrCopy := CopyArr(arr)
		quickSort3(arr)
		sort.Ints(arrCopy)
		if !EqualArr(arr, arrCopy) {
			PrintArr(arr)
			PrintArr(arrCopy)
			panic("shit bro sort is error")
		}
	}
	fmt.Println("test good")

}

func quickSort2(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	process2(arr, 0, len(arr)-1)
}

func process2(arr []int, l, r int) {
	if l >= r {
		return
	}
	equalArea := netherlandsFlag(arr, l, r)
	process2(arr, l, equalArea[0]-1)
	process2(arr, equalArea[1]+1, r)
}

func quickSort3(arr []int) {
	if arr == nil || len(arr) <= 1 {
		return
	}
	process3(arr, 0, len(arr)-1)
}

func process3(arr []int, l, r int) {
	if l >= r {
		return
	}
	// 快排3.0就是随机选择一个位置与R位置做交换
	// 其他的与2.0版本一样
	swap(arr, l+rand.Intn(r-l+1), r)
	// 使用的还是荷兰国旗问题的最后一个值进行划分
	// 但是最后一个值的选取在之前做了随机选择
	equalArea := netherlandsFlag(arr, l, r)
	process3(arr, l, equalArea[0]-1)
	process3(arr, equalArea[1]+1, r)
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
