package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 1, 0, 4}
	// false
	fmt.Println(canJump(nums))
}

// https://leetcode.com/problems/jump-game/

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	toMax := nums[0]
	for i := 1; i < len(nums); i++ {
		// 过滤条件
		if toMax >= len(nums)-1 {
			return true
		}
		if i > toMax {
			return false
		}
		toMax = max(toMax, i+nums[i])
	}
	return true
}
