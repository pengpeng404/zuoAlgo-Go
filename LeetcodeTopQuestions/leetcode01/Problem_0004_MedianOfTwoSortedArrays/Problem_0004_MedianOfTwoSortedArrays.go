package main

import "fmt"

func main() {
	arr1 := []int{1, 3, 4, 7, 9}
	arr2 := []int{3, 6, 7, 8, 10}
	// 1 3 3 4 6 7 7 8 9 10
	// (6 + 7) / 2 = 6.5
	fmt.Println(findMedianSortedArrays(arr1, arr2))
}

// https://leetcode.com/problems/median-of-two-sorted-arrays/
/*
重要的算法原型 两个等长的有序数组中 找到上中位数 getUpMedian
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	length := m + n
	even := (m+n)%2 == 0
	if m != 0 && n != 0 {
		if even {
			// 如果是偶数
			return (float64(findKthNum(nums1, nums2, length/2)) + float64(findKthNum(nums1, nums2, length/2+1))) / 2
		} else {
			// 如果是奇数
			return float64(findKthNum(nums1, nums2, (length+1)/2))
		}
	} else if m != 0 {
		if even {
			// 如果是偶数
			return (float64(nums1[m/2]) + float64(nums1[(m/2)-1])) / 2
		} else {
			// 如果是奇数
			return float64(nums1[(m-1)/2])
		}
	} else if n != 0 {
		if even {
			// 如果是偶数
			return (float64(nums2[n/2]) + float64(nums2[(n/2)-1])) / 2
		} else {
			// 如果是奇数
			return float64(nums2[(n-1)/2])
		}
	} else {
		return 0
	}
}

func findKthNum(a1, a2 []int, Kth int) int {
	long := a1
	shot := a2
	l := len(a1)
	s := len(a2)
	if l < s {
		shot = a1
		long = a2
		l = len(a2)
		s = len(a1)
	}
	if Kth <= s {
		return getUpMedian(a1, a2, 0, Kth-1, 0, Kth-1)
	} else if Kth > s && Kth <= l {
		if long[Kth-s-1] >= shot[s-1] {
			return long[Kth-s-1]
		} else {
			return getUpMedian(shot, long, 0, s-1, Kth-s, Kth-1)
		}
	} else {
		if shot[Kth-l-1] >= long[l-1] {
			return shot[Kth-l-1]
		} else if long[Kth-s-1] >= shot[s-1] {
			return long[Kth-s-1]
		} else {
			return getUpMedian(shot, long, Kth-l, s-1, Kth-s, l-1)
		}
	}
}

func getUpMedian(a1, a2 []int, s1, e1, s2, e2 int) int {
	mid1 := 0
	mid2 := 0
	for s1 < e1 {
		mid1 = (s1 + e1) / 2
		mid2 = (s2 + e2) / 2
		if a1[mid1] == a2[mid2] {
			return a1[mid1]
		}
		if (e1-s1)%2 == 1 {
			// 偶数
			if a1[mid1] > a2[mid2] {
				e1 = mid1
				s2 = mid2 + 1
			} else {
				s1 = mid1 + 1
				e2 = mid2
			}
		} else {
			// 奇数
			if a1[mid1] > a2[mid2] {
				if a2[mid2] >= a1[mid1-1] {
					return a2[mid2]
				} else {
					e1 = mid1 - 1
					s2 = mid2 + 1
				}
			} else {
				if a1[mid1] >= a2[mid2-1] {
					return a1[mid1]
				} else {
					s1 = mid1 + 1
					e2 = mid2 - 1
				}
			}
		}
	}
	return min(a1[s1], a2[s2])
}
