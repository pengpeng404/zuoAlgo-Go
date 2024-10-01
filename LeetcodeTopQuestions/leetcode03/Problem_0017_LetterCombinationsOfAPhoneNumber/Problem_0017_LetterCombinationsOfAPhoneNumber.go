package main

import "fmt"

func main() {
	fmt.Print(letterCombinations("23"))
}

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}
	phone := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}
	ans := make([]string, 0)
	path := make([]byte, len(digits))
	process(digits, 0, &ans, path, phone)
	return ans
}

func process(digits string, index int, ans *[]string, path []byte, phone [][]byte) {
	if index == len(digits) {
		*ans = append(*ans, string(path))
		return
	}
	indexByte := phone[digits[index]-'2']
	for i := 0; i < len(indexByte); i++ {
		path[index] = indexByte[i]
		process(digits, index+1, ans, path, phone)
	}
}
