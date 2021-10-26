package binary_tree

// func maxPathSum(root *TreeNode) int {
// 	maxSum := math.MinInt32
// 	var maxPath func(*TreeNode) int
// 	maxPath = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}

// 		leftGain := max(maxPath(node.Left), 0)
// 		rightGain := max(maxPath(node.Right), 0)

// 		priceNewPath := node.Val + leftGain + rightGain

// 		maxSum = max(maxSum, priceNewPath)

// 		return node.Val + max(leftGain, rightGain)
// 	}
// 	maxPath(root)
// 	return maxSum
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}

// 	return a
// }
