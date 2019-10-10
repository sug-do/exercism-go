package isogram

import "unicode"

// IsIsogram will return true if a word is an isogram and false if it isn't.
func IsIsogram(wordIn string) bool {
	isoMap := make(map[rune]bool)
	for _, value := range wordIn {
		if unicode.IsLetter(value) {
			upValue := unicode.ToUpper(value)
			if _, exist := isoMap[upValue]; exist {
				return false
			}
			isoMap[upValue] = false
		}
	}
	return true
}
