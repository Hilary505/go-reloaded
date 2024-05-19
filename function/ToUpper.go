package function

import (
	"strconv"
	"strings"
)
//function to uppercase the previous  word
func ToUpper(s string) string {
	contents := strings.Fields(s)
	for i, word := range contents {
		if word == "(up)"  { 
			if i-1 >= 0 {
			contents[i-1] = strings.ToUpper(contents[i-1])
		}
			contents = append(contents[:i], contents[i+1:]...)
		}
//function to uppercase the previous  word depending on the specifications
		if word == "(up," {
			myWd := strings.Trim(contents[i+1], contents[i+1][1:])
			num, _ := strconv.Atoi(myWd)
			for j := 0; j <= num; j++ {
				contents[i-j] = strings.ToUpper(contents[i-j])
			}
			contents = append(contents[:i], contents[i+2:]...)
		}
	}
	return strings.Join(contents, " ")
}
