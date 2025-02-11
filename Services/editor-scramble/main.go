package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func scrambleText(text string) string {
	letters := []rune(text)
	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})
	return string(letters)
}

func scrambleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w,
			`{"error": true, "string": "Missing 'text' parameter", "answer": ""}`,
			http.StatusBadRequest)
		return
	}
	scrambledText := scrambleText(text)
	res := map[string]interface{}{
		"error":  false,
		"string": "Scrambled text",
		"answer": scrambledText,
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/scramble", scrambleHandler)
	fmt.Println("Scramble Service is running on port 8010")
	http.ListenAndServe(":8010", nil)
}
