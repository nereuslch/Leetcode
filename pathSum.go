package Leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return pathSum(root, 0, sum)

}

func pathSum(node *TreeNode, cur, sum int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return cur+node.Val == sum
	}

	return pathSum(node.Left, cur+node.Val, sum) || pathSum(node.Right, cur+node.Val, sum)

}
