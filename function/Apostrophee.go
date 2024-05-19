package function 

import (
     "regexp"
	 "strings"
)

func Punc(text string) string {
	pattern := `'(\s*.*?\s*)'`

    // Compile the regex pattern
    match := regexp.MustCompile(pattern)

    // Find all matches in the input string
    matches := match.FindAllStringSubmatch(text, -1)

    // Iterate over matches
    for _, match := range matches {
        // Remove leading and trailing whitespace
        word := strings.TrimSpace(match[1])

        // Replace the match in the input string
        text = strings.Replace(text, match[0], "'"+word+"'", 1)
    }
	return text
}