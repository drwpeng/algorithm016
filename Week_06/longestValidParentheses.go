package main

//最长有效括号

//动态规划
func longestValidParentheses(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i - 2] + 2
				} else {
					dp[i] = 2
				}
			} else if i - dp[i - 1] > 0 && s[i - dp[i - 1] - 1] == '(' {
				if i - dp[i - 1] >= 2 {
					dp[i] = dp[i - 1] + dp[i - dp[i - 1] - 2] + 2
				} else {
					dp[i] = dp[i - 1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}

//双向遍历求最大值，单向遍历回漏掉某些值 比如："(()"
func longestValidParentheses2(s string) int {
	left, right, lmax, rmax := 0, 0, 0, 0
	for _, v := range s {
		if v == '(' {
			left++
		} else {
			right++
		}
		if right == left  {
			if lmax < 2 * left{
				lmax = 2 * left
			}
		} else if left < right{
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if left == right {
			if rmax < 2 * right {
				rmax = 2 * right
			}
		}else if left > right{
			left, right = 0, 0
		}
	}
	return max(lmax, rmax)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
