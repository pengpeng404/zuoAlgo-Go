package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{}(())[][]"))
}

// https://leetcode.com/problems/valid-parentheses/description/

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if opening, exists := pairs[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
