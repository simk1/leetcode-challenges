# LeetCode Solutions

This repository contains solutions to various LeetCode problems that I solve as part of my daily coding practice. The problems are categorized by difficulty programming languages.

## How to Use

Each solution file includes:
- A brief description of the problem.
- The solution implemented in the given language.

## Why LeetCode?

Solving LeetCode problems helps in sharpening problem-solving skills and understanding different data structures and algorithms. Regular practice on these problems prepares for technical interviews and improves coding proficiency.

## Examples

### Two Sum

```go
// Problem: https://leetcode.com/problems/two-sum/
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

package main

import "fmt"

func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if idx, found := hashMap[complement]; found {
            return []int{idx, i}
        }
        hashMap[num] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println(result)  // Output: [0, 1]
}
```
Feel free to browse through the solutions and contribute by suggesting improvements or adding new solutions.
