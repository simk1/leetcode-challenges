// Problem: https://leetcode.com/problems/middle-of-the-linked-list/description/
// Given the head of a singly linked list, return the middle node of the linked list. If there are two middle nodes, return the second middle node.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    tortoise := head
    bunny := head

    for bunny != nil && bunny.Next != nil {
        tortoise = tortoise.Next
        bunny = bunny.Next.Next
    }

    return tortoise
}