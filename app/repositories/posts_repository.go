package repositories

import (
	"github.com/alexiscrack3/blog-go/app/models"
	"math/rand"
	"strconv"
)

type PostsRepository struct {
	Posts []models.Post
}

func NewPostsRepository() *PostsRepository {
	var posts []models.Post
	for i := 1; i < 4; i++ {
		p := models.Post {
			ID: i,
			Title: "title" + strconv.Itoa(i),
			Body: nil,
			UserID: 1 + rand.Intn(100),
		}
		posts = append(posts, p)
	}
	return &PostsRepository {
		Posts: posts,
	}
}

func (postsRepository PostsRepository) GetPosts() []models.Post {
	posts := postsRepository.Posts
	return posts
}
