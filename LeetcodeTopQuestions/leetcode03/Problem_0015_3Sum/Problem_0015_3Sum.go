package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, -1, -4, 5, -3, 6}
	fmt.Println(threeSum(nums))
}

// https://leetcode.com/problems/3sum/description/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return threeSumK(nums, 0)
}

func threeSumK(nums []int, k int) [][]int {
	var ans [][]int
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] != nums[i] {
			// 此处返回两数之和为某个值的结果
			iAns := twoSumK(nums, k-nums[i], i+1)
			for j := 0; j < len(iAns); j++ {
				ans = append(ans, []int{nums[i], iAns[j][0], iAns[j][1]})
			}
		}
	}
	return ans
}

func twoSumK(nums []int, k, index int) [][]int {
	if index >= len(nums)-1 {
		return make([][]int, 0)
	}
	var ans [][]int
	l := index
	r := len(nums) - 1
	for l < r {
		if nums[l]+nums[r] > k {
			r--
		} else if nums[l]+nums[r] < k {
			l++
		} else {
			if l == index || nums[l] != nums[l-1] {
				ans = append(ans, []int{nums[l], nums[r]})
				l++
			} else {
				l++
			}
		}
	}
	return ans
}
