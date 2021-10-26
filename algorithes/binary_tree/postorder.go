package binding_tree

import "reflect"

func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var postorder func(root *TreeNode)

	reflect.DeepEqual()
	postorder = func(root *TreeNode) {
		if root != nil {
			postorder(root.Left)
			postorder(root.Right)
			ret = append(ret, root.Val)
		}
	}

	postorder(root)

	return ret
}
