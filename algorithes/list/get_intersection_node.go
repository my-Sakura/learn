package list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	if headA == headB {
		return headA
	}

	hashMap := make(map[*ListNode]bool)
	for headA != nil {
		hashMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := hashMap[headB]; ok {
			return headB
		}
		hashMap[headB] = true
		headB = headB.Next
	}

	return nil
}

// two pointers
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
//     if headA == nil || headB == nil {
//         return nil
//     }
//     p1, p2 := headA, headB

//     for {
//         if p1 == nil && p2 == nil {
//             return nil
//         }

//         if p1 == nil {
//             p1 = headB
//         }
//         if p2 == nil {
//             p2 = headA
//         }

//         if p1 == p2 {
//             return p1
//         }

//         p1 = p1.Next
//         p2 = p2.Next
//     }
// }
