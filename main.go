package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	log.Fatalln(http.ListenAndServe(":8080", router))
}

func test(w http.ResponseWriter, _ *http.Request) {
	log.Fatalln(w.Write([]byte("This is a test")))
}