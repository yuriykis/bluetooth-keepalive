package util

import "strings"

func ClearString(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.TrimSpace(s)
	s = strings.TrimLeft(s, "\n")
	s = strings.TrimRight(s, "\n")
	return s
}
