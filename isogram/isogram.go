package isogram

import (
	"strings"
)

// IsIsogram will return true if a word is an isogram and false if it isn't.
func IsIsogram(wordIn string) bool {
	wordInUp := strings.ToUpper(wordIn)
	for _, letter := range wordInUp {
		if letter >= 65 && letter <= 90 {
			matchLet := 0
			for _, wordLet := range wordInUp {
				switch wordLet == letter {
				case true:
					matchLet++
				}
				if matchLet > 1 {
					return false
				}
			}
		}
	}
	return true
}
