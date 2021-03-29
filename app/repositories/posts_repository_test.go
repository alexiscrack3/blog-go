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
	post, _ := testObject.GetPostById(1)
	assert.NotNil(t, post, "there should be 1 post")
}
