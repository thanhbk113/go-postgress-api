package handler

import (
	"fmt"
	gincustom "thanhbk113/internal/gin"

	dto "thanhbk113/pkg/admin/dto/request"
	i "thanhbk113/pkg/admin/interface"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostService i.PostInterface
}

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
