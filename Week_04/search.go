package main

import "fmt"

func main() {
	k := search([]int{4,5,6,7,0,1,2}, 0)
	fmt.Println(k)
}

//搜索旋转排序数组
func search(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n-1

	for low <= high {
		if nums[low] == target {
			return low
		}
		if nums[high] == target {
			return high
		}
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[n-1] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}