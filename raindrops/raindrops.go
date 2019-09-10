package raindrops

import (
	"strconv"
	"strings"
)

// Convert will check if a number has 3, 5, or 7 as a factor
// define the strings for each factorial. then create the variable for building the string result
// integer array of factorials to use. then create variable used to count instances of non-factorial.
// finally, add string if factorial. and, return the combined string if factorial match(es) exist
// return original number if no factorial match
func Convert(inNum int) string {
	output := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	var message strings.Builder
	factNum := []int{3, 5, 7}
	for _, num := range factNum {
		if inNum%num == 0 {
			message.WriteString(output[num])
		}
	}
	if message.String() == "" {
		return strconv.Itoa(inNum)
	}
	return message.String()
}
