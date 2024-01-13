package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	port := 8080

	http.HandleFunc("/api/message", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		message := Message{Text: "Hello, World!"}
		json.NewEncoder(w).Encode(message)
	})

	fmt.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
