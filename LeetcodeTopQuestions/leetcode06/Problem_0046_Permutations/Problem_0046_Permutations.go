package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
	numsWithRepeat := []int{1, 1, 3}
	fmt.Println(permuteWithRepeat(numsWithRepeat))
}

// https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	process(&ans, 0, nums)
	return ans
}

func process(ans *[][]int, index int, nums []int) {
	if index == len(nums) {
		res := make([]int, len(nums))
		for i := 0; i < len(nums); i++ {
			res[i] = nums[i]
		}
		*ans = append(*ans, res)
		return
	}
	for i := index; i < len(nums); i++ {
		swap(nums, index, i)
		process(ans, index+1, nums)
		swap(nums, index, i)
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func permuteWithRepeat(nums []int) [][]int {
	ans := make([][]int, 0)
	processWithRepeat(&ans, 0, nums)
	return ans
}

func processWithRepeat(ans *[][]int, index int, nums []int) {
	if index == len(nums) {
		res := make([]int, len(nums))
		for i := 0; i < len(nums); i++ {
			res[i] = nums[i]
		}
		*ans = append(*ans, res)
		return
	}
	isExist := make(map[int]int)
	for i := index; i < len(nums); i++ {
		_, exist := isExist[nums[i]]
		if !exist {
			isExist[nums[i]] = nums[i]
			swap(nums, index, i)
			processWithRepeat(ans, index+1, nums)
			swap(nums, index, i)
		}
	}
}
