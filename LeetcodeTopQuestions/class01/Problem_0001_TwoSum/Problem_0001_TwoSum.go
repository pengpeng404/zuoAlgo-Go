package main

import "fmt"

func main() {
	arr := []int{3, 5, 7, 9}
	target := 10
	fmt.Println(twoSum(arr, target))
}

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	hp := make(map[int]int)
	for i := range nums {
		val, exist := hp[target-nums[i]]
		if exist {
			return []int{val, i}
		} else {
			hp[nums[i]] = i
		}
	}
	return []int{-1, -1}
}
