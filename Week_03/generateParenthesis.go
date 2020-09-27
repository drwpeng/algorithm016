package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var res []string
	generate(0, 0, n, "", &res)
	return res
}

func generate(left, right, max int, s string, res *[]string) {
	if right == max {
		*res = append(*res, s)
		return
	}
	if left < max {
		generate(left+1, right, max, s+"(", res)
	}
	if right < left {
		generate(left, right+1, max, s+")", res)
	}
}



func main() {
	s := generateParenthesis(3)
	fmt.Println(s)
}
