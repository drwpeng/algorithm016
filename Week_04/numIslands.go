package main

func numIslands(grid [][]byte) int {
	islands := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '1' {
				islands++
				dfs(grid, r, c)
			}
		}
	}
	return islands
}
func dfs(grid [][]byte, r, c int) {
	nr := len(grid)
	nc := len(grid[0])

	grid[r][c] = '0'
	if r + 1 < nr && grid[r+1][c] == '1' {
		dfs(grid, r+1, c)
	}
	if r - 1 >= 0 && grid[r-1][c] == '1' {
		dfs(grid, r-1, c)
	}
	if c + 1 < nc && grid[r][c + 1] == '1' {
		dfs(grid, r, c+1)
	}
	if c - 1 >= 0 && grid[r][c-1] == '1' {
		dfs(grid, r, c-1)
	}
}