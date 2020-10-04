package main

func main() {
	var matrix [][]int
	for i := 0; i < 2; i++{
		inline := make([]int, 1)
		matrix = append(matrix, inline)
	}
	matrix[0][0] = 1
	matrix[1][0] = 3
	searchMatrix(matrix, 2)
}
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if matrix == nil || n == 0 {
		return false
	}
	left, high := 0, len(matrix[0])*n - 1
	for left <= high {
		mid := (left + high) / 2
		temp := matrix[mid/2][mid%2]
		if temp == target {
			return true
		} else if temp >  target{
			high = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
