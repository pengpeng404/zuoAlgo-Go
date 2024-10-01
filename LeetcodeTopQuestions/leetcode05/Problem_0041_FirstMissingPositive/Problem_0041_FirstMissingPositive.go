package main

import "fmt"

func main() {
	nums := []int{1, 2, -1, -4}
	// 3
	fmt.Println(firstMissingPositive(nums))
}

// https://leetcode.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	l := 0
	r := len(nums)
	for l < r {
		if nums[l] == l+1 {
			l++
		} else if nums[l] <= l || nums[l] > r || nums[nums[l]-1] == nums[l] {
			swap(nums, l, r-1)
			r--
		} else {
			swap(nums, l, nums[l]-1)
		}
	}
	return l + 1
}

func swap(nums []int, a, b int) {
	tmp := nums[a]
	nums[a] = nums[b]
	nums[b] = tmp
}
