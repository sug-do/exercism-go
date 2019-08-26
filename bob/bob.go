package bob

import (
	"regexp"
	"strings"
)

var IsLetter = regexp.MustCompile(`[a-zA-Z]+`).MatchString
var IsPunc = regexp.MustCompile(`[?!.]$`).MatchString
var IsNumber = regexp.MustCompile(`[0-9]`).MatchString
var IsValid = regexp.MustCompile(`[0-9a-zA-Z?!.]+`).MatchString
var q, o, u bool
var w string

func Hey(s string) string {
	w := ""
	a := "Sure."
	b := "Whoa, chill out!"
	c := "Calm down, I know what I'm doing!"
	d := "Fine. Be that way!"
	e := "Whatever."
	switch {
	//determine if a valid string
	case string_or_nah(s):
		//trim whitespace from valid string
		s := strings.TrimSpace(s)
		switch {
		//determine if letters are in the sentence
		case IsLetter(s):
			switch {
			//determine if all upper case
			case up_or_nah(s):
				switch {
				//if all upper case, is it a question?
				case ques_or_nah(s):
					w = c
				default:
					w = b
				}
			//if not all upper case, is it a question?
			case ques_or_nah(s):
				w = a
			default:
				w = e
			}
		//if no letters, is it a question?
		case ques_or_nah(s):
			w = a
		default:
			w = e
		}
	//everything else
	default:
		w = d
	}
	return w
}

//determines statement not empty
func string_or_nah(s string) bool {
	switch {
	case IsValid(s):
		q = true
	default:
		q = false
	}
	return q
}

//determines if remark is all caps
func up_or_nah(s string) bool {
	switch {
	case strings.ToUpper(s) == s:
		u = true
	default:
		u = false
	}
	return u
}

//determines if remark is a question
func ques_or_nah(s string) bool {
	switch {
	case s[len(s)-1:] == "?":
		o = true
	default:
		o = false
	}
	return o
}
