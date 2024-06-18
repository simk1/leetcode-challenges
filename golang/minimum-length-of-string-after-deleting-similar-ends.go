// Problem: https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/description/
/*
Given a string s consisting only of characters 'a', 'b', and 'c'. You are asked to apply the following algorithm on the string any number of times:

Pick a non-empty prefix from the string s where all the characters in the prefix are equal.
Pick a non-empty suffix from the string s where all the characters in this suffix are equal.
The prefix and the suffix should not intersect at any index.
The characters from the prefix and suffix must be the same.
Delete both the prefix and the suffix.
Return the minimum length of s after performing the above operation any number of times (possibly zero times).
*/

func minimumLength(s string) int {
    if len(s) == 2 {
        if s[0] != s[1] {
            return 2
        } else {
            return 0
        }
    } else if len(s) <= 1 {
        return len(s)
    }
    first := 0
    last := len(s) - 1

    for first < last {
        if s[first] != s[last] {
            return last - first + 1
        }
        for first < last && s[first] == s[first+1] {
            first++
        }
        for last > first && s[last] == s[last-1] {
            last--
        }
        if first == last {
            return 0
        }
        first++
        last--
    }
    return last - first + 1
}