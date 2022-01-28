package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type post struct {
	Title	string	`json:"title"`
	Body	string	`json:"body"`
	Author	string	`json:"author"`
}

var posts []post

func main() {
	fmt.Println("server running on port 5000")
	router := mux.NewRouter()
	router.HandleFunc("/add", addItem).Methods("POST")
	log.Fatalln(http.ListenAndServe(":5000", router))
}

func addItem(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newPost post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		log.Fatalln(err)
	}
	posts = append(posts, newPost)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Fatalln(err, "error parsing json data")
	}
}