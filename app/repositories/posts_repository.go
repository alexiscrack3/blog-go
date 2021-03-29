package repositories

import (
    "errors"
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

func (postsRepository PostsRepository) GetPostById(id int) (*models.Post, error) {
    for _, post := range postsRepository.Posts {
        if post.ID == id {
            return &post, nil
        }
    }
    return nil, errors.New("Post was not found")
}
