package leetcode

/*
* 105. 从前序与中序遍历序列构造二叉树
* 根据一棵树的前序遍历与中序遍历构造二叉树。
* 注意:
* 你可以假设树中没有重复的元素。
 */

//TODO: 思路: 对于任意一颗树而言，
//* 前序遍历的形式总是
//* [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
//* 即根节点总是前序遍历中的第一个节点。
//* 而中序遍历的形式总是
//* [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

/*
* 106. 从中序与后序遍历序列构造二叉树
* 根据一棵树的中序遍历与后序遍历构造二叉树。

* 注意:
* 你可以假设树中没有重复的元素。

* 例如，给出

* 中序遍历 inorder = [9,3,15,20,7]
* 后序遍历 postorder = [9,15,7,20,3]
* 返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/

//TODO： 思路同上极其类似
func buildTree_1(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}

	root.Left = buildTree(inorder[:i], postorder[:len(inorder[:i])])
	root.Right = buildTree(inorder[i+1:], postorder[len(inorder[:i])+1:len(postorder)-1])

	return root
}
