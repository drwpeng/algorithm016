package Week_09

func longestPalindrome(s string) string {
	var index int = 0
	var size int = len(s)
	var maxLen = 0
	var res = ""
	for index < size{
		var pointerL = index
		var pointerR = index
		for pointerL >= 0 && s[pointerL] == s[index]{
			pointerL--
		}
		for pointerR < size && s[pointerR] == s[index]{
			pointerR++
		}
		var nextPoint = pointerR
		for pointerL >= 0 && pointerR < size && s[pointerR] == s[pointerL]{
			pointerL--
			pointerR++
		}
		var tmpMaxLen = pointerR - pointerL - 1
		if tmpMaxLen > maxLen{
			res = s[pointerL + 1:pointerR]
			maxLen = tmpMaxLen
		}

		index = nextPoint
	}
	return res
}

