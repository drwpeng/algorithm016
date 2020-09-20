package Week_02

func groupAnagrams(strs []string) [][]string {
	team := make(map[string]int)
	res := make([][]string, 0)

	for _, str := range strs {
		b := []byte{}
		count := make([]int, 26)
		for _, v := range str {
			count[v-'a']++
		}
		for i, v := range count {
			if v != 0 {
				b = append(b, byte(i)+'a', byte(v)+'0')
			}
		}
		t := string(b)
		if i, ok := team[t]; !ok {
			team[t] = len(res)
			res = append(res, []string{str})
		} else {
			res[i] = append(res[i], str)
		}
	}
	return res
}
