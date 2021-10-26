package list

func reverseKGroup(head *ListNode, k int) *ListNode {
	pre := head
	count := k
	for head != nil {
		k--
		if k == 0 {
			tmp := head.Next
			head.Next = nil
			newList := reverse(pre)
			pre.Next = reverseKGroup(tmp, count)
			return newList
		}

		head = head.Next
	}

	return pre
}

func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}
