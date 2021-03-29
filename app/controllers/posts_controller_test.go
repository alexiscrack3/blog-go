package controllers

import (
    // "encoding/json"
    // "github.com/alexiscrack3/blog-go/app/repositories"
    // "net/http"
    // "net/http/httptest"
    // "testing"
)

// func Test_GetPosts_ShouldReturnPosts(t *testing.T) {
//     postsRepository := repositories.NewPostsRepository()
//     posts := postsRepository.GetPosts()
//     testObject := NewPostsController(postsRepository)

// 	req, err := http.NewRequest("GET", "/posts", nil)
// 	if err != nil {
//         t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(testObject.GetPosts)
//     handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
//         t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	// Check the response body is what we expect.
//     bytes, _ := json.Marshal(posts)
//     expected := string(bytes)
// 	if rr.Body.String() != expected {
//         t.Errorf("handler returned unexpected body: got %v want %v",
//         rr.Body.String(), expected)
// 	}
// }
