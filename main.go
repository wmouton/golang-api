package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var data []string

func main() {
	fmt.Println("server running on port 5000")
	router := mux.NewRouter()
	router.HandleFunc("/test", test).Methods("GET")
	router.HandleFunc("/add/{item}", addItem).Methods("POST")
	log.Fatalln(http.ListenAndServe(":5000", router))
}

func test(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(struct { ID	string}{ID: "555"})
	if err != nil {
	log.Fatalln(err, "error writing header")
	}
}

func addItem(w http.ResponseWriter, r *http.Request) {
	routeVar := mux.Vars(r)["item"]
	data = append(data, routeVar)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalln(err, "error parsing json data")
	}
}