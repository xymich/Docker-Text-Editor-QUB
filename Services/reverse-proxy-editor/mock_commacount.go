package main

import (
	"encoding/json"
	"net/http"
)

func commacountMockHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w,
			`{"error": true, "string": "Missing 'text' parameter", "answer": ""}`,
			http.StatusBadRequest)
		return
	}
	commaCount := 1 // Simplified for testing
	res := map[string]interface{}{
		"error":  false,
		"string": "Comma count",
		"answer": commaCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func startMockCommacountService() {
	http.HandleFunc("/commacount", commacountMockHandler)
	go http.ListenAndServe(":8040", nil)
}
