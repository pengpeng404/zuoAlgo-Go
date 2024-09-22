package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3456))
}

// https://leetcode.com/problems/integer-to-roman/description/
// 1 - 3999

func intToRoman(num int) string {
	chart := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}
	sb := strings.Builder{}
	sb.WriteString(chart[3][num/1000%10])
	sb.WriteString(chart[2][num/100%10])
	sb.WriteString(chart[1][num/10%10])
	sb.WriteString(chart[0][num%10])
	return sb.String()
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
