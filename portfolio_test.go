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

func TestGetUserInputDecimal(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("\n"))

	_, err := getUserInputDecimal(&stdin)
	if err != nil {
		t.Error(err)
	}
}
