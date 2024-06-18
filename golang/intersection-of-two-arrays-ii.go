// Problem: https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

func intersect(nums1 []int, nums2 []int) []int {
    intersection := []int{}
	memo := make(map[int]int)

	for _, num := range nums1 {
		memo[num]++
	}

    for _, num := range nums2 {
        if memo[num] > 0 {
            intersection = append(intersection, num)
            memo[num]--
        }
    }
    return intersection
}