package accumulate

// Accumulate will pass the word to the function for conversion and return the end result of the
// appended result in a string slice format
func Accumulate(testIn []string, convIn func(string) string) []string {
	var output []string
	for _, wordP := range testIn {
		output = append(output, convIn(wordP))
	}
	return (output)
}
