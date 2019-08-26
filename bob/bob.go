package bob

import (
	"regexp"
	"strings"
)

var IsLetter = regexp.MustCompile(`[a-zA-Z]+`).MatchString
var IsNumber = regexp.MustCompile(`[0-9]`).MatchString
var IsValid = regexp.MustCompile(`[0-9a-zA-Z?!.]+`).MatchString
var w string

func Hey(s string) string {
	w := ""
	switch {
	//determine if a valid string
	case IsValid(s):
		//trim whitespace from valid string
		s := strings.TrimSpace(s)
		switch {
		//
		case !(s[len(s)-1:] == "?"):
			switch {
			case strings.ToUpper(s) == s && IsLetter(s):
				w = "Whoa, chill out!"
			default:
				w = "Whatever."
			}
		case s[len(s)-1:] == "?":
			switch {
			case IsNumber(s):
				w = "Sure."
			case strings.ToUpper(s) == s && IsLetter(s):
				w = "Calm down, I know what I'm doing!"
			default:
				w = "Sure."
			}
		}
	default:
		w = "Fine. Be that way!"
	}
	return w
}
