package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Post struct represents the post data
type Post struct {
	Title	string	`json:"title"`
	Info	string	`json:"Info"`
	Author	*Author	`json:"author"`
}

// Author struct represents user data
type Author struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var posts []Post

func main() {
	fmt.Println("server running on port 5000")
	router := mux.NewRouter()
	router.HandleFunc("/add", addItem).Methods("POST")
	log.Fatalln(http.ListenAndServe(":5000", router))
}

func addItem(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newPost Post
	_ = json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Fatalln(err, "error parsing json data")
	}
}