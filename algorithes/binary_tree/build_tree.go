package binary_tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	var i int
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	leftLength := len(inorder[:i])
	root.Left = buildTree(preorder[1:leftLength+1], inorder[:leftLength])
	root.Right = buildTree(preorder[leftLength+1:], inorder[leftLength+1:])

	return root
}

// func buildTree(inorder []int, postorder []int) *TreeNode {
//     if len(inorder) == 0 {
//         return nil
//     }

//     root := &TreeNode{postorder[len(postorder) - 1 ], nil, nil}
//     var i int
//     for ; i < len(inorder); i++ {
//         if inorder[i] == postorder[len(postorder) - 1] {
//             break
//         }
//     }

//     root.Left = buildTree(inorder[:i], postorder[:i])
//     root.Right = buildTree(inorder[i + 1:], postorder[i:len(postorder) - 1])

//     return root
// }
