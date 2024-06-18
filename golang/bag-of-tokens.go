// Problem: https://leetcode.com/problems/bag-of-tokens/description/
// You start with an initial power of power, an initial score of 0, and a bag of tokens given as an integer array tokens, where each tokens[i] denotes the value of token_i.

func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    c_pow := power
    left_token := 0
    right_token := len(tokens) - 1
    score := 0
    maxScore := 0

    for left_token <= right_token {
        if c_pow >= tokens[left_token] {
            c_pow -= tokens[left_token]
            left_token++
            score++
            if score > maxScore {
                maxScore = score
            }
        } else if score > 0 {
            c_pow += tokens[right_token]
            right_token--
            score--
        } else {
            break
        }
    }
    return maxScore
}