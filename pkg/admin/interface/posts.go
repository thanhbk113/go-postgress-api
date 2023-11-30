package interfaces

import (
	"context"
	"thanhbk113/internal/query"
	dtoreq "thanhbk113/pkg/admin/dto/request"
	dtores "thanhbk113/pkg/admin/dto/response"
)

type PostInterface interface {
	CreatePost(ctx context.Context, dto dtoreq.CreatePostRequest) error
	DeletePostById(ctx context.Context, postId string) error
	GetPostById(ctx context.Context, postId string) (dtores.PostResponse, error)
	TransactionLikePost(ctx context.Context, postId string) error
	TransactionDisLikePost(ctx context.Context, postId string) error
	GetPosts(ctx context.Context, query query.CommonQuery) (dtores.PostResposeAll, error)
}
