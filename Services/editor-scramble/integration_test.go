package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// test server setup
func setup() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/scramble", scrambleHandler)
	return httptest.NewServer(mux)
}

// integration tests
func TestScrambleHandler_NoText(t *testing.T) {
	server := setup()
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/scramble", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", res.StatusCode)
	}

	expectedMap := map[string]interface{}{
		"error":  true,
		"string": "Missing 'text' parameter",
		"answer": "",
	}
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	if !isEqual(result, expectedMap) {
		t.Errorf("Expected %v, got %v", expectedMap, result)
	}
}

func TestScrambleHandler_EmptyText(t *testing.T) {
	server := setup()
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/scramble?text=", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", res.StatusCode)
	}

	expectedMap := map[string]interface{}{
		"error":  true,
		"string": "Missing 'text' parameter",
		"answer": "",
	}
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	if !isEqual(result, expectedMap) {
		t.Errorf("Expected %v, got %v", expectedMap, result)
	}
}

func TestScrambleHandler_ValidText(t *testing.T) {
	server := setup()
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/scramble?text=Hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", res.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	expectedMap := map[string]interface{}{
		"error":  false,
		"string": "Scrambled text",
	}
	if result["error"] != expectedMap["error"] || result["string"] != expectedMap["string"] || !hasSameChars(result["answer"].(string), "Hello") {
		t.Errorf("Expected %v, got %v", expectedMap, result)
	}
}

// compare two maps for equality
func isEqual(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// check if two strings contain the same characters
func hasSameChars(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, char := range a {
		if !strings.ContainsRune(b, char) {
			return false
		}
	}
	return true
}
