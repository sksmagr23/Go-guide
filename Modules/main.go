package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mod in golang")
	gretter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func gretter() {
	fmt.Println("hello mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) { // first parameter is writer and second is reader
	w.Write([]byte("<h1>Welcome to Golang</h1>"))
}

// go mod tidy
// go mod graph
// go mod edit
// go mod vendor