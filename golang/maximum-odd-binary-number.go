// Problem: https://leetcode.com/problems/maximum-odd-binary-number/description/
/*
You are given a binary string s that contains at least one '1'.

You have to rearrange the bits in such a way that the resulting binary number is the maximum odd binary number that can be created from this combination.

Return a string representing the maximum odd binary number that can be created from the given combination.

Note that the resulting string can have leading zeros.
*/

func maximumOddBinaryNumber(s string) string {
    countOnes := strings.Count(s, "1")
    zeros := ""
    ones := ""
    for i:=0; i < len(s) - countOnes; i++ {
        zeros = zeros + "0"
    }
    for i:=0; i < countOnes-1; i++ {
        ones = ones + "1"
    }
    return ones + zeros + "1"
}