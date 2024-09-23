package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	testTimes := 5000
	length := 2000
	for i := 0; i < testTimes; i++ {
		//fmt.Println("this is ", i, " time test")
		arr := generateRandomArray(length)
		w := rand.Intn(1500) + 1
		ans := slidingWindow(arr, w)
		rightAns := right(arr, w)
		if !isSame(ans, rightAns) {
			fmt.Println("Oops!!!!!!!")
			break
		}
	}
	fmt.Println("test finish")
}

/*
固定大小窗口滑过数组
每次滑动返回当前窗口的最大值
*/

func slidingWindow(arr []int, w int) []int {
	length := len(arr)
	ans := make([]int, length-w+1)
	var dq []int
	for i := 0; i < w; i++ {
		for len(dq) > 0 && arr[i] >= arr[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}
	ans[0] = arr[dq[0]]
	for i := w; i < length; i++ {
		for len(dq) > 0 && arr[i] >= arr[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if dq[0] == i-w {
			dq = dq[1:]
		}
		ans[i-w+1] = arr[dq[0]]
	}
	return ans
}

func right(arr []int, w int) []int {
	length := len(arr)
	ans := make([]int, length-w+1)
	for i := 0; i < length-w+1; i++ {
		res := math.MinInt32
		for j := i; j < i+w; j++ {
			res = max(res, arr[j])
		}
		ans[i] = res
	}
	return ans
}

func generateRandomArray(n int) []int {
	// 定义数组
	arr := make([]int, n)
	// 填充数组，元素范围为 -10^9 <= a_i <= 10^9
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(2*10000) - 10000
	}
	return arr
}

// 判断两个切片是否相同
func isSame(a, b []int) bool {
	// 如果长度不同，直接返回false
	if len(a) != len(b) {
		return false
	}

	// 逐个元素比较
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
