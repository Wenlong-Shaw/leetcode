package leetcode

/*
236. 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，
最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1：
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1  【注：数组为层次遍历结果】
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree

*/
//TODO: 思路：由于 p != q; 因此，p，q的最近父节点可以收到左、右子树搜索到p，q的节点信息。
//TODO: 于是，只需要比对第一
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//* 找到p，或q之后立马返回该点。
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	//* 接着遍历左右子树，是否有返回目的节点信息
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//* 搜索到了p，q节点，左、右子树就不会为空
	if left != nil && right != nil {
		return root
	}
	//* 没有搜索到节点就会是在另一边，如果另一边也是nil，则应返回nil，可以直接返回另一边的nil。
	if left == nil {
		return right
	}
	return left

}
