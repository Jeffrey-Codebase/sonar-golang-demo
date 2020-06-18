package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != helloString {
		t.Fatal("Test fail")
	}
}
