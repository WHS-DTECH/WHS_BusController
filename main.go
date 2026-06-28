package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	fileServer := http.FileServer(http.Dir("."))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve index.html for root and fall back to normal static file handling otherwise.
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
			return
		}
		fileServer.ServeHTTP(w, r)
	})

	log.Printf("serving static site on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
