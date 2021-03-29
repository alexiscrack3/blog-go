package repositories

import (
    "errors"
    "github.com/alexiscrack3/blog-go/app/models"
    "math/rand"
    "strconv"
)

const NOT_FOUND = -1

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

func (postsRepository *PostsRepository) DeletePostByID(id int) error {
    if id > 0 {
        posts := postsRepository.Posts
        index := getIndex(posts, id)
        if index == NOT_FOUND {
            return errors.New("Post was not found")
        } else {
            postsRepository.Posts = removePostByIndex(posts, index)
            return nil
        }
    } else {
        return errors.New("PostID should be greater than zero")
    }
}

func getIndex(posts []models.Post, postID int) int {
    for i := 0; i < len(posts); i++ {
        post := posts[i]
        if post.ID == postID {
            return i
        }
    }
    return NOT_FOUND
}

func removePostByIndex(posts []models.Post, index int) []models.Post {
    return append(posts[:index], posts[index + 1:]...)
}
