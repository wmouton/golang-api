package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type item struct {
	Data	string
}

var data []item

func main() {
	fmt.Println("server running on port 5000")
	router := mux.NewRouter()
	router.HandleFunc("/add", addItem).Methods("POST")
	log.Fatalln(http.ListenAndServe(":5000", router))
}

func addItem(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newItem item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		log.Fatalln(err)
	}
	data = append(data, newItem)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalln(err, "error parsing json data")
	}
}