package controllers

import (
    "encoding/json"
    "github.com/alexiscrack3/blog-go/app/repositories"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "strconv"
)

type UsersController struct {
    PostsRepository *repositories.PostsRepository
}

func NewUsersController(postsRepository *repositories.PostsRepository) *UsersController {
    return &UsersController {
        PostsRepository: postsRepository,
    }
}

func (usersController UsersController) GetPosts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    log.Printf("User ID = %s\n", params["id"])
    userID, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Println("User ID is not a string")
    } else {
        posts := usersController.PostsRepository.GetPostsByUserID(userID)
        json.NewEncoder(w).Encode(posts)
    }
}
