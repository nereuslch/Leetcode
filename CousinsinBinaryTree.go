package Leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	var queue, tmp []*TreeNode
	var existx, existy bool

	queue = append(queue, root)
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]

		if top.Left != nil && top.Right != nil {
			if (top.Left.Val == x && top.Right.Val == y) ||
				(top.Left.Val == y && top.Right.Val == x) {
				return false
			}
		}

		if top.Left != nil {
			if x == top.Left.Val {
				existx = true
			}
			if y == top.Left.Val {
				existy = true
			}
			tmp = append(tmp, top.Left)
		}

		if top.Right != nil {
			if x == top.Right.Val {
				existx = true
			}
			if y == top.Right.Val {
				existy = true
			}
			tmp = append(tmp, top.Right)
		}

		if existx && existy {
			return true
		}

		if len(queue) == 0 {
			queue = append(queue, tmp...)
			tmp = tmp[:0]
			existy, existx = false, false
		}
	}
	return false
}
