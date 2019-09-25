package scrabble

import (
	"strings"
)

// Score will add up the value assigned to each letter of a word and return the value
func Score(wordIn string) int {
	allPossible := [][]rune{
		[]rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'},
		[]rune{'D', 'G'},
		[]rune{'B', 'C', 'M', 'P'},
		[]rune{'F', 'H', 'V', 'W', 'Y'},
		[]rune{'K'},
		[]rune{'J', 'X'},
		[]rune{'Q', 'Z'},
	}
	totalScore := 0
	wordInR := []rune(strings.ToUpper(wordIn))
	for _, letter := range wordInR {
		for i, runedList := range allPossible {
			for _, letterList := range runedList {
				if letterList == letter {
					switch i {
					case 0:
						totalScore++
					case 1:
						totalScore += 2
					case 2:
						totalScore += 3
					case 3:
						totalScore += 4
					case 4:
						totalScore += 5
					case 5:
						totalScore += 8
					case 6:
						totalScore += 10
					}
				}
			}
		}
	}
	return totalScore
}
