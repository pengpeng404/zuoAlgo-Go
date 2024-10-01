package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(7832452))
}

// https://leetcode.com/problems/reverse-integer/description/

func reverse(x int) int {
	//neg := ((x >> 31) & 1) == -1
	neg := x < 0
	m := math.MinInt32 / 10
	o := math.MinInt32 % 10
	if !neg {
		x = -x
	}
	res := 0
	for x != 0 {
		if res < m || (res == m && x%10 < o) {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}
	if !neg {
		return -res
	}
	return res
}
