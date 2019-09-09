package raindrops

import (
	"strconv"
	"strings"
)

// Convert will check if a number has 3, 5, or 7 as a factor
func Convert(inNum int) string {
	// define the strings for each factorial
	output := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	// create the variable for building the string result
	var message strings.Builder
	// integer array of factorials to use
	factNum := []int{3, 5, 7}
	// create variable used to count instances of non-factorial
	var notFact int
	for _, num := range factNum {
		// add string if factorial
		if inNum%num == 0 {
			message.WriteString(output[num])
		} else {
			// increase notFact if not factorial
			notFact++
		}
	}
	// return number if no factorial match
	if notFact == 3 {
		return strconv.Itoa(inNum)
	}
	// return the combined string if factorial match(es) exist
	return message.String()
}
