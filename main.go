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

	log.Fatal(http.ListenAndServe(":3000", router))
}
