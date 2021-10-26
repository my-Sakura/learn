package binary_tree

// recursion
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root != nil {
			inorder(root.Left)
			ret = append(ret, root.Val)
			inorder(root.Right)
		}
	}
	inorder(root)

	return ret
}

// iteration
// func inorderTraversal(root *TreeNode) []int {
//     ret := make([]int, 0)
//     stack := make([]*TreeNode, 0)
//     for root != nil || len(stack) != 0 {
//         if root != nil {
//             stack = append(stack, root)
//             root = root.Left
//         } else {
//             root = stack[len(stack) - 1]
//             stack = stack[:len(stack) - 1]
//             ret = append(ret, root.Val)
//             root = root.Right
//         }
//     }

//     return ret
// }
