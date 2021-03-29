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

func (postsRepository PostsRepository) GetPostByID(id int) (*models.Post, error) {
    for _, post := range postsRepository.Posts {
        if post.ID == id {
            return &post, nil
        }
    }
    return nil, errors.New("Post was not found")
}

func (postsRepository *PostsRepository) CreatePost(reqBody map[string]string) (*models.Post, error) {
    length := len(postsRepository.Posts)
    newID := postsRepository.Posts[length - 1].ID + 1
    title := reqBody["title"]
    body := reqBody["body"]
    p := models.Post {
        ID: newID,
        Title: title,
        Body: &body,
        UserID: 1 + rand.Intn(100),
    }
    postsRepository.Posts = append(postsRepository.Posts, p)
    return &p, nil
}
