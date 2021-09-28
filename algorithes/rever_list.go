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

func reverseList2(head *node) {
	newList := new(node)
	for head != nil {
		mid := head.next
		head.next = newList.next
		newList.next = head
		head = mid
	}

	for newList.next != nil {
		fmt.Println(newList.next.val)
		newList.next = newList.next.next
	}
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

func main() {
	n1 := node{val: "1"}
	n2 := node{val: "2"}
	n3 := node{val: "3"}
	n1.next = &n2
	n2.next = &n3
	fmt.Println(n1.val, n2.val, n3.val)

}
