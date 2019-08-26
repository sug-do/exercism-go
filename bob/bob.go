package bob

import (
	"regexp"
	"strings"
)

var isLetter = regexp.MustCompile(`[a-zA-Z]+`).MatchString
var isValid = regexp.MustCompile(`[0-9a-zA-Z?!.]+`).MatchString

func Hey(s string) string {
	if !isValid(s) {
		return "Fine. Be that way!"
	}

	w := ""
	s = strings.TrimSpace(s)
	switch {
	case strings.ToUpper(s) == s && isLetter(s):
		if !(s[len(s)-1:] == "?") {
			w = "Whoa, chill out!"
		} else {
			w = "Calm down, I know what I'm doing!"
		}
	case s[len(s)-1:] == "?":
		w = "Sure."
	default:
		w = "Whatever."
	}
	return w
}
