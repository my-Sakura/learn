package list

func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	pre := dummyNode

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}

	return dummyNode.Next
}
