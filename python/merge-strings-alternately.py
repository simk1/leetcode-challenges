# Problem: https://leetcode.com/problems/merge-strings-alternately/description/
"""
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.
"""

class Solution(object):
    def mergeAlternately(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: str
        """
        l1 = len(word1)
        l2 = len(word2)

        word3 = ""
        for i in range(0, min(l1, l2)):
            word3 = word3 + word1[i] + word2[i]

        if l1 > l2:
            word3 = word3 + word1[i + 1:]
        elif l2 > l1:
            word3 = word3 + word2[i + 1:]

        return word3