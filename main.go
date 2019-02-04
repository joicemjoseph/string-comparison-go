package exam

import (
	"regexp"
)

// Exam is
func Exam(a, b string) int {
	r, _ := regexp.Compile(b)
	return len(r.FindAllString(a, -1))
}
