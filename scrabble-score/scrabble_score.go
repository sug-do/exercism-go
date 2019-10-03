package scrabble

import "unicode"

// Score will add up the value assigned to each letter of a word and return the value
func Score(wordIn string) int {
	totalScore := 0
	for _, letter := range wordIn {
		switch unicode.ToUpper(letter) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			totalScore++
		case 'D', 'G':
			totalScore += 2
		case 'B', 'C', 'M', 'P':
			totalScore += 3
		case 'F', 'H', 'V', 'W', 'Y':
			totalScore += 4
		case 'K':
			totalScore += 5
		case 'J', 'X':
			totalScore += 8
		case 'Q', 'Z':
			totalScore += 10

		}
	}
	return totalScore
}
