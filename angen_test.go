package main

import "testing"

func TestGenerate(t *testing.T) {
	password := Generate()
	if len(password) == 0 {
		t.Error("password is empty")
	}
}
