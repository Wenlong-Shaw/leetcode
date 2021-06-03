package leetcode

func maxDepth(root *TreeNode) int {
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		childdepthnum := max(depth(node.Left), depth(node.Right))
		depthnum := childdepthnum + 1
		return depthnum
	}
	return depth(root)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
