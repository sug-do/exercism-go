package hamming

import (
	"errors"
	"strings"
)

var err error

// Distance will compare each letter of strings of equal length and record the number of
// letters that don't match in each "slot"
func Distance(a, b string) (int, error) {
	diff := 0
	// checking if both strings match first nets performance gain. keeping.
	if a == b {
		return diff, err
	}
	// make sure length matches before comparing
	if len(a) == len(b) {
		// convert each string to slice to compare letter for letter
		aSplit := strings.Split(a, "")
		bSplit := strings.Split(b, "")
		// compare each letter of the given converted strings of a and b
		for i := range aSplit {
			// increase counter if two corresponding letters from each slice don't match
			if aSplit[i] != bSplit[i] {
				diff++
			}
		}
		// give main function the result of the comparison
		return diff, err
	}
	// return an error to the main function when lengths don't match
	err = errors.New("length a (" + a + ") and b (" + b + ") DO NOT match")
	return diff, err
}
