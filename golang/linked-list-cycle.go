// Problem: https://leetcode.com/problems/linked-list-cycle/description/
// Given head, the head of a linked list, determine if the linked list has a cycle in it.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    list := make(map[*ListNode]int)
    for head != nil {
        if _, exists := list[head]; exists {
            return true
        } else {
            list[head] = 1
            head = head.Next
        }
    }
    return false
}