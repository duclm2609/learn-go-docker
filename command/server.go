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
		_, _ = fmt.Fprintf(w, "You've hit %v", os.Hostname())
	})

	fmt.Println("Server listening!")
	_ = http.ListenAndServe(":80", router)
}
