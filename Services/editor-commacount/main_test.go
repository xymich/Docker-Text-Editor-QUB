package main

import (
	"strings"
	"testing"
)

func TestCountCommas(t *testing.T) {
	input := "Hello, World!"
	count := strings.Count(input, ",")
	if count != 1 {
		t.Errorf("Expected 1 comma, got %d", count)
	}
}

func TestCountNoCommas(t *testing.T) {
	input := "Hello World!"
	count := strings.Count(input, ",")
	if count != 0 {
		t.Errorf("Expected 0 commas, got %d", count)
	}
}

func TestCountMultipleCommas(t *testing.T) {
	input := "Hello,,, World!!!"
	count := strings.Count(input, ",")
	if count != 3 {
		t.Errorf("Expected 3 commas, got %d", count)
	}
}

func TestCountEmptyString(t *testing.T) {
	input := ""
	count := strings.Count(input, ",")
	if count != 0 {
		t.Errorf("Expected 0 commas, got %d", count)
	}
}

func TestCountLongString(t *testing.T) {
	input := "#########helloworld####helloohleootes tingonehello    ohleootestingonehelloohleootestingonehelloohleootestingonehelloohleootestingonehelloohleootestingonehelloohleootestingonehelloohleootestingonehelloohleootestingonehelloohleootestingonehelloohleootestingonehelloo         jhh b b   hleootestingone,wadawdadada"
	count := strings.Count(input, ",")
	if count != 1 {
		t.Errorf("Expected 1 commas, got %d", count)
	}
}
