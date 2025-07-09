package api

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error writing response: %v", err)
		return

	}
}

func About(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "About")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error writing response: %v", err)
		return
	}
}
