package binary_tree

func sortedListToBST(head *ListNode) *TreeNode {
	slice := make([]int, 0)

	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}

	return arrayToBST(slice)
}

func arrayToBST(array []int) *TreeNode {
	if len(array) == 0 {
		return nil
	}

	if len(array) == 1 {
		return &TreeNode{array[0], nil, nil}
	}

	mid := (len(array) - 1) >> 1
	root := &TreeNode{array[mid], nil, nil}
	root.Left = arrayToBST(array[:mid])
	root.Right = arrayToBST(array[mid+1:])

	return root
}

var globalHead *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	globalHead = head
	length := getLength(head)
	return buildTree(0, length-1)
}

func getLength(head *ListNode) int {
	ret := 0
	for ; head != nil; head = head.Next {
		ret++
	}
	return ret
}
