package main

import "fmt"

func main() {
	root := newNode(1)
	root.Next = newNode(2)
	root.Next.Next = newNode(3)
	root.Next.Next.Next = newNode(4)
	root.Next.Next.Next.Next = newNode(5)
	root.Next.Next.Next.Next.Next = newNode(6)
	printListNode(removeNthFromEnd(root, 1))
	fmt.Println()
	printListNode(removeNthFromEnd(root, 1))
	fmt.Println()
	printListNode(removeNthFromEnd(root, 1))
	fmt.Println()
	printListNode(removeNthFromEnd(root, 1))
}

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var enough bool
	var pre *ListNode
	cur := head
	for cur != nil {
		n--
		if n <= 0 {
			if n == 0 {
				enough = true
			} else if n == -1 {
				pre = head
			} else {
				pre = pre.Next
			}
		}
		cur = cur.Next
	}
	if !enough {
		return head
	}
	if pre == nil {
		return head.Next
	}
	pre.Next = pre.Next.Next
	return head
}

/*************************  *************************/

type ListNode struct {
	Val  int
	Next *ListNode
}

func newNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func printListNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
}
