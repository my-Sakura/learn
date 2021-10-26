package list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	tail := &ListNode{Next: head}
	mid, cur := head, head
	n := 0
	for cur != nil {
		n++
		tail = tail.Next
		cur = cur.Next
	}

	k = k % n
	for i := 0; i < n-k-1; i++ {
		mid = mid.Next
	}

	tail.Next = head
	head = mid.Next
	mid.Next = nil

	return head
}
