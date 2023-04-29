package app2

import (
	"fmt"
	"regexp"
)

// testemail_test.go
// testemail.go

func IsEmail(s string) (string, error) {
	// Used a raw string here so I didn't have
	// to double backslashes
	r, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,3}`)
	if r.MatchString(s) {
		return "Valid Email", nil
	} else {
		return "", fmt.Errorf("Not a valid email")
	}
}

func main() {
}
