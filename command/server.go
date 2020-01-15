package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Welcome to this life-changing API.")
	})

	fmt.Println("Server listening!")
	_ = http.ListenAndServe(":80", router)
}
