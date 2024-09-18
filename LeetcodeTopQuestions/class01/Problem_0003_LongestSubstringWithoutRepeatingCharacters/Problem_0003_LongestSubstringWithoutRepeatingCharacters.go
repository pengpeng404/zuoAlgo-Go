package main

import "fmt"

func main() {
	str := "pengpeng404"
	// peng40 6
	fmt.Println(lengthOfLongestSubstring(str))
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
/*
以i位置为结尾的子串的最长无重复字符的字串长度
使用哈希表存储之前遇到的当前位置字符的位置
*/

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	length := len(s)
	dp := make([]int, length)
	dp[0] = 1
	hp := make(map[byte]int)
	hp[s[0]] = 0
	ans := 1
	for i := 1; i < length; i++ {
		if val, exist := hp[s[i]]; exist {
			dp[i] = min(dp[i-1]+1, i-val)
		} else {
			dp[i] = dp[i-1] + 1
		}
		ans = max(ans, dp[i])
		hp[s[i]] = i
	}
	return ans
}
