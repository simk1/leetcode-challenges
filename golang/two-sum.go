// Problem: https://leetcode.com/problems/two-sum/description/
/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
    res := []int{}
    for ix := 0; ix < len(nums)-1; ix++ {
        for jx := ix+1; jx < len(nums); jx++ {
            if nums[ix] + nums[jx] == target {
                res = append(res, ix, jx)
                break
            }
        }
    }
    return res
}