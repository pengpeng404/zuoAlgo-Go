package main

import "fmt"

func main() {
	nums := []int{3, 1, 7, 1, 2, 2, 1, 1, 2, 1}
	// 2
	fmt.Println(jump(nums))
}

// https://leetcode.com/problems/jump-game-ii/description/

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	step := 0
	cur := 0
	next := nums[0]
	length := len(nums) - 1
	for i := 1; i <= length; i++ {
		if i > cur {
			cur = next
			step++
		}
		// 过滤条件
		if cur >= length {
			return step
		}
		next = max(next, i+nums[i])
	}
	return step
}
