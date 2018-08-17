package wf

import (
	"fmt"
	"regexp"
	"strings"
)

// print on stdout
func Π(s string) {
	fmt.Println(s)
}

// convert to upper case
func Σ(s string) string {
	return strings.ToUpper(s)
}

// find all numbers in string
func Δ(s string) string {
	re := regexp.MustCompile("[0-9]+")
	return strings.Join(re.FindAllString(s, -1), ",")
}
