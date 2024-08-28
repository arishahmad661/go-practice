package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	messages := Message{Text: "Hello World"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func main() {
	http.HandleFunc("/message", getMessages)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
