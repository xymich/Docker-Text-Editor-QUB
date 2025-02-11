package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func commacountHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w,
			`{"error": true, "string": "Missing 'text' parameter", "answer": ""}`,
			http.StatusBadRequest)
		return
	}
	commaCount := strings.Count(text, ",")
	res := map[string]interface{}{
		"error":  false,
		"string": "Comma count",
		"answer": commaCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/commacount", commacountHandler)
	log.Println("Commacount service is running on port 8040")
	log.Fatal(http.ListenAndServe(":8040", nil))
}
