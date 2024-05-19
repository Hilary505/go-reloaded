package function

import (

	"regexp"
)

func Pun(text string) string {

	punctuation := `[.,!?;:]`
	groupPunctuation := `[..]|[!?]|[,,]|[! :]`

	// Replace punctuations followed by space with just the punctuation
	text = regexp.MustCompile(`(^\w)`+punctuation+`\s`).ReplaceAllString(text, "$1$2")

	// Replace punctuations preceded by space with just the punctuation
	text = regexp.MustCompile(`\s(`+punctuation+`)(\w)`).ReplaceAllString(text, "$1 $2")

	// Format group punctuations with no space after them
	text = regexp.MustCompile(`\s(`+groupPunctuation+`)`).ReplaceAllString(text, "$1")

	return text
}
