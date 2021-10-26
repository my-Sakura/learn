package list

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	font, mid := midNode(head)
	back := reverseList(mid)
	var tail *ListNode

	for font != nil && back != nil {
		aTmp := font.Next
		bTmp := back.Next
		back.Next = aTmp
		font.Next = back
		tail = back
		font = aTmp
		back = bTmp
	}

	if back != nil {
		tail.Next = back
	}
}

func midNode(head *ListNode) (font, back *ListNode) {
	pre := &ListNode{Next: head}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		pre = pre.Next
	}
	pre.Next = nil

	return head, slow
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}

	return pre
}
