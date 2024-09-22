package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MMVI"))
}

// https://leetcode.com/problems/roman-to-integer/description/

func romanToInt(s string) int {
	nums := make([]int, len(s))
	for i := range s {
		if s[i] == 'I' {
			nums[i] = 1
		} else if s[i] == 'V' {
			nums[i] = 5
		} else if s[i] == 'X' {
			nums[i] = 10
		} else if s[i] == 'L' {
			nums[i] = 50
		} else if s[i] == 'C' {
			nums[i] = 100
		} else if s[i] == 'D' {
			nums[i] = 500
		} else if s[i] == 'M' {
			nums[i] = 1000
		}
	}
	ans := 0
	for i := 0; i < len(s)-1; i++ {
		if nums[i] < nums[i+1] {
			ans -= nums[i]
		} else {
			ans += nums[i]
		}
	}
	return ans + nums[len(s)-1]
}

/*
I	1
V	5
X	10
L	50
C	100
D	500
M	1000
*/
