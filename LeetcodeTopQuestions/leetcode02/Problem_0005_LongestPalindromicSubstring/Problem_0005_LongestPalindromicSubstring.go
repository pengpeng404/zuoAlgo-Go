package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestPalindrome("1abac"))
}

// https://leetcode.com/problems/longest-palindromic-substring/description/
// manacher

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	charArr := manacherString(s)
	pArr := make([]int, len(charArr))
	R := -1
	C := -1
	mid := 0
	maxL := math.MinInt32
	for i := 0; i < len(charArr); i++ {
		if i < R {
			pArr[i] = min(R-i, pArr[2*C-i])
		} else {
			pArr[i] = 1
		}
		for i+pArr[i] < len(charArr) && i-pArr[i] > -1 {
			if charArr[i+pArr[i]] == charArr[i-pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}
		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}
		if maxL < pArr[i] {
			maxL = pArr[i]
			mid = i
		}
	}
	// 返回字符串
	mid = (mid - 1) / 2
	maxL -= 1
	if maxL&1 == 0 {
		return s[mid-(maxL/2)+1 : mid+maxL/2+1]
	} else {
		return s[mid-(maxL/2) : mid+maxL/2+1]
	}
}

func manacherString(s string) []byte {
	charArr := make([]byte, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		charArr[2*i] = '#'
		charArr[2*i+1] = s[i]
	}
	charArr[len(charArr)-1] = '#'
	return charArr
}
