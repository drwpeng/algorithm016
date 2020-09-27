package main

import "math"

func isValidBST(root *TreeNode) bool {
	head := math.MinInt64
	for root != nil {
		if root.Left != nil {
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else {
				if root.Val > head {
					head = root.Val
				} else{
					return false
				}
				predecessor.Right = nil
				root = root.Right
			}
		} else { // 没有左子树
			if root.Val > head {
				head = root.Val
			} else{
				return false
			}
			root = root.Right
		}
	}
	return true
}