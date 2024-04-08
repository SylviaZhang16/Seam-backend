package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"postId"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var (
	posts    []Post    = []Post{}
	comments []Comment = []Comment{}
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts := []string{"Post 1", "Post 2"}
	json.NewEncoder(w).Encode(posts)
}

func getPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = len(posts) + 1 // Mock ID
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
}

func updatePostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if strconv.Itoa(item.ID) == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = item.ID
			posts = append(posts, post)
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	http.NotFound(w, r)
}

func deletePostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if strconv.Itoa(item.ID) == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var postComments []Comment
	for _, item := range comments {
		if strconv.Itoa(item.PostID) == params["postId"] {
			postComments = append(postComments, item)
		}
	}
	json.NewEncoder(w).Encode(postComments)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	postId, _ := strconv.Atoi(params["postId"])
	var comment Comment
	_ = json.NewDecoder(r.Body).Decode(&comment)
	comment.ID = len(comments) + 1
	comment.PostID = postId
	comments = append(comments, comment)
	json.NewEncoder(w).Encode(comment)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPostById).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", updatePostById).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePostById).Methods("DELETE")
	router.HandleFunc("/posts/{postId}/comments", getComments).Methods("GET")
	router.HandleFunc("/posts/{postId}/comments", createComment).Methods("POST")

	log.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
