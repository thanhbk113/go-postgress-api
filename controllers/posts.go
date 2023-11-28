package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"thanhbk113/schemas"
	"time"

	db "thanhbk113/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostController struct {
	db  *db.Queries
	ctx context.Context
}

func NewPostController(db *db.Queries, ctx context.Context) *PostController {
	return &PostController{db, ctx}
}

// [...] Create post handler
func (ac *PostController) CreatePost(ctx *gin.Context) {
	var payload *schemas.CreatePost

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	args := &db.CreatePostParams{
		Title:     payload.Title,
		Category:  payload.Category,
		Content:   payload.Content,
		Image:     payload.Image,
		CreatedAt: now,
		UpdatedAt: now,
	}

	post, err := ac.db.CreatePost(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "post": post})
}

// [...] Update post handler
func (ac *PostController) UpdatePost(ctx *gin.Context) {
	var payload *schemas.UpdatePost
	postId := ctx.Param("postId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()

	args := db.UpdatePostParams{
		ID:        uuid.MustParse(postId),
		Title:     sql.NullString{String: payload.Title, Valid: payload.Title != ""},
		Category:  sql.NullString{String: payload.Category, Valid: payload.Category != ""},
		Content:   sql.NullString{String: payload.Content, Valid: payload.Content != ""},
		Image:     sql.NullString{String: payload.Image, Valid: payload.Image != ""},
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
	}

	post, err := ac.db.UpdatePost(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "post": post})
}

// Get a single post handler
func (ac *PostController) GetPostById(ctx *gin.Context) {
	postId := ctx.Param("postId")

	post, err := ac.db.GetPostById(ctx, uuid.MustParse(postId))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "post": post})
}

// Get all posts handler
func (ac *PostController) GetAllPosts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	args := &db.ListPostsParams{
		Limit:  int32(intLimit),
		Offset: int32(offset),
	}

	posts, err := ac.db.ListPosts(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if posts == nil {
		posts = []db.Post{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "posts": posts})
}

// Delete a single post handler
func (ac *PostController) DeletePostById(ctx *gin.Context) {
	postId := ctx.Param("postId")

	err := ac.db.DeletePost(ctx, uuid.MustParse(postId))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

}