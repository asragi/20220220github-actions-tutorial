package main

import "testing"

func TestGreeting(t *testing.T) {
	expect := "Hello"
	actual := Greeting()
	if expect != actual {
		t.Errorf("got %q want %q", actual, expect)
	}
}
