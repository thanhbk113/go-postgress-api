package handler

import (
	"fmt"
	gincustom "thanhbk113/internal/gin"
	"thanhbk113/internal/query"
	"thanhbk113/internal/util"

	dto "thanhbk113/pkg/admin/dto/request"
	i "thanhbk113/pkg/admin/interface"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostService i.PostInterface
}

// Create godoc
// @tags Post
// @summary Create Post
// @id create-post
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body dto.CreatePostRequest true "Payload"
// @success 200 {object} nil
// @router /posts [post]
func (p *PostHandler) CreatePost(c *gin.Context) {

	var (
		payload dto.CreatePostRequest
		gg      = gincustom.GinGetCustomCtx(c)
		ctx     = gg.GetRequestCtx()
	)

	if err := gg.ShouldBindJSON(&payload); err != nil {
		gg.Response200(nil, err.Error())
	}

	err := p.PostService.CreatePost(ctx, payload)
	if err != nil {
		gg.Response400(nil, err.Error())
		return
	}
	gg.Response200(nil, "")
}

// Create godoc
// @tags Post
// @summary Delete Post
// @id delete-post
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Post Id"
// @success 200 {object} nil
// @router /posts/{id} [delete]
func (p *PostHandler) DeletePostById(c *gin.Context) {

	var (
		gg  = gincustom.GinGetCustomCtx(c)
		ctx = gg.GetRequestCtx()
	)

	postId := c.Param("postId")

	err := p.PostService.DeletePostById(ctx, postId)
	if err != nil {
		gg.Response400(nil, err.Error())
		return
	}
	gg.Response200(nil, "")

}

// Create godoc
// @tags Post
// @summary GetPostById
// @id get-post-by-id
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Post Id"
// @success 200 {object} nil
// @router /posts/{id} [get]
func (p *PostHandler) GetPostById(c *gin.Context) {

	var (
		gg  = gincustom.GinGetCustomCtx(c)
		ctx = gg.GetRequestCtx()
	)

	postId := c.Param("postId")

	fmt.Println(postId)

	post, err := p.PostService.GetPostById(ctx, postId)
	if err != nil {
		gg.Response400(nil, err.Error())
		return
	}
	gg.Response200(post, "")
}

// Create godoc
// @tags Post
// @summary GetListPost
// @id get-posts
// @security ApiKeyAuth
// @accept json
// @produce json
// @success 200 {object} nil
// @param payload query dto.GetPostsRequest true "Query"
// @router /posts [get]
func (p *PostHandler) GetPosts(c *gin.Context) {

	var (
		gg    = gincustom.GinGetCustomCtx(c)
		ctx   = gg.GetRequestCtx()
		query = query.CommonQuery{
			Limit: util.StringToInt(c.Query("limit")),
			Page:  util.StringToInt(c.Query("page")),
		}
	)

	post, err := p.PostService.GetPosts(ctx, query)
	if err != nil {
		gg.Response400(nil, err.Error())
		return
	}
	gg.Response200(post, "")
}

// Create godoc
// @tags Post
// @summary LikePost
// @id like-post
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Post Id"
// @success 200 {object} nil
// @router /post/{id} [patch]
func (p *PostHandler) LikePost(c *gin.Context) {

	var (
		gg  = gincustom.GinGetCustomCtx(c)
		ctx = gg.GetRequestCtx()
	)

	postId := c.Param("postId")

	err := p.PostService.TransactionLikePost(ctx, postId)
	if err != nil {
		gg.Response400(nil, err.Error())
		return
	}
	gg.Response200(nil, "")
}

func (p *PostHandler) DislikePost(c *gin.Context) {

	var (
		gg  = gincustom.GinGetCustomCtx(c)
		ctx = gg.GetRequestCtx()
	)

	postId := c.Param("postId")

	err := p.PostService.TransactionDisLikePost(ctx, postId)
	if err != nil {
		gg.Response400(nil, err.Error())
		return
	}
	gg.Response200(nil, "")
}
