package list

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	back := reverse(slow.Next)
	for back != nil {
		if back.Val != head.Val {
			return false
		}
		back = back.Next
		head = head.Next
	}

	return true
}
