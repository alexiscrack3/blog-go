package repositories

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_GetPosts_ShouldReturnPosts(t *testing.T) {
    testObject := NewPostsRepository()
    posts := testObject.GetPosts()
    assert.Equal(t, len(posts), 3, "there should be 3 posts")
}

func Test_GetPostByID_ShouldReturnPost(t *testing.T) {
    testObject := NewPostsRepository()
    post, _ := testObject.GetPostByID(1)
    assert.NotNil(t, post, "post should not be nil")
}

func Test_CreatePost_ShouldCreatePost(t *testing.T) {
    testObject := NewPostsRepository()
    reqBody := map[string]string {
        "title": "new",
        "body": "lorem ipsum",
    }
    post, err := testObject.CreatePost(reqBody)
    assert.NotNil(t, post, "post should not be nil")
    assert.Nil(t, err, "err is not nil")
}

func Test_UpatePostByID_ShouldUpdatePost(t *testing.T) {
    testObject := NewPostsRepository()
    title := "new"
    body := "lorem ipsum"
    reqBody := map[string]string {
        "title": title,
        "body": body,
    }
    post, err := testObject.UpdatePostByID(1, reqBody)
    assert.Equal(t, post.Title, title, "title was not updated")
    assert.Equal(t, *post.Body, body, "body was not updated")
    assert.Nil(t, err, "err is not nil")
}

func Test_UpatePostByID_ShouldReturnError_WhenIdIsLessThanOrEqualToZero(t *testing.T) {
    testObject := NewPostsRepository()
    reqBody := make(map[string]string)
    _, err := testObject.UpdatePostByID(-1, reqBody)
    assert.NotNil(t, err, "err is nil")
}

func Test_UpatePostByID_ShouldReturnError_WhenIdIsNotFound(t *testing.T) {
    testObject := NewPostsRepository()
    reqBody := make(map[string]string)
    _, err := testObject.UpdatePostByID(100, reqBody)
    assert.NotNil(t, err, "err is nil")
}

func Test_DeletePost_ShouldDeletePost(t *testing.T) {
    testObject := NewPostsRepository()
    err := testObject.DeletePostByID(2)
    posts := testObject.GetPosts()
    assert.Equal(t, len(posts), 2, "there should be 2 posts")
    assert.Nil(t, err, "err is not nil")
}

func Test_DeletePost_ShouldReturnError_WhenIdIsLessThanOrEqualToZero(t *testing.T) {
    testObject := NewPostsRepository()
    err := testObject.DeletePostByID(-1)
    assert.NotNil(t, err, "err is nil")
}

func Test_DeletePost_ShouldReturnError_WhenIdIsNotFound(t *testing.T) {
    testObject := NewPostsRepository()
    err := testObject.DeletePostByID(100)
    assert.NotNil(t, err, "err is nil")
}
