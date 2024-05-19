package function

// function to check for articles "an" and "An" before vowels
func CheckVowel(s []string) []string {

	// Define a slice that contains all Uppercase and Lowercase vowels including "h and H"
	vowel := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	// loop through each element (val) and its index (i) in the input slice (s)
	for i, val := range s {
		for _, letter := range vowel {
			if val == "a" && string(s[i+1][0]) == letter {
				// replace "a" with "an" before a vowel
				s[i] = "an"
			} else if val == "A" && string(s[i+1][0]) == letter {
				// Replace "A" with "An" before a vowel
				s[i] = "An"
			}
		}
	}
	return s
}
