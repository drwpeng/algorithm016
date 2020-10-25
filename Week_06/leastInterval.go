package main

import "fmt"

func main() {
	tasks := []byte{'A','A','A','B','B','B'}
	out := leastInterval(tasks, 2)
	fmt.Println(out)
}

func leastInterval(tasks []byte, n int) int {
	queue := [26]int{0}
	max, last := 0, 0

	for _, v := range tasks {
		queue[v-'A']++
		if max < queue[v-'A'] {
			max = queue[v-'A']
			last = 1
		} else if max == queue[v-'A'] {
			last++
		}
	}
	if (max-1) * (n+1) + last < len(tasks) {
		return len(tasks)
	}

	return (max-1) * (n + 1) + last
}