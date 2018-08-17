package wf

import (
	"fmt"
	"regexp"
	"strings"
)

func Π(s string) {
	fmt.Println(s)
}
func Σ(s string) string {
	return strings.ToUpper(s)
}

func Δ(s string) string {
	re := regexp.MustCompile("[0-9]+")
	return strings.Join(re.FindAllString(s, -1), ",")
}
