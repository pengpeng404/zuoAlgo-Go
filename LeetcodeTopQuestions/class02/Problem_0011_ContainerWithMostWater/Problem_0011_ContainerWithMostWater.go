package main

import (
	"fmt"
	"math"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// 49
	fmt.Println(maxArea(height))
}

/*
这一题有一个隐藏的情况 看笔记
*/
// https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	if len(height) == 1 {
		return 0
	}
	l := 0
	r := len(height) - 1
	ans := math.MinInt32
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] >= height[r] {
			r--
		} else {
			l++
		}
	}
	return ans
}
