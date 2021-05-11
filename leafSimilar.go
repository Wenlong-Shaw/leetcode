package leetcode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafList1 := make([]int, 0)
	leafList2 := make([]int, 0)
	traverseBT(root1, &leafList1)
	traverseBT(root2, &leafList2)
	if len(leafList1) != len(leafList2) {
		return false
	}
	for i := range leafList1 {
		if leafList1[i] != leafList2[i] {
			return false
		}
	}
	return true
}

func traverseBT(root *TreeNode, leafList *([]int)) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leafList = append(*leafList, root.Val)
	}
	traverseBT(root.Left, leafList)
	traverseBT(root.Right, leafList)
	return
}
