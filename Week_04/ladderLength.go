package main

import "fmt"

func main() {
	wl := []string{"hot","dot","dog","lot","log","cog"}
	n := ladderLength("cog", "hit", wl)
	fmt.Println(n)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wdict, used := make(map[string][]int), make(map[int]bool)
	level := 0
	for key, val := range wordList {
		for i := range val {
			w := val[0:i] + "*" +val[i+1:]
			if _, ok := wdict[w]; !ok {
				wdict[w] = []int{}
			}
			wdict[w] = append(wdict[w], key)
		}
	}
	wordList = append(wordList, beginWord)
	q := []int{len(wordList)-1}
	for len(q) != 0 {
		nextQ := []int{}
		level++
		for _, qid := range q {
			word := wordList[qid]
			if word == endWord {
				return level
			}
			for i, _ := range word {
				w := word[0:i] + "*" + word[i+1:]
				for _, id := range wdict[w] {
					if !used[id] {
						used[id] = true
						nextQ = append(nextQ, id)
					}
				}
			}
		}
		q = nextQ
	}
	return 0
}

//
//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	wdict, usedbegin, usedend := make(map[string][]int), make(map[int]bool), make(map[int]bool)
//	for key, val := range wordList {
//		for i := range val {
//			w := val[0:i] + "*" +val[i+1:]
//			if _, ok := wdict[w]; !ok {
//				wdict[w] = []int{}
//			}
//			wdict[w] = append(wdict[w], key)
//		}
//	}
//	wordList = append(wordList, beginWord, endWord)
//	begin := []int{len(wordList)-2}
//	end := []int{len(wordList)-1}
//	for len(begin) > 0 && len(end) > 0 {
//		bn := len(begin)- 1
//		en := len(end) - 1
//		if begin[bn] == end[en] {
//			return bn + en + 1
//		}
//		if len(begin) > len(end) {
//			end = help(end, wordList, wdict, usedend)
//		} else {
//			begin = help(begin, wordList, wdict, usedbegin)
//		}
//	}
//	return 0
//}
//func help(begin []int, wordList []string, wdict map[string][]int, used map[int]bool) []int {
//	nextQ := []int{}
//	for _, id := range begin {
//		word := wordList[id]
//		for i := range word {
//			w := word[0:i] + "*" + word[i+1:]
//			for _, key := range wdict[w] {
//				if !used[key] {
//					used[key] = true
//					nextQ = append(nextQ, key)
//				}
//			}
//		}
//	}
//	return nextQ
//}
