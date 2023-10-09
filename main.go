package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Hello struct {
	Title string
	Desc  string
	Ver   string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Hello{Title: "Hello", Desc: "BuildPack", Ver: "0.0.1"})
	})
	port := os.Getenv("PORT")

	log.Printf("Listening on %v", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
