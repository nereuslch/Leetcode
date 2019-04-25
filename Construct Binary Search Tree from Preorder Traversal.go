package Leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var (
		rightidx int
	)

	for rightidx = 0; rightidx < len(preorder); rightidx ++ {
		if preorder[rightidx] > preorder[0] {
			break
		}
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:rightidx]),
		Right: bstFromPreorder(preorder[rightidx:]),
	}
}
