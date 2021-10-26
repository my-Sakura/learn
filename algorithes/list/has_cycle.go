package list

func hasCycle(head *ListNode) bool {
	hashMap := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := hashMap[head]; ok {
			return true
		}
		hashMap[head] = true
		head = head.Next
	}

	return false
}

// fast && slow pointers
// func hasCycle(head *ListNode) bool {
//     if head == nil || head.Next == nil {
//         return false
//     }

//     slow, fast := head, head.Next
//     for fast != slow {
//         if fast == nil || fast.Next == nil {
//             return false
//         }
//         fast = fast.Next.Next
//         slow = slow.Next
//     }

//     return true
// }
