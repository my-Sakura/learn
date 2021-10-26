package binary_tree

// recursion
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var preorder func(root *TreeNode)

	preorder = func(root *TreeNode) {
		if root != nil {
			ret = append(ret, root.Val)
			preorder(root.Left)
			preorder(root.Right)
		}
	}

	preorder(root)

	return ret
}

// iteration
// func preorderTraversal(root *TreeNode) []int {
//     ret := make([]int, 0)
//     stack := make([]*TreeNode, 0)

//     for root != nil || len(stack) != 0 {
//         if root != nil {
//             ret = append(ret, root.Val)
//             stack = append(stack, root)
//             root = root.Left
//         } else {
//             root = stack[len(stack) - 1]
//             stack = stack[:len(stack) - 1]
//             root = root.Right
//         }
//     }

//     return ret
// }
