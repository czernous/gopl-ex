package main

import (
	"fmt"
	"lissajous/lj"
	"log"
	"net/http"
	"strconv"
)

var severAddr = "0.0.0.0:6969"

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Server is running on http://%s\n", severAddr)

	// Start the server
	err := http.ListenAndServe(severAddr, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	cyclesQuery := r.Form.Get("cycles")

	cNum, err := strconv.Atoi(cyclesQuery)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid 'cycles' parameter"))
		return
	}

	lj.Lissajous(w, cNum)
}
