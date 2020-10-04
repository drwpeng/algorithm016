package main


func canJump(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}
	flag := true
	step, last := 0, 0
	for j := n-1; j >= 0; j-- {
		if !flag {
			step++
			if nums[j] > step {
				flag = true
				step = 0
			} else if nums[j] == step {
				if last == n-1 {
					step = 0
					flag = true
				}
			}
		}
		if nums[j] == 0 && flag {
			flag = false
			last = j
		}
	}

	return flag
}