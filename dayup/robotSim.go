package main

import "fmt"

func main()  {
	com := []int{4, -1, 4, -2, 4}
	obs := [][]int{{2,4}}
	n := robotSim(com, obs)
	fmt.Println(n)
}

func robotSim(commands []int, obstacles [][]int) int {
	step := 0
	if len(obstacles) == 0 {
		for _, v := range commands {
			if v != -1 && v != -2 {
				step += v * v
			}
		}
		return step
	}
	type point struct {
		x int
		y int
	}
	obMap := make(map[point]bool)
	for _, v := range obstacles {
		obMap[point{v[0], v[1]}] = true
	}

	dir := 3   //东南西北 0123
	tx, ty := 0, 0
	for _, v := range commands {
		if v == -1 {
			dir = (dir + 1) % 4
		} else if v == -2 {
			dir = (dir + 3) % 4
		} else {
			temp := 0
			switch dir {
			case 0:
				for i := 1; i <= v; i++ {
					x := tx+1
					p := point{x, ty}
					if _, ok := obMap[p]; ok {
						break
					}
					temp++
					tx++
				}
			case 1:
				for i := 1; i <= v; i++ {
					y := ty - 1
					p := point{tx, y}
					if _, ok := obMap[p]; ok {
					}
					temp++
					ty--
				}
			case 2:
				for i := 1; i <= v; i++ {
					x := tx-1
					p := point{x, ty}
					if _, ok := obMap[p]; ok {
						break
					}
					temp++
					tx--
				}
			case 3:
				for i := 1; i <= v; i++ {
					y := ty + 1
					p := point{tx, y}
					if _, ok := obMap[p]; ok {
						break
					}
					temp++
					ty++
				}
			}
			step += temp * temp
		}
	}
	return tx * tx + ty * ty
}
