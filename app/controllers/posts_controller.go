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

func (postsController PostsController) GetPostByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    log.Printf("Post ID = %s\n", params["id"])
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Println("PostID is not a string")
    } else {
        post, err := postsController.PostsRepository.GetPostByID(id)
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusNotFound)
        } else {
            json.NewEncoder(w).Encode(post)
        }
    }
}

func (postsController PostsController) CreatePost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var reqBody map[string]string
    err := json.NewDecoder(r.Body).Decode(&reqBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    log.Println(reqBody)

    post, err := postsController.PostsRepository.CreatePost(reqBody)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(post)
    }
}

func (postsController PostsController) UpdatePost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    log.Printf("Post ID = %s\n", params["id"])

    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Println("PostID is not a string")
    } else {
        var reqBody map[string]string
        err := json.NewDecoder(r.Body).Decode(&reqBody)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        post, err := postsController.PostsRepository.UpdatePostByID(id, reqBody)
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
        } else {
            json.NewEncoder(w).Encode(post)
        }
    }
}

func (postsController PostsController) DeletePost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    log.Printf("Post ID = %s\n", params["id"])

    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Println("PostID is not a string")
    } else {
        err := postsController.PostsRepository.DeletePostByID(id)
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
        } else {
            w.WriteHeader(http.StatusNoContent)
        }
    }
}
