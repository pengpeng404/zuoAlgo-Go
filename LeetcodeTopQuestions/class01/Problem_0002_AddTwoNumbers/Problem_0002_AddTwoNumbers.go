package main

func main() {

}

// https://leetcode.com/problems/add-two-numbers/

/*
方法一：直接在链表上操作
方法二：使用容器
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// =============================== 复制以下内容
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ca := 0
	Val1 := 0
	Val2 := 0
	n := 0
	node1 := l1
	node2 := l2
	var pre *ListNode
	var node *ListNode
	for node1 != nil || node2 != nil {
		Val1 = 0
		if node1 != nil {
			Val1 = node1.Val
		}
		Val2 = 0
		if node2 != nil {
			Val2 = node2.Val
		}
		n = Val1 + Val2 + ca
		ca = n / 10
		node = &ListNode{Val: n % 10}
		node.Next = pre
		pre = node
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
	}
	if ca == 1 {
		node = &ListNode{Val: 1}
		node.Next = pre
		pre = node
	}
	return reverseListNode(pre)
}

func reverseListNode(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func addTwoNumbersContainer(l1 *ListNode, l2 *ListNode) *ListNode {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	ans := make([]int, 0)
	node1 := l1
	node2 := l2
	for node1 != nil || node2 != nil {
		if node1 != nil {
			arr1 = append(arr1, node1.Val)
			node1 = node1.Next
		}
		if node2 != nil {
			arr2 = append(arr2, node2.Val)
			node2 = node2.Next
		}

	}
	ca := 0
	len1 := len(arr1)
	len2 := len(arr2)
	maxLen := 0
	minLen := 0
	var maxLenArr []int
	if len1 > len2 {
		maxLenArr = arr1
		maxLen = len1
		minLen = len2
	} else {
		maxLenArr = arr2
		maxLen = len2
		minLen = len1
	}
	for i := 0; i < minLen; i++ {
		n := arr1[i] + arr2[i] + ca
		ans = append(ans, n%10)
		ca = n / 10
	}
	for i := minLen; i < maxLen; i++ {
		n := maxLenArr[i] + ca
		ans = append(ans, n%10)
		ca = n / 10
	}
	if ca == 1 {
		ans = append(ans, 1)
	}
	head := &ListNode{Val: ans[0]}
	node := head
	for i := 1; i < len(ans); i++ {
		next := &ListNode{Val: ans[i]}
		node.Next = next
		node = next
	}
	return head
}
