package main

import "fmt"

func main() {
	height := []int{3, 2, 5, 1, 6}
	fmt.Println(trap(height))
}

// https://leetcode.com/problems/trapping-rain-water/

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	left := height[0]
	right := height[len(height)-1]
	L := 1
	R := len(height) - 2
	water := 0
	for L <= R {
		if left <= right {
			water += max(0, left-height[L])
			left = max(left, height[L])
			L++
		} else {
			water += max(0, right-height[R])
			right = max(right, height[R])
			R--
		}
	}
	return water
}
