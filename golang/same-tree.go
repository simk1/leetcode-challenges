// Problem: https://leetcode.com/problems/same-tree/description/
/*
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
    /*if p == nil && q == nil {
        return true
    } else if p != nil && q == nil || p == nil && q != nil {
        return false
    } else {
        if p.Val != q.Val {
            return false
        }
        return isSameTree(p.Left, q.Left)
        return isSameTree(p.Right, q.Right)
    }
    return true
    */
    var result bool = false

    if p == nil && q == nil {
        result = true
    } else if p != nil && q != nil {
        if p.Val == q.Val {
            result = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
        }
    }

    return result
}