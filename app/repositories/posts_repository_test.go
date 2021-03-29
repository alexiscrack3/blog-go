package repositories

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_GetPosts_ShouldReturn_Posts(t *testing.T) {
	testObject := NewPostsRepository()
	posts := testObject.GetPosts()
	assert.Equal(t, len(posts), 3, "there should be 3 posts")
}
