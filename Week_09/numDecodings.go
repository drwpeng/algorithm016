package Week_09

func numDecodings(s string) int {
	pre, cur := 1, 1
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}
			cur = pre
		} else {
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
				cur, pre = cur + pre, cur
			} else {
				pre = cur
			}
		}
	}
	return cur
}

