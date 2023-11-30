package posts

import (
	"context"
	"encoding/json"
	"fmt"
	db "thanhbk113/db/sqlc"
	"thanhbk113/internal/constant"
	"thanhbk113/internal/module/redis"
	"thanhbk113/internal/query"

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

// GetPosts
func (p *postImpl) GetPosts(ctx context.Context, query query.CommonQuery) (dtores.PostResposeAll, error) {
	var (
		postResponseAll = dtores.PostResposeAll{
			PostResponse: make([]dtores.PostResponse, query.Limit),
			Total:        0,
			Limit:        query.Limit,
		}
	)

	//check db in redis or not
	cacheKey := constant.CachePosts

	cacheData, err := redis.GetValue(cacheKey)

	if err != nil {
		fmt.Println("err get data cache redis: ", err)
	}

	if cacheData != "" {
		err = json.Unmarshal([]byte(cacheData), &postResponseAll)

		if err != nil {
			return postResponseAll, err
		}

		return postResponseAll, nil
	}

	args := &db.ListPostsParams{
		Limit:  query.Limit,
		Offset: query.Page,
	}

	posts, err := initialize.GetDB().ListPosts(ctx, *args)

	if err != nil {
		return postResponseAll, err
	}

	total, err := initialize.GetDB().CountPosts(ctx)

	if err != nil {
		return postResponseAll, err
	}

	postResponseAll.Total = int(total)

	if len(posts) == 0 {
		return postResponseAll, nil
	}

	for i, post := range posts {
		postResponse := dtores.PostResponse{
			Title:     post.Title,
			Category:  post.Category,
			Content:   post.Content,
			Image:     post.Image,
			CreatedAt: post.CreatedAt.String(),
		}

		postResponseAll.PostResponse[i] = postResponse

	}

	//cache data to redis
	err = redis.SetKeyValue(cacheKey, postResponseAll, 15)

	if err != nil {
		return postResponseAll, err
	}

	return postResponseAll, nil
}
