package main

import "fmt"

func main() {
	root := newNode()
	root.insert("abc")
	root.insert("abdd")
	root.insert("abcc")
	root.insert("abcda")
	root.delete("abc")
	fmt.Println(root.search("abdd"))

}

type node struct {
	pass int
	end  int
	next []*node
}

func newNode() *node {
	return &node{
		pass: 0,
		end:  0,
		next: make([]*node, 26),
	}
}

func (n *node) insert(word string) {
	cnode := n
	cnode.pass++
	bytesArr := []byte(word)
	for _, bt := range bytesArr {
		if cnode.next[int(bt-'a')] == nil {
			cnode.next[int(bt-'a')] = newNode()
		}
		cnode = cnode.next[int(bt-'a')]
		cnode.pass++
	}
	cnode.end++
}

func (n *node) search(word string) int {
	if word == "" {
		return 0
	}
	cnode := n
	bytesArr := []byte(word)
	for _, bt := range bytesArr {
		if cnode.next[int(bt-'a')] == nil {
			return 0
		}
		cnode = cnode.next[int(bt-'a')]
	}
	return cnode.end
}

func (n *node) delete(word string) {
	if word == "" {
		return
	}
	if n.search(word) == 0 {
		return
	}
	// delete
	cnode := n
	cnode.pass--
	bytesArr := []byte(word)
	for _, bt := range bytesArr {
		cnode.next[int(bt-'a')].pass--
		if cnode.next[int(bt-'a')].pass == 0 {
			cnode.next[int(bt-'a')] = nil
			return
		}
		cnode = cnode.next[int(bt-'a')]
	}
	cnode.end--
}

func (n *node) prefixNumber(word string) int {
	if word == "" {
		return 0
	}
	cnode := n
	bytesArr := []byte(word)
	for _, bt := range bytesArr {
		if cnode.next[int(bt-'a')] == nil {
			return 0
		}
		cnode = cnode.next[int(bt-'a')]
	}
	return cnode.pass
}
