package main

import "fmt"

func main() {
	root := newNode(1)
	root.next = newNode(2)
	//root.next.next = newNode(3)
	fmt.Println(midOrUpMidNode(root))
}

/*
单链表 仅用几个变量 实现
1、奇数个节点 返回唯一中点 偶数个节点 返回上中点
2、奇数个节点 返回唯一中点 偶数个节点 返回下中点
3、奇数个节点 返回中点前一个 偶数个节点 返回上中点前一个
4、奇数个节点 返回中点前一个 偶数个节点 返回下中点前一个
*/

type node struct {
	val  int
	next *node
}

func newNode(val int) *node {
	return &node{
		val: val,
	}
}

// midOrUpMidNode 函数 1
func midOrUpMidNode(head *node) *node {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}
	slow := head
	fast := head
	// 为什么这个地方的 fast 要检查两步？
	// 因为 fast 要跳两步
	if fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

func midOrDownMidNode(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	slow := head.next
	fast := head.next
	if fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func midOrUpMidPreNode(head *node) *node {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}
	slow := head
	fast := head.next.next
	// 为什么这个地方的 fast 要检查两步？
	// 因为 fast 要跳两步
	if fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

func midOrDownMidPreNode(head *node) *node {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}
	slow := head
	fast := head.next.next
	// 为什么这个地方的 fast 要检查两步？
	// 因为 fast 要跳两步
	if fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}
