package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	rs := Response{
		Message: "Hi there!",
	}
	out, _ := json.Marshal(rs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(out)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
