package main

import (
    "github.com/alexiscrack3/blog-go/app/controllers"
    "github.com/alexiscrack3/blog-go/app/repositories"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
    router := mux.NewRouter()

    postsRepository := repositories.NewPostsRepository()
    postsControllers := controllers.NewPostsController(postsRepository)
    router.HandleFunc("/posts", postsControllers.GetPosts).Methods("GET")
    router.HandleFunc("/posts/{id}", postsControllers.GetPostByID).Methods("GET")
    router.HandleFunc("/posts", postsControllers.CreatePost).Methods("POST")

    log.Fatal(http.ListenAndServe(":3000", router))
}
