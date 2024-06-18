// Problem: https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
// Given the head of a linked list, remove the nth node from the end of the list and return its head.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    listLen := 0
    for current := head; current != nil; current = current.Next {
		listLen++
	}
    if listLen == n {
        return head.Next
    }
    prev := head
    for i := 0; i < listLen - n - 1; i++ {
        prev = prev.Next
    }

    if prev.Next == nil {
        return nil
    }
    toBeRemoved := prev.Next
    prev.Next = toBeRemoved.Next
    toBeRemoved = nil
    return head
}