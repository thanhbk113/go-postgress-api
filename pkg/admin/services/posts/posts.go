package posts

import (
	"context"
	db "thanhbk113/db/sqlc"

	dto "thanhbk113/pkg/admin/dto/request"
	dtores "thanhbk113/pkg/admin/dto/response"
	i "thanhbk113/pkg/admin/interface"
	"thanhbk113/pkg/admin/server/initialize"

	"github.com/google/uuid"
)

func NewPostsService(ctx context.Context) i.PostInterface {
	return &postImpl{
		ctx: ctx,
	}
}

type postImpl struct {
	ctx context.Context
}

// CreatePost
func (p *postImpl) CreatePost(ctx context.Context, payload dto.CreatePostRequest) error {

	args := &db.CreatePostParams{
		Title:    payload.Title,
		Category: payload.Category,
		Content:  payload.Content,
		Image:    payload.Image,
	}

	_, err := initialize.GetDB().CreatePost(ctx, *args)

	if err != nil {
		return err
	}

	return nil
}

// DeletePostById
func (p *postImpl) DeletePostById(ctx context.Context, postId string) error {

	err := initialize.GetDB().DeletePost(ctx, uuid.MustParse(postId))

	if err != nil {
		return err
	}

	return nil
}

// GetPostById
func (p *postImpl) GetPostById(ctx context.Context, postId string) (dtores.PostResponse, error) {
	var (
		postResponse dtores.PostResponse
	)
	post, err := initialize.GetDB().GetPostById(ctx, uuid.MustParse(postId))

	if err != nil {
		return postResponse, err
	}

	postResponse = dtores.PostResponse{
		Title:     post.Title,
		Category:  post.Category,
		Content:   post.Content,
		Image:     post.Image,
		CreatedAt: post.CreatedAt.String(),
	}

	return postResponse, nil
}

// TransactionLikePost
func (p *postImpl) TransactionLikePost(ctx context.Context, postId string) error {

	db := initialize.GetDB()
	tx, err := initialize.GetSQLDB().BeginTx(ctx, nil)

	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := db.WithTx(tx)

	err = qtx.LikePost(ctx, uuid.MustParse(postId))

	if err != nil {
		return err
	}

	return tx.Commit()
}

// TransactionDislikePost
func (p *postImpl) TransactionDisLikePost(ctx context.Context, postId string) error {

	db := initialize.GetDB()
	tx, err := initialize.GetSQLDB().BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	qtx := db.WithTx(tx)

	err = qtx.DislikePost(ctx, uuid.MustParse(postId))

	if err != nil {
		return err
	}

	return tx.Commit()
}
