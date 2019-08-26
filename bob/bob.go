package bob

import (
	"regexp"
	"strings"
)

var IsLetter = regexp.MustCompile(`[a-zA-Z]+`).MatchString
var IsNumber = regexp.MustCompile(`[0-9]`).MatchString
var IsValid = regexp.MustCompile(`[0-9a-zA-Z?!.]+`).MatchString

func Hey(s string) string {
	if !IsValid(s) {
		return "Fine. Be that way!"
	}

	w := ""
	s = strings.TrimSpace(s)
	switch {
	case strings.ToUpper(s) == s && IsLetter(s) && !(s[len(s)-1:] == "?"):
		w = "Whoa, chill out!"
	case strings.ToUpper(s) == s && IsLetter(s):
		w = "Calm down, I know what I'm doing!"
	case s[len(s)-1:] == "?":
		w = "Sure."
	default:
		w = "Whatever."
	}
	return w
}
