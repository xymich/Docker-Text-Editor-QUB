package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/commacount?text=Hello,,World!", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	commacountHandler(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"answer":2}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestHandlerNoCommas(t *testing.T) {
	req, err := http.NewRequest("GET", "/commacount?text=HelloWorld!", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	commacountHandler(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"answer":0}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}

func TestHandlerMissingText(t *testing.T) {
	req, err := http.NewRequest("GET", "/countCommas", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	commacountHandler(w, req)

	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error":"Missing 'text' parameter"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", w.Body.String(), expected)
	}
}
