package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	// "github.com/deliveryhero/basicdam"

	_ "github.com/lib/pq"
)

type data struct {
	AvailableCount int `json:"available_count"`
}
type output struct {
	Code int  `json:"status_code"`
	Data data `json:"data"`
}

var timeout int

func handler(w http.ResponseWriter, r *http.Request) {
	o := output{
		Code: 200,
		Data: data{
			AvailableCount: 10,
		},
	}
	out, _ := json.Marshal(o)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(out)
	time.Sleep(time.Microsecond * time.Duration(timeout))
}

func main() {
	timeoutPtr := flag.Int("timeout", 0, "timeout in microsecond")
	flag.Parse()
	timeout = *timeoutPtr
	http.HandleFunc("/", handler)

	db := connectDB()
	defer db.Close()

	initUser(db)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
