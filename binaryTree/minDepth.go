package leetcode

func minDepth(root *TreeNode) int {
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return 1
		} else if node.Left != nil && node.Right != nil {
			return min(depth(node.Left), depth(node.Right)) + 1
		} else if node.Left != nil {
			return depth(node.Left) + 1
		} else if node.Right != nil {
			return depth(node.Right) + 1
		}
		return 0
	}
	return depth(root)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
