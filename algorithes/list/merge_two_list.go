package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := &ListNode{}
	ret := newList
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			newList.Next = l2
			newList = newList.Next
			l2 = l2.Next
		} else {
			newList.Next = l1
			newList = newList.Next
			l1 = l1.Next
		}
	}

	if l1 != nil {
		newList.Next = l1
	}

	if l2 != nil {
		newList.Next = l2
	}

	return ret.Next
}
