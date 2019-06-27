package Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow,fast := head,head.Next.Next
	for  {
		if slow == fast {
			return true
		}

		slow = slow.Next
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		fast = fast.Next.Next

	}

	return false

}