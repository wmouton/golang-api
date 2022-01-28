package main

import (
	// Standard library imports
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// Third party imports
	"github.com/gorilla/mux"
)

// Post struct represents the post data
type Post struct {
	Title	string	`json:"title"`
	Body	string	`json:"body"`
	Author	*Author	`json:"author"`
}

// Author struct represents user data
type Author struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var posts []Post

// Entry Point
func main() {
	// Indicates that the server is running on port 5000.
	fmt.Println("server running on port 5000")
	// Create a new mux router.
	router := mux.NewRouter()
	// Handler functions
	router.HandleFunc("/add", addPost).Methods("POST")
	// Listen and serve on port 5000 - log the errors
	log.Fatalln(http.ListenAndServe(":5000", router))
}

// addPost adds a new post
func addPost(w http.ResponseWriter, r *http.Request) {
	// Get the Post value from the JSON body
	var newPost Post
	// Decode the JSON data
	_ = json.NewDecoder(r.Body).Decode(&newPost)
	// Append each new post to the posts array
	posts = append(posts, newPost)
	// Send data back as JSON in the header
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Fatalln(err, "error parsing json data")
	}
}