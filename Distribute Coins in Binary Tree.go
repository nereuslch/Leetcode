package Leetcode


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	var ans int = 0
	balance(root,&ans)
	return ans
}

func balance(root *TreeNode,ans *int) int {
	if root == nil {
		return 0
	}
	left := balance(root.Left,ans)
	right := balance(root.Right,ans)

	*ans += abs(left) + abs(right)

	return root.Val - 1 + left + right
}

func abs(a int) int {
	if a >0 {
		return a
	}

	return -1 * a
}