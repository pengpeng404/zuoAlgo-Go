package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 4, 4, 5, 5, 5, 5, 6, 6, 7}
	target := 4
	fmt.Println(searchRange(nums, target))
}

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	if nums == nil || len(nums) < 1 {
		return ans
	}
	ans[0] = findLeft(nums, target)
	ans[1] = findRight(nums, target)
	return ans
}

func findLeft(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	ans := -1
	mid := 0
	for l <= r {
		mid = l + (r-l)>>1
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			ans = mid
			r = mid - 1
		}
	}
	return ans
}

func findRight(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	ans := -1
	mid := 0
	for l <= r {
		mid = l + (r-l)>>1
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			ans = mid
			l = mid + 1
		}
	}
	return ans
}
