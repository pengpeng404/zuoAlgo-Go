package main

import (
	"fmt"
	"math/rand"
)

/*
start 出发位置
end 结束位置
下标从 1 开始 只能向左跳 nums[i-1] 或者向右跳 nums[i-1]
返回到达目标的最小步

最优解 宽度优先遍历
*/

/*
在暴力递归中引入 K 改不了动态规划 只能暴力递归
因为 zuo 引入这个 K 的时候 整个暴力递归的函数定义不清晰
*/

func main() {

	maxN := 20
	maxV := 10
	testTimes := 200
	fmt.Println("test begin")

	for i := 0; i < testTimes; i++ {
		arr := getRandomArray(maxN, maxV)
		N := len(arr)
		start := rand.Intn(N) + 1
		end := rand.Intn(N) + 1
		ans1 := jumpBFS(N, start, end, arr)
		ans2 := jumpRight(N, start, end, arr)

		if ans1 != ans2 {
			printArray(arr)
			fmt.Printf("start : %d\n", start)
			fmt.Printf("end : %d\n", end)
			fmt.Printf("ans1 : %d\n", ans1)
			fmt.Printf("ans2 : %d\n", ans2)
			fmt.Println("Oops!")
			break
		}
	}
	fmt.Println("test end")
}

func jumpBFS(N, start, end int, arr []int) int {
	if start < 1 || start > N || end < 1 || end > N {
		return -1
	}
	level := make(map[int]int)
	level[start] = 0
	queue := make([]int, 0)
	queue = append(queue, start)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		curLevel, _ := level[cur]
		if cur == end {
			return curLevel
		}
		left := cur - arr[cur-1]
		right := cur + arr[cur-1]
		if _, exist := level[left]; left >= 1 && !exist {
			queue = append(queue, left)
			level[left] = curLevel + 1
		}
		if _, exist := level[right]; right <= N && !exist {
			queue = append(queue, right)
			level[right] = curLevel + 1
		}
	}
	return -1
}

// jumpRight 暴力递归
func jumpRight(N, start, end int, arr []int) int {
	if start < 1 || start > N || end < 1 || end > N {
		return -1
	}
	return process(arr, N, end, start, 0)
}

// process
func process(arr []int, N, end, cur, k int) int {
	if cur < 1 {
		return -1
	}
	if cur > N {
		return -1
	}
	if cur == end {
		return k
	}
	// k 最多跳 N 步
	if k > N {
		return -1
	}
	left := cur - arr[cur-1]
	right := cur + arr[cur-1]
	ans := -1
	ans1 := process(arr, N, end, left, k+1)
	ans2 := process(arr, N, end, right, k+1)
	if ans1 != -1 && ans2 != -1 {
		ans = min(ans1, ans2)
	} else if ans1 != -1 {
		ans = ans1
	} else {
		ans = ans2
	}

	return ans
}

// 生成随机数组
func getRandomArray(N, R int) []int {
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(R) // 生成 [0, R) 范围内的随机数
	}
	return arr
}

// 打印数组
func printArray(arr []int) {
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
