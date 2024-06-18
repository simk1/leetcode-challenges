// Problem: https://leetcode.com/problems/squares-of-a-sorted-array/description/
// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

import "sort"

func sortedSquares(nums []int) []int {
    var squares []int
    for _, v := range nums {
        squares = append(squares, v*v)
    }
    sort.Ints(squares)
    return squares
}