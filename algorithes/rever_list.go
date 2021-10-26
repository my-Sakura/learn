package main

import "fmt"

type node struct {
	next *node
	val  string
}

func reverseList1(head *node) []*node {
	newNode := make([]*node, 0)

	if head != nil {
		newNode = reverseList1(head.next)
		newNode = append(newNode, head)
	}

	return newNode
}

func reverseList2(head *node) *node {
	newList := new(node)
	for head != nil {
		mid := head.next
		head.next = newList
		newList = head
		head = mid
	}

	return newList
}

func reverseList3(head *node) {
	stack := make([]*node, 0)

	for head != nil {
		stack = append(stack, head)
		head = head.next
	}

	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i].val)
	}
}
