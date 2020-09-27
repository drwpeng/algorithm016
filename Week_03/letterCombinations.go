package main

var phone = map[string]string{
	"2":"abc",
	"3":"def",
	"4":"ghi",
	"5":"jkl",
	"6":"mno",
	"7":"pqrs",
	"8":"tuv",
	"9":"wxyz",
}
var res []string
func letterCombinations(digits string) []string {
	res = []string{}
	if digits == ""{
		return res
	}
	combinations(digits, 0, "")
	return res
}

func combinations(digits string, index int, combina string) {
	if index == len(digits) {
		res = append(res, combina)
		return
	} else {
		ch := digits[index]
		letters := phone[string(ch)]
		for i := 0; i < len(letters); i++ {
			combinations(digits, index+1, combina+string(letters[i]))
		}
	}
}

func main() {
	letterCombinations("23")
}
