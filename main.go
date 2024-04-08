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
	posts []Post = []Post{
		{ID: 1, Title: "Sample Post 1", Content: "This is the first sample post.", Author: "Author 1"},
		{ID: 2, Title: "Sample Post 2", Content: "This is the second sample post.", Author: "Author 2"},
	}
	comments []Comment = []Comment{
		{ID: 1, PostID: 1, Content: "This is a sample comment on the first post.", Author: "Commenter One"},
	}
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(posts) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "No posts found"})
		return
	}
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
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Post not found"})
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if post.Title == "" || post.Content == "" || post.Author == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing required post fields (title, content, author)"})
		return
	}

	maxID := 0
	for _, p := range posts {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	post.ID = maxID + 1

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
			// This check is usually done in the frontend
			if post.Title == "" || post.Content == "" || post.Author == "" {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "Missing required post fields (title, content, author)"})
				return
			}
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
	postId, err := strconv.Atoi(params["postId"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	postExists := false
	for _, p := range posts {
		if p.ID == postId {
			postExists = true
			break
		}
	}
	if !postExists {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var postComments []Comment
	for _, item := range comments {
		if item.PostID == postId {
			postComments = append(postComments, item)
		}
	}

	if len(postComments) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"message": "No comments found for this post"})
		return
	}

	json.NewEncoder(w).Encode(postComments)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	postId, err := strconv.Atoi(params["postId"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	var comment Comment

	// This check is usually done in the frontend
	// if comment.Content == "" || comment.Author == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(map[string]string{"error": "Missing required post fields (content, author)"})
	// 	return
	// }
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	maxID := 0
	for _, c := range comments {
		if c.ID > maxID {
			maxID = c.ID
		}
	}
	comment.ID = maxID + 1
	comment.PostID = postId

	comments = append(comments, comment)
	json.NewEncoder(w).Encode(comment)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "This route does not exist."})
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{"error": "This method is not allowed for this route."})
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

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)

	log.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
