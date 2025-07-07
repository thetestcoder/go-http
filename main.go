package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8888"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	log.Println("Server start on port " + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}
