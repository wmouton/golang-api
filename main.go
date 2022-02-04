package main

import (
	// Standard library imports
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	router.HandleFunc("/posts", addPost).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getSinglePost).Methods("GET")
	router.HandleFunc("/posts/{id}", updateSinglePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	// Listen and serve on port 5000 - log the errors
	log.Fatalln(http.ListenAndServe(":5000", router))
}

func getSinglePost(w http.ResponseWriter, r *http.Request) {
	// Get the ID of the a single post
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		log.Fatalln(w.Write([]byte("Could not convert ID to integer")))
		return
	}
	// Error checking on id
	if id >= len(posts) {
		w.WriteHeader(400)
		log.Fatalln(w.Write([]byte("No post found with specified ID")))
	}
	post := posts[id]
	// Set header content type
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(post); err != nil {
		log.Fatalln(err, "Could not encode JSON data")
	}
}

// getAllPosts returns all of the posts
func getAllPosts(w http.ResponseWriter, _ *http.Request) {
	// Set the header content type to get the data back in JSON format
	w.Header().Set("Content-Type", "application/json")
	// Handle error when encoding data
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Fatalln(err, "could not encode json data")
	}
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
	_ =  json.NewEncoder(w).Encode(posts)
}

func updateSinglePost(w http.ResponseWriter, r *http.Request) {
	// Get the ID of the post from the route parameters
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		log.Fatalln(w.Write([]byte("ID could not be converted to integer")))
		return
	}
	// Error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		log.Fatalln(w.Write([]byte("Post not found with specified ID")))
		return
	}
	// Get the value from the JSON body
	var updatedPost Post
	_ = json.NewDecoder(r.Body).Decode(&updatedPost)
	posts[id] = updatedPost
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedPost); err != nil {
		log.Fatalln(err)
	}
}

func patchPost(w http.ResponseWriter, r *http.Request) {
	
}