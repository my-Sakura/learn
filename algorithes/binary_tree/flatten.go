package binary_tree

func flatten(root *TreeNode) {
	list := make([]*TreeNode, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			list = append(list, root)
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}

	for i := 1; i < len(list); i++ {
		pre, next := list[i-1], list[i]
		pre.Left, pre.Right = nil, next
	}
}
