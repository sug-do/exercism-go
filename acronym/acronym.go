package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Abbreviate will take words and turn them into acronyms

func Abbreviate(s string) string {
	// use regexp to filter for valid word characters. include one space as valid
	reg, err := regexp.Compile("[^a-zA-Z'\\s]+")
	// log if error
	if err != nil {
		log.Fatal(err)
	}
	// use the regexp to replace any characters not matching the regexp pattern
	filteredString := reg.ReplaceAllString(s, " ")
	// declare variables before loop
	var w, letters string = "", ""
	// turn string into slice so that we can deal with each word
	var splitWord = strings.Fields(filteredString)
	// begin loop and feed in each word of the array with range into w
	for _, w = range splitWord {
		// take the 1st letter of the word assigned to w, convert it to string, capitalize, and add each consecutive to the string "letters"
		letters += strings.ToUpper(string(w[0]))
	}
	// return letters as the string value to the main function as the desired acronym
	return letters
}
