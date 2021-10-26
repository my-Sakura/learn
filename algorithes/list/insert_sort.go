package list

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyNode := &ListNode{Next: head}
	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Val < pre.Val {
			tmp := dummyNode
			for cur.Val > tmp.Next.Val {
				tmp = tmp.Next
			}
			pre.Next = cur.Next
			cur.Next = tmp.Next
			tmp.Next = cur
		} else {
			pre = cur
		}
		cur = pre.Next
	}

	return dummyNode.Next
}
