package main

func main() {
	ladderLength("hit", "cog", []string{"hot","dot","dog","lot","log","cog"})
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict, used := make(map[string][]string), make(map[string]bool)
	level := 0
	for _, v := range wordList {
		for j := 0; j < len(v); j++ {
			temp := v[:j] + "*" + v[j+1:]
			dict[temp] = append(dict[temp], v)
		}
	}
	fqueue := map[string]bool{beginWord: true}
	equeue := map[string]bool{endWord: true}
	for len(fqueue) > 0 || len(equeue) > 0 {
		level++
		next := make(map[string]bool)
		if len(fqueue) > len(equeue) {
			fqueue, equeue = equeue, fqueue
		}
		for word, _ := range fqueue {
			for i := 0; i < len(word); i++ {
				temp := word[:i] + "*" + word[i+1:]
				if _, ok := dict[temp]; !ok {
					continue
				}
				for _, v := range dict[temp] {
					if _, ok := used[v]; ok {
						continue
					}
					if _, ok := equeue[v]; ok {
						return level+1
					}
					next[v] = true
					used[v] = true
				}
			}
		}
		fqueue = next
	}
	return 0
}