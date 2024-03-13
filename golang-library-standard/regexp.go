package main

import (
	"fmt"
	"regexp"
)

// regular expression (regexp)
func main() {
	var regex = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("eko, edo, egi, ego, eto, e1o, eKo", 10))
}
