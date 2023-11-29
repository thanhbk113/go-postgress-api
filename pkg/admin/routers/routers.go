package routers

import (
	"context"
	"thanhbk113/internal/middleware"
	"thanhbk113/pkg/admin/services/posts"

	"github.com/gin-gonic/gin"
)

func Init(c *gin.Engine) {
	ctx := context.Background()
	c.Use(middleware.CORSConfig())

	r := c.Group("/api")

	postsv := posts.NewPostsService(ctx)

	PostRoute(r, postsv)

}
