package Week_07

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict, used := make(map[string][]string), make(map[string]bool)
	levels := 0
	for _, word := range wordList {
		for i := range word{
			val := word[:i] + "*" + word[i+1:]
			dict[val] = append(dict[val], word)
		}
	}

	//广度优先
	queue := []string{beginWord}
	for len(queue) > 0 {
		levels++
		next := []string{}
		for _, word := range queue {
			if word == endWord {
				return levels
			}
			for i := range word {
				val := word[:i] + "*" + word[i+1:]
				if _, ok := dict[val]; ok {
					for _, v := range dict[val] {
						if _, ok := used[v]; !ok {
							next = append(next, v)
							used[v] = true
						}
					}
				}
			}
		}

		queue = next
	}
	return 0
}
