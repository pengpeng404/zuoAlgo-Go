package main

import "math"

func main() {

}

// 单调栈 找到左边比 i 小的 右边比 i 小的 之间的范围就是大于等于 i 位置数的范围

// https://leetcode.com/problems/largest-rectangle-in-histogram

func largestRectangleArea(heights []int) int {
	length := len(heights)
	stack := make([]int, length)
	stackSize := 0
	ans := math.MinInt32
	for i := 0; i < length; i++ {
		for stackSize != 0 && heights[stack[stackSize-1]] >= heights[i] {
			j := stack[stackSize-1]
			stackSize--
			if stackSize == 0 {
				ans = max(ans, heights[j]*i)
			} else {
				ans = max(ans, heights[j]*(i-1-stack[stackSize-1]))
			}
		}
		stack[stackSize] = i
		stackSize++
	}
	for stackSize != 0 {
		j := stack[stackSize-1]
		stackSize--
		if stackSize == 0 {
			ans = max(ans, heights[j]*length)
		} else {
			ans = max(ans, heights[j]*(length-1-stack[stackSize-1]))
		}
	}
	return ans
}
