package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(8, -7))
}

// https://leetcode.com/problems/divide-two-integers/description/

func divide(dividend int, divisor int) int {
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		} else {
			return 0
		}
	}
	// 被除数是系统最小
	if dividend == math.MinInt32 {
		if divisor == negNum(1) {
			return math.MaxInt32
		}
		res := div(add(dividend, 1), divisor)
		res = add(res, div(minus(dividend, multi(res, divisor)), divisor))
		return res
	}

	// 都不是系统最小
	return div(dividend, divisor)
}

func div(a, b int) int {
	neg := false
	if a < 0 {
		neg = !neg
		a = negNum(a)
	}
	if b < 0 {
		neg = !neg
		b = negNum(b)
	}
	res := 0
	for i := 31; i >= 0; i-- {
		if (a >> i) >= b {
			res |= (1 << i)
			a = minus(a, b<<i)
		}
	}
	if neg {
		return negNum(res)
	} else {
		return res
	}
}

/*
位运算实现 加减乘除
*/

func add(a, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
	return sum
}

func minus(a, b int) int {
	return add(a, negNum(b))
}

func negNum(a int) int {
	return add(^a, 1)
}

func multi(a, b int) int {
	negative := false

	// 如果 b 是负数，转换为正数，标记结果为负
	if b < 0 {
		b = negNum(b)
		negative = !negative
	}

	// 如果 a 是负数，转换为正数，标记结果为负
	if a < 0 {
		a = negNum(a)
		negative = !negative
	}

	res := 0
	for b != 0 {
		if b&1 != 0 {
			res = add(res, a)
		}
		a <<= 1
		b >>= 1
	}

	// 根据符号返回正负结果
	if negative {
		return negNum(res)
	}
	return res
}
