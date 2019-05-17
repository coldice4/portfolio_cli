package main

import (
	"bytes"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("AT4524GDFSDF\n"))

	result := getUserInput(&stdin)
	if result != "AT4524GDFSDF" {
		t.Fail()
	}
}
