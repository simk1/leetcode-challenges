# Problem: https://leetcode.com/problems/repeated-substring-pattern/description/
# Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

class Solution(object):
    def repeatedSubstringPattern(self, s):
        """
        :type s: str
        :rtype: bool
        """
        ls = len(s)
        if ls <= 1:
            return False
        elif s[:ls / 2] == s[ls / 2:]:
            return True

        pattern = s[0]
        i = 1
        for i in range(1, ls / 2):
            if s[i:i + len(pattern)] == pattern and ls % len(pattern) == 0:
                break
            pattern = pattern + s[i]

        if len(s) % len(pattern) != 0:
            return False

        i = 0
        l = len(pattern)
        for i in range(0, ls, l):
            if i > ls:
                return False
            if pattern != s[i:i + l]:
                return False
        return True