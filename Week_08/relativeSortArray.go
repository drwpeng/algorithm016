package Week_08

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	hash := make(map[int]int)
	notexist := make([]int, 0)
	res := make([]int, 0)
	for _, v := range arr2 {
		hash[v] = 0
	}
	for _, v := range arr1 {
		if _, ok := hash[v]; ok {
			hash[v]++
		} else {
			notexist = append(notexist, v)
		}
	}

	if len(notexist) > 1 {
		sort.Ints(notexist)
	}
	for _, v := range arr2 {
		for i := 0; i < hash[v]; i++ {
			res = append(res, v)
		}
	}
	res = append(res, notexist...)
	return res
}
