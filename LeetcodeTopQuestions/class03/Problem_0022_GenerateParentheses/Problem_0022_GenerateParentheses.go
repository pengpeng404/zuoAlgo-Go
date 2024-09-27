package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

// https://leetcode.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	if n == 0 {
		return make([]string, 0)
	}
	ans := make([]string, 0)
	path := make([]byte, 2*n)
	process(0, &ans, 0, n, path)
	return ans
}

func process(index int, ans *[]string, leftMinusRight int, leftRest int, path []byte) {
	if index == len(path) {
		*ans = append(*ans, string(path))
		return
	}
	if leftRest > 0 {
		// 说明还有左括号可以加
		path[index] = '('
		process(index+1, ans, leftMinusRight+1, leftRest-1, path)
	}
	if leftMinusRight > 0 {
		// 说明可以加右括号
		path[index] = ')'
		process(index+1, ans, leftMinusRight-1, leftRest, path)
	}
}
