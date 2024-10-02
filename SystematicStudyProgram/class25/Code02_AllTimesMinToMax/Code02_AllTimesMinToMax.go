package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 3, 1, 2}
	// 18
	fmt.Println(maxSumMinProduct(nums))
}

/*
单调栈结构 既然找到了左边比 i 小的 右边比 i 小的 那么中间的位置都是不小于 i 的
*/

// https://leetcode.com/problems/maximum-subarray-min-product/description/

func maxSumMinProduct(nums []int) int {
	length := len(nums)
	maxA := math.MinInt64
	stack := make([]int, length)
	stackSize := 0
	preSum := make([]int, length)
	preSum[0] = nums[0]
	for i := 1; i < length; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	for i := 0; i < length; i++ {
		for stackSize != 0 && nums[stack[stackSize-1]] >= nums[i] {
			// 弹出栈顶值
			j := stack[stackSize-1]
			stackSize--
			// 计算此时的结果
			if stackSize == 0 {
				maxA = max(maxA, nums[j]*preSum[i-1])
			} else {
				maxA = max(maxA, nums[j]*(preSum[i-1]-preSum[stack[stackSize-1]]))
			}
		}
		stack[stackSize] = i
		stackSize++
	}
	for stackSize != 0 {
		// 弹出栈顶值
		j := stack[stackSize-1]
		stackSize--
		// 计算此时的结果
		if stackSize == 0 {
			maxA = max(maxA, nums[j]*preSum[length-1])
		} else {
			maxA = max(maxA, nums[j]*(preSum[length-1]-preSum[stack[stackSize-1]]))
		}
	}
	return int(maxA % 1000000007)
}
