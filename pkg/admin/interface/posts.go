package interfaces

import (
	"context"
	dtoreq "thanhbk113/pkg/admin/dto/request"
	dtores "thanhbk113/pkg/admin/dto/response"
)

type PostInterface interface {
	CreatePost(ctx context.Context, dto dtoreq.CreatePostRequest) error
	DeletePostById(ctx context.Context, postId string) error
	GetPostById(ctx context.Context, postId string) (dtores.PostResponse, error)
}