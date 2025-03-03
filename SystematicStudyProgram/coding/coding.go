package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */

func main() {
	arr := []int{1, 3, 2, 4, 1}
	fmt.Println(InversePairs(arr))
}

func InversePairs(nums []int) int {
	// write code here
	if len(nums) < 2 {
		return 0
	}
	return process(nums, 0, len(nums)-1)
}

func process(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)>>1
	num1 := process(arr, l, mid)
	num2 := process(arr, mid+1, r)
	num3 := merge(arr, l, mid, r)
	return num1 + num2 + num3
}

func merge(arr []int, l, mid, r int) int {
	help := make([]int, r-l+1)
	index := 0
	p1 := l
	p2 := mid + 1
	res := 0
	for p1 <= mid && p2 <= r {
		if arr[p1] > arr[p2] {
			help[index] = arr[p2]
			res += r - p2 + 1
			p2++
		} else {
			help[index] = arr[p1]
			p1++
		}
		index++
	}
	for p1 <= mid {
		help[index] = arr[p1]
		p1++
		index++
	}
	for p2 <= r {
		help[index] = arr[p2]
		p2++
		index++
	}
	for i := 0; i < len(help); i++ {
		arr[l+i] = help[i]
	}
	return res
}
