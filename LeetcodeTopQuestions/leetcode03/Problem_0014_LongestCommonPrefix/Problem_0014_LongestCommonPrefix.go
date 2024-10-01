package main

import "fmt"

func main() {
	strs := []string{
		"abc",
		"abde",
		"abac",
		"abbe",
	}
	fmt.Print(longestCommonPrefix(strs))
}

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minIndex := len(strs[0])

	for i := 1; i < len(strs); i++ {
		index := 0
		nextArr := []byte(strs[i])
		// 使用 min 限制比较长度
		maxCompareLength := min(minIndex, len(nextArr))

		for index < maxCompareLength && strs[0][index] == nextArr[index] {
			index++
		}
		minIndex = min(minIndex, index)

		if minIndex == 0 {
			return ""
		}
	}
	return strs[0][:minIndex]
}
