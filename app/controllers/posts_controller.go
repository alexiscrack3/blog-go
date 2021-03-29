package controllers

import (
	"encoding/json"
	"github.com/alexiscrack3/blog-go/app/repositories"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type PostsController struct {
	PostsRepository *repositories.PostsRepository
}

func NewPostsController(postsRepository *repositories.PostsRepository) *PostsController {
	return &PostsController {
		PostsRepository: postsRepository,
	 }
}

func (postsController PostsController) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts := postsController.PostsRepository.GetPosts()
	json.NewEncoder(w).Encode(posts)
}

func (postsController PostsController) GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Printf("Post ID = %s\n", params["id"])
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("PostID is not a string")
	} else {
		post, err := postsController.PostsRepository.GetPostById(id)
		if err != nil {
			log.Println(err)
			json.NewEncoder(w).Encode(nil)
		} else {
			json.NewEncoder(w).Encode(post)
		}
	}
}
