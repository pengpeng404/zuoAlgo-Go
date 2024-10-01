package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5, 6}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) < 2 {
		return len(nums)
	}
	done := 0
	for cur := 1; cur < len(nums); cur++ {
		if nums[done] != nums[cur] {
			done++
			nums[done] = nums[cur]
		}
	}
	return done + 1
}
