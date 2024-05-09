package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := HelloWithError("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`HelloWithError("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := HelloWithError("")
	if msg != "" || err == nil {
		t.Fatalf(`HelloWithError("") = %q, %v, want "", error`, msg, err)
	}
}
