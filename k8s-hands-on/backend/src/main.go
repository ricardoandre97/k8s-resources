package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := HandsOn{
		Time:     time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":9090", nil)
}
