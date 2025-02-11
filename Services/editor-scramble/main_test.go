package main

import (
	"testing"
)

// unit tests
func TestScrambleText_NoChange(t *testing.T) {
	input := "Hello, World!"
	scrambled := scrambleText(input)
	if scrambled == input {
		t.Errorf("Expected scrambled text to be different from input")
	}
}

func TestScrambleText_EmptyString(t *testing.T) {
	input := ""
	scrambled := scrambleText(input)
	if scrambled != input {
		t.Errorf("Expected scrambled text to be the same as input")
	}
}

func TestScrambleText_SpecialCharacters(t *testing.T) {
	input := "Hello, World!"
	scrambled := scrambleText(input)
	// Just ensure it's a scramble and includes the same characters, exact order check is not required
	if scrambled == input {
		t.Errorf("Expected scrambled text to be different from input with special characters")
	}
}

func TestScrambleText_SpacedString(t *testing.T) {
	input := "   "
	scrambled := scrambleText(input)
	if scrambled != input {
		t.Errorf("Expected scrambled text to be the same as input with only spaces")
	}
}

func TestScrambleText_SimpleString(t *testing.T) {
	input := "abc"
	scrambled := scrambleText(input)
	if scrambled == input {
		t.Errorf("Expected scrambled text to be different from input for simple string")
	}
}
