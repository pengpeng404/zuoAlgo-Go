package main

import (
	"fmt"
	"math/rand"
)

func main() {
	testTimes := 50000
	for i := 0; i < testTimes; i++ {
		nums := generateRandomArray()
		target := rand.Intn(30) + 10
		if bestSolve(nums, target) != right(nums, target) {
			fmt.Println("Oops!!!!!")
			fmt.Println(bestSolve(nums, target))
			fmt.Println(right(nums, target))
			break
		}
	}
}

/*
给定数组 nums 和 target
返回所有满足条件的子数组数量
	条件： subArr.max - subArr.min <= target

暴力解 ：o(N^3)
窗口内最大值最小值更新结构 ：o(N)
*/

func bestSolve(nums []int, target int) int {
	maxQueue := make([]int, 0)
	minQueue := make([]int, 0)
	r := 0
	count := 0
	for l := 0; l < len(nums); l++ {
		for r < len(nums) {
			for len(maxQueue) != 0 && nums[r] >= nums[maxQueue[len(maxQueue)-1]] {
				maxQueue = maxQueue[:len(maxQueue)-1]
			}
			maxQueue = append(maxQueue, r)
			for len(minQueue) != 0 && nums[r] <= nums[minQueue[len(minQueue)-1]] {
				minQueue = minQueue[:len(minQueue)-1]
			}
			minQueue = append(minQueue, r)
			if nums[maxQueue[0]]-nums[minQueue[0]] > target {
				break
			} else {
				r++
			}
		}
		count += r - l
		if maxQueue[0] == l {
			maxQueue = maxQueue[1:]
		}
		if minQueue[0] == l {
			minQueue = minQueue[1:]
		}
	}
	return count
}

// right 暴力解
func right(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			maxNum := nums[i]
			minNum := nums[i]
			for k := i + 1; k <= j; k++ {
				if nums[k] > maxNum {
					maxNum = nums[k]
				}
				if nums[k] < minNum {
					minNum = nums[k]
				}
			}
			if maxNum-minNum <= target {
				count++
			}
		}
	}
	return count
}

func generateRandomArray() []int {
	// 随机生成数组长度，范围在 10 到 50
	length := rand.Intn(41) + 10
	// 初始化数组
	arr := make([]int, length)
	// 填充数组，值范围在 10 到 100 之间
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(91) + 10
	}
	return arr
}
