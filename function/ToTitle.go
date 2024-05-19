package function

import (
	"strconv"
	"strings"
)

func ToTitle(s string) string {
	arr := []rune(s)
	for i, c := range arr {
		if isNumber(c) {
			if i == 0 || !isNumber(c) {
				if c >= 'a' && c <= 'z' {
					arr[i] = c - 32
				}
			}
		} else {
			if c >= 'A' && c <= 'Z' {
				arr[i] = c + 32
			}
		}
	}
	return string(arr)
}
func isNumber(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}
//function to capitalise the first letter of  a word
func Title(s string) string {
	contents := strings.Fields(s)
	for i, word := range contents {
		if word == "(cap)" {
			if i-1 >= 0 {
				contents[i-1] = ToTitle(contents[i-1])
			}
			contents = append(contents[:i], contents[i+1:]...)
		}
//function to capitalise the first letter of  a word depending on specifications
		if word == "(cap," {
			newWord := strings.Trim(string(contents[i+1]), contents[i+1][1:])
			number, _ := strconv.Atoi(string(newWord))
			for j := 1; j <= number; j++ {
				contents[i-j] = ToTitle(contents[i-j])
			}
			contents = append(contents[:i], contents[i+2:]...)
		}

	}
	return strings.Join(contents, " ")
}
