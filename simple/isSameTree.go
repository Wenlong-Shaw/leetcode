package leetcode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	ans := true
	var dfs func(p *TreeNode, q *TreeNode)
	dfs = func(p *TreeNode, q *TreeNode) {
		if p == nil && q == nil {
			return
		}
		if p != nil && q != nil {
			if p.Val != q.Val {
				ans = false
				return
			}
			dfs(p.Left, q.Left)
			dfs(p.Right, q.Right)
		}
		if (p != nil && q == nil) || (p == nil && q != nil) {
			ans = false
		}
		return
	}
	dfs(p, q)
	return ans
}

// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	if p == nil && q == nil {
// 		return true
// 	}
// 	if p != nil && q != nil {
// 		if p.Val == q.Val {
// 			return true
// 		}
// 	}
// 	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
// }
