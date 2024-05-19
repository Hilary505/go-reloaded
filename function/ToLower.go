package function

import (
	"strconv"
	"strings"
)
//function to lowercase the previous  word
func ToLower(s string) string {
	contents := strings.Fields(s)
	for i, word := range contents {
		if word == "(low)" {
			if i-1 >= 0 {
				contents[i-1] = strings.ToLower(contents[i-1])
			}
			contents = append(contents[:i], contents[i+1:]...)
		}
//function to lowercase the previous  word depending on specifications
		if word == "(low," {
			newWord := strings.Trim(contents[i+1], contents[i+1][1:])
			number, _ := strconv.Atoi(newWord)
			for j := 1; j <= number && i-j >= 0; j++ {
				if i-j >= 0 {
					contents[i-j] = strings.ToLower(contents[i-j])
				}
			}
			contents = append(contents[:i], contents[i+2:]...)
		}
	}
	return strings.Join(contents, " ")
}
