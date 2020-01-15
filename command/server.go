package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		_, _ = fmt.Fprintf(w, "You've hit %s", hostname)
	})

	fmt.Println("Server listening!")
	_ = http.ListenAndServe(":80", router)
}
