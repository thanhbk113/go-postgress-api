package request

import "thanhbk113/pkg/admin/schemas"

type CreatePostRequest struct {
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Image    string `json:"image" binding:"required"`
}

func (p CreatePostRequest) MapCreatePostRequestToSchema() schemas.CreatePost {
	return schemas.CreatePost{
		Title:    p.Title,
		Category: p.Category,
		Content:  p.Content,
		Image:    p.Image,
	}
}

type GetPostsRequest struct {
	Page  int `json:"page" binding:"required"`
	Limit int `json:"limit" binding:"required"`
}
