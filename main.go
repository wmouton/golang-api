package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("server running on port 5000")
	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	log.Fatalln(http.ListenAndServe(":5000", router))
}

func test(w http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(w).Encode(struct { ID	string}{ID: "555"})
	if err != nil {
	log.Fatalln(err, "error writing header")
	}
}