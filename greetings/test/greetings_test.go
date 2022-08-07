package test

import (
	"example/greetings"
	"regexp"
	"testing"
)

func TestHelloNameMatch(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg := greetings.Hello(name)
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("%s") = %q, want match for %#q, nil`, name, msg, want)
	}
}

func TestHelloWithErrorNameMatch(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := greetings.HelloWithError(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`HelloWithError("%s") = %q, %v, want %#q, nil`, name, msg, err, want)
	}
}

func TestHelloWithErrorEmpty(t *testing.T) {
	msg, err := greetings.HelloWithError("")
	if msg != "" || err == nil {
		t.Fatalf(`HelloWithError("") = %q, %v, want match for "", error`, msg, err)
	}
}
