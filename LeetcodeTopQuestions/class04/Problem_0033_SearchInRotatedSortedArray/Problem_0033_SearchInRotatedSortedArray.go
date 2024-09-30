package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 1, 2, 3}
	fmt.Println(search(nums, 3))

}

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
// search 二分
func search(nums []int, target int) int {
	L := 0
	R := len(nums) - 1
	M := 0
	for L <= R {
		M = L + (R-L)>>1
		if nums[M] == target {
			return M
		}
		// 如果 L M R 相等
		if nums[L] == nums[M] && nums[M] == nums[R] {
			for L != M && nums[L] == nums[M] {
				L++
			}
			if L == M {
				L++
				continue
			}
		}
		// L M R 不全相等
		if nums[L] != nums[M] {
			if nums[L] > nums[M] {
				if target > nums[M] && target <= nums[R] {
					L = M + 1
				} else {
					R = M - 1
				}
			} else {
				if target >= nums[L] && target < nums[M] {
					R = M - 1
				} else {
					L = M + 1
				}
			}
		} else {
			if nums[M] > nums[R] {
				if target >= nums[L] && target < nums[M] {
					R = M - 1
				} else {
					L = M + 1
				}
			} else {
				if target > nums[M] && target <= nums[R] {
					L = M + 1
				} else {
					R = M - 1
				}
			}
		}
	}
	return -1
}
