func climbStairs(n int) int {
    if n < 1 {
        return 0
    }
    pre, cur := 0, 1
    for i := 0; i < n; i++ {
        cur, pre = cur + pre, cur
    }
    return cur
}