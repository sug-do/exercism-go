package etl

import (
	"strings"
)

//Transform takes a map of one type and converts it to another type
func Transform(inData map[int][]string) map[string]int {
	outData := make(map[string]int)
	for key, value := range inData {
		for _, letter := range value {
			outData[strings.ToLower(letter)] = key
		}
	}
	return outData
}
