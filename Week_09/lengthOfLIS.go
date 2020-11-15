package Week_09

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	res := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
