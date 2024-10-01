package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
}

// https://leetcode.com/problems/count-and-say/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	lastString := countAndSay(n - 1)
	sb := strings.Builder{}
	times := 1
	for i := 1; i < len(lastString); i++ {
		if lastString[i] == lastString[i-1] {
			times++
		} else {
			sb.WriteByte(byte('0' + times))
			sb.WriteByte(lastString[i-1])
			times = 1
		}
	}
	sb.WriteByte(byte('0' + times))
	sb.WriteByte(lastString[len(lastString)-1])
	return sb.String()
}
