package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestProxy() *httptest.Server {
	mux := http.NewServeMux()
	services := map[string]string{
		"/commacount": "http://localhost:8040/",
	}

	for route, target := range services {
		proxy := createReverseProxy(target)
		mux.Handle(route, proxy)
	}

	return httptest.NewServer(mux)
}

func TestProxyCommacount(t *testing.T) {
	startMockCommacountService()
	proxy := startTestProxy()
	defer proxy.Close()

	resp, err := http.Get(proxy.URL + "/commacount?text=hello,world")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	expected := map[string]interface{}{
		"error":  false,
		"string": "Comma count",
		"answer": float64(1),
	}

	if res["error"] != expected["error"] ||
		res["string"] != expected["string"] ||
		res["answer"] != expected["answer"] {
		t.Fatalf("Expected %v, got %v", expected, res)
	}
}
