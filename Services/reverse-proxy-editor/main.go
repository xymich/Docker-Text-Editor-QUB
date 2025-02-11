package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func createReverseProxy(target string) http.Handler {
	targetURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s", r.Method, r.URL.String())
		log.Printf("Original Path: %s", r.URL.Path)

		// set CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// handle preflight reqs
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// remove any trailing "/" from URL path
		if strings.HasSuffix(r.URL.Path, "/") {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		// set the proxy URL path to include query parameters
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
		r.URL.RawPath = targetURL.Path + "?" + r.URL.RawQuery

		log.Printf("Modified Path: %s", r.URL.Path)
		log.Printf("Query: %s", r.URL.RawQuery)

		// serve request
		proxy.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// service endpoints and URLs
	services := map[string]string{
		"/scramble":   "http://localhost:8010/",
		"/reverse":    "http://localhost:8020/",
		"/duplicate":  "http://localhost:8030/",
		"/commacount": "http://localhost:8040/",
		"/wordcount":  "http://wordcount.editor.qpc.qubcloud.uk",
		"/charcount":  "http://charcount.editor.qpc.qubcloud.uk",
	}

	// make reverse proxy for each service
	for route, target := range services {
		proxy := createReverseProxy(target)
		log.Printf("Proxying %s to %s", route, target)
		mux.Handle(route, proxy)
	}

	log.Println("Reverse proxy server is running on port 7777")
	log.Fatal(http.ListenAndServe(":7777", mux))
}
