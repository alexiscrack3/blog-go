package controllers

import (
	"encoding/json"
	"github.com/alexiscrack3/blog-go/app/repositories"
	"net/http"
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
