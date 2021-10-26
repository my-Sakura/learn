package main

func recoverTree(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	var x, y, pre *TreeNode

	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && root.Val < pre.Val {
				y = root
				if x == nil {
					x = pre
				} else {
					break
				}
			}

			pre = root
			root = root.Right
		}
	}

	x.Val, y.Val = y.Val, x.Val
}
