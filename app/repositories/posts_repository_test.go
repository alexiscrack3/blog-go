package repositories

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_GetPosts_ShouldReturn_Posts(t *testing.T) {
    testObject := NewPostsRepository()
    posts := testObject.GetPosts()
    assert.Equal(t, len(posts), 3, "there should be 3 posts")
}

func Test_GetPostByID_ShouldReturn_Post(t *testing.T) {
    testObject := NewPostsRepository()
    post, _ := testObject.GetPostByID(1)
    assert.NotNil(t, post, "post should not be nil")
}

func Test_CreatePost_ShouldCreate_Post(t *testing.T) {
    testObject := NewPostsRepository()
    reqBody := map[string]string {
        "title": "new",
        "body": "lorem ipsum",
    }
    post, err := testObject.CreatePost(reqBody)
    assert.NotNil(t, post, "post should not be nil")
    assert.Nil(t, err, "err is not nil")
}
